package Loader

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

type Flag struct {
	List bool
}

var fg Flag

func GetFlag() *Flag {
	return &fg
}

func ParseFlag() error {
	flagSet := flag.NewFlagSet("Zentao", flag.ContinueOnError)
	flagSet.StringVarP(&GetConfig().Account, "account", "a", GetConfig().Account, "禅道账号")
	flagSet.StringVarP(&GetConfig().Password, "password", "p", GetConfig().Password, "禅道密码")
	flagSet.BoolVarP(&GetFlag().List, "list", "l", false, "查看配置列表")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("flagSet.Parse -> %v", err)
	}
	return nil
}

func SaveFlag() {
	SaveConfig()
}
