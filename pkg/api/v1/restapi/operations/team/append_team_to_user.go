// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AppendTeamToUserHandlerFunc turns a function with the right signature into a append team to user handler
type AppendTeamToUserHandlerFunc func(AppendTeamToUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AppendTeamToUserHandlerFunc) Handle(params AppendTeamToUserParams) middleware.Responder {
	return fn(params)
}

// AppendTeamToUserHandler interface for that can handle valid append team to user params
type AppendTeamToUserHandler interface {
	Handle(AppendTeamToUserParams) middleware.Responder
}

// NewAppendTeamToUser creates a new http.Handler for the append team to user operation
func NewAppendTeamToUser(ctx *middleware.Context, handler AppendTeamToUserHandler) *AppendTeamToUser {
	return &AppendTeamToUser{Context: ctx, Handler: handler}
}

/*AppendTeamToUser swagger:route POST /teams/{teamID}/users team appendTeamToUser

Assign a user to team

*/
type AppendTeamToUser struct {
	Context *middleware.Context
	Handler AppendTeamToUserHandler
}

func (o *AppendTeamToUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAppendTeamToUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
