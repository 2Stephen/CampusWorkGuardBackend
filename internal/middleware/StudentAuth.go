package middleware

import (
	"CampusWorkGuardBackend/internal/model"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
)

func CHSIAuth(vcode string) (*model.CHSIStudentInfo, error) {
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
		return nil, err
	}

	// 解析html
	var cHSIStudentInfo *model.CHSIStudentInfo
	cHSIStudentInfo, err = parseStudentInfo(html)
	if err != nil {
		log.Println("解析学生信息错误：", err)
		return nil, err
	}
	return cHSIStudentInfo, nil
}

func parseStudentInfo(html string) (*model.CHSIStudentInfo, error) {
	var cHSIStudentInfo model.CHSIStudentInfo
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	// 判断学信网验证码是否已过期
	if strings.Contains(doc.Text(), "在线验证报告已过期") {
		return nil, fmt.Errorf("学信网验证码已过期")
	}
	// 判断学信网验证码是否不存在
	if strings.Contains(doc.Text(), "此在线验证码无效") {
		return nil, fmt.Errorf("学信网验证码无效")
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

	return &cHSIStudentInfo, nil
}
