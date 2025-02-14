// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// SendEmailVerificationReader is a Reader for the SendEmailVerification structure.
type SendEmailVerificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendEmailVerificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendEmailVerificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendEmailVerificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSendEmailVerificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSendEmailVerificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSendEmailVerificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSendEmailVerificationOK creates a SendEmailVerificationOK with default headers values
func NewSendEmailVerificationOK() *SendEmailVerificationOK {
	return &SendEmailVerificationOK{}
}

/* SendEmailVerificationOK describes a response with status code 200, with default header values.

Email updated
*/
type SendEmailVerificationOK struct {
	Payload *mono_models.Message
}

func (o *SendEmailVerificationOK) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/send][%d] sendEmailVerificationOK  %+v", 200, o.Payload)
}
func (o *SendEmailVerificationOK) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SendEmailVerificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEmailVerificationBadRequest creates a SendEmailVerificationBadRequest with default headers values
func NewSendEmailVerificationBadRequest() *SendEmailVerificationBadRequest {
	return &SendEmailVerificationBadRequest{}
}

/* SendEmailVerificationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SendEmailVerificationBadRequest struct {
	Payload *mono_models.Message
}

func (o *SendEmailVerificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/send][%d] sendEmailVerificationBadRequest  %+v", 400, o.Payload)
}
func (o *SendEmailVerificationBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SendEmailVerificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEmailVerificationForbidden creates a SendEmailVerificationForbidden with default headers values
func NewSendEmailVerificationForbidden() *SendEmailVerificationForbidden {
	return &SendEmailVerificationForbidden{}
}

/* SendEmailVerificationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SendEmailVerificationForbidden struct {
	Payload *mono_models.Message
}

func (o *SendEmailVerificationForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/send][%d] sendEmailVerificationForbidden  %+v", 403, o.Payload)
}
func (o *SendEmailVerificationForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SendEmailVerificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEmailVerificationNotFound creates a SendEmailVerificationNotFound with default headers values
func NewSendEmailVerificationNotFound() *SendEmailVerificationNotFound {
	return &SendEmailVerificationNotFound{}
}

/* SendEmailVerificationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SendEmailVerificationNotFound struct {
	Payload *mono_models.Message
}

func (o *SendEmailVerificationNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/send][%d] sendEmailVerificationNotFound  %+v", 404, o.Payload)
}
func (o *SendEmailVerificationNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SendEmailVerificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEmailVerificationInternalServerError creates a SendEmailVerificationInternalServerError with default headers values
func NewSendEmailVerificationInternalServerError() *SendEmailVerificationInternalServerError {
	return &SendEmailVerificationInternalServerError{}
}

/* SendEmailVerificationInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SendEmailVerificationInternalServerError struct {
	Payload *mono_models.Message
}

func (o *SendEmailVerificationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/send][%d] sendEmailVerificationInternalServerError  %+v", 500, o.Payload)
}
func (o *SendEmailVerificationInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SendEmailVerificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
