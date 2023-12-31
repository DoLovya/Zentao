package UnitTest

import (
	"ChanDaoScript/Api"
	"ChanDaoScript/TerminalUI"
	"fmt"
	"testing"
)

func TestBugResolvedBuild_Exec(t *testing.T) {
	if err := Api.LoadToken("", ""); err != nil {
		fmt.Println(err)
		return
	}
	bugUI := TerminalUI.BugResolvedBuild{}
	var x string
	if err := bugUI.Exec(61, &x); err != nil {
		fmt.Println(err)
		x = ""
		return
	}
}
