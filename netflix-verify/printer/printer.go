package printer

import (
	"encoding/json"
	"fmt"
	"github.com/sjlleo/netflix-verify/verify"
	"io/ioutil"
	"net/http"
)

const (
	AUTHOR        = "@ivan"
	VERSION       = "v3.0"
	RED_PREFIX    = "\033[1;31m"
	GREEN_PREFIX  = "\033[1;32m"
	YELLOW_PREFIX = "\033[1;33m"
	BLUE_PREFIX   = "\033[1;34m"
	PURPLE_PREFIX = "\033[1;35m"
	CYAN_PREFIX   = "\033[1;36m"
	RESET_PREFIX  = "\033[0m"
	White_PREFIX  = "\033[1;37m"
)

func Print(fr verify.FinalResult) {
	fmt.Println()
	fmt.Println(White_PREFIX + "////////////////////////////////////////////////////////////")
	ip, region, country, timezone, err := getIP()
	if err != nil {
		fmt.Println(RED_PREFIX+"获取 IP 地址时出错:", err)
		return
	}
	printVersion()
	fmt.Println(White_PREFIX + "國家: " + country + ", 地區: " + region + ", 時區: " + timezone + ", IP: " + ip)
	fmt.Println(White_PREFIX + "////////////////////////////////////////////////////////////")
	printResult("4", fr.Res[1])
	fmt.Println()
	printResult("6", fr.Res[2])
	fmt.Println()
}

func printVersion() {
	fmt.Println("**NetFlix 解锁检测小工具 " + VERSION + " By " + CYAN_PREFIX + AUTHOR + RESET_PREFIX + "**")
}

// IPInfo 结构体用于存储解析后的 JSON 数据
type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func getIP() (string, string, string, string, error) {
	// 发送 HTTP GET 请求
	resp, err := http.Get("https://ipinfo.io/?format=json")
	if err != nil {
		return "", "", "", "", fmt.Errorf("发送请求时出错: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", fmt.Errorf("读取响应时出错: %v", err)
	}

	// 解析 JSON 数据
	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return "", "", "", "", fmt.Errorf("解析 JSON 数据时出错: %v", err)
	}

	// 返回解析后的 IP 地址
	return ipInfo.IP, ipInfo.Region, ipInfo.Country, ipInfo.Timezone, nil
}

func printResult(ipVersion string, vResponse verify.VerifyResponse) {
	fmt.Printf("[IPv%s]\n", ipVersion)
	switch code := vResponse.StatusCode; {
	case code < -1:

		fmt.Println(RED_PREFIX + "您的网络可能没有正常配置IPv" + ipVersion + "，或者没有IPv" + ipVersion + "网络接入" + RESET_PREFIX)
	case code == -1:
		fmt.Println(RED_PREFIX + "Netflix在您的出口IP所在的国家不提供服务" + RESET_PREFIX)
	case code == 0:
		fmt.Println(RED_PREFIX + "Netflix在您的出口IP所在的国家提供服务，但是您的IP疑似代理，无法正常使用服务" + RESET_PREFIX)
		fmt.Println(CYAN_PREFIX + "NF所识别的IP地域信息：" + vResponse.CountryName + RESET_PREFIX)
	case code == 1:
		fmt.Println(YELLOW_PREFIX + "您的出口IP可以使用Netflix，但仅可看Netflix自制剧" + RESET_PREFIX)
		fmt.Println(CYAN_PREFIX + "NF所识别的IP地域信息：" + vResponse.CountryName + RESET_PREFIX)
	case code == 2:
		fmt.Println(GREEN_PREFIX + "您的出口IP完整解锁Netflix，支持非自制剧的观看" + RESET_PREFIX)
		fmt.Println(CYAN_PREFIX + "NF所识别的IP地域信息：" + vResponse.CountryName + RESET_PREFIX)
	case code == 3:
		fmt.Println(YELLOW_PREFIX + "您的出口IP无法观看此电影" + RESET_PREFIX)
	case code == 4:
		fmt.Println(GREEN_PREFIX + "您的出口IP可以观看此电影" + RESET_PREFIX)
		fmt.Println(CYAN_PREFIX + "NF所识别的IP地域信息：" + vResponse.CountryName + RESET_PREFIX)
	default:
		fmt.Println(YELLOW_PREFIX + "解锁检测失败，请稍后重试" + RESET_PREFIX)
	}
}
