package main

import (
	"ChanDaoScript/Api"
	"ChanDaoScript/Loader"
	"ChanDaoScript/TerminalUI"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	// 加载配置项
	if err := Loader.LoadConfig(); err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		if firstArg == "config" {
			if err := Loader.ParseFlag(); err != nil {
				fmt.Println(err)
				return
			}
			if Loader.GetFlag().List {
				fmt.Println(Loader.GetConfig())
				return
			}
			Loader.SaveFlag()

		} else if firstArg == "run" {
			if err := Api.LoadToken(Loader.GetConfig().Account, Loader.GetConfig().Password); err != nil {
				fmt.Println(err)
				return
			}
			productUI := TerminalUI.NewProducts()
			var productId int
			if err := productUI.Exec(&productId); err != nil {
				fmt.Println(err)
				return
			}
			bugResolvedBuild := TerminalUI.NewBugResolvedBuild()
			var resolvedBuild string
			if err := bugResolvedBuild.Exec(productId, &resolvedBuild); err != nil {
				fmt.Println(err)
				return
			}
			var bugs []Api.Bug
			if err := Api.GetBugs(productId, resolvedBuild, &bugs); err != nil {
				fmt.Println(err)
				return
			}
			formattedStr := ""
			for i, bug := range bugs {
				str := fmt.Sprintf("[%02d] %d %s\n", i+1, bug.Id, bug.Title)
				fmt.Printf(str)
				formattedStr += str
			}
			if err := clipboard.WriteAll(formattedStr); err != nil {
				fmt.Println("无法将文本复制到剪贴板 -> ", err)
				return
			}
		}
	}
}
