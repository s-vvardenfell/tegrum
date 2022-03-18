package clouds

type Uploader interface {
	UploadFile(filename string) string
}

type Downloader interface {
	DownLoadFile(fileId, dst string)
}

// описание для функций
// общий интерфейс для пакета
// реализация для яндекса
// тесты! - исп-ть fileIdByName и fileNameById для проверки появился ли файл
