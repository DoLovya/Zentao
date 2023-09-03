package Loader

import (
	"fmt"
	"github.com/gdamore/tcell/encoding"
	"os/exec"
	"syscall"
)

func loadChinese() {
	encoding.Register()
}

func loadUtf8() error {
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CmdLine: "cmd /c chcp 65001",
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Failed to start executable: %v", err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Executable finished with error: %v", err)
	}
	return nil
}

func LoadConfig() error {
	loadChinese()
	if err := loadUtf8(); err != nil {
		return fmt.Errorf("loadUtf8() -> %v", err)
	}
	return nil
}
