package dfx

import (
	"errors"
	"strings"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gstones/moke-kit/server/pkg/sfx"
)

// Author is a demo auth service
type Author struct {
}

// Auth is a demo auth method
func (d *Author) Auth(token string) (string, error) {
	if !strings.Contains(token, "test") {
		return "", errors.New("token error")
	}
	return token, nil
}

// AuthModule is a demo auth module
// you can implement your own auth service or auth function
// Auth will check every rpc/http request,
// if you want to disable it for a service, add `utility.WithoutAuth` in struct of your service
//
//	type service struct {
//		utility.WithoutAuth
//	}
var AuthModule = fx.Provide(
	func(
		l *zap.Logger,
	) (out sfx.AuthServiceResult, err error) {
		out.AuthService = &Author{}
		return
	},
)
