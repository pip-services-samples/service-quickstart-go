package quickstart

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type HelloWorldServiceFactory struct {
	cbuild.Factory
}

func NewHelloWorldServiceFactory() *HelloWorldServiceFactory {
	c := HelloWorldServiceFactory{}
	c.Factory = *cbuild.NewFactory()
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
