package main

type InputData struct {
	Data []string `json:"data"`
}
type OutputData struct {
	Data []Output `json:"YashOju"`
}
type Output struct {
	Text   string `json:"text"`
	Entity string `json:"entity"`
	Types  string `json:"types"`
}
