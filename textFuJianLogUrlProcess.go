package main

import (
	"fmt"
	"github.com/swgloomy/go-common"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	fileArray, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var filePathArray []string
	for _, value := range fileArray {
		if strings.HasPrefix(value.Name(), "INFO-") {
			filePathArray = append(filePathArray, value.Name())
		}
	}
	fileWrite, err := common.AppendFileOpen(fmt.Sprintf("./gloomy-%s", common.DateFormat(time.Now(), "yyyyMMdd")))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, value := range filePathArray {
		common.RWFileByWhere(fmt.Sprintf("./%s", value), fileWrite, whereFun)
	}
	fileWrite.Close()
}

func whereFun(content string, fileWrite *os.File) {
	var (
		arrayList []string
		urlObj    *url.URL
		err       error
	)
	if strings.Index(content, "feedback: uuid=") > 0 {
		arrayList = strings.Split(content, "referer: ")
		if len(arrayList) != 2 {
			return
		}
		urlObj, err = url.Parse(arrayList[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err = fileWrite.Write([]byte(fmt.Sprintf("%s\n", urlObj.Host)))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
