package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SummarizeURL(url string) (string, error) {
	// 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("获取网页失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("网页返回状态码异常: %d", resp.StatusCode)
	}

	// 解析 HTML
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("网页内容读取失败")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(bodyBytes)))
	if err != nil {
		return "", errors.New("HTML 解析失败")
	}

	// 提取 <p> 标签文字
	var content strings.Builder
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if len(text) > 20 {
			content.WriteString(text + "\n")
		}
	})

	rawText := content.String()
	if len(rawText) < 100 {
		return "", errors.New("网页内容太少，无法生成摘要")
	}

	// 调用已有模型摘要逻辑
	return CallDeepSeekAPI(rawText)
}
