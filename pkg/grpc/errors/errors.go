package grpcerrors

import (
	"fmt"

	"connectrpc.com/connect"
)

var (
	ErrReadOnly = connect.NewError(connect.CodeInternal, fmt.Errorf("data-control-center is in read-only mode"))
)
