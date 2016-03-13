package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if _, err := req.Cookie("auth"); err == http.ErrNoCookie {
		// not authenticated
		rw.Header().Set("Location", "/login")
		rw.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		// some other error...
		panic(err.Error())
	} else {
		//success - call the next handler
		h.next.ServeHTTP(rw, req)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
