package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func HandleError(err error, errMsg string) bool {
	if err != nil {
		fmt.Println(errMsg)
		fmt.Println(err)
		return false
	}
	return true
}
 
func ReadJson(file string) InputData {
	fileData, err := os.Open(file)
	HandleError(err, file+" file not found.")
	defer fileData.Close()
	byteData, err := ioutil.ReadAll(fileData)
	HandleError(err, "ERROR: File not opened")
	jsonData := ByteToJson(byteData)
	return jsonData
}
//Parsing
func ByteToJson(byteData []byte) InputData {
	var jsonData InputData
	err := json.Unmarshal(byteData, &jsonData)
	HandleError(err, "ERROR: Parsing")
	return jsonData
}

func CreateJson(newData []Output) {
	file, err := json.MarshalIndent(OutputData{newData}, "", " ")
	HandleError(err, "Unable to marshal")
	err = ioutil.WriteFile(submissionFileLoc, file, 0644)
	HandleError(err, "Unable to create a file")
}
