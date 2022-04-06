# tegrum
Backuper written in go<br>

The common way to use is type command `store`, select zip/tar archive type, multiple- or one-file mode, storages to upload<br>
Uploaded files id can be saved and files can be retrieved with the command `retrieve`

## store
The following flags are available for `store`<br>
Specify archiver:<br>
`--zip` - use Zip archiver<br>
`--tar` - use Tar archiver and Gzip compressor<br>

Specify mode:<br>
`-o` - "one-file" mode means that tegrum will archive and upload file that will be specified in `-s` flag<br>
`-m` - "multiple-file" mode means that tegrum will attempt to read the specified file as a list of files to archive and send to storages<br>

Specify where file id records will be placed:<br>
`--csv` - use csv-file to store and read<br>
If no recorder specified, no data will be store, and will be no data to use with `retrieve`<br>

To specify source and destination dirs<br>
`-s` for source (file or list of files to archive and upload)<br>
`-d` for destination (dir where the archive will appear)<br>

Storages can be used:<br>
`-g` Google Drive<br>
`-y` Yandex Disk (not implemented yet)<br>
`-t` Telegram<br>
`-e` Email<br>

#### Examples
Using one-file mode, just zip important_doc.pdf, upload archive to GDrive, YaDisk, Telegram, send to email, store file id returned from all storages exept email to csv-file<br>
```
tegrum backup -o --zip --csv -s D:/resources/important_doc.pdf -d D:/result -g -y -t -e
```
Using multi-file mode, tar all files that listed in important_docs_list.json, upload archive to GDrive<br>
```
tegrum backup -m --tar -s D:/resources/important_docs_list.json -d D:/result -g
```
## retrieve
Specify from where file id records will be read:<br>
`--csv` - use csv-file to read<br>
If no recorder specified, no data will be read, an error<br>

To specify destination dir<br>
`-d` for destination (dir where the archive will appear)<br>

Storages can be used:<br>
`-g` Google Drive<br>
`-y` Yandex Disk (not implemented yet)<br>
`-t` Telegram<br>

#### Example
tegrum tries to download archives that was uploaded earlier using file id, stored in csv-file<br>
```
tegrum retrieve --csv -d D:/result -g -y -t
```


