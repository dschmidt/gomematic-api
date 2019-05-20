// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// LoginUserOKCode is the HTTP code returned for type LoginUserOK
const LoginUserOKCode int = 200

/*LoginUserOK A generated token with expire

swagger:response loginUserOK
*/
type LoginUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthToken `json:"body,omitempty"`
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {

	return &LoginUserOK{}
}

// WithPayload adds the payload to the login user o k response
func (o *LoginUserOK) WithPayload(payload *models.AuthToken) *LoginUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user o k response
func (o *LoginUserOK) SetPayload(payload *models.AuthToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LoginUserUnauthorizedCode is the HTTP code returned for type LoginUserUnauthorized
const LoginUserUnauthorizedCode int = 401

/*LoginUserUnauthorized Unauthorized if wrong credentials

swagger:response loginUserUnauthorized
*/
type LoginUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewLoginUserUnauthorized creates LoginUserUnauthorized with default headers values
func NewLoginUserUnauthorized() *LoginUserUnauthorized {

	return &LoginUserUnauthorized{}
}

// WithPayload adds the payload to the login user unauthorized response
func (o *LoginUserUnauthorized) WithPayload(payload *models.GeneralError) *LoginUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user unauthorized response
func (o *LoginUserUnauthorized) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*LoginUserDefault Some error unrelated to the handler

swagger:response loginUserDefault
*/
type LoginUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewLoginUserDefault creates LoginUserDefault with default headers values
func NewLoginUserDefault(code int) *LoginUserDefault {
	if code <= 0 {
		code = 500
	}

	return &LoginUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the login user default response
func (o *LoginUserDefault) WithStatusCode(code int) *LoginUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the login user default response
func (o *LoginUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the login user default response
func (o *LoginUserDefault) WithPayload(payload *models.GeneralError) *LoginUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user default response
func (o *LoginUserDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
