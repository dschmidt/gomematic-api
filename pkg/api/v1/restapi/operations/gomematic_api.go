// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/auth"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/profile"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/team"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/user"
)

// NewGomematicAPI creates a new Gomematic instance
func NewGomematicAPI(spec *loads.Document) *GomematicAPI {
	return &GomematicAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		TeamAppendTeamToUserHandler: team.AppendTeamToUserHandlerFunc(func(params team.AppendTeamToUserParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamAppendTeamToUser has not yet been implemented")
		}),
		UserAppendUserToTeamHandler: user.AppendUserToTeamHandlerFunc(func(params user.AppendUserToTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation UserAppendUserToTeam has not yet been implemented")
		}),
		TeamCreateTeamHandler: team.CreateTeamHandlerFunc(func(params team.CreateTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamCreateTeam has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserCreateUser has not yet been implemented")
		}),
		TeamDeleteTeamHandler: team.DeleteTeamHandlerFunc(func(params team.DeleteTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamDeleteTeam has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserDeleteUser has not yet been implemented")
		}),
		UserDeleteUserFromTeamHandler: user.DeleteUserFromTeamHandlerFunc(func(params user.DeleteUserFromTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation UserDeleteUserFromTeam has not yet been implemented")
		}),
		TeamDelteTeamFromUserHandler: team.DelteTeamFromUserHandlerFunc(func(params team.DelteTeamFromUserParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamDelteTeamFromUser has not yet been implemented")
		}),
		TeamListTeamUsersHandler: team.ListTeamUsersHandlerFunc(func(params team.ListTeamUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamListTeamUsers has not yet been implemented")
		}),
		TeamListTeamsHandler: team.ListTeamsHandlerFunc(func(params team.ListTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamListTeams has not yet been implemented")
		}),
		UserListUserTeamsHandler: user.ListUserTeamsHandlerFunc(func(params user.ListUserTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation UserListUserTeams has not yet been implemented")
		}),
		UserListUsersHandler: user.ListUsersHandlerFunc(func(params user.ListUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation UserListUsers has not yet been implemented")
		}),
		AuthLoginUserHandler: auth.LoginUserHandlerFunc(func(params auth.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthLoginUser has not yet been implemented")
		}),
		TeamPermitTeamUserHandler: team.PermitTeamUserHandlerFunc(func(params team.PermitTeamUserParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamPermitTeamUser has not yet been implemented")
		}),
		UserPermitUserTeamHandler: user.PermitUserTeamHandlerFunc(func(params user.PermitUserTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation UserPermitUserTeam has not yet been implemented")
		}),
		AuthRefreshAuthHandler: auth.RefreshAuthHandlerFunc(func(params auth.RefreshAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthRefreshAuth has not yet been implemented")
		}),
		ProfileShowProfileHandler: profile.ShowProfileHandlerFunc(func(params profile.ShowProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileShowProfile has not yet been implemented")
		}),
		TeamShowTeamHandler: team.ShowTeamHandlerFunc(func(params team.ShowTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamShowTeam has not yet been implemented")
		}),
		UserShowUserHandler: user.ShowUserHandlerFunc(func(params user.ShowUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserShowUser has not yet been implemented")
		}),
		ProfileTokenProfileHandler: profile.TokenProfileHandlerFunc(func(params profile.TokenProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileTokenProfile has not yet been implemented")
		}),
		ProfileUpdateProfileHandler: profile.UpdateProfileHandlerFunc(func(params profile.UpdateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation ProfileUpdateProfile has not yet been implemented")
		}),
		TeamUpdateTeamHandler: team.UpdateTeamHandlerFunc(func(params team.UpdateTeamParams) middleware.Responder {
			return middleware.NotImplemented("operation TeamUpdateTeam has not yet been implemented")
		}),
		UserUpdateUserHandler: user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUpdateUser has not yet been implemented")
		}),
		AuthVerifyAuthHandler: auth.VerifyAuthHandlerFunc(func(params auth.VerifyAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthVerifyAuth has not yet been implemented")
		}),
	}
}

