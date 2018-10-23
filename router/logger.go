/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package router

import (
	"net/http"
	"time"

	"github.com/VuliTv/api/libs/logging"
)

var log = logging.GetProdLog()

// Logger --
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Infow("api_call",
			"type", w.Header()["Content-Type"],
			"request", r.Method,
			"uri", r.RequestURI,
			"method", name,
			"response", time.Since(start),
		)
	})
}
