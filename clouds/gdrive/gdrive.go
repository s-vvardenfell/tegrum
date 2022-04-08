package gdrive

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// go get -u google.golang.org/api/drive/v3
// go get -u golang.org/x/oauth2/google

var ErrFileNotFound = errors.New("file not found")

const ext = "gdrive"

type GDrive struct {
	Srv       *drive.Service
	extension string
}

func NewGDrive(credentials string) *GDrive {
	ctx := context.Background()

	b, err := ioutil.ReadFile(credentials)
	if err != nil {
		logrus.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logrus.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return &GDrive{Srv: srv, extension: ext}
}

func (gd *GDrive) Extension() string {
	return gd.extension
}

// Download file by filename in directory dst
func (gd *GDrive) DownLoadFile(filename, dst string) error {
	if fileId, err := gd.fileIdByName(filename); err == nil {
		dst = filepath.Join(dst, filename)

		f, err := gd.Srv.Files.Get(fileId).Download()
		if err != nil {
			return fmt.Errorf("unable to download file: %v", err)
		}

		defer func() { _ = f.Body.Close() }()

		respByte, err := ioutil.ReadAll(f.Body)
		if err != nil {
			return fmt.Errorf("unable to read from file: %v", err)
		}

		if err = os.WriteFile(dst, respByte, 0666); err != nil {
			return fmt.Errorf("unable to save data to file: %v", err)
		}
	} else {
		return fmt.Errorf("cannot get file id by given name for downloading")
	}
	return nil
}

// Uploads file to disk and returns its id if success
// Deletes file with the same name
func (gd *GDrive) UploadFile(filename string) (string, error) {
	if id, err := gd.fileIdByName(filename); err == nil {
		err = gd.deleteFile(id)
		if err != nil {
			return "", fmt.Errorf("cannot delete old file by id %s", filename)
		}
	} else {
		return "", fmt.Errorf("cannot get file id by name %s", filename)
	}

	baseMimeType, err := mimetype.DetectFile(filename)
	if err != nil {
		return "", fmt.Errorf("cannot detect mime type: %v", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}

	fileInf, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("unable get file's stats: %v", err)
	}

	defer func() { _ = file.Close() }()

	f := &drive.File{Name: filename}
	res, err := gd.Srv.Files. //TODO use not deprecated
					Create(f).
					ResumableMedia(context.Background(), file, fileInf.Size(), baseMimeType.String()).
					ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
					Do()

	if err != nil {
		return "", fmt.Errorf("error while uploading file: %v", err)
	}
	return res.Id, nil
}

func (gd *GDrive) ShowFilesList() error {
	r, err := gd.Srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve files: %v", err)
	}

	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
	return nil
}

// Delete file by given file id
func (gd *GDrive) deleteFile(fileId string) error {
	f := gd.Srv.Files.Delete(fileId)

	if err := f.Do(); err != nil {
		return fmt.Errorf("error while deleting file: %v", err)
	}
	return nil
}

// Returns file id if present or empty string if no such file found
func (gd *GDrive) fileIdByName(filename string) (string, error) {
	r, err := gd.Srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve files: %v", err)
	}

	if len(r.Files) == 0 {
		return "", ErrFileNotFound
	} else {
		for _, i := range r.Files {
			if i.Name == filename {
				return i.Id, nil
			}
		}
	}
	return "", ErrFileNotFound
}

// Returns filename if present or empty string if no such file found
func (gd *GDrive) fileNameById(fileid string) (string, error) {
	r, err := gd.Srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve files: %v", err)
	}

	if len(r.Files) == 0 {
		return "", ErrFileNotFound
	} else {
		for _, i := range r.Files {
			if i.Id == fileid {
				return i.Name, nil
			}
		}
	}
	return "", ErrFileNotFound
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		logrus.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		logrus.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		logrus.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer func() { _ = f.Close() }()
	json.NewEncoder(f).Encode(token)
}