/*GomematicAPI API definition for Gomematic */
type GomematicAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// TeamAppendTeamToUserHandler sets the operation handler for the append team to user operation
	TeamAppendTeamToUserHandler team.AppendTeamToUserHandler
	// UserAppendUserToTeamHandler sets the operation handler for the append user to team operation
	UserAppendUserToTeamHandler user.AppendUserToTeamHandler
	// TeamCreateTeamHandler sets the operation handler for the create team operation
	TeamCreateTeamHandler team.CreateTeamHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// TeamDeleteTeamHandler sets the operation handler for the delete team operation
	TeamDeleteTeamHandler team.DeleteTeamHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// UserDeleteUserFromTeamHandler sets the operation handler for the delete user from team operation
	UserDeleteUserFromTeamHandler user.DeleteUserFromTeamHandler
	// TeamDelteTeamFromUserHandler sets the operation handler for the delte team from user operation
	TeamDelteTeamFromUserHandler team.DelteTeamFromUserHandler
	// TeamListTeamUsersHandler sets the operation handler for the list team users operation
	TeamListTeamUsersHandler team.ListTeamUsersHandler
	// TeamListTeamsHandler sets the operation handler for the list teams operation
	TeamListTeamsHandler team.ListTeamsHandler
	// UserListUserTeamsHandler sets the operation handler for the list user teams operation
	UserListUserTeamsHandler user.ListUserTeamsHandler
	// UserListUsersHandler sets the operation handler for the list users operation
	UserListUsersHandler user.ListUsersHandler
	// AuthLoginUserHandler sets the operation handler for the login user operation
	AuthLoginUserHandler auth.LoginUserHandler
	// TeamPermitTeamUserHandler sets the operation handler for the permit team user operation
	TeamPermitTeamUserHandler team.PermitTeamUserHandler
	// UserPermitUserTeamHandler sets the operation handler for the permit user team operation
	UserPermitUserTeamHandler user.PermitUserTeamHandler
	// AuthRefreshAuthHandler sets the operation handler for the refresh auth operation
	AuthRefreshAuthHandler auth.RefreshAuthHandler
	// ProfileShowProfileHandler sets the operation handler for the show profile operation
	ProfileShowProfileHandler profile.ShowProfileHandler
	// TeamShowTeamHandler sets the operation handler for the show team operation
	TeamShowTeamHandler team.ShowTeamHandler
	// UserShowUserHandler sets the operation handler for the show user operation
	UserShowUserHandler user.ShowUserHandler
	// ProfileTokenProfileHandler sets the operation handler for the token profile operation
	ProfileTokenProfileHandler profile.TokenProfileHandler
	// ProfileUpdateProfileHandler sets the operation handler for the update profile operation
	ProfileUpdateProfileHandler profile.UpdateProfileHandler
	// TeamUpdateTeamHandler sets the operation handler for the update team operation
	TeamUpdateTeamHandler team.UpdateTeamHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler
	// AuthVerifyAuthHandler sets the operation handler for the verify auth operation
	AuthVerifyAuthHandler auth.VerifyAuthHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *GomematicAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GomematicAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GomematicAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GomematicAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GomematicAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GomematicAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GomematicAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GomematicAPI
