package web

import (
	"fmt"
	"forum/users"
	"log"
	"net/http"
)

func OpenServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Secure!")
	})

	http.HandleFunc("/login", users.LoginHandler)
	http.HandleFunc("/loginauth", users.LoginAuthHandler)
	http.HandleFunc("/logout", users.LogoutHandler)

	http.HandleFunc("/register/", users.RegisterUserHandler)
	http.HandleFunc("/registerauth", users.RegisterAuthHandler)
	log.Fatal(http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", nil))
}
