package main

import (
	"bufio"
	"fmt"
	"github.com/swgloomy/go-common"
	"io"
	"os"
	"strings"
)

var arrayMap map[string]string = make(map[string]string)

func main() {
	err := ReadFileByLine("./gloomy-20170322")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var one []byte
	for key, _ := range arrayMap {
		if len(strings.TrimSpace(key)) != 0 {
			one = append(one, []byte(fmt.Sprintf("%s\n", key))...)
		}
	}
	err = common.FileCreateAndWrite(&one, "./gloomy-20170322123", false)
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
文件读取逐行进行读取
创建人:邵炜
创建时间:2016年9月20日10:23:41
输入参数: 文件路劲
输出参数: 字符串数组(数组每一项对应文件的每一行) 错误对象
*/
func ReadFileByLine(filePath string) error {
	var (
		readAll  = false
		readByte []byte
		line     []byte
		err      error
	)
	fs, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {
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
			arrayMap[string(line)] = ""
			line = line[:0]
		}
	}
	return nil
}
