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

package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for health check.
func Handler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
