package middleware

import (
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func NewHandler() http.Handler {

	return alice.New(recoverHandler, authHandler).ThenFunc(logicHandler)

}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Panic("Recoverd!! Panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func authHandler(next http.Handler) http.Handler {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/restricted":
	case "/login":
		switch r.Method {
		case "GET":
		case "POST":
		default:

		}
	case "/register":
		switch r.Method {
		case "GET":
		case "POST":
		default:
		}
	case "/logout":
	case "/deleteuser":
	default:

	}
}

func nullifyTokenCookies(w *http.ResponseWriter, r *http.Request) {

}

func setAuthAndRefreshCookies() {

}

func grabCsrfFromReq(r *http.Request) string {

}
