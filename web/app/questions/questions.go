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

func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.String()

		// Strip the query part.
		if strings.Index(url, "?") > -1 {
			url = url[:strings.Index(url, "?")]
		}

		// Poor man’s directory traversal prevention.
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

		if url == "/warm-up/" || url == "/warm-up" {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		if url == "/about/" || url == "/about" {
			ctx.Redirect(http.StatusSeeOther, "/about/doc.go.html")
			return
		}

		if url == "/about/" || url == "/about" {
			ctx.Redirect(http.StatusSeeOther, "/about/doc.go.html")
			return
		}

		if url == "/pro/" || url == "/pro" {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}

		ctx.File("/usr/local/share/fizz/dist" + url)
	}
}
