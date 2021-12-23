/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package auth

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"os"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func (a *Authenticator) VerifyIdToken(
	ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIdToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIdToken)
}

func New() (*Authenticator, error) {
	auth0Domain := os.Getenv("FIZZ_WEB_AUTH0_DOMAIN")
	if auth0Domain == "" {
		return nil, errors.New("auth0Domain not defined")
	}
	auth0ClientId := os.Getenv("FIZZ_WEB_AUTH0_CLIENT_ID")
	if auth0ClientId == "" {
		return nil, errors.New("auth0ClientId not defined")
	}
	auth0ClientSecret := os.Getenv("FIZZ_WEB_AUTH0_CLIENT_SECRET")
	if auth0ClientSecret == "" {
		return nil, errors.New("auth0ClientSecret not defined")
	}
	auth0CallbackUrl := os.Getenv("FIZZ_WEB_AUTH0_CALLBACK_URL")
	if auth0CallbackUrl == "" {
		return nil, errors.New("auth0CallbackUrl not defined")
	}

	provider, err := oidc.NewProvider(
		context.Background(), "https://"+auth0Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     auth0ClientId,
		ClientSecret: auth0ClientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "email", "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}
