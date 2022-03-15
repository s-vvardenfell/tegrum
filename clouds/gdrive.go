package clouds

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// go get -u google.golang.org/api/drive/v3
// go get -u golang.org/x/oauth2/google

type GDrive struct {
	Srv *drive.Service
}

//TODO уточнить, как правильно делать такие конструкторы, что возвращать
func NewGDrive() *GDrive {
	ctx := context.Background()

	b, err := ioutil.ReadFile("resources/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return &GDrive{Srv: srv}
}

// Download file by filename in directory dst
// TODO CUT FILEPATH
func (gd *GDrive) DownLoadFile(filename, dst string) {

	if fileId := gd.fileIdByName(filename); filename != "" {
		dst = dst + filename

		f, err := gd.Srv.Files.Get(fileId).Download()

		if err != nil {
			log.Fatalf("Unable to download file: %v", err)
		}

		defer f.Body.Close()
		respByte, err := ioutil.ReadAll(f.Body)

		if err != nil {
			log.Fatalf("Unable to read from file: %v", err)
		}

		if err = os.WriteFile(dst, respByte, 0666); err != nil {
			log.Fatalf("Unable to save data to file: %v", err)
		}
	} else {
		log.Fatal("Cannot get file id by given name for downloading")
	}
}

// Uploads file to disk and returns its id if success
// Deletes file with the same name
func (gd *GDrive) UploadFile(filename string) string {
	if id := gd.fileIdByName(filename); id != "" {
		gd.deleteFile(id)
	}

	baseMimeType := "text/plain" //TODO сделать авто-определение mime type

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	fileInf, err := file.Stat()
	if err != nil {
		log.Fatalf("Unable get file's stats: %v", err)
	}

	defer file.Close()

	f := &drive.File{Name: filename}
	res, err := gd.Srv.Files. //TODO исп-ть не устаревший метод
					Create(f).
					ResumableMedia(context.Background(), file, fileInf.Size(), baseMimeType). //TODO у имени файла убрать путь
					ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
					Do()

	if err != nil {
		log.Fatalf("Error while uploading file: %v", err)
	}
	return res.Id
}

func (gd *GDrive) ShowFilesList() {
	r, err := gd.Srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	fmt.Println("Files:")

	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}

// Delete file by given file id
func (gd *GDrive) deleteFile(fileId string) {
	f := gd.Srv.Files.Delete(fileId)

	if err := f.Do(); err != nil {
		log.Fatalf("Error while deleting file: %v", err)
	}
}

// Returns file id if present or empty string if no such file found
func (gd *GDrive) fileIdByName(filename string) string {
	r, err := gd.Srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	if len(r.Files) == 0 {
		return ""
	} else {
		for _, i := range r.Files {
			if i.Name == filename {
				return i.Id
			}
		}
	}
	return ""
}

// Returns filename if present or empty string if no such file found
func (gd *GDrive) fileNameById(fileid string) string {
	r, err := gd.Srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	if len(r.Files) == 0 {
		return ""
	} else {
		for _, i := range r.Files {
			if i.Id == fileid {
				return i.Name
			}
		}
	}
	return ""
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
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
