package main

import (
	"fmt"
	//"io"
	"os"
)
//一个成功的输入输出样例
func main(){

	reader,_ := os.Open("1")
	writer,_ := os.Create("2")
	var buffer = make([]byte,1024)
	for true{
		n,_ := reader.Read(buffer)
		fmt.Println("start!")
		if n == 0 {
			break
		}
			writer.Write(buffer[:n])

	}

	defer func(){
		writer.Close()
		reader.Close()
	}()
}