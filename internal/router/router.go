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

package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-web/internal/authenticator"
	"github.com/zerotohero-dev/fizz-web/internal/middleware"
	"github.com/zerotohero-dev/fizz-web/web/app/callback"
	"github.com/zerotohero-dev/fizz-web/web/app/healthz"
	"github.com/zerotohero-dev/fizz-web/web/app/home"
	"github.com/zerotohero-dev/fizz-web/web/app/houston"
	"github.com/zerotohero-dev/fizz-web/web/app/login"
	"github.com/zerotohero-dev/fizz-web/web/app/logout"
	"github.com/zerotohero-dev/fizz-web/web/app/questions"
	"github.com/zerotohero-dev/fizz-web/web/app/subscribe"
	"os"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	mongoUrl := os.Getenv("FIZZ_WEB_MONGODB_CONNECTION_STRING")
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		log.Fatal("Problem connecting Mongo: " + err.Error())
		return nil
	}

	c := session.DB("fizz").C("sessions")
	// 2592000 == ~1month (in seconds)
	store := mongo.NewStore(
		c, 2592000, true, []byte(os.Getenv("FIZZ_WEB_SESSION_SECRET")),
	)

	router.Use(sessions.Sessions("auth-session", store))

	router.LoadHTMLGlob("web/template/*")

	// Home.
	router.GET(
		"/",
		middleware.Canonical,
		home.Handler(),
	)

	// Liveness probe.
	router.GET(
		"/healthz",
		healthz.Handler(),
	)

	// Gumroad and Auth0 can redirect here upon failure.
	router.GET(
		"/error",
		middleware.Canonical,
		houston.Handler(),
	)

	// User and identity management (auth0 integration)
	router.GET(
		"/auth/callback",
		middleware.Canonical,
		callback.Handler(auth),
	)
	router.GET(
		"/login",
		middleware.Canonical,
		login.Handler(auth),
	)
	router.GET(
		"/logout",
		middleware.Canonical,
		logout.Handler(),
	)

	// Gumroad integration.
	router.GET(
		"/subscribe",
		middleware.Canonical,
		middleware.IsAuthenticated,
		middleware.IsNotSubscribed,
		subscribe.Handler(),
	)

	// Free routes.
	router.GET(
		"/about/*path",
		middleware.Canonical,
		questions.Handler(),
	)
	router.GET(
		"/concepts/*path",
		middleware.Canonical,
		questions.Handler(),
	)
	router.GET(
		"/warm-up/*path",
		middleware.Canonical,
		questions.Handler(),
	)

	// Premium routes.
	router.GET(
		"/pro/*path",
		middleware.Canonical,
		middleware.IsAuthenticated,
		middleware.IsSubscribed,
		questions.Handler(),
	)

	return router
}
