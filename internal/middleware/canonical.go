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

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
)

// Canonical redirects www routes to their non-www version.
// (i.e, `www.fizzbuzz.pro/yeet` yeets to `fizzbuzz.pro/yeet`.)
// This fixes some Oauth0 and Gumroad redirection edge cases.
func Canonical(ctx *gin.Context) {
	if strings.Index(ctx.Request.Host, "www.") != -1 {
		ctx.Redirect(
			http.StatusSeeOther,
			"https://"+path.Join("fizzbuzz.pro", ctx.Request.RequestURI),
		)
		return
	}

	ctx.Next()
}
