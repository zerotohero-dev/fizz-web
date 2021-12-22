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

package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	// TODO:
	// Home is a public path, you don’t need session verification.
	// but a “welcome $user | logout" and a "login" would be nice
	// in the header of not only home but also all other paths.

	ctx.HTML(http.StatusOK, "home.html", nil)
}
