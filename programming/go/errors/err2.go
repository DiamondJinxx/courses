/*в этой программе инкапсулируем работу с ошибками*/

package main

import (
	"os"
	"io"
	"fmt"
)

type safeWriter struct{
	w io.Writer
	err error
}

func (sw *safeWriter) writeln(s string){
	if sw.err != nil{
		return
	}
	_, sw.err = fmt.Fprintf(sw.w, s)
}

func proverbs(name string) error{
	f, err := os.Create(name)
	if err != nil{
		return err
	}
	// sw := safeWriter{w:f}
	var sw safeWriter 
	sw.w = f
	sw.writeln("Errors are values.")
    sw.writeln("Don't just check errors, handle them gracefully.")
    sw.writeln("Don't panic.")
    sw.writeln("Make the zero value useful.")
    sw.writeln("The bigger the interface, the weaker the abstraction.")
    sw.writeln("interface{} says nothing.")
    sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
    sw.writeln("Documentation is for users.")
    sw.writeln("A little copying is better than a little dependency.")
    sw.writeln("Clear is better than clever.")
    sw.writeln("Concurrency is not parallelism.")
    sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
    sw.writeln("Channels orchestrate; mutexes serialize.")
    return sw.err
}

func main(){
	err := proverbs("слишкоминтересно")
	if err != nil{
		fmt.Printf("Ошибка %v\n", err)
		os.Exit(1)
	}
}
