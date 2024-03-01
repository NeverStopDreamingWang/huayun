package main

import (
	"fmt"
	"github.com/NeverStopDreamingWang/huayun"
	"github.com/spf13/cobra"
	"os"
)

// 根命令
var HuaYunCmd = &cobra.Command{
	Use:   "huayun",
	Short: `华运面板，一款更适合中华的运维面板`,
	RunE: func(cmd *cobra.Command, args []string) error {
		huayun.Start()
		return nil
	},
}

var Help = `华运面板 version（版本）: %s
使用"huayun help <command>"获取命令的更多信息。

Usage（用法）: 
  huayun <command> [arguments]
The commands are（命令如下）: 
  version		查看 华运面板 版本信息
  info			查看 华运面板 服务信息
  port			修改 华运面板 服务端口
  domain		修改 华运面板 服务域名
  entrance		修改 华运面板 服务安全入口
  
  user
  	username	修改 华运面板 服务管理员用户
  	password	修改 华运面板 服务管理员密码
  
  listen-ip
    ipv4		修改 华运面板 服务监听 IPv4
  	ipv6		修改 华运面板 服务监听 IPv6
  
  ssl
  	enable		开启 华运面板 服务 SSL
  	disable		关闭 华运面板 服务 SSL

`

// 根命令
var HelpCmd = &cobra.Command{
	Use:   "help",
	Short: `help 帮助`,
	RunE: func(cmd *cobra.Command, args []string) error {
		help_txt := fmt.Sprintf(Help, "103")
		fmt.Print(help_txt)
		return nil
	},
}

func main() {
	HuaYunCmd.AddCommand(HelpCmd) // 帮助信息
	if err := HuaYunCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
