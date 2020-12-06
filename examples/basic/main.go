package main

import (
	"github.com/schollz/progressbar/v3"
	"time"
)

func main() {
	bar := progressbar.Default(100)
	//bar.RenderBlank()
	for i := 0; i < 100; i++ {
		bar.WriteLog("HelloWorld")

		//bar.RenderBlank()

		//bar.Add(1)
		bar.Add(1)
		time.Sleep(time.Second)
		//time.Sleep(1 * time.Second)

		//time.Sleep(time.Second*1)
		//err := bar.Add(1)
		//if err!=nil{
		//	fmt.Println(err)
		//}
		//time.Sleep(40 * time.Millisecond)
	}
}
