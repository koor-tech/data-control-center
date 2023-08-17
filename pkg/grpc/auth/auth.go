package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
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

const (
	PermSuperUser = "SuperUser"
	PermAny       = "Any"
)

var UserInfoKey struct{}

var (
	ErrNoToken          = status.Errorf(codes.Unauthenticated, "errors.pkg-auth.ErrNoToken")
	ErrInvalidToken     = status.Error(codes.Unauthenticated, "errors.pkg-auth.ErrInvalidToken")
	ErrCheckToken       = status.Error(codes.Unauthenticated, "errors.pkg-auth.ErrCheckToken")
	ErrUserNoPerms      = status.Error(codes.PermissionDenied, "errors.pkg-auth.ErrUserNoPerms")
	ErrNoUserInfo       = status.Error(codes.Unauthenticated, "errors.pkg-auth.ErrNoUserInfo")
	ErrPermissionDenied = status.Errorf(codes.PermissionDenied, "errors.pkg-auth.ErrPermissionDenied")
)

type GRPCAuth struct {
	tm *TokenMgr
}

func NewGRPCAuth(tm *TokenMgr) *GRPCAuth {
	return &GRPCAuth{
		tm: tm,
	}
}

func (g *GRPCAuth) GRPCAuthFunc(ctx context.Context, fullMethod string) (context.Context, error) {
	t, err := GetTokenFromGRPCContext(ctx)
	if err != nil {
		return nil, err
	}

	if t == "" {
		return nil, ErrNoToken
	}

	// Parse token only returns the token info when the token is still valid
	tInfo, err := g.tm.ParseWithClaims(t)
	if err != nil {
		return nil, ErrInvalidToken
	}

	ctx = logging.InjectFields(ctx, logging.Fields{
		AuthSubCtxTag, tInfo.Subject,
		AuthAccIDCtxTag, tInfo.AccID,
	})

	return context.WithValue(ctx, UserInfoKey, tInfo), nil
}

func (g *GRPCAuth) GRPCAuthFuncWithoutUserInfo(ctx context.Context, fullMethod string) (context.Context, error) {
	t, err := GetTokenFromGRPCContext(ctx)
	if err != nil {
		return nil, err
	}

	if t == "" {
		return nil, ErrNoToken
	}

	// Parse token only returns the token info when the token is still valid
	tInfo, err := g.tm.ParseWithClaims(t)
	if err != nil {
		return nil, ErrInvalidToken
	}

	ctx = logging.InjectFields(ctx, logging.Fields{
		AuthSubCtxTag, tInfo.Subject,
		AuthAccIDCtxTag, tInfo.AccID,
	})

	return ctx, nil
}

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

func MustGetUserInfoFromContext(ctx context.Context) *userinfo.UserInfo {
	userInfo, ok := FromContext(ctx)
	if !ok {
		panic(ErrNoUserInfo)
	}
	return userInfo
}
