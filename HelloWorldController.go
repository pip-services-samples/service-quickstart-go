package quickstart

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

type HelloWorldController struct {
	defaultName string
}

func NewHelloWorldController() *HelloWorldController {
	c := HelloWorldController{}
	c.defaultName = "Pip User"
	return &c
}

func (c *HelloWorldController) Configure(config *cconf.ConfigParams) {
	c.defaultName = config.GetAsStringWithDefault("default_name", c.defaultName)
}

func (c *HelloWorldController) Greeting(name string) (result string, err error) {
	if name == "" {
		name = c.defaultName
	}
	return "Hello, " + name + "!", nil
}
