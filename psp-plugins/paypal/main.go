package main

import "fmt"

type plugin struct {
}

func (p plugin) Test() string {
	fmt.Println("Plug-in plug-out wasaaaaaaaaaa")
	return "ups"
}

var Plugin plugin
