package goProject01

import (
	"os"
)

// Copy 一个成功的拷贝方法
func Copy(source string,target string)(info int){
	info = 1
	reader,_ := os.Open(source)
	writer,_ := os.Create(target)
	var buffer = make([]byte,1024)
	for true{
		n,_ := reader.Read(buffer)
		if n == 0 {
			break
		}
			writer.Write(buffer[:n])

	}
	info = 0
	defer func(){
		writer.Close()
		reader.Close()
	}()
	return
}
