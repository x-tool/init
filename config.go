package tool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadConfigFromFile(model interface{}) (error) {
	// Open config file
	file, err := os.Open("config.json")
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
