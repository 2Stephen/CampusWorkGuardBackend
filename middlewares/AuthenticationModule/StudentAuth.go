package middlewares

import (
	"CampusWorkGuardBackend/models/request"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
)

type CHSIStudentInfo struct {
	Name         string
	Gender       string
	Birthday     string
	Nation       string
	School       string
	Level        string
	Major        string
	Duration     string
	DegreeType   string
	StudyMode    string
	College      string
	Department   string
	EntranceDate string
	Status       string
	ExpectedGrad string
	Vcode        string
	StudentID    string
	Email        string
}

func (cHSIStudentInfo *CHSIStudentInfo) StudentAuth(params request.StudentAuthParams) bool {
	// http服务调用chrome，chrome打开学信网接口进行认证
	log.Println("Starting student authentication for ID:", params.ID)
	cHSIStudentInfo.CHSIAuth(params.Vcode)
	// 对比信息
	// 输出获取到的学生信息
	log.Printf("Retrieved Student Info: %+v\n", cHSIStudentInfo)
	if cHSIStudentInfo.School == params.School && cHSIStudentInfo.Vcode == params.Vcode {
		log.Println("Student authentication successful for StudentID:", params.ID)
		cHSIStudentInfo.StudentID = params.ID
		cHSIStudentInfo.Email = params.Email
		// 调用DTO存储认证信息

		return true
	} else {
		log.Println("Student authentication failed for StudentID:", params.ID)
		return false
	}
}

func (cHSIStudentInfo *CHSIStudentInfo) CHSIAuth(vcode string) {
	// 1. 获取真实浏览器的 websocket 地址
	wsURL := os.Getenv("CHROME_WS_URL")

	// 2. 使用 RemoteAllocator 连接到真实 Chrome
	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), wsURL)
	defer cancel()

	// 3. 创建 chromedp 上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var html string
	url := fmt.Sprintf("https://www.chsi.com.cn/xlcx/bg.do?vcode=%s&srcid=bgcx", vcode)

	// 4. 执行浏览器操作（真实 Chrome）
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &html),
	)

	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	// 解析html
	err = cHSIStudentInfo.ParseStudentInfo(html)
	if err != nil {
		fmt.Println("解析错误：", err)
		return
	}
}

func (cHSIStudentInfo *CHSIStudentInfo) ParseStudentInfo(html string) error {
	log.Println(html)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return err
	}

	// 遍历每个 .report-info-item
	doc.Find(".report-info-item").Each(func(i int, s *goquery.Selection) {
		label := strings.TrimSpace(s.Find(".label").Text())
		value := strings.TrimSpace(s.Find(".value").Text())

		switch label {
		case "姓名":
			cHSIStudentInfo.Name = value
		case "性别":
			cHSIStudentInfo.Gender = value
		case "出生日期":
			cHSIStudentInfo.Birthday = value
		case "民族":
			cHSIStudentInfo.Nation = value
		case "学校名称":
			cHSIStudentInfo.School = value
		case "层次":
			cHSIStudentInfo.Level = value
		case "专业":
			cHSIStudentInfo.Major = value
		case "学制":
			cHSIStudentInfo.Duration = value
		case "学历类别":
			cHSIStudentInfo.DegreeType = value
		case "学习形式":
			cHSIStudentInfo.StudyMode = value
		case "分院":
			cHSIStudentInfo.College = value
		case "系所":
			cHSIStudentInfo.Department = value
		case "入学日期":
			cHSIStudentInfo.EntranceDate = value
		case "学籍状态":
			cHSIStudentInfo.Status = value
		case "预计毕业日期":
			cHSIStudentInfo.ExpectedGrad = value
		}
	})

	// 获取在线验证码
	cHSIStudentInfo.Vcode = strings.TrimSpace(doc.Find(".yzm").Text())

	return nil
}
