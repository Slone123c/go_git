package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init [git repository]",
	Short: "Init a git repository.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Adding files:", args)
		// 调用实现添加文件到暂存区的函数
	},
}
