package auth

import (
	"context"
)

type Authentication struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (a Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	//TODO implement me
	//panic("implement me")
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (a Authentication) RequireTransportSecurity() bool {
	//TODO implement me
	//panic("implement me")
	return false
}
