package clouds

type Downloader interface {
	DownLoadFile(fileId, dst string)
}

type Uploader interface {
	UploadFile(filename string) string
}

// описание для функций
// общий интерфейс для пакета
// реализация для яндекса
// тесты! - исп-ть fileIdByName и fileNameById для проверки появился ли файл
