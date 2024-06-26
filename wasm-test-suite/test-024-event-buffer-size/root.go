package main

import (
	"log"
	"math/rand"

	"github.com/vugu/vugu"
)

type Option struct {
	value  string
	output string
}
type Root struct {
	OptionsList []Option
}

func (c *Root) Init(ctx vugu.InitCtx) {
	foobar := Option{}

	for a := 0; a < 1024; a++ {
		s := randStringBytes(50)
		foobar.value = s
		foobar.output = s
		c.OptionsList = append(c.OptionsList, foobar)
	}
}
func (c *Root) OutputEvent(e vugu.DOMEvent) {
	log.Println(e.PropString("target", "value"), "target value")
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
