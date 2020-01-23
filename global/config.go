package global

// AppConfig app配置结构体
type AppConfig struct {
	SubscribeURL string `json:"subscribe_url"` // 订阅地址
	SSLBinary    string `json:"ss_local_exec"` // ss-local可执行文件地址
	AppHome      string `json:"app_home"`      // 本程序安装目录
}

// SSConfigFile 程序生成的ss-local配置文件名
var SSConfigFile = "/ss-local-config.json"

// RestartSSExec 切换ss-local连接节点的bash文件名
var RestartSSExec = "/restart-ss-local.sh"

// AppConfigPath app 配置文件路径  ~/.switch-ssr.conf
var AppConfigPath = "/.switch-ssr.conf"

// AppConf 配置文件内容全局对象
var AppConf = &AppConfig{}

// Store 保存配置
func (ac *AppConfig) Store(c AppConfig) {
	ac = &c
}

// Retrive 获取配置
func (ac *AppConfig) Retrive() *AppConfig {
	return ac
}
