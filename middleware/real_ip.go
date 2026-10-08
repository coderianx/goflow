package middleware

import "net/http"

// RealIP sets the remote address of the request to the value of the
// X-Real-IP header, if present, or the remote address of the request.
//
// It is intended to be used as a middleware that wraps the next handler.
func RealIP() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ip := r.Header.Get("X-Real-IP")

			if ip == "" {
				ip = r.RemoteAddr
			}

			r.RemoteAddr = ip
			next.ServeHTTP(w, r)
		})
	}
}
