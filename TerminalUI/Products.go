package TerminalUI

import (
	"ChanDaoScript/Api"
)

type Products struct {
	list *List
}

func NewProducts() *Products {
	products := &Products{}
	products.list = NewList()
	return products
}

func (this *Products) Exec(productId *int) error {
	this.list.SetTitle("选择产品")
	var allProduct []Api.Product
	if err := Api.GetAllProduct(&allProduct); err != nil {
		return err
	}
	var options []string
	for _, product := range allProduct {
		options = append(options, product.Name)
	}
	this.list.SetOptions(options)
	var selectedIndex int
	if err := this.list.Exec(&selectedIndex); err != nil {
		return err
	}
	*productId = allProduct[selectedIndex].Id
	return nil
}
