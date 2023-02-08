package main

import (
	"fmt"
	"learngo/interface/mock"
	"learngo/interface/real"
)

type Receiver interface {
	GetRes(s string) string
}
type ReceiverPoster interface {
	GetRes(s string) string
	GetResPoster(s string) string
}

func response(r Receiver) string {
	return r.GetRes("https://www.imooc.com")
}

func main() {
	var r Receiver
	r = mock.Receiver{Content: "fake interface"}
	fmt.Println(response(r))
	r = real.Receiver{}
	fmt.Println(response(r))
}
