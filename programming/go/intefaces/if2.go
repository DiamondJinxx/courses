package main

import "fmt"

type Reader interface{
	read() 
}

type Writer interface{
	write(text string)
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

// реализуем интерфейс Reader
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

func (f *Folder) write(text string){
	fmt.Println("Запись в папку текста",text)
}

func main(){
	var(
		file *File = &File{}	
		folder *Folder = &Folder{}
	)

	writeToStream(file, "Новая стрчока")
	readerFromStream(file)
	writeToStream(folder, "Новая стрчока в папке")
	readerFromStream(folder)
}