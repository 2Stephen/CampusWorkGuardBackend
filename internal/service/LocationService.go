package service

import (
	"CampusWorkGuardBackend/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type DistrictResponse struct {
	Status    string     `json:"status"`
	Info      string     `json:"info"`
	Infocode  string     `json:"infocode"`
	Districts []District `json:"districts"`
}

type District struct {
	Adcode    string     `json:"adcode"`
	Name      string     `json:"name"`
	Center    string     `json:"center"`
	Level     string     `json:"level"`
	Districts []District `json:"districts"`
}

func GetLocationList(keywords string) (*model.Location, error) {
	var (
		key         = os.Getenv("AMAP_KEY")
		subdistrict = "1"
		baseURL     = "https://restapi.amap.com/v3/config/district?"
	)

	if key == "" {
		return nil, errors.New("请先设置环境变量 AMAP_KEY")
	}
	// 默认查询全国
	if keywords == "" {
		keywords = "中华人民共和国"
	} else {
		last := []rune(keywords)[len([]rune(keywords))-1]
		if last != '省' && last != '市' && last != '区' {
			return nil, errors.New("关键词必须以省、市、区结尾")
		}
	}

	// 拼接 URL
	url := fmt.Sprintf("%skey=%s&subdistrict=%s&keywords=%s", baseURL, key, subdistrict, keywords)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭响应体错误：", err)
		}
	}(resp.Body)

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析 JSON
	var data DistrictResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// 处理返回结果
	if data.Status != "1" {
		return nil, fmt.Errorf("请求失败: %s", data.Info)
	}
	var res model.Location
	if len(data.Districts) == 0 || len(data.Districts[0].Districts) == 0 {
		return nil, errors.New("未找到对应的行政区划信息")
	}
	res.Name = data.Districts[0].Name
	res.Adcode = data.Districts[0].Adcode
	res.Level = data.Districts[0].Level
	for _, district := range data.Districts[0].Districts {
		var dist model.District
		dist.Name = district.Name
		dist.Adcode = district.Adcode
		dist.Level = district.Level
		res.Districts = append(res.Districts, dist)
	}

	return &res, nil
}
