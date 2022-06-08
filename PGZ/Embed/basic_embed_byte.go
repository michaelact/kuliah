package main

import (
	"io/ioutil"
	"io/fs"
	_ "embed"
)

//go:embed logo.jpg
var file []byte

func main() {
	err := ioutil.WriteFile("logo_new.jpg", file, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
