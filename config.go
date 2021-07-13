package auth

import (
	g "github.com/go-ginger/ginger"
	"github.com/mohammadMghi/auth/base"
	"github.com/mohammadMghi/auth/password"
)

type Config struct {
	base.Config

	AuthRouters []*g.RouterGroup
	Router      *g.RouterGroup
	Password    *password.Config
}
