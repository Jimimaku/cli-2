// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetLogoutReader is a Reader for the GetLogout structure.
type GetLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetLogoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogoutNoContent creates a GetLogoutNoContent with default headers values
func NewGetLogoutNoContent() *GetLogoutNoContent {
	return &GetLogoutNoContent{}
}

/* GetLogoutNoContent describes a response with status code 204, with default header values.

Success
*/
type GetLogoutNoContent struct {
}

func (o *GetLogoutNoContent) Error() string {
	return fmt.Sprintf("[GET /logout][%d] getLogoutNoContent ", 204)
}

func (o *GetLogoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogoutInternalServerError creates a GetLogoutInternalServerError with default headers values
func NewGetLogoutInternalServerError() *GetLogoutInternalServerError {
	return &GetLogoutInternalServerError{}
}

/* GetLogoutInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetLogoutInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetLogoutInternalServerError) Error() string {
	return fmt.Sprintf("[GET /logout][%d] getLogoutInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLogoutInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetLogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
