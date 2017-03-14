package api

import (
	"encoding/json"
	"github.com/gocraft/web"
	"github.com/gorilla/context"
	"github.com/plutov/go-docker-compose-skeleton/pkg/env"
	"github.com/plutov/go-docker-compose-skeleton/pkg/shared"
	"io/ioutil"
	"net/http"
)

// Context struct
type Context struct {
	Env env.Context
}

// AjaxResponse struct
type AjaxResponse struct {
	Code  int
	Error string
	Data  interface{}
}

// Ajax func
func (c *Context) Ajax(w web.ResponseWriter, r *web.Request, response interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resultJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
}

// ListenAndServe func
func ListenAndServe(e env.Context) {
	ctx := new(Context)
	ctx.Env = e

	r := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Get("/swagger.json", ctx.swagger)

	shared.LogErr(http.ListenAndServe(ctx.Env.Config.Addr, context.ClearHandler(r)))
}

func (c *Context) swagger(w web.ResponseWriter, r *web.Request) {
	f, err := ioutil.ReadFile("swagger.json")
	shared.LogErr(err)
	if err == nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(f)
	}
}
