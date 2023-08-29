package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth/userinfo"
)

func FromContext(ctx context.Context) (*userinfo.UserInfo, bool) {
	c, ok := ctx.Value(UserInfoKey).(*userinfo.UserInfo)
	return c, ok
}

func GetTokenFromGRPCContext(ctx context.Context) (string, error) {
	return grpc_auth.AuthFromMD(ctx, "bearer")
}

func GetUserInfoFromContext(ctx context.Context) (*userinfo.UserInfo, bool) {
	return FromContext(ctx)
}
