package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

//func init() {
//	tokenAuth = jwtauth.New("HS256", []byte(viper.GetString("secret.key")), nil)
//}

func Routes(c OrderController) chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.CreateOrder)
	return r
}
