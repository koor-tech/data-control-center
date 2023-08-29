package auth

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/koor-tech/data-control-center/pkg/grpc/auth/userinfo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	AuthAccIDCtxTag              = "auth.accid"
	AuthActiveCharIDCtxTag       = "auth.chrid"
	AuthActiveCharJobCtxTag      = "auth.chrjob"
	AuthActiveCharJobGradeCtxTag = "auth.chrjobg"
	AuthSubCtxTag                = "auth.sub"
)

var UserInfoKey struct{}

var (
	ErrNoToken      = status.Errorf(codes.Unauthenticated, "No token given, please login!")
	ErrInvalidToken = status.Error(codes.Unauthenticated, "Invalid token! Please login again.")
)

type GRPCAuth struct {
	tm *TokenMgr
}

func NewGRPCAuth(tm *TokenMgr) *GRPCAuth {
	return &GRPCAuth{
		tm: tm,
	}
}

func (g *GRPCAuth) NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			t := req.Header().Get("authorization")

			if t == "" {
				return nil, ErrNoToken
			}

			t = strings.TrimPrefix(t, "Bearer ")
			// Parse token only returns the token info when the token is still valid
			tInfo, err := g.tm.ParseWithClaims(t)
			if err != nil {
				return nil, ErrInvalidToken
			}

			ctx = context.WithValue(ctx, UserInfoKey, &userinfo.UserInfo{
				AccId:    tInfo.AccID,
				Username: tInfo.Username,
			})

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
