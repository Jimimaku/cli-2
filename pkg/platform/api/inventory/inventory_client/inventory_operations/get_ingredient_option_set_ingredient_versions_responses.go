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

// GetIngredientOptionSetIngredientVersionsReader is a Reader for the GetIngredientOptionSetIngredientVersions structure.
type GetIngredientOptionSetIngredientVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIngredientOptionSetIngredientVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIngredientOptionSetIngredientVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIngredientOptionSetIngredientVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIngredientOptionSetIngredientVersionsOK creates a GetIngredientOptionSetIngredientVersionsOK with default headers values
func NewGetIngredientOptionSetIngredientVersionsOK() *GetIngredientOptionSetIngredientVersionsOK {
	return &GetIngredientOptionSetIngredientVersionsOK{}
}

/* GetIngredientOptionSetIngredientVersionsOK describes a response with status code 200, with default header values.

A paginated list of ingredient versions
*/
type GetIngredientOptionSetIngredientVersionsOK struct {
	Payload *inventory_models.IngredientVersionAndUsageTypePagedList
}

func (o *GetIngredientOptionSetIngredientVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/ingredient-option-sets/{ingredient_option_set_id}/ingredient-versions][%d] getIngredientOptionSetIngredientVersionsOK  %+v", 200, o.Payload)
}
func (o *GetIngredientOptionSetIngredientVersionsOK) GetPayload() *inventory_models.IngredientVersionAndUsageTypePagedList {
	return o.Payload
}

func (o *GetIngredientOptionSetIngredientVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientVersionAndUsageTypePagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIngredientOptionSetIngredientVersionsDefault creates a GetIngredientOptionSetIngredientVersionsDefault with default headers values
func NewGetIngredientOptionSetIngredientVersionsDefault(code int) *GetIngredientOptionSetIngredientVersionsDefault {
	return &GetIngredientOptionSetIngredientVersionsDefault{
		_statusCode: code,
	}
}

/* GetIngredientOptionSetIngredientVersionsDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetIngredientOptionSetIngredientVersionsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get ingredient option set ingredient versions default response
func (o *GetIngredientOptionSetIngredientVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetIngredientOptionSetIngredientVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ingredient-option-sets/{ingredient_option_set_id}/ingredient-versions][%d] getIngredientOptionSetIngredientVersions default  %+v", o._statusCode, o.Payload)
}
func (o *GetIngredientOptionSetIngredientVersionsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetIngredientOptionSetIngredientVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
