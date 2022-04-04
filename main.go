package main

import "github.com/asciifaceman/incremental-emails/pkg/core"

var version = "localdev"

func main() {
	c := core.New(version)
	c.Init()
}
