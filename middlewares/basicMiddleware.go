package middlewares

import (
	"log"
	"net/http"
)

//Simple middleware - just for printing
func BasicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		}) // if you place ) at next line, you will get an error.
}

// The above ) problem is a go Parser syntax limitation. You need to add , after }
/*  	next.ServeHTTP(w, r)
	},
)
*/
