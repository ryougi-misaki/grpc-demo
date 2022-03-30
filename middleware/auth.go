package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientIdKey     = "client-id"
	ClientSecretKey = "client-secret"
)

func Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	// 从 ctx 获取到了 metadata 信息
	md, _ := metadata.FromIncomingContext(ctx)

	// 认证请求合法性
	cids := md.Get(ClientIdKey)
	sids := md.Get(ClientSecretKey)

	if len(cids) == 0 || len(sids) == 0{
		err = status.Errorf(codes.Unauthenticated, "需要认证")
		return
	}

	clientId := cids[0]
	clientSecret := sids[0]

	if !(clientId == "admin" && clientSecret == "123456") {
		err = status.Errorf(codes.Unauthenticated, "认证无效")
		return
	}

	// 认证通过, 请求交个后面的handler处理
	return handler(ctx, req)
}

func NewAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

type Authentication struct {
	clientID     string
	clientSecret string
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (
	map[string]string, error) {
	return map[string]string{
		ClientIdKey:     a.clientID,
		ClientSecretKey: a.clientSecret,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

