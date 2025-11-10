package jwtauth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthFactory interface {
		SignToken() string
	}

	Claims struct {
		Id       string `json:"id"`
		RoleCode int    `json:"role_code"`
	}

	AuthMapClaims struct {
		*Claims
		jwt.RegisteredClaims
	}

	authConcrete struct {
		Secret []byte
		Claims *AuthMapClaims `json:"claims"`
	}

	accessToken struct{ *authConcrete }

	refreshToken struct{ *authConcrete }

	apiKey struct{ *authConcrete }
)

func NewAuthFactory(factory AuthFactory) AuthFactory {
	return factory
}

func (a *authConcrete) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	ss, _ := token.SignedString(a.Secret)
	return ss
}
func now() time.Time {
	loc, _ := time.LoadLocation("Asia/Bankok")
	return time.Now().In(loc)
}

// t is second unit
func jwtTimeDurationCal(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(now().Add(time.Duration(t) * time.Second))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func NewAccessToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &accessToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "go-shop",
					Subject:   "access-token",
					Audience:  jwt.ClaimStrings{"go-shop.com"},
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "go-shop",
					Subject:   "refresh-token",
					Audience:  []string{"go-shop.com"},
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ReloadToken(secret string, expireAt int64, claims *Claims) string {
	obj := &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "go-shop",
					Subject:   "refresh-token",
					Audience:  []string{"go-shop.com"},
					ExpiresAt: jwtTimeRepeatAdapter(expireAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
	return obj.SignToken()
}

func NewApiKey(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &apiKey{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "go-shop",
					Subject:   "api-key",
					Audience:  []string{"go-shop.com"},
					ExpiresAt: jwtTimeDurationCal(31560000),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenString string) (*AuthMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthMapClaims{}, func(t *jwt.Token) (interface{}, error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: unexpecting signing method")
		}
		return []byte(secret) , nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		}else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil , errors.New("error: token is expired")
		}else {
			return nil, errors.New("error: unable to parse token")
		}
	}

	if claims, ok := token.Claims.(*AuthMapClaims): ok {
		return claims, nil
	}else {
		return nil, errors.New("error: token is invalid")
	}


}
