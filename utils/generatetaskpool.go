package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	branchAndEmailMap map[string]string
	result            = make([]map[string]string, 0)
	singleMap         = make(map[string]string, 0)
)

func GenerateTaskPool() []map[string]string {
	file, err := os.Open("participant.json")
	if err != nil {
		log.Println(time.Now().String()+":Error opening file:", err)
		return nil
	}
	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(time.Now().String()+":Error reading file:", err)
		return nil
	}
	err = json.Unmarshal(jsonBytes, &branchAndEmailMap)
	if err != nil {
		log.Println(time.Now().String()+":Error decoding JSON:", err)
		return nil
	}
	taskPool := turnMapToArraymap(branchAndEmailMap)
	return taskPool
}

func turnMapToArraymap(mapu map[string]string) []map[string]string {

	for key, value := range mapu {
		singleMap[key] = value
		result = append(result, singleMap)
	}
	return result
}
