package main

import "github.com/s-vvardenfell/tegrum/cmd"

const source = "W:/Golang/src/Backuper/resources/test_file_20kb.txt"
const target = "W:/Golang/src/Backuper/result/"

func main() {
	// r := csv_record.CsvRecorderRetriever{}

	// file, err := os.Open("W:/Golang/src/Backuper/resources/data.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(r.Retrieve(file, "telegram"))
	cmd.Execute()

	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8888")
}

/*
TODO
парсинг файла со списком файлов/папок для архивации - упростить json-файл, убрать название "dirs:"
добавить флаги --csv и тд - для реализации полиморфизма хранилищ - добавить 1 нереализованное хранилище
readme с инструкциями

#Архивация
-добавление неск файлов в 1 архив, не создавать папки-пути: в итоговом архиве дб только файлы(?)
retrieve - распаковка - после

#Общее
Сообщения об успешной отправке/загрузке везде + логи logrus
Readme сделать, все нужные инструкции
тесты для всех пакетов
getTokenFromWeb и прочие возврат ошибки сделать?

Clouds
описание для функций
реализация для яндекса
тесты! - исп-ть fileIdByName и fileNameById для проверки появился ли файл - проверить всё для google

-задачи из todo.txt сюда

# Получение
tergum retrieve -g -y -t //скачивает последние бекап-архивы
*/
