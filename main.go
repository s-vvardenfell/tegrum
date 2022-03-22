package main

import "github.com/s-vvardenfell/Backuper/cmd"

func main() {
	cmd.Execute()

	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8888")

	// t := &archiver.Zip{}
	// someArchive(t)

	// if err := t.Archive("resources/test_file_48kb.txt", "resources"); err != nil {
	// 	log.Fatal(err)
	// }

	// archiver.Gzip("resources/test_file_48kb.tar", "resources")

	// if err := t.Extract("resources/test_file_48kb.tar.gz", "resources"); err != nil {
	// 	log.Fatal(err)
	// }
}

/*
TODO

#Архивация
-общий код в tar и zip - рефакторинг
-рефакторинг архивации - дб возвр значение мб - имя/путь сформированного архива
-Пароль для архивов
-сбор нескольких архивов в 1 для tar и zip (сейчас архивация неск-х архивов)

На счет Uploader и Downloader интерфейсов - точно ли они должны возвращать значение? почему не ошибку? как используется это значение?
мб нужно вызывать интерфейс Store который будет сохранять куда-то данные от Uploader
UploadFile подумать чтобы возвращал ошибку, а не просто пустую строку как в telegram!
должны ли вообще почта и тг следовать этим интерфейсам?

#Общее
Сообщения об успешной отправке/загрузке везде
Шифрование для smtp? Или хватит пароля на архив?
Readme сделать, все нужные инструкции
logrus
работа как фоновый процесс?
yandex реализация


dst + "/" + filename
везде где есть "/" нужно заменить вызовом ф-ии из пакета http/filepath, которая конкатенирует корректно
сделать тоже для параметров

-переименовать в tegrum + команды
-задачи из todo.txt сюда

# Очистка
tegrum clean //удалит архивы старше, чем опция

# Получение
tergum retrieve -g -y -t //скачивает последние бекап-архивы
и мб даже разархивация в нужные пути но хз, долго
*/
