package middleware

import (
	"github.com/google/uuid"
	"github.com/jmc-quetzal/api/handlers"
	"github.com/jmc-quetzal/api/models"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func WithAuth (sessionStore models.SessionService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sid")
		if err != nil {
			if err == http.ErrNoCookie {
				handlers.WriteJSON(w,401,"No cookie")
				return
			}
			handlers.WriteJSON(w,401,"Error with cookie")
			return
		}
		sessionToken := c.Value
		uid, err := uuid.Parse(sessionToken)
		exists, err  := sessionStore.Get(uid)
		if exists == false {
			handlers.WriteJSON(w,401,"Unauthorized")
			return
		}
		next(w,r);
	}
}