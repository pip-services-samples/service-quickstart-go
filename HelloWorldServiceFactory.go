package quickstart

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type HelloWorldServiceFactory struct {
	*cbuild.Factory
}

func NewHelloWorldServiceFactory() *HelloWorldServiceFactory {
	c := HelloWorldServiceFactory{}
	c.Factory = cbuild.NewFactory()
	c.RegisterType(
		cref.NewDescriptor("hello-world", "controller", "default", "*", "1.0"),
		NewHelloWorldController,
	)
	c.RegisterType(
		cref.NewDescriptor("hello-world", "service", "http", "*", "1.0"),
		NewHelloWorldRestService,
	)
	return &c
}
