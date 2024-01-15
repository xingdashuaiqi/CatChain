package main

import (
	"fmt"
	"regexp"
	"gopkg.in/gomail.v2"
	"time"
	"math/rand"
	"github.com/gocolly/colly"
)
var isSetBuyMarketR bool // 用于存储状态
var isSendmessage bool // 是否发送信息到我的邮箱
var leixin string
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15",
	// 添加更多User-Agent...
}
// 随机选择一个User-Agent
func getRandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(userAgents))
	return userAgents[index]
}
func sendEmail(txHash string,leixin string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "2192303400@qq.com")
	m.SetHeader("To", "2192303400@qq.com")
	m.SetHeader("Subject", fmt.Sprintf("New transaction: %s",leixin))
	m.SetBody("text/html", fmt.Sprintf("正在有人进行买入: %s", txHash))

	d := gomail.NewDialer("smtp.qq.com", 587, "2192303400@qq.com", "oflpbdwpgodqeaja")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("邮件发送失败:", err)
	} else {
		fmt.Println("邮件发送成功")
	}
}
func StartM() {

c := colly.NewCollector()

	// 设置 User-Agent 和请求速率限制
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", getRandomUserAgent())
	})

	// 定义处理 HTML 元素的回调函数
	c.OnHTML("tbody.align-middle.text-nowrap tr:nth-child(20) td:nth-child(3) span.d-block.badge", func(e *colly.HTMLElement) {
		name := e.Text
		fmt.Println("Transaction name is:", name)
		isSetBuyMarketR = name == "Swap Exact ETH F..."
		isSendmessage = name == "Swap Exact ETH F..."
		leixin = name
	})

	c.OnHTML("tbody.align-middle.text-nowrap tr:nth-child(20) td:nth-child(2) a.hash-tag", func(e *colly.HTMLElement) {
		if e.Index == 0 && isSetBuyMarketR {
			// 提取链接中的文本信息
			txHash := e.Text
			fmt.Println("Transaction Hash is:", txHash)
			StartMtwo(txHash)       
			sendEmail(txHash,leixin)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// 这里写需要监控页的地址
	if err:= c.Visit("https://bscscan.com/address/0x10ed43c718714eb63d5aa57b78b54704e256024e");
	err != nil {
	fmt.Println("Error visiting the website:", err)
	}
}
func StartMtwo(txHash string) {
	c := colly.NewCollector()
	fmt.Println("txHash", txHash)
	// 设置 User-Agent 和请求速率限制
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	})

	// 定义处理 HTML 元素的回调函数
	c.OnHTML("span.d-inline-flex.flex-wrap.align-items-center a", func(e *colly.HTMLElement) {
		// 提取 a 标签的 href 属性值
		href := e.Attr("href")
		fmt.Printf("Href value: %s\n", href)
	})
	// 定义处理 HTML 元素的回调函数
	c.OnHTML("span[class='me-1']:not([id])", func(e *colly.HTMLElement) {
		// 提取 a 标签的 href 属性值
		if e.DOM.Find("a").Length() == 0 {
			Value := e.Text
			match, _ := regexp.MatchString("^[^a-zA-Z]+$", Value)
			if match{
				fmt.Printf("value: %s\n", Value)
			}
	// 		// 去除逗号
	//     valueWithoutComma := strings.ReplaceAll(value, ",", "")
	//     // 将字符串转换为浮点数
	//    resultFloat, err := strconv.ParseFloat(valueWithoutComma, 64)
	//     if err != nil {
	// 	fmt.Println("转换失败:", err)
	// 	return
	//     }
	// 		if resultFloat > 10000.0{

	// 		}
			
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	txHashURL := fmt.Sprintf("https://bscscan.com/tx/%s", txHash)
	// 这里写需要监控页的地址
	c.Visit(txHashURL)
}
func main(){
	// 利用 time.Tick 来实现定时调用
	ticker := time.Tick(6 * time.Second)
    nonoce := 1 
	for range ticker {
		nonoce ++
		fmt.Println("Visiting:", nonoce)
		StartM()
	}
}
