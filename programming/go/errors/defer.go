package main

import(
	"fmt"
	"os"
)


func proverbs(name string) error{
	f, err := os.Create(name)
	if err != nil{
		return err
	}
	defer f.Close() // у нас тут гарантировано будет исполнение функции до возврата из функции

	_, err = fmt.Fprintf(f, "Какой-то текст, очень интересно.")
	if err != nil{
		return err
	}
	_, err = fmt.Fprintf(f, "еще один очень инетересный текст.")
	return err
}

func main(){
	e := proverbs("интересно")
	if e != nil{
		fmt.Println("Чот фигня, Максимыч.")
		os.Exit(1)

	}
}