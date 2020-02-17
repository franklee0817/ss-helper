package main

import (
	"encoding/json"
	"fmt"
	"os"
	"ss-helper/global"
	"ss-helper/pathloader"
	"ss-helper/server"
	"strings"
)

func main() {
	// 获取用户Home目录下的配置文件内容
	userHome, _ := pathloader.Home()
	fileName := userHome + global.AppConfigPath
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(global.AppConf)
	if err != nil {
		str := fmt.Sprintf("找不到默认配置文件: %s%s", userHome, global.AppConfigPath)
		fmt.Println(str)
		return
	}

	var cl server.ConnectivityList
	if len(global.AppConf.SubscribeURL) > 0 {
		cl = fromSubscribe(global.AppConf.SubscribeURL)
	} else {
		str := fmt.Sprintf("配置文件为空： %s%s", userHome, global.AppConfigPath)
		fmt.Println(str)
		return
	}

	// 尝试切换服务器
	// 订阅列表中尝尝会有中国大陆的”不可用“节点，先过滤掉包含中国字眼的节点
	if global.AppConf.EnableSwitch == true {
		notCnServs, _ := filterCnNode(cl)
		if len(notCnServs) > 0 {
			serv := notCnServs[0]
			fmt.Println(fmt.Sprintf("设置当前VPN节点为：[%s] %s", serv.Group, serv.Remarks))
			slc := &server.SSLocalConfig{}
			slc.StoreConf(serv)
			slc.Reconnect()
		}
	}

}

// filterCnNode 过滤中国大陆节点
func filterCnNode(cl server.ConnectivityList) ([]server.Server, server.ConnectivityList) {
	if len(cl) <= 0 {
		return nil, nil
	}
	var notCnServs []server.Server
	var filteredCl server.ConnectivityList
	for _, connRes := range cl {
		cnIdx := strings.Index(connRes.Name, "中国")
		if cnIdx < 0 {
			filteredCl = append(filteredCl, connRes)
			notCnServs = append(notCnServs, connRes.ServConf)
		}
	}

	return notCnServs, filteredCl
}

// fromSubscribe 从订阅地址获取服务器列表并检测连通性
func fromSubscribe(url string) server.ConnectivityList {
	servers := server.PullSubscribe(url)
	if len(servers) <= 0 {
		fmt.Println(servers)
		panic("获取订阅节点失败")
	}
	cl := server.CheckServers(servers)
	printCl(cl)

	return cl
}

func printCl(cl server.ConnectivityList) {
	_, filteredCl := filterCnNode(cl)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("===========================================================================")
	fmt.Println("延迟最低的前20的节点为.")
	i := 0
	for _, v := range filteredCl {
		i++
		str := fmt.Sprintf("%d. %s 延迟: %d ms", i, v.Name, v.Delay)
		fmt.Println(str)
		if i >= 20 {
			break
		}
	}
}
