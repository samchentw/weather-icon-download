package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DownloadFile(URL, fileName string, isDay bool) error {
	url := URL
	response, e := http.Get(url)

	data, err := ioutil.ReadAll(response.Body)

	folder := "早上"

	if !isDay {
		folder = "晚上"
	}

	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	ioutil.WriteFile("images/"+folder+"/"+fileName, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success!")
	return nil
}

func CreateFileOrRead(folder string) {
	path := "images/" + folder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}
