package client

import (
	"bytes"
	"fmt"
	"net/http"
)

type Endpoint struct {
	Url           CephApiEndpoint
	ApiVersion    CephApiVersion
	Method        string
	Payload       *bytes.Buffer
	HeadersAccept string
}

func (e *Endpoint) Headers() map[string]string {
	return map[string]string{
		"Accept":       fmt.Sprintf("application/%s", e.ApiVersion),
		"Content-Type": "application/json",
	}
}

func NewEndpointAuth(payload []byte) *Endpoint {
	return &Endpoint{
		Url:        CephApiEndpointAuth,
		ApiVersion: CephApiVersionV10,
		Method:     http.MethodPost,
		Payload:    bytes.NewBuffer(payload),
	}
}

func NewEndpointHealthFull() *Endpoint {
	return &Endpoint{
		Url:        CephApiEndpointHealthFull,
		ApiVersion: CephApiVersionV10,
		Method:     http.MethodGet,
	}
}

func NewEndpointBlockImage() *Endpoint {
	return &Endpoint{
		Url:        CephApiEndpointBlockImage,
		ApiVersion: CephApiVersionV20,
		Method:     http.MethodGet,
	}
}

func NewEndpointUsers() *Endpoint {
	return &Endpoint{
		Url:        CephApiEndpointUsers,
		ApiVersion: CephApiVersionV10,
		Method:     http.MethodGet,
	}
}

func NewPostEndpointUsers(payload []byte) *Endpoint {
	return &Endpoint{
		Url:        CephApiEndpointUsers,
		ApiVersion: CephApiVersionV10,
		Method:     http.MethodPost,
		Payload:    bytes.NewBuffer(payload),
	}
}

func NewEndpointDeleteUser(username string) *Endpoint {
	return &Endpoint{
		Url:        CephApiEndpoint(fmt.Sprintf("%s/%s", CephApiEndpointUsers, username)),
		ApiVersion: CephApiVersionV10,
		Method:     http.MethodDelete,
	}
}
