package main

import "go-examples/captcha"

func main() {
	c := captcha.Create()
	println(c.Captcha())
}
