package php

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Basename(file_path string) string {
	return filepath.Base(file_path)
}

func File_get_contents(url string) string {

	if strings.Index(url, "http") == 0 {
		return string(file_get_contents_url(url))
	}

	if strings.Index(url, "https") == 0 {
		return string(file_get_contents_url(url))
	}

	return string(file_get_contents_file(url))
	// (strings.Split("a,b,c,d,e,f,g", ",")) // [a b c d e f g]
}

func file_get_contents_url(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func file_get_contents_file(url string) []byte {

	file, err := os.Open(url)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return body
}

func File_put_contents(fileName string, write_data string) int {
	file, _ := os.Create(fileName)
	defer file.Close()

	wrote_byte, _ := file.Write([]byte(write_data))
	file.Sync()

	return wrote_byte
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func InArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
