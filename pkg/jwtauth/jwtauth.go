package jwtauth

type (
	AuthFactory interface {
		SignToken() string
	}

	authConcrete struct {

	}

	accessToken struct {}

	refreshToken struct {}

	apiKey struct {}


)