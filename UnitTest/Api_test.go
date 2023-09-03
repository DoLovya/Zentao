package UnitTest

import (
	"ChanDaoScript/Api"
	"fmt"
	"testing"
)

func TestGetAllResolvedBuild(t *testing.T) {
	if err := Api.LoadToken("", ""); err != nil {
		fmt.Println(err)
		return
	}
	var allResolvedBuild []string
	if err := Api.GetAllResolvedBuild(61, &allResolvedBuild); err != nil {
		fmt.Println(err)
		return
	}
	for _, resolvedBuild := range allResolvedBuild {
		fmt.Println(resolvedBuild)
	}
}
