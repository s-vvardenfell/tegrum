package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Tar(src string, writers ...io.Writer) error {

	// ensure the src actually exists before trying to tar it
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("Unable to tar files - %v", err.Error())
	}

	mw := io.MultiWriter(writers...)

	gzw := gzip.NewWriter(mw)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	// walk path
	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {

		// return on any error
		if err != nil {
			return err
		}

		// return on non-regular files (thanks to [kumo](https://medium.com/@komuw/just-like-you-did-fbdd7df829d3) for this suggested update)
		if !fi.Mode().IsRegular() {
			return nil
		}

		// create a new dir/file header
		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		// update the name to correctly reflect the desired destination when untaring
		header.Name = strings.TrimPrefix(strings.Replace(file, src, "", -1), string(filepath.Separator))

		// write the header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		// manually close here after each file operation; defering would cause each file close
		// to wait until all operations have completed.
		f.Close()

		return nil
	})
}

const target = "W:/Golang/src/Backuper/resources/test_file_48kb.txt"
const writer = "W:/Golang/src/Backuper/resources/map.tar"

func main() {
	// cmd.Execute()

	// os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8888")

	file, err := os.OpenFile(writer, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if err := Tar(target, file); err != nil {
		log.Fatal(err)
	}

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
