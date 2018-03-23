package main

import "github.com/automationbroker/bundle-controller/pkg/controller"

func main() {
	c := controller.CreateController()
	c.Start()
}
