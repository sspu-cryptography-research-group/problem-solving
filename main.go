package main

import (
	"problem-solving/service"
	"problem-solving/utils"
	"time"
)

func main() {

	// 获取今天的日期，格式为GitHub API使用的日期（例如：2023-06-12）
	today := time.Now().UTC().Format("2006-01-02")
	taskPool := utils.GenerateTaskPool("participant.json")
	service.Checkout(taskPool, today)
}
