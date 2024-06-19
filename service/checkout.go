package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"problem-solving/utils"
	"sync"
	"time"
)

var (
	resetContinuesDay bool
	taskNum           int
	wg                sync.WaitGroup
	branch            string
	email             string
)

func Checkout(taskPool []map[string]string, today string) error {
	repo, token, dbconf := utils.GetConfig()
	taskNum = len(taskPool)

	for i := 0; i < taskNum; i++ {
		wg.Add(1)
		go RunTask(&wg, taskPool[i], repo, token, dbconf, today)
	}
	wg.Wait()
	return nil
}

func RunTask(wg *sync.WaitGroup, task map[string]string, repo, token string, dbconf utils.MysqlConf, today string) {
	defer wg.Done()

	for key, value := range task {
		branch = key
		email = value
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits?sha=%s&since=%sT00:00:00Z", repo, branch, today)

	// 创建HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(time.Now().String()+":Error creating request:", err)
	}

	// 设置请求头，添加认证信息
	req.Header.Add("Authorization", "Bearer "+token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(time.Now().String()+":Error sending request:", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(time.Now().String()+":Error reading response body:", err)
		}

		// 解析响应体
		var commits []map[string]interface{}
		err = json.Unmarshal(body, &commits)
		if err != nil {
			log.Println(time.Now().String()+":Error decoding JSON:", err)
		}

		// 检查是否有提交
		if len(commits) == 0 {
			log.Printf(time.Now().String()+":No commits for branch %s today.\n", branch)
			//今日未提交，发送邮件
			utils.SendEmail(branch, email)
			resetContinuesDay = true
		} else {
			resetContinuesDay = false

		}
		utils.RecordStatus(branch, email, 1, resetContinuesDay, dbconf)
	} else {
		log.Printf(time.Now().String()+":Failed to fetch commits, status code: %d\n", resp.StatusCode)
	}
}
