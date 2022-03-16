package main

import "github.com/s-vvardenfell/Backuper/cmd"

type DirsToBackup struct {
	Dirs []string `json:"dirs"`
}

func main() {
	cmd.Execute()

	// archiver.Tarr("resources", "resources/")
	// archiver.Untar("resources/resources.tar", "resources/")

	// tar := archiver.Tar{}
	// tar.Archive("resources/sum.txt", "resources/")

	// zip := archiver.Zip{}
	// if err := zip.Archive("resources/sum.txt", "resources/"); err != nil {
	// 	log.Fatal(err)
	// }

	// archiver.Gzip("resources/sum.txt", "resources/")

}

/*
общий код в tar и zip - рефакторинг

мб стоит сделать для backup дочернюю команду, которая отправляет архивы

gzip для tar-архива/архивов
сбор нескольких архивов в 1 для tar и zip

CLI
tergum	daemon	start -p
		stop
-p - устанавливает периодичность

	backup	-g -y -t -e


-g - бекап в gdrive
-y - бекап в ya-disk
-t - бекап в tg
-e - бекап на почту

сделать чтобы файл-конфиг и файл выход подавались в аргументах

config читается в любом случае
директорию, куда сохранять архив - указать в конфиге?


нужны еще команды/опции tar или zip!
режим работы "конфиг" и "cli" вот решение мб

tegrum! поменять имя

tergum backup -a tar -g -y -t -e
tergum backup -a zip -g -y -t -e

tergum retrieve -g -y -t

ФУНКЦИОНАЛ ЗАГРУЗКИ бекап-файлов обратно

почта должна следовать интерфейсу cloud? она не сможет "загружать" файлы
по сигнатуре тоже не подходит даже под Upload, надо либо в папку clouds перенести, переименовать на storages,
и интерфейс тоже, либо оставить в пакете почты и не юзать интерфейс, в таком случае можно снова объединить интерфейс Cloud в 1

*/
