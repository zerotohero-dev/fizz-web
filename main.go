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
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-web/platform/authenticator"
	"github.com/zerotohero-dev/fizz-web/platform/router"
	"net/http"
	"os"
)

func sanitize() {
}

func main() {
	log.Init(log.InitParams{
		IsDevEnv:       false,
		LogDestination: os.Getenv("FIZZ_WEB_PAPERTRAIL_LOG_DESTINATION"),
		SanitizeFn:     sanitize,
		AppName:        "fizz-web",
	})

	auth, err := authenticator.New()
	if err != nil {
		log.Fatal("Failed to initialize the authenticator")
		return
	}

	rtr := router.New(auth)

	log.Info("fizz-web is listening on http://localhost:8888/")
	if err := http.ListenAndServe("0.0.0.0:8888", rtr); err != nil {
		log.Info("There was an error with the http server: %v", err)
	}
}
