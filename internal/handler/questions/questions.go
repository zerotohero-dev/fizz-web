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

package questions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
)

func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.String()

		// Strip the querystring: We don’t need it.
		if strings.Index(url, "?") > -1 {
			url = url[:strings.Index(url, "?")]
		}

		// Only extensions we allow are `.go.html` and `.html`.
		// Replace all `go.html` and `.html`. If there are still `.`s
		// remaining after the replacement, then it is a malformed url.
		if strings.Index(
			strings.Replace(
				strings.Replace(url, ".go.html", "", 1),
				".html", "", 1,
			),
			".",
		) > -1 {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		// Do not process anything you don’t understand.
		if strings.Index(url, "/warm-up") != 0 &&
			strings.Index(url, "/about") != 0 &&
			strings.Index(url, "/concepts") != 0 &&
			strings.Index(url, "/pro") != 0 {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		// Disallow directory listing.
		if url == "/warm-up/" || url == "/warm-up" {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		// Disallow directory listing.
		if url == "/about/" || url == "/about" {
			ctx.Redirect(http.StatusPermanentRedirect, "/about/doc.go.html")
			return
		}

		// Disallow directory listing.
		if url == "/concepts/" || url == "/concepts" {
			ctx.Redirect(http.StatusPermanentRedirect, "/concepts/doc.go.html")
			return
		}

		// Disallow directory listing.
		if url == "/pro/" || url == "/pro" {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		// All the questions and source code end with .go.html
		// Anything else is likely some mangled url.
		if !strings.HasSuffix(url, ".go.html") {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		ctx.File(path.Join("/usr/local/share/fizz/dist", url))
	}
}
