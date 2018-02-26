package tool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadConfigFromFile(fileName string, model interface{}) error {
	var _fileName string
	if fileName == "" {
		_fileName = "config.json"
	}
	// Open config file
	file, err := os.Open(_fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonStruct, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Json handle
	err = json.Unmarshal(jsonStruct, &model)
	if err != nil {
		log.Fatal("LOAD CONFIG ERROR:", err)
		panic("")
	}
	return err
}
