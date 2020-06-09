package quickstart

import (
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rpcbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
)

type HelloWorldProcess struct {
	cproc.ProcessContainer
}

func NewHelloWorldProcess() *HelloWorldProcess {
	c := HelloWorldProcess{}
	c.ProcessContainer = *cproc.NewProcessContainer("hello-world", "HelloWorld microservice")
	c.SetConfigPath("./config.yml")
	c.AddFactory(NewHelloWorldServiceFactory())
	c.AddFactory(rpcbuild.NewDefaultRpcFactory())
	return &c
}
