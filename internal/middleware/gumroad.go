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

type GumroadSubscriber struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	UserEmail   string `json:"user-email,omitempty"`
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}

type GumroadResponse struct {
	Success     bool                `json:"success"`
	Subscribers []GumroadSubscriber `json:"subscribers"`
}
