package io

import (
	"os"
)

func WriteFile(name string, path string, content string) {
	file, err := os.OpenFile(path+name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	file.WriteString(content + "\n")
}
