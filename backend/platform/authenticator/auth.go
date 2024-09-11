package authenticator

import (
	"context"
	"errors"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+os.Getenv("AUTH0_DOMAIN")+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

// Example of the profile object
//
// map[
//     aud:                "..."
//     exp:                1726120262
//     family_name:        "Rabello"
//     given_name:         "Ryan"
//     iat:                1726084262
//     iss:                "https://dev-er8olzbkc0c5ok16.us.auth0.com/"
//     name:               "Ryan Rabello"
//     nickname:           "ryan.s.rabello"
//     picture:            "https://lh3.googleusercontent.com/a/ACg8ocJT45KLPNHqcnYXL8JFBD9hq2dGTUkKRdlx1_0MTtPTjimkO0dx0Q=s96-c"
//     sid:                "......"
//     sub:                "google-oauth2|109002322269706807378"
//     updated_at:         "2024-09-11T19:50:55.408Z"
// ]
