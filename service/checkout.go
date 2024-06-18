package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"problem-solving/utils"
	"time"
)

var (
	Token             = "this is a secret token,if you want to run main code,please generate your own token and replace it!"
	resetContinuesDay bool
)

func Checkout(branchAndEmailMap map[string]string, today, repo string) error {
	for branch, email := range branchAndEmailMap {
		url := fmt.Sprintf("https://api.github.com/repos/%s/commits?sha=%s&since=%sT00:00:00Z", repo, branch, today)

		// 创建HTTP请求
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(time.Now().String()+":Error creating request:", err)
			return err
		}

		// 设置请求头，添加认证信息
		req.Header.Add("Authorization", "Bearer "+Token)

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(time.Now().String()+":Error sending request:", err)
			return err
		}
		defer resp.Body.Close()

		// 检查响应状态码
		if resp.StatusCode == http.StatusOK {
			// 读取响应体
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(time.Now().String()+":Error reading response body:", err)
				return err
			}

			// 解析响应体
			var commits []map[string]interface{}
			err = json.Unmarshal(body, &commits)
			if err != nil {
				log.Println(time.Now().String()+":Error decoding JSON:", err)
				return err
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
			utils.RecordStatus(branch, email, 1, resetContinuesDay)
		} else {
			log.Printf(time.Now().String()+":Failed to fetch commits, status code: %d\n", resp.StatusCode)
		}
	}
	return nil
}
