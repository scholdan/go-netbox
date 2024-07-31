// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasTagsDeleteReader is a Reader for the ExtrasTagsDelete structure.
type ExtrasTagsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasTagsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /extras/tags/{id}/] extras_tags_delete", response, response.Code())
	}
}

// NewExtrasTagsDeleteNoContent creates a ExtrasTagsDeleteNoContent with default headers values
func NewExtrasTagsDeleteNoContent() *ExtrasTagsDeleteNoContent {
	return &ExtrasTagsDeleteNoContent{}
}

/*
ExtrasTagsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasTagsDeleteNoContent extras tags delete no content
*/
type ExtrasTagsDeleteNoContent struct {
}

// IsSuccess returns true when this extras tags delete no content response has a 2xx status code
func (o *ExtrasTagsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras tags delete no content response has a 3xx status code
func (o *ExtrasTagsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags delete no content response has a 4xx status code
func (o *ExtrasTagsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras tags delete no content response has a 5xx status code
func (o *ExtrasTagsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags delete no content response a status code equal to that given
func (o *ExtrasTagsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the extras tags delete no content response
func (o *ExtrasTagsDeleteNoContent) Code() int {
	return 204
}

func (o *ExtrasTagsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/tags/{id}/][%d] extrasTagsDeleteNoContent", 204)
}

func (o *ExtrasTagsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/tags/{id}/][%d] extrasTagsDeleteNoContent", 204)
}

func (o *ExtrasTagsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
