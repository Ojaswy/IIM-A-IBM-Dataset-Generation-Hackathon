package main

import (
	"gopkg.in/jdkato/prose.v2"
	"fmt"
	"math/rand"
	"time"
)

var dataFileLoc string = "data/dataset.json"
var typeFileLoc string = "data/types.json"
var submissionFileLoc string = "output.json"

func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	fmt.Println("Process begins\n")
	defer fmt.Println("\nProcess Terminated")
	datasetJson := ReadJson(dataFileLoc)
	typesJson := ReadJson(typeFileLoc)
	var s []Output
	patterns := []string{
		"XXXX XXXX XXXX XXXX",
		"XXXX XXXX XXXX",
		"XXXX XXXX'",
		"XXXX",
		"XX/XX/XXXX",
		"XX/XX/"}

	for i := 0; i < len(datasetJson.Data); i++ {
		var obj Output
		var entity string
		var types string
		found := false
		entry := datasetJson.Data[i]
		fmt.Println("Entry No: ", i)

		doc, _ := prose.NewDocument(entry)
		for _, ent := range doc.Entities() {

			if InArray(ent.Text, patterns) {

				switch ent.Label {
				case "PERSON":
					types = typesJson.Data[5]
					entity = RandomStr(typesJson.Data[5])
				case "ORGANIZATION":
					types = typesJson.Data[19] + " " + typesJson.Data[20]
					entity = RandomStr(typesJson.Data[19])
				case "GPE":
					types = typesJson.Data[19] + " " + typesJson.Data[20]
					entity = RandomStr(typesJson.Data[19])
				}
				found = true
			}
		}

		if !found {
			ran := rand.Intn(len(typesJson.Data))
			types = typesJson.Data[ran]
			entity = RandomStr(types)
		}

		obj.Text = Replace(entry, entity)
		obj.Entity = entity
		obj.Types = types
		s = append(s, obj)
	}

	CreateJson(s)
}
