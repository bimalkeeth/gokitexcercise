package account

import (
	"context"
	httpTrans "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context,endpoints EndPoints) http.Handler{
    r:=mux.NewRouter()
    r.Use(commonMiddleWare)
    r.Methods("POST").Path("/user").Handler(httpTrans.NewServer(
    	 endpoints.CreateUser,
    	 decodeUserReq,
    	 encodeResponse,
    	))

	r.Methods("GET").Path("/user/{id}").Handler(httpTrans.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))
    return r
}
func commonMiddleWare(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		w.Header().Add("Content-Type","application/json")
		next.ServeHTTP(w,r)
	})
}
