package clouds

type Cloud interface {
	DownLoadFile(fileId, dst string)
	UploadFile(filename string) string
}

// описание для функций
// общий интерфейс для пакета
// реализация для яндекса
// тесты! - исп-ть fileIdByName и fileNameById для проверки появился ли файл
