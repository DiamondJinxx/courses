package main

import(
	"os"
	"fmt"
)

func proverbs(name string) error{
	f, err := os.Create(name)
	if err != nil{
		return err
	}
	_, err = fmt.Fprintf(f, "Текст про ошибки.")
	if err != nil{
		f.Close()
		return err
	}
	_, err = fmt.Fprintf(f, "Еще одна строчка про ошибки.")
	f.Close()
	return err

}	

func main(){
	err := proverbs("text.txt")
	if err != nil{
		fmt.Println("Чот фигня, Максимыч.")
		os.Exit(1)
	}
}