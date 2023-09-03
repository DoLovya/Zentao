package UnitTest

import (
	"ChanDaoScript/Api"
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	_ = Api.LoadToken("", "")
}

func TestGetBugSet(t *testing.T) {
	TestGetToken(t)

	var bugSet Api.BugSet
	err := Api.GetBugSet(61, Api.BugParam{
		Page:   1,
		Status: "all",
		Limit:  1,
	}, &bugSet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bugSet)
}

func TestGetBugAmount(t *testing.T) {
	TestGetToken(t)

	var bugAmount int
	if err := Api.GetBugAmount(61, &bugAmount); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bugAmount)
}

func TestGetAllBug(t *testing.T) {
	TestGetToken(t)

	var allBug []Api.Bug
	if err := Api.GetAllBug(61, &allBug); err != nil {
		fmt.Println(err)
		return
	}

	for _, bug := range allBug {
		fmt.Printf("%d\n%s\n%s\n", bug.Id, bug.Title, bug.Status)
	}
}
