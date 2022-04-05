package main

import "github.com/s-vvardenfell/tegrum/cmd"

const source = "W:/Golang/src/Backuper/resources/test_file_20kb.txt"
const target = "W:/Golang/src/Backuper/result/"

func main() {

	cmd.Execute()

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
переписать backup.go
проблема с именованием - storage и repositories, archives etc - record все еще не очень нравится название
исправить все тесты - генерация тест-файла, os.GetWd для пути

#Архивация
-рефакторинг архивации - дб возвр значение мб - имя/путь сформированного архива
-тесты для архиваторов
-добавление неск файлов в 1 архив, не создавать папки-пути: в итоговом архиве дб только файлы(?)

#Общее
Сообщения об успешной отправке/загрузке везде + логи logrus
Шифрование для smtp? Или хватит пароля на архив?
Readme сделать, все нужные инструкции
тесты для всех пакетов
getTokenFromWeb и прочие возврат ошибки сделать?
LOGRUS везде

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
