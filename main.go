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

package main

import (
	"fmt"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-web/internal/auth"
	"github.com/zerotohero-dev/fizz-web/internal/router"
	"net/http"
	"os"
)

func initLogging() {
	sanitize := func() {}
	log.Init(log.InitParams{
		IsDevEnv:       false,
		LogDestination: os.Getenv("FIZZ_WEB_PAPERTRAIL_LOG_DESTINATION"),
		SanitizeFn:     sanitize,
		AppName:        "fizz-web",
	})
}

func initAuth() *auth.Authenticator {
	a, err := auth.New()
	if err != nil {
		log.Fatal(
			fmt.Sprintf("Failed to initialize the authenticator: %s", err.Error()),
		)
		return nil
	}

	return a
}

func initRoutes(a *auth.Authenticator) {
	rtr := router.New(a)

	err := http.ListenAndServe("0.0.0.0:8888", rtr)
	if err != nil {
		log.Fatal(
			fmt.Sprintf("Failed to initialize the server: %s", err.Error()),
		)
	}

	log.Info("Started serving at 0.0.0.0:8888")
}

func main() {
	initLogging()
	a := initAuth()
	initRoutes(a)
}
