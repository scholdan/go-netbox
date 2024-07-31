// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasConfigContextsDeleteReader is a Reader for the ExtrasConfigContextsDelete structure.
type ExtrasConfigContextsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasConfigContextsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /extras/config-contexts/{id}/] extras_config-contexts_delete", response, response.Code())
	}
}

// NewExtrasConfigContextsDeleteNoContent creates a ExtrasConfigContextsDeleteNoContent with default headers values
func NewExtrasConfigContextsDeleteNoContent() *ExtrasConfigContextsDeleteNoContent {
	return &ExtrasConfigContextsDeleteNoContent{}
}

/*
ExtrasConfigContextsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasConfigContextsDeleteNoContent extras config contexts delete no content
*/
type ExtrasConfigContextsDeleteNoContent struct {
}

// IsSuccess returns true when this extras config contexts delete no content response has a 2xx status code
func (o *ExtrasConfigContextsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts delete no content response has a 3xx status code
func (o *ExtrasConfigContextsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts delete no content response has a 4xx status code
func (o *ExtrasConfigContextsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts delete no content response has a 5xx status code
func (o *ExtrasConfigContextsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts delete no content response a status code equal to that given
func (o *ExtrasConfigContextsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the extras config contexts delete no content response
func (o *ExtrasConfigContextsDeleteNoContent) Code() int {
	return 204
}

func (o *ExtrasConfigContextsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/{id}/][%d] extrasConfigContextsDeleteNoContent", 204)
}

func (o *ExtrasConfigContextsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/{id}/][%d] extrasConfigContextsDeleteNoContent", 204)
}

func (o *ExtrasConfigContextsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
