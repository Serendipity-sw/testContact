package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/smtc/glog"
	"fmt"
)

func main() {
	aesMob,err:=AesEncrypt("18551734732")
	if err != nil {
		fmt.Printf("aes err! err: %s \n",err.Error())
		return
	}
	fmt.Println(aesMob)
	enMob:="FrUW5i7w3CeMkqePA3OMIw=="
	mob,err:=AesDecrypt(enMob)
	if err != nil {
		fmt.Printf("unAes err! err: %s \n",err.Error())
		return
	}
	fmt.Println(mob)
}

var AESKEY = []byte("sfe023f_9fd&fwfl")

/**
字符串加密
创建人:邵炜
创建时间:2016年3月18日09:50:36
输入参数: 需要加密的字符串
输出参数: 加密后字符串 错误对象
*/
func AesEncrypt(origData string) (string, error) {
	origDataByte := []byte(origData)
	block, err := aes.NewCipher([]byte(AESKEY))
	if err != nil {
		glog.Error("aes newcipher is error, err: %s \n", err.Error())
		return "", err
	}
	blockSize := block.BlockSize()
	origDataByte = PKCS5Padding(origDataByte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, AESKEY[:blockSize])
	crypted := make([]byte, len(origDataByte))
	blockMode.CryptBlocks(crypted, origDataByte)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

/**
字符串解密
创建人:邵炜
创建时间:2016年3月18日09:56:20
输入参数: 需要解密的字符串  解密后字符串长度
输出参数: 解密后字符串  错误对象
*/
func AesDecrypt(crypted string) (string, error) {
	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	block, err := aes.NewCipher(AESKEY)
	if err != nil {
		glog.Error("aes newcipher is error, err: %s \n", err.Error())
		return "", err
	}
	blockSize := block.BlockSize()

	if len(cryptedByte)%blockSize != 0 {
		glog.Error("crypto/cipher: input not full blocks , crypted: %s  \n", crypted)
		return "", errors.New("crypto/cipher: input not full blocks")
	}

	blockMode := cipher.NewCBCDecrypter(block, AESKEY[:blockSize])
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)
	origData = PKCS5UnPadding(origData)

	return string(origData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
