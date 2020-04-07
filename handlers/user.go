package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jmc-quetzal/api/models"
	"github.com/jmc-quetzal/api/common"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

//CreateUserHandler takes in a service and returns a http handler
func CreateUserHandler(service models.UserService, sessions models.SessionService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := models.UserRegister{}
		err := json.NewDecoder(r.Body).Decode(&user)
		bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		pwd := string(bytes)
		if err != nil {
			panic(err)
		}
		err, pid := service.CreateUser(user.Username, user.Email, pwd, user.BirthDate)
		if err != nil {
			WriteError(w,422,err.Error())
			return
		}
		user.ID = pid
		sid := uuid.New()
		http.SetCookie(w, &http.Cookie{
			Name:    "sid",
			Path:     "/",
			Value:   sid.String(),
			Expires: time.Now().AddDate(0, 0, 7),
			HttpOnly: true,
		})
		err = sessions.Set(&models.Session{
			SessionId: sid,
			User:      &models.User{
				ID: pid,
			},
		})
		if err != nil {
			 log.Println("Error with Session in login handler")
		}
		WriteJSON(w, http.StatusCreated, user.User)
	})
}

func LoginHandler(service models.UserService, sessions models.SessionService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := models.UserRegister{}
		_ = json.NewDecoder(r.Body).Decode(&u)
		user, rErr := service.Login(u.Email,u.Password)
		if rErr != nil {
			print(rErr.Error())
			switch  e := rErr.(type) {
			case *common.DataError:
					WriteError(w,e.Code,e.Error())
					return
			}
		}
		if user == nil {
			WriteError(w, http.StatusUnauthorized, "Incorrect login")
		} else {
			sid := uuid.New()
			http.SetCookie(w, &http.Cookie{
				Name:    "sid",
				Path:     "/",
				HttpOnly: true,
				Value:   sid.String(),
				Expires: time.Now().AddDate(0, 0, 7),
			})

			err := sessions.Set(&models.Session{
				SessionId: sid,
				User:      user,
			})
			if err != nil {
				log.Println("Error with Session in login handler")
			}
			WriteJSON(w,200,user)
		}
	})
}


func LogoutHandler(sessions models.SessionService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		sid, err := r.Cookie("sid")
		if err != nil {
			WriteError(w,400,"internal error")
			return
		}
		sessionToken := sid.Value
		str, err := uuid.Parse(sessionToken)
		if err != nil {
			print("err is ", err.Error())
			WriteError(w,400, "Internal Error")
			return
		}
		err = sessions.Delete(str)
		if err != nil {
			WriteError(w,http.StatusInternalServerError,"Session Error")
			return
		}
		http.SetCookie(w,&http.Cookie{
			Name:"sid",
			Value: "",
			MaxAge: 0,
			HttpOnly: true,
			Path: "/",
		});
		WriteJSON(w,200,"Logged out")
	})
}
