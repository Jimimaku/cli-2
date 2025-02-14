// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LinkGithubAccountReader is a Reader for the LinkGithubAccount structure.
type LinkGithubAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LinkGithubAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewLinkGithubAccountFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLinkGithubAccountFound creates a LinkGithubAccountFound with default headers values
func NewLinkGithubAccountFound() *LinkGithubAccountFound {
	return &LinkGithubAccountFound{}
}

/* LinkGithubAccountFound describes a response with status code 302, with default header values.

Found
*/
type LinkGithubAccountFound struct {
}

func (o *LinkGithubAccountFound) Error() string {
	return fmt.Sprintf("[GET /oauth/github/link][%d] linkGithubAccountFound ", 302)
}

func (o *LinkGithubAccountFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
