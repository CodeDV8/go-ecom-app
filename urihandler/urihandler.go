package urihandler

import (
	EComBase "github.com/codedv8/go-ecom-base"
	"github.com/gin-gonic/gin"
)

type URIHandler struct {
	URIList *EComBase.LinkedTree
}

func (uh *URIHandler) Init() {
	uh.URIList = &EComBase.LinkedTree{}
}

func (uh *URIHandler) AddURI(uri string, handler func(ctx *gin.Context) (bool, error)) {
	uh.URIList.Add(uri, handler)
}

func (uh *URIHandler) HandleURI(uri string, c *gin.Context) (bool, error) {
	node, err := uh.URIList.FindNode(uri)
	if err != nil {
		return false, err
	}
	if node == nil {
		return false, nil
	}
	if node.Data == nil {
		return false, nil
	}
	call := (node.Data).(func(*gin.Context) (bool, error))
	return call(c)
}
