package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	contetArrayList, err := readFileByLine("../bcyReptile/coserPhoto.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	htmlFile, err := os.OpenFile("./coserPhoto.html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer htmlFile.Close()
	htmlFile.Write([]byte("<!DOCTYPE html> <html lang=\"en\"> <head> <meta charset=\"UTF-8\"> <title>Title</title> </head> <body>"))
	var array []string
	for _, value := range *contetArrayList {
		array = strings.Split(value, " ")
		htmlFile.Write([]byte(fmt.Sprintf("<div><a href=\"%s\">%s</a></div>", array[0], array[1])))
	}
	htmlFile.Write([]byte("</body> </html>"))
}

/**
文件读取逐行进行读取
创建人:邵炜
创建时间:2016年9月20日10:23:41
输入参数: 文件路劲
输出参数: 字符串数组(数组每一项对应文件的每一行) 错误对象
*/
func readFileByLine(filePath string) (*[]string, error) {
	var (
		readAll     = false
		readByte    []byte
		line        []byte
		err         error
		contentLine []string
	)
	fs, err := os.Open(filePath)
	if err != nil {
		//glog.Error("readFileByLine open error! filePath: %s err: %s \n", filePath, err.Error())
		return nil, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {
			//glog.Error("readFileByLine read error! err: %s \n", err.Error())
		}
		if readAll {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
		} else {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
			if len(strings.TrimSpace(string(line))) == 0 {
				continue
			}
			contentLine = append(contentLine, string(line))
			line = line[:0]
		}
	}
	//glog.Info("readFileByLine run success! filePath: %s fileContent: %v \n", filePath, contentLine)
	return &contentLine, nil
}
