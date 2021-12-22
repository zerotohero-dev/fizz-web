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

package houston

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	ctx.HTML(http.StatusMethodNotAllowed, "houston.html", nil)
}
