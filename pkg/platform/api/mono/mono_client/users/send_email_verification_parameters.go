// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSendEmailVerificationParams creates a new SendEmailVerificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSendEmailVerificationParams() *SendEmailVerificationParams {
	return &SendEmailVerificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSendEmailVerificationParamsWithTimeout creates a new SendEmailVerificationParams object
// with the ability to set a timeout on a request.
func NewSendEmailVerificationParamsWithTimeout(timeout time.Duration) *SendEmailVerificationParams {
	return &SendEmailVerificationParams{
		timeout: timeout,
	}
}

// NewSendEmailVerificationParamsWithContext creates a new SendEmailVerificationParams object
// with the ability to set a context for a request.
func NewSendEmailVerificationParamsWithContext(ctx context.Context) *SendEmailVerificationParams {
	return &SendEmailVerificationParams{
		Context: ctx,
	}
}

// NewSendEmailVerificationParamsWithHTTPClient creates a new SendEmailVerificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewSendEmailVerificationParamsWithHTTPClient(client *http.Client) *SendEmailVerificationParams {
	return &SendEmailVerificationParams{
		HTTPClient: client,
	}
}

/* SendEmailVerificationParams contains all the parameters to send to the API endpoint
   for the send email verification operation.

   Typically these are written to a http.Request.
*/
type SendEmailVerificationParams struct {

	/* Email.

	   email address to change
	*/
	Email string

	/* Username.

	   username of desired User
	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the send email verification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendEmailVerificationParams) WithDefaults() *SendEmailVerificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the send email verification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendEmailVerificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the send email verification params
func (o *SendEmailVerificationParams) WithTimeout(timeout time.Duration) *SendEmailVerificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send email verification params
func (o *SendEmailVerificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send email verification params
func (o *SendEmailVerificationParams) WithContext(ctx context.Context) *SendEmailVerificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send email verification params
func (o *SendEmailVerificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send email verification params
func (o *SendEmailVerificationParams) WithHTTPClient(client *http.Client) *SendEmailVerificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send email verification params
func (o *SendEmailVerificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the send email verification params
func (o *SendEmailVerificationParams) WithEmail(email string) *SendEmailVerificationParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the send email verification params
func (o *SendEmailVerificationParams) SetEmail(email string) {
	o.Email = email
}

// WithUsername adds the username to the send email verification params
func (o *SendEmailVerificationParams) WithUsername(username string) *SendEmailVerificationParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the send email verification params
func (o *SendEmailVerificationParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *SendEmailVerificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
