package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Service interface {
	RegisterService(g *gin.RouterGroup)
}

// AsService annotates the given constructor to state that
// it provides a Connect service to the "connectServices" group.
func AsService(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Service)),
		fx.ResultTags(`group:"connectServices"`),
	)
}
