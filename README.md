# 介绍

这是一个自用的简单工具，用来自动获取订阅的 SS 配置，并以连接性最好的节点配置重启 ss-local，暂时不支持 PAC 和 HTTP 代理等高级功能。

## 使用指南

### 1. 程序安装

1. 编译程序，将程序放到 `/usr/local/ss-helper` 下。
2. 将 `restart-ss-local.sh` 放到 `/usr/local/ss-helper` 下。程序需要通过这个 shell 文件重启 ss-local，注意检查脚本是否有执行权限
3. 在 `/usr/local/bin` 下创建本程序的软连接，以软连接的名字作为命令入口
4. 在 `/usr/local/ss-helper` 下创建 `ss-local-config.json` 保证当前用户有编辑权限

### 2. 配置文件：

使用前需要在用户目录下创建配置文件 `~/.switch-ssr.conf` 内容如下:

```json
{
  "subscribe_url": "xxx",
  "ss_local_exec": "/usr/local/bin/ss-local",
  "enable_auto_switch": false,
  "app_home": "/usr/local/ss-helper"
}
```

`subscribe_url`代表 SS 订阅地址，程序会自动根据订阅地址拉取订阅配置，然后测试节点连通状态，以连通性最好的节点来重启 ss-local。

`ss_local_exec`代表 ss-local 的可执行文件路径

`app_home`代表本程序的安装目录
