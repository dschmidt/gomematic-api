// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/gomematic/gomematic-api/pkg/api/v1/models"
)

// TokenProfileOKCode is the HTTP code returned for type TokenProfileOK
const TokenProfileOKCode int = 200

/*TokenProfileOK The unlimited auth token

swagger:response tokenProfileOK
*/
type TokenProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.AuthToken `json:"body,omitempty"`
}

// NewTokenProfileOK creates TokenProfileOK with default headers values
func NewTokenProfileOK() *TokenProfileOK {

	return &TokenProfileOK{}
}

// WithPayload adds the payload to the token profile o k response
func (o *TokenProfileOK) WithPayload(payload *models.AuthToken) *TokenProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token profile o k response
func (o *TokenProfileOK) SetPayload(payload *models.AuthToken) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TokenProfileForbiddenCode is the HTTP code returned for type TokenProfileForbidden
const TokenProfileForbiddenCode int = 403

/*TokenProfileForbidden User is not authorized

swagger:response tokenProfileForbidden
*/
type TokenProfileForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTokenProfileForbidden creates TokenProfileForbidden with default headers values
func NewTokenProfileForbidden() *TokenProfileForbidden {

	return &TokenProfileForbidden{}
}

// WithPayload adds the payload to the token profile forbidden response
func (o *TokenProfileForbidden) WithPayload(payload *models.GeneralError) *TokenProfileForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token profile forbidden response
func (o *TokenProfileForbidden) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenProfileForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TokenProfileInternalServerErrorCode is the HTTP code returned for type TokenProfileInternalServerError
const TokenProfileInternalServerErrorCode int = 500

/*TokenProfileInternalServerError Failed to generate a token

swagger:response tokenProfileInternalServerError
*/
type TokenProfileInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTokenProfileInternalServerError creates TokenProfileInternalServerError with default headers values
func NewTokenProfileInternalServerError() *TokenProfileInternalServerError {

	return &TokenProfileInternalServerError{}
}

// WithPayload adds the payload to the token profile internal server error response
func (o *TokenProfileInternalServerError) WithPayload(payload *models.GeneralError) *TokenProfileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token profile internal server error response
func (o *TokenProfileInternalServerError) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenProfileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TokenProfileDefault Some error unrelated to the handler

swagger:response tokenProfileDefault
*/
type TokenProfileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewTokenProfileDefault creates TokenProfileDefault with default headers values
func NewTokenProfileDefault(code int) *TokenProfileDefault {
	if code <= 0 {
		code = 500
	}

	return &TokenProfileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the token profile default response
func (o *TokenProfileDefault) WithStatusCode(code int) *TokenProfileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the token profile default response
func (o *TokenProfileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the token profile default response
func (o *TokenProfileDefault) WithPayload(payload *models.GeneralError) *TokenProfileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token profile default response
func (o *TokenProfileDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenProfileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
