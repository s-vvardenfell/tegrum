# tegrum
Backuper written in go<br>

The common way to use is type `store` command, select archive type zip/tar, multi- or one-file mode, storages to upload<br>
After uploading archives their id can be saved and archives can be retrieved with the command `retrieve`

## store
The following flags are available for `store`<br>
Specify archiver:<br>
`--zip` - use Zip archiver-compressor<br>
`--tar` - use Tar archiver and Gzip compressor<br>

Specify mode:<br>
`-o` - "one-file" mode means that tegrum will archive and upload the file that will be specified in `-s` flag<br>
`-m` - "multi-file" mode means that tegrum will try to read the specified file as a list of files for archiving and sending to storages<br>
For this mode, specify the files and directories for archiving by the list in the file of the form
```json
{
	"dirs":[	
	"C:/Some_important_dir",
	"D:/Another_important_dir"
	]
}
```

Specify where file id records will be placed:<br>
`--csv` - use csv-file to store<br>
If no recorder is specified, the data will not be saved and cannot be used to `retrieve` archives<br>

To specify source and destination dirs<br>
`-s` for source (file or list of files in json format for archiving and downloading)<br>
`-d` for destination (directory where the archive will be located)<br>

Storages that can be used:<br>
`-g` Google Drive<br>
`-y` Yandex Disk (not implemented yet)<br>
`-t` Telegram<br>
`-e` Email<br>

#### Examples
Using the single-file mode, we simply archive Important_doc.pdf, upload the archive to GDrive, YaDisk, Telegram, send it via email, save the file ids received from all storages, except email, to a csv file<br>
```
tegrum backup -o --zip --csv -s D:/resources/important_doc.pdf -d D:/result -g -y -t -e
```
Using multi-file mode, we archive all files from the Important_docs_list.json list, upload the archive to GDrive<br>
```
tegrum backup -m --tar -s D:/resources/important_docs_list.json -d D:/result -g
```
## retrieve
The following flags are available for `retrieve`<br>
Specify from where file id records will be read:<br>
`--csv` - use csv-file to read<br>
If the recorder is not specified, the data will not be read, an error will occur<br>

To specify destination dir<br>
`-d` - dir where the archive will appear<br>

Storages that can be used:<br>
`-g` Google Drive<br>
`-y` Yandex Disk (not implemented yet)<br>
`-t` Telegram<br>

#### Example
tegrum tries to get archives that were downloaded earlier using the file id stored in the csv file<br>
```
tegrum retrieve --csv -d D:/result -g -y -t
```
## requirements
To use repositories, you need to prepare them<br>
#### Google Drive
Using [this](https://developers.google.com/drive/api/guides/about-sdk) and [this](https://developers.google.com/workspace/guides/get-started) get credentials.json file from Google, place it to resources/credentials.json<br>
```json
{
    "installed": {
        "client_id": "CLIENT_ID",
        "project_id": "PROJECT_ID",
        "auth_uri": "https://accounts.google.com/o/oauth2/auth",
        "token_uri": "https://oauth2.googleapis.com/token",
        "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
        "client_secret": "CLIENT_SECRET",
        "redirect_uris": [
            "urn:ietf:wg:oauth:2.0:oob",
            "http://localhost"
        ]
    }
}
```
#### Yandex Disk
Not implemented<br>
#### Telegram
Create telegram-bot using [botfather]<br>
(https://core.telegram.org/bots)
Record the received token and the ID of the selected chat in the telegram configuration file (resources/telegram.json)
```json
{
    "token":"TOKEN",
    "chat_id":"CHAT ID"
}
```
#### Email
To use sending by email, fill in the email configuration file (resources/email.json)<br>
```json
{
    "sender": "EMAIL",
    "user": "EMAIL",
    "passw": "PASSWORD",
    "address": "smtp.<EMAIL_SERVER_ADDRESS>:25",
    "host": "smtp.<EMAIL_SERVER_ADDRESS>",
    "to": [
        "MAIL_TO_1",
        "MAIL_TO_2"
    ],
    "subject": "SUBJECT",
    "body": "BODY"
}
```
