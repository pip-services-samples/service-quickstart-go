package quickstart

import (
	"net/http"

	crefer "github.com/pip-services3-go/pip-services3-commons-go/refer"
	rpc "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type HelloWorldRestService struct {
	*rpc.RestService
	controller *HelloWorldController
}

func NewHelloWorldRestService() *HelloWorldRestService {
	c := HelloWorldRestService{}
	c.RestService = rpc.NewRestService()
	c.RestService.IRegisterable = &c
	c.BaseRoute = "/hello_world"
	c.DependencyResolver.Put("controller", crefer.NewDescriptor("hello-world", "controller", "*", "*", "1.0"))
	return &c
}

func (c *HelloWorldRestService) SetReferences(references crefer.IReferences) {

	c.RestService.SetReferences(references)
	depRes, depErr := c.DependencyResolver.GetOneRequired("controller")
	if depErr == nil && depRes != nil {
		c.controller = depRes.(*HelloWorldController)
	}
}

func (c *HelloWorldRestService) greeting(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	result, err := c.controller.Greeting(name)
	c.SendResult(res, req, result, err)

}

func (c *HelloWorldRestService) Register() {
	c.RegisterRoute("get", "/greeting", nil, c.greeting)
}
