package main

import (
	"os"
	"log"
	"path"
	"io/ioutil"
)

var (
	root string
)

func getRoot() string {
	if root == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		root = dir
	}
	return root
}

func getAbsPath(customPath string) (absPath string) {
	absPath = path.Join(root, customPath)
	return
}

func ReadFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Fatal(err)
	}
	return data
}

func WriteFile(filePath string, data []byte) {
	ioutil.WriteFile(filePath, data, 0644)
}
