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
внутри архиватора полем можно хранить его расширение
внутри типов telegram, email можно хранить поле- название типа
переписать backup.go
добавить флаги --csv и тд - для реализации полиморфизма хранилищ - добавить 1 нереализованное хранилище

#Архивация
-добавление неск файлов в 1 архив, не создавать папки-пути: в итоговом архиве дб только файлы(?)

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
