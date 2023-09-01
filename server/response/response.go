package response

import (
	"connectrpc.com/connect"
	"github.com/koor-tech/data-control-center/gen/go/api/services/response"
)

func NewSuccessResponse(statusCode int32, resp *response.Response_ClusterStatusResponse) *connect.Response[response.Response] {
	return &connect.Response[response.Response]{
		Msg: &response.Response{
			StatusCode: statusCode,
			Message:    "Cluster stats retrieved successfully",
			Response:   resp,
		},
	}
}

func NewFailureResponse(statusCode int32, message error) *connect.Response[response.Response] {
	return &connect.Response[response.Response]{
		Msg: &response.Response{
			StatusCode: statusCode,
			Message:    message.Error(),
			//Response:   resp,
		},
	}
}
