package main

import (
	"github.com/NeverStopDreamingWang/goi"
	"github.com/NeverStopDreamingWang/huayun/huayun"
	"github.com/spf13/cobra"
	"path"
)

func init() {
	HuaYunCmd.AddCommand(InfoCmd)     // 查看面板信息
	HuaYunCmd.AddCommand(PortCmd)     // 修改面板端口
	HuaYunCmd.AddCommand(DomainCmd)   // 修改访问绑定域名
	HuaYunCmd.AddCommand(EntranceCmd) // 修改面板安全入口
	HuaYunCmd.AddCommand(ListenIpCmd) // 切换面板监听 IP
	{
		ListenIpCmd.AddCommand(Ipv4Cmd) // 切换面板监听 IPv4
		ListenIpCmd.AddCommand(Ipv6Cmd) // 切换面板监听 IPv6
	}
	HuaYunCmd.AddCommand(SSLCmd) // 切换面板 SSl
	{
		SSLCmd.AddCommand(EnableSSLCmd)  // 开启面板 SSL
		SSLCmd.AddCommand(DisableSSLCmd) // 关闭面板 SSL
	}
	HuaYunCmd.AddCommand(VersionCmd) // 查看面板版本信息
}

// 查看面板信息
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看面板版本信息",
	RunE:  Info,
}

func Info(cmd *cobra.Command, args []string) error {

	return nil
}

// 修改面板端口
var PortCmd = &cobra.Command{
	Use:   "port",
	Short: "修改面板端口",
	RunE:  Port,
}

func Port(cmd *cobra.Command, args []string) error {

	return nil
}

// 修改访问绑定域名
var DomainCmd = &cobra.Command{
	Use:   "domain",
	Short: "修改访问绑定域名",
	RunE:  Domain,
}

func Domain(cmd *cobra.Command, args []string) error {

	return nil
}

// 修改面板安全入口
var EntranceCmd = &cobra.Command{
	Use:   "entrance",
	Short: "修改面板安全入口",
	RunE:  Entrance,
}

func Entrance(cmd *cobra.Command, args []string) error {

	return nil
}

// 切换面板监听 IP
var ListenIpCmd = &cobra.Command{
	Use:   "listen-ip",
	Short: "查看面板版本信息",
}

// 切换面板监听 IPv4
var Ipv4Cmd = &cobra.Command{
	Use:   "ipv4",
	Short: "切换面板监听 IPv4",
	RunE:  Ipv4,
}

func Ipv4(cmd *cobra.Command, args []string) error {

	return nil
}

// 切换面板监听 IPv6
var Ipv6Cmd = &cobra.Command{
	Use:   "ipv6",
	Short: "切换面板监听 IPv6",
	RunE:  Ipv6,
}

func Ipv6(cmd *cobra.Command, args []string) error {

	return nil
}

// 切换面板 SSl
var SSLCmd = &cobra.Command{
	Use:   "ssl",
	Short: "切换面板 SSl",
}

// 开启面板 SSL
var EnableSSLCmd = &cobra.Command{
	Use:   "enable",
	Short: "开启面板 SSL",
	RunE:  EnableSSL,
}

func EnableSSL(cmd *cobra.Command, args []string) error {
	huayun.Server.Settings.SSL = goi.MetaSSL{
		STATUS:    true, // SSL 开关
		CERT_PATH: path.Join(huayun.Server.Settings.BASE_DIR, "ssl/huayun.crt"),
		KEY_PATH:  path.Join(huayun.Server.Settings.BASE_DIR, "ssl/huayun.key"),
		CSR_PATH:  path.Join(huayun.Server.Settings.BASE_DIR, "ssl/huayun.csr"), // 可选
	}
	return nil
}

// 关闭面板 SSL
var DisableSSLCmd = &cobra.Command{
	Use:   "disable",
	Short: "关闭面板 SSL",
	RunE:  DisableSSL,
}

func DisableSSL(cmd *cobra.Command, args []string) error {
	huayun.Server.Settings.SSL = goi.MetaSSL{
		STATUS:    false, // SSL 开关
		CERT_PATH: path.Join(huayun.Server.Settings.BASE_DIR, "ssl/huayun.crt"),
		KEY_PATH:  path.Join(huayun.Server.Settings.BASE_DIR, "ssl/huayun.key"),
		CSR_PATH:  path.Join(huayun.Server.Settings.BASE_DIR, "ssl/huayun.csr"), // 可选
	}
	return nil
}

// 查看面板版本信息
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看面板版本信息",
	RunE:  Version,
}

func Version(cmd *cobra.Command, args []string) error {

	return nil
}