func (o *GomematicAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TeamAppendTeamToUserHandler == nil {
		unregistered = append(unregistered, "team.AppendTeamToUserHandler")
	}

	if o.UserAppendUserToTeamHandler == nil {
		unregistered = append(unregistered, "user.AppendUserToTeamHandler")
	}

	if o.TeamCreateTeamHandler == nil {
		unregistered = append(unregistered, "team.CreateTeamHandler")
	}

	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}

	if o.TeamDeleteTeamHandler == nil {
		unregistered = append(unregistered, "team.DeleteTeamHandler")
	}

	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}

	if o.UserDeleteUserFromTeamHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserFromTeamHandler")
	}

	if o.TeamDelteTeamFromUserHandler == nil {
		unregistered = append(unregistered, "team.DelteTeamFromUserHandler")
	}

	if o.TeamListTeamUsersHandler == nil {
		unregistered = append(unregistered, "team.ListTeamUsersHandler")
	}

	if o.TeamListTeamsHandler == nil {
		unregistered = append(unregistered, "team.ListTeamsHandler")
	}

	if o.UserListUserTeamsHandler == nil {
		unregistered = append(unregistered, "user.ListUserTeamsHandler")
	}

	if o.UserListUsersHandler == nil {
		unregistered = append(unregistered, "user.ListUsersHandler")
	}

	if o.AuthLoginUserHandler == nil {
		unregistered = append(unregistered, "auth.LoginUserHandler")
	}

	if o.TeamPermitTeamUserHandler == nil {
		unregistered = append(unregistered, "team.PermitTeamUserHandler")
	}

	if o.UserPermitUserTeamHandler == nil {
		unregistered = append(unregistered, "user.PermitUserTeamHandler")
	}

	if o.AuthRefreshAuthHandler == nil {
		unregistered = append(unregistered, "auth.RefreshAuthHandler")
	}

	if o.ProfileShowProfileHandler == nil {
		unregistered = append(unregistered, "profile.ShowProfileHandler")
	}

	if o.TeamShowTeamHandler == nil {
		unregistered = append(unregistered, "team.ShowTeamHandler")
	}

	if o.UserShowUserHandler == nil {
		unregistered = append(unregistered, "user.ShowUserHandler")
	}

	if o.ProfileTokenProfileHandler == nil {
		unregistered = append(unregistered, "profile.TokenProfileHandler")
	}

	if o.ProfileUpdateProfileHandler == nil {
		unregistered = append(unregistered, "profile.UpdateProfileHandler")
	}

	if o.TeamUpdateTeamHandler == nil {
		unregistered = append(unregistered, "team.UpdateTeamHandler")
	}

	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}

	if o.AuthVerifyAuthHandler == nil {
		unregistered = append(unregistered, "auth.VerifyAuthHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GomematicAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GomematicAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *GomematicAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *GomematicAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *GomematicAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *GomematicAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the gomematic API
func (o *GomematicAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GomematicAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teams/{teamID}/users"] = team.NewAppendTeamToUser(o.context, o.TeamAppendTeamToUserHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/{userID}/teams"] = user.NewAppendUserToTeam(o.context, o.UserAppendUserToTeamHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teams"] = team.NewCreateTeam(o.context, o.TeamCreateTeamHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teams/{teamID}"] = team.NewDeleteTeam(o.context, o.TeamDeleteTeamHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{userID}"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{userID}/teams"] = user.NewDeleteUserFromTeam(o.context, o.UserDeleteUserFromTeamHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teams/{teamID}/users"] = team.NewDelteTeamFromUser(o.context, o.TeamDelteTeamFromUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{teamID}/users"] = team.NewListTeamUsers(o.context, o.TeamListTeamUsersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams"] = team.NewListTeams(o.context, o.TeamListTeamsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userID}/teams"] = user.NewListUserTeams(o.context, o.UserListUserTeamsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = user.NewListUsers(o.context, o.UserListUsersHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/login"] = auth.NewLoginUser(o.context, o.AuthLoginUserHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teams/{teamID}/users"] = team.NewPermitTeamUser(o.context, o.TeamPermitTeamUserHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{userID}/teams"] = user.NewPermitUserTeam(o.context, o.UserPermitUserTeamHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/refresh"] = auth.NewRefreshAuth(o.context, o.AuthRefreshAuthHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile/self"] = profile.NewShowProfile(o.context, o.ProfileShowProfileHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{teamID}"] = team.NewShowTeam(o.context, o.TeamShowTeamHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{userID}"] = user.NewShowUser(o.context, o.UserShowUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile/token"] = profile.NewTokenProfile(o.context, o.ProfileTokenProfileHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/profile/self"] = profile.NewUpdateProfile(o.context, o.ProfileUpdateProfileHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teams/{teamID}"] = team.NewUpdateTeam(o.context, o.TeamUpdateTeamHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{userID}"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/verify/{token}"] = auth.NewVerifyAuth(o.context, o.AuthVerifyAuthHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GomematicAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *GomematicAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GomematicAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GomematicAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
