package main

import (
	"log"
	"strconv"

	"github.com/s-vvardenfell/tegrum/archiver"
)

const source = "W:/Golang/src/Backuper/resources/test_file_20kb.txt"
const target = "W:/Golang/src/Backuper/result/"

func GenerateWords(btns []string) [][]string {
	cnt := len(btns) * len(btns)
	words := make([][]string, cnt)
	// var temp string
	for j := 0; j < len(btns); j++ {
		for i := range words {
			words[i] = append(words[i], strconv.Itoa(i))
		}
	}

	return words
}

func main() {

	t := archiver.Tar{}
	if err := t.Archive(source, target); err != nil {
		log.Fatal(err)
	}

	// cmd.Execute()

	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8888")

	// file, err := os.OpenFile(writer, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := Tar(target, file); err != nil {
	// 	log.Fatal(err)
	// }

	// data := []string{"some", "data", "tostore"}
	// st := &storages.CsvStorage{}
	// temp := fmt.Sprintf("%T", st)
	// fmt.Println(temp[strings.Index(temp, ".")+1:])

	// err = st.Store(file, data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()

	// file, err = os.OpenFile("result/data.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := st.Retrieve(file, "some")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)
	// file.Close()
}

/*
TODO
проблема с именоанием - storage и repositories, archives etc

#Архивация
-общий код в tar и zip - рефакторинг - можно объединить в 1 ф-ю; сделать gzip внутри tar
-рефакторинг архивации - дб возвр значение мб - имя/путь сформированного архива
-сбор нескольких архивов в 1 для tar и zip (сейчас архивация неск-х архивов)

#Общее
Сообщения об успешной отправке/загрузке везде + логи logrus
Шифрование для smtp? Или хватит пароля на архив?
Readme сделать, все нужные инструкции
тесты для всех пакетов
getTokenFromWeb и прочие возврат ошибки сделать?

Clouds
описание для функций
реализация для яндекса
тесты! - исп-ть fileIdByName и fileNameById для проверки появился ли файл



dst + "/" + filename
везде где есть "/" нужно заменить вызовом ф-ии из пакета http/filepath, которая конкатенирует корректно
сделать тоже для параметров

-переименовать в tegrum + команды
-задачи из todo.txt сюда

# Получение
tergum retrieve -g -y -t //скачивает последние бекап-архивы
и мб даже разархивация в нужные пути но хз, долго
*/
