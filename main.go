package main

import (
	"fmt"
	"log"
	"os"

	"github.com/s-vvardenfell/tegrum/storages"
)

func main() {
	// cmd.Execute()

	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8888")

	file, err := os.OpenFile("result/data.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	data := []string{"some", "data", "tostore"}
	st := storages.CsvStorage{}

	err = st.Store(file, data)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.OpenFile("result/data.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	res, err := st.Retrieve(file, "some")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	file.Close()
}

/*
TODO
в backup вообще не исп-ся archiveName, т.е. возвращаемое имя архива не нужно функции?
мб исправить gzipping tar'ов - с этой переменной логика будет проще
удалять архив .tar после того как создался .gz

#Архивация
-общий код в tar и zip - рефакторинг - мб можно объединить в 1 ф-ю
-рефакторинг архивации - дб возвр значение мб - имя/путь сформированного архива
-Пароль для архивов
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
