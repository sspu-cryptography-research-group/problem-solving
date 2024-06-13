package main

import (
	"daily/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	branchAndEmailMap map[string]string
	repo              string = "sspu-cryptography-research-group/problem-solving"
	//repo        = "OrangeDou/QHTLC"
	githubToken = "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"
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

	//人多了就改成并发处理
	for branch, email := range branchAndEmailMap {
		url := fmt.Sprintf("https://api.github.com/repos/%s/commits?sha=%s&since=%sT00:00:00Z", repo, branch, today)

		// 创建HTTP请求
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(time.Now().String()+":Error creating request:", err)
			return
		}

		// 设置请求头，添加认证信息
		req.Header.Add("Authorization", "Bearer "+githubToken)

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(time.Now().String()+":Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		// 检查响应状态码
		if resp.StatusCode == http.StatusOK {
			// 读取响应体
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(time.Now().String()+":Error reading response body:", err)
				return
			}

			// 解析响应体
			var commits []map[string]interface{}
			err = json.Unmarshal(body, &commits)
			if err != nil {
				log.Println(time.Now().String()+":Error decoding JSON:", err)
				return
			}

			// 检查是否有提交
			if len(commits) == 0 {
				log.Printf(time.Now().String()+":No commits for branch %s today.\n", branch)
				utils.SendEmail(branch, email)
				//utils.SendEmail2(email)
			} else {
				log.Printf(time.Now().String()+":Commits found for branch %s today: %+v\n", branch, commits)
			}
		} else {
			log.Printf(time.Now().String()+":Failed to fetch commits, status code: %d\n", resp.StatusCode)
		}
	}
}