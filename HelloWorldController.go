package quickstart

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type HelloWorldController struct {
	defaultName string
}

func NewHelloWorldController() *HelloWorldController {
	c := HelloWorldController{}
	c.defaultName = "Pip User"
	return &c
}

func (c *HelloWorldController) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.defaultName = config.GetAsStringWithDefault("default_name", c.defaultName)
}

func (c *HelloWorldController) Greeting(ctx context.Context, name string) (result string, err error) {
	if name == "" {
		name = c.defaultName
	}
	return "Hello, " + name + "!", nil
}
