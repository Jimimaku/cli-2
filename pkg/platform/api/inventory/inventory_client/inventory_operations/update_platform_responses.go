// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// UpdatePlatformReader is a Reader for the UpdatePlatform structure.
type UpdatePlatformReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlatformReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePlatformOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePlatformBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdatePlatformDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePlatformOK creates a UpdatePlatformOK with default headers values
func NewUpdatePlatformOK() *UpdatePlatformOK {
	return &UpdatePlatformOK{}
}

/* UpdatePlatformOK describes a response with status code 200, with default header values.

The updated state of the platform
*/
type UpdatePlatformOK struct {
	Payload *inventory_models.Platform
}

func (o *UpdatePlatformOK) Error() string {
	return fmt.Sprintf("[PUT /v1/platforms/{platform_id}][%d] updatePlatformOK  %+v", 200, o.Payload)
}
func (o *UpdatePlatformOK) GetPayload() *inventory_models.Platform {
	return o.Payload
}

func (o *UpdatePlatformOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlatformBadRequest creates a UpdatePlatformBadRequest with default headers values
func NewUpdatePlatformBadRequest() *UpdatePlatformBadRequest {
	return &UpdatePlatformBadRequest{}
}

/* UpdatePlatformBadRequest describes a response with status code 400, with default header values.

If the platform update in invalid
*/
type UpdatePlatformBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *UpdatePlatformBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/platforms/{platform_id}][%d] updatePlatformBadRequest  %+v", 400, o.Payload)
}
func (o *UpdatePlatformBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *UpdatePlatformBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlatformDefault creates a UpdatePlatformDefault with default headers values
func NewUpdatePlatformDefault(code int) *UpdatePlatformDefault {
	return &UpdatePlatformDefault{
		_statusCode: code,
	}
}

/* UpdatePlatformDefault describes a response with status code -1, with default header values.

If there is an error processing the request
*/
type UpdatePlatformDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the update platform default response
func (o *UpdatePlatformDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePlatformDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/platforms/{platform_id}][%d] updatePlatform default  %+v", o._statusCode, o.Payload)
}
func (o *UpdatePlatformDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *UpdatePlatformDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
