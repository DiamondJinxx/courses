package main

import "fmt"

type Stream interface{
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string){
	stream.write(text)
}

func closeStream(stream Stream){
	stream.close()

}

type File struct{
	name string
	text string
}

type Folder struct{}

func  newFile(fileName string) (*File) {
	file := new(File)
	file.name = fileName
	return file
}

func (f File) read() string{
	return f.text
}
/* такой прикол, что реализация методов для типа *File не предусматривает автоматическую реализация для типа File
	что есть забавно. При попытке реализовать методы для типа File - компилятор выдаст ошибку переопределения методов.
	Но все же можно создать статический объект(тут такой терминологии нет) и передать адрес,
	тогда считается, что интер
*/
func (f File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл", f.name, " строки: ", message)
}

func (f File) close() {
	fmt.Println("Файл закрыт")
}

func (f *Folder) close() {
	fmt.Println("Папка закрыта")
}


func main(){
	file := newFile("Димин файл")
	var f File
	f.name = "Новый файл"
	// file.write("Новая запись в файл")
	fmt.Printf("%T\n", file)
	fmt.Printf("%T\n", f)
	writeToStream(file, "Новая запись в файл")
	writeToStream(f, "Новая запись в другой файл")
	f.write("")
	closeStream(file)
}