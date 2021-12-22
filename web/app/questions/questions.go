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
	"strings"
)

func Handler(ctx *gin.Context) {
	url := ctx.Request.URL.String()

	// Strip the querystring: We don’t need it.
	// Plus, this will eliminate a whole family of injection attacks.
	if strings.Index(url, "?") > -1 {
		url = url[:strings.Index(url, "?")]
	}

	// Only extensions we allow are `.go.html` and `.html`.
	// If, when we replace those extensions “once”, there is still
	// `.` in the url, then that is a URL that we don’t recognize.
	// Yeet the user to the web root.
	//
	// TODO: maybe do a regex matcher instead.
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

	// These are the context roots that this handler understands.
	// Anything else will result in a yeet to the web root.
	if strings.Index(url, "/warm-up") != 0 &&
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
	if url == "/concepts/" || url == "/concepts" {
		ctx.Redirect(http.StatusSeeOther, "/concepts/doc.go.html")
		return
	}

	// Disallow directory listing.
	if url == "/pro/" || url == "/pro" {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	// Serve the file.
	ctx.File("/usr/local/share/fizz/dist" + url)
}
