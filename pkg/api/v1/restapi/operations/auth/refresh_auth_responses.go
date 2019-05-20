// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// RefreshAuthOKCode is the HTTP code returned for type RefreshAuthOK
const RefreshAuthOKCode int = 200

/*RefreshAuthOK A refreshed token with expire

swagger:response refreshAuthOK
*/
type RefreshAuthOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthToken `json:"body,omitempty"`
}

// NewRefreshAuthOK creates RefreshAuthOK with default headers values
func NewRefreshAuthOK() *RefreshAuthOK {

	return &RefreshAuthOK{}
}

// WithPayload adds the payload to the refresh auth o k response
func (o *RefreshAuthOK) WithPayload(payload *models.AuthToken) *RefreshAuthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the refresh auth o k response
func (o *RefreshAuthOK) SetPayload(payload *models.AuthToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RefreshAuthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RefreshAuthUnauthorizedCode is the HTTP code returned for type RefreshAuthUnauthorized
const RefreshAuthUnauthorizedCode int = 401

/*RefreshAuthUnauthorized Unauthorized if failed to generate

swagger:response refreshAuthUnauthorized
*/
type RefreshAuthUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewRefreshAuthUnauthorized creates RefreshAuthUnauthorized with default headers values
func NewRefreshAuthUnauthorized() *RefreshAuthUnauthorized {

	return &RefreshAuthUnauthorized{}
}

// WithPayload adds the payload to the refresh auth unauthorized response
func (o *RefreshAuthUnauthorized) WithPayload(payload *models.GeneralError) *RefreshAuthUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the refresh auth unauthorized response
func (o *RefreshAuthUnauthorized) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RefreshAuthUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*RefreshAuthDefault Some error unrelated to the handler

swagger:response refreshAuthDefault
*/
type RefreshAuthDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewRefreshAuthDefault creates RefreshAuthDefault with default headers values
func NewRefreshAuthDefault(code int) *RefreshAuthDefault {
	if code <= 0 {
		code = 500
	}

	return &RefreshAuthDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the refresh auth default response
func (o *RefreshAuthDefault) WithStatusCode(code int) *RefreshAuthDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the refresh auth default response
func (o *RefreshAuthDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the refresh auth default response
func (o *RefreshAuthDefault) WithPayload(payload *models.GeneralError) *RefreshAuthDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the refresh auth default response
func (o *RefreshAuthDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RefreshAuthDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
