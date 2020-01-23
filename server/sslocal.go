package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"ss-helper/global"
)

// SSLocalConfig 本地ss-local配置
// macOS ~/Library/Application Support/ShadowsocksX-NG
type SSLocalConfig struct {
	Method        string `json:"method"`
	Timeout       int    `json:"timeout"`
	Protocol      string `json:"protocol"`
	LocalAddress  string `json:"local_address"`
	ProtocolParam string `json:"protocol_param"`
	ObfsParam     string `json:"obfs_param"`
	LocalPort     int    `json:"local_port"`
	ServerPort    int    `json:"server_port"`
	Password      string `json:"password"`
	Obfs          string `json:"obfs"`
	Server        string `json:"server"`
}

// StoreConf 将server配置写入应用配置文件中
func (sc *SSLocalConfig) StoreConf(serv Server) {
	// 初始化配置
	sc.Timeout = 60
	sc.LocalAddress = "127.0.0.1"
	sc.LocalPort = 1080
	sc.Method = serv.Method
	sc.Protocol = serv.Protocol
	sc.ProtocolParam = serv.ProtocolParam
	sc.ObfsParam = serv.ObfsParam
	sc.ServerPort = serv.ServerPort
	sc.Password = serv.Password
	sc.Obfs = serv.Obfs
	sc.Server = serv.Server

	confPath := global.AppConf.AppHome + global.SSConfigFile
	jsonBytes, err := json.Marshal(sc)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(confPath, jsonBytes, 0666)
	if err != nil {
		panic(err)
	}
}

// Reconnect 使用新配置启动ss-local
func (sc *SSLocalConfig) Reconnect() {
	// 调switch脚本重启ss-local
	newConf := global.AppConf.AppHome + global.SSConfigFile
	switchyExec := global.AppConf.AppHome + global.RestartSSExec
	doSwitch := exec.Command(switchyExec, global.AppConf.SSLBinary, newConf)
	doSwitch.Stdout = os.Stdout
	err := doSwitch.Run()
	if err != nil {
		panic(err)
	}
}
