// middleware/optional_jwt.go
package middleware

import (
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// “任意認証”ミドルウェアのための自作ミドルウェア
func OptionalAuth() (echo.MiddlewareFunc, error) {
	issuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	issuerURL, err := url.Parse(issuer)
	if err != nil {
		return nil, err
	}

	// Auth0 推奨の JWKS キャッシュ付きプロバイダ
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	v, err := validator.New(
		provider.KeyFunc,
		validator.RS256,                       // Auth0 既定の RS256 を想定
		issuer,                                // iss 検証
		[]string{os.Getenv("AUTH0_AUDIENCE")}, // aud 検証（Auth0 の API Identifier）
		validator.WithCustomClaims(func() validator.CustomClaims { return &CustomClaims{} }),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			h := c.Request().Header.Get("Authorization")
			if h == "" || !strings.HasPrefix(h, "Bearer ") {
				// トークン無し → そのまま通す（匿名）
				return next(c)
			}
			tokStr := strings.TrimPrefix(h, "Bearer ")

			// トークン検証（失敗しても 401 を返さない）
			tok, err := v.ValidateToken(c.Request().Context(), tokStr)
			if err == nil && tok != nil {
				if vt, ok := tok.(*validator.ValidatedClaims); ok && vt != nil {
					// 必要最小限だけ context に差し込む
					c.Set("sub", vt.RegisteredClaims.Subject)
					if cc, ok := vt.CustomClaims.(*CustomClaims); ok && cc != nil {
						c.Set("scope", cc.Scope)
					}
					// 必要なら tok 自体も保持
					c.Set("token", vt)
				}
			}
			// 成功/失敗に関わらず通す
			return next(c)
		}
	}, nil
}
