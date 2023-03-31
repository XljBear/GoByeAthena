package main

import (
	"GoByeAthena/structs"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("请输入您的SessData：")
	sessData := ""
	fmt.Scanf("%s\n", &sessData)
	master := GetMasterInfo(sessData)
	if master == nil || master.Code != 0 {
		fmt.Println("登录失败，请确认SessData是否正确。")
		time.Sleep(time.Second * 5)
		return
	}
	if master.Data.Strength < 5 {
		fmt.Printf("%s，您的剩余体力已不足以外出探索了，感谢您这段时间的陪伴，期待新的活动与您再相见。Athena\n", master.Data.CallName)
		time.Sleep(time.Second * 5)
		return
	}
	fmt.Printf("%s，Athena即将在4月初下线，您的剩余的[%d]体力将会消失，是否现在需要我为您进行自动探索来将体力转化为金币呢？(y/n)：", master.Data.CallName, master.Data.Strength)
	choice := "n"
	fmt.Scanf("%s\n", &choice)
	if choice == "y" || choice == "Y" {
		fmt.Println("自动探索将在3秒后开始...")
		time.Sleep(3 * time.Second)
		adventureTime := 1
		adventureTimeCount := 0
		startCoin := master.Data.Amount
		for {
			oldCoin := master.Data.Amount
			oldStrength := master.Data.Strength
			adventureCount := 0
			if master.Data.Strength >= 30 {
				adventureCount = 5
			} else if master.Data.Strength < 6 {
				break
			} else {
				adventureCount = int(math.Floor(float64(master.Data.Strength) / 6.0))
			}
			fmt.Println("---------------------------------------------")
			fmt.Printf("[探索%d] 进行一次%d连探索，消耗%d点体力...", adventureTime, adventureCount, adventureCount*6)
			adventureinfo := AutoAdventure(sessData, adventureCount)
			if adventureinfo == nil {
				fmt.Printf("失败！请稍后再试\n")
				time.Sleep(time.Second * 5)
				return
			} else if adventureinfo.Code != 0 {
				fmt.Printf("失败！失败原因[%s]\n", adventureinfo.Message)
				time.Sleep(time.Second * 5)
				return
			}
			master = GetMasterInfo(sessData)
			gotCoin := master.Data.Amount - oldCoin
			gotStrength := master.Data.Strength - (oldStrength - adventureCount*6)
			fmt.Printf("成功！获得[%d枚金币]，[%d点体力]\n", gotCoin, gotStrength)
			fmt.Printf("目前剩余体力：%d\n", master.Data.Strength)
			adventureTime++
			adventureTimeCount += adventureCount
			time.Sleep(time.Millisecond * 100)
		}
		fmt.Printf("自动探索结束！一共进行了%d次探索，共收获金币%d枚。感谢您这段时间的陪伴，期待新的活动与您再相见。Athena\n", adventureTimeCount, master.Data.Amount-startCoin)
		time.Sleep(time.Second * 5)
		return
	} else {
		fmt.Printf("好的，那么需要我时再启动我吧，再见%s。Athena\n", master.Data.CallName)
		time.Sleep(time.Second * 5)
		return
	}
}
func AutoAdventure(sessData string, count int) *structs.AdventureInfo {
	url := "https://show.bilibili.com/api/activity/athena/adventure/start"
	method := "POST"

	payload := strings.NewReader(`{"count":` + strconv.Itoa(count) + `,"placeId":"","propIds":[]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil
	}
	req.Header.Add("Host", "show.bilibili.com")
	req.Header.Add("Cookie", "SESSDATA="+sessData+";")
	req.Header.Add("User-Agent", "bili-universal/7.23.0 (iPhone; iOS 16.4; Scale/3.00) BiliSmallApp/3.99.0")
	req.Header.Add("Mobi-App", "iphone")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Language", "zh-Hans-CN;q=1")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Device", "phone")

	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	adventureInfo := &structs.AdventureInfo{}
	err = json.Unmarshal(body, adventureInfo)
	if err != nil {
		return nil
	}
	return adventureInfo
}
func GetMasterInfo(sessData string) *structs.MasterInfo {
	url := "https://show.bilibili.com/api/activity/athena/home?isInit=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil
	}
	req.Header.Add("Host", "show.bilibili.com")
	req.Header.Add("Cookie", "SESSDATA="+sessData+";")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "bili-universal/7.23.0 (iPhone; iOS 16.4; Scale/3.00) BiliSmallApp/3.99.0")

	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	masterInfo := &structs.MasterInfo{}
	err = json.Unmarshal(body, masterInfo)
	if err != nil {
		return nil
	}
	return masterInfo
}
