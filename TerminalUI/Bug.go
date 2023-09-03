package TerminalUI

import (
	"ChanDaoScript/Api"
)

type BugResolvedBuild struct {
	list *List
}

func NewBugResolvedBuild() *BugResolvedBuild {
	bugResolvedBuild := &BugResolvedBuild{}
	bugResolvedBuild.list = NewList()
	return bugResolvedBuild
}

func (this *BugResolvedBuild) Exec(productId int, resolvedBuild *string) error {
	this.list.SetTitle("选择Bug的解决版本")
	var allResolvedBuild []string
	if err := Api.GetAllResolvedBuild(productId, &allResolvedBuild); err != nil {
		return err
	}
	this.list.SetOptions(allResolvedBuild)

	var selectedIndex int
	if err := this.list.Exec(&selectedIndex); err != nil {
		return err
	}
	*resolvedBuild = allResolvedBuild[selectedIndex]
	return nil
}
