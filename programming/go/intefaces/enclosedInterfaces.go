package main

import "fmt"

type Reader interface{
	read() 
}

type Writer interface{
	write(text string)
}
/* для соответствия этом интерфейсу тип должен реализовывать
   методы для обоих интерфейсов
*/
type ReadWrite interface{ // тут типа вложенные интерфейсы
	Reader
	Writer
}

func writeToStream(w Writer, text string){
	w.write(text)
}

func readerFromStream(r Reader) {
	r.read()
}

type File struct{
	name string
	text string
}

type Folder struct{
	name string
}

func (f *File) read() {
	fmt.Println(f.text)

}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Текст", message, "записан в файл")
}

func (f *Folder) close(){
	fmt.Println("Папка закрыта")
}

func (f *Folder) write(){
	fmt.Println("Запись в папку")
}

func main(){
	var(
		file ReadWrite = &File{}	

	)
	writeToStream(file, "Новая стрчока")
	readerFromStream(file)
}