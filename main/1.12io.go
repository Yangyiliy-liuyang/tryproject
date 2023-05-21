package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	//1.打开文件 Open只读文件
	fileObj, err := os.Open("src/gocode/project01/main/xx.txt")
	if err != nil {
		fmt.Printf("Open file failed,err:%v\n", err)
		return
	}
	//读取文件的内容
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed,err:%v\n", err)
			return
		}
	}(fileObj) //关闭文件
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("Read file failed,err:%v\n", err)
		return
	}
	fmt.Printf("read %d bytes from file.\n", n)
	fmt.Println(string(tmp))
}
func readAll() {
	//1.打开文件 Open只读文件
	fileObj, err := os.Open("src/gocode/project01/main/xx.txt")
	if err != nil {
		fmt.Printf("Open file failed,err:%v\n", err)
		return
	}
	//读取文件的内容
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed,err:%v\n", err)
			return
		}
	}(fileObj) //关闭文件
	for true {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { //end of file
			//把
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("Read file failed,err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp))
	}
}
func readByBuff() {
	//1.打开文件 Open只读文件
	fileObj, err := os.Open("src/gocode/project01/main/xx.txt")
	if err != nil {
		fmt.Printf("Open file failed,err:%v\n", err)
		return
	}
	//读取文件的内容
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed,err:%v\n", err)
			return
		}
	}(fileObj) //关闭文件
	reader := bufio.NewReader(fileObj)
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF { //end of file
			//读到最后，把最后一行读出来。
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Printf("read file filed,err%v\n", err)
			return
		}
		fmt.Print(line)
	}
}
func readByIoutil() {
	content, err := ioutil.ReadFile("src/gocode/project01/main/xx.txt")
	if err != nil {
		fmt.Printf("read file filed,err%v\n", err)
		return
	}
	fmt.Println(string(content))
}
func write() {
	fileobj, err := os.OpenFile("src/gocode/project01/main/xx.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("read file filed,err%v\n", err)
		return
	}
	defer fileobj.Close()
	str := "vuejs开发"
	fileobj.Write([]byte(str))
	fileobj.WriteString("\nhello,zhangsan")
}
func writeByBufio() {
	fileobj, err := os.OpenFile("src/gocode/project01/main/xx.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("read file filed,err%v\n", err)
		return
	}
	defer fileobj.Close()
	writer := bufio.NewWriter(fileobj)
	writer.WriteString("计算机科学与技术") //本操作将内容写入缓冲区
	writer.Flush()                         //将缓冲区的内容写入磁盘
}
func writeByIoutil() {
	ioutil.WriteFile("src/gocode/project01/main/xx.txt", []byte("nihao"), 0644)

}
func main() {
	writeByIoutil()
}
