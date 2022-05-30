package lexer

import (
	"io/ioutil"
	"unsafe"
)

// 将 .c 的内容读入为字符串

func ReadFile(name string, path string) string {
	content, err := ioutil.ReadFile(path + name)
	if err != nil {
		panic(err)
	}

	return bytes2String(content)
}

func bytes2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
