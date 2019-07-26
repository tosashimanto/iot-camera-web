package util

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateImageFile(fileName string, m []byte) {

	out, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	ioutil.WriteFile(fileName, m, os.ModePerm)

}
