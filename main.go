package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"problem-solving/service"
	"time"
)

var (
	branchAndEmailMap map[string]string
	repo              string = "sspu-cryptography-research-group/problem-solving"
	//repo        = "OrangeDou/QHTLC"
)

func main() {
	//Read the participant.json
	file, err := os.Open("participant.json")
	if err != nil {
		log.Println(time.Now().String()+":Error opening file:", err)
		return
	}
	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(time.Now().String()+":Error reading file:", err)
		return
	}
	err = json.Unmarshal(jsonBytes, &branchAndEmailMap)
	if err != nil {
		log.Println(time.Now().String()+":Error decoding JSON:", err)
		return
	}

	//token := os.Getenv("TOKEN")
	// 获取今天的日期，格式为GitHub API使用的日期（例如：2023-06-12）
	today := time.Now().UTC().Format("2006-01-02")

	service.Checkout(branchAndEmailMap, today, repo)
}
