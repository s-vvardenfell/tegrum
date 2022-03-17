package main

import (
	"os"

	"github.com/s-vvardenfell/Backuper/telegram"
)

type TelegramUploadResponse struct {
	Ok     bool                 `json:"ok"`
	Result TelegramUploadResult `json:"result"`
}

type TelegramDownloadResponse struct {
	Ok     bool                     `json:"ok"`
	Result TelegramResponseDocument `json:"result"`
}

type TelegramUploadResult struct {
	Document TelegramResponseDocument `json:"document"`
}

type TelegramResponseDocument struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FilePath     string `json:"file_path"`
}

type TelegramDownloadError struct {
	Ok        bool    `json:"ok"`
	ErrorCode float64 `json:"error_code"`
	Descr     string  `json:"description"`
}

func main() {
	// cmd.Execute()

	// fmt.Println(time.Now().Format("02.Jan.2006_15:04:05"))
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8888")
	tg := telegram.NewTelegram("resources/telegram.json")
	id := tg.UploadFile("resources/test_file.txt")
	tg.DownLoadFile(id, "resources/")

	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.Size())
	// }
}

/*
TODO сделать это todo на темы

#Telegram:
-тесты!

#Архивация
-создание папки с именем архива! и послед-я ее архивация
-общий код в tar и zip - рефакторинг
-gzip для tar-архива/архивов

#Mail
вместо конфига - newMail(config string)

-сбор нескольких архивов в 1 для tar и zip
пока сделан только сбор в одну папку нескольких архивов и послед-я архивация
если я создаю архив с уникальным именем, то они будут копиться в хранилищах
т.о. надо либо сделать команду на удаление(+) либо просто делать 1 статичное имя и хранить 1 экз в хранилищах (-)

На счет Uploader и Downloader интерфейсов - точно ли они должны возвращать значение? почему не ошибку? как используется это значение?
мб нужно вызывать интерфейс Store который будет сохранять куда-то данные от Uploader
UploadFile подумать чтобы возвращал ошибку, а не просто пустую строку как в telegram!

#Общее

dst + "/" + filename
везде где есть "/" нужно заменить вызовом ф-ии из пакета http, которая конкатенирует корректно
сделать тоже для параметров

возможно, download и upload все же должны возвращать ошибку

-переименовать в tegrum + команды
-задачи из todo.txt сюда

tegrum clean //удалит архивы старше, чем опция
tergum retrieve -g -y -t

ФУНКЦИОНАЛ ЗАГРУЗКИ бекап-файлов обратно
и мб даже разархивация в нужные пути но хз, долго

почта должна следовать интерфейсу cloud? она не сможет "загружать" файлы
по сигнатуре тоже не подходит даже под Upload, надо либо в папку clouds перенести, переименовать на storages,
и интерфейс тоже, либо оставить в пакете почты и не юзать интерфейс, в таком случае можно снова объединить интерфейс Cloud в 1

Рефакторинг - конспект

исп в отдельном файле для ответов от апи
назыв "entities"
type TelegramResponseResult struct {
	FileId   string `json:"file_id"`
	FilePath string `json:"file_path"`
}
type TelegramResponse struct {
	Ok     bool                   `json:"ok"`
	Result TelegramResponseResult `json:"result"`
}

выносить все адреса и тд в константы
url := fmt.Sprintf("%s%s/getFile?file_id=%s", BASE_URL, botToken, fileId) использовать


Config - структура
исп filepath.Abs("config.yaml")

*/
