package urihandler

import (
	ecombase "github.com/codedv8/go-ecom-base"
	"github.com/gin-gonic/gin"
)

// URIHandler - Struct managing uri's registered to the system
type URIHandler struct {
	URIList *ecombase.LinkedTree
}

// Init - Initialize the URIHandler
func (uh *URIHandler) Init() {
	uh.URIList = &ecombase.LinkedTree{}
}

// AddURI - Add a new uri and a handler for it
func (uh *URIHandler) AddURI(uri string, handler func(ctx *gin.Context) (bool, error)) {
	uh.URIList.Add(uri, handler)
}

// HandleURI - Find if threr is a registered uri that matches the argument and call it
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
