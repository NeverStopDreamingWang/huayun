package main

import (
	"github.com/spf13/cobra"
)

func init() {
	HuaYunCmd.AddCommand(UserCmd) // 修改面板用户信息
	{
		UserCmd.AddCommand(UsernameCmd) // 修改用户名
		UserCmd.AddCommand(PasswordCmd) // 修改用户密码
	}
}

// 修改面板用户信息
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "修改面板用户信息",
}

// 修改用户名
var UsernameCmd = &cobra.Command{
	Use:   "username",
	Short: "查看面板版本信息",
	RunE:  Username,
}

func Username(cmd *cobra.Command, args []string) error {

	return nil
}

// 修改用户密码
var PasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "查看面板版本信息",
	RunE:  Password,
}

func Password(cmd *cobra.Command, args []string) error {

	return nil
}
