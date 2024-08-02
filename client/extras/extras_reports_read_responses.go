// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasReportsReadReader is a Reader for the ExtrasReportsRead structure.
type ExtrasReportsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasReportsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasReportsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/reports/{id}/] extras_reports_read", response, response.Code())
	}
}

// NewExtrasReportsReadOK creates a ExtrasReportsReadOK with default headers values
func NewExtrasReportsReadOK() *ExtrasReportsReadOK {
	return &ExtrasReportsReadOK{}
}

/*
ExtrasReportsReadOK describes a response with status code 200, with default header values.

ExtrasReportsReadOK extras reports read o k
*/
type ExtrasReportsReadOK struct {
}

// IsSuccess returns true when this extras reports read o k response has a 2xx status code
func (o *ExtrasReportsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras reports read o k response has a 3xx status code
func (o *ExtrasReportsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras reports read o k response has a 4xx status code
func (o *ExtrasReportsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras reports read o k response has a 5xx status code
func (o *ExtrasReportsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras reports read o k response a status code equal to that given
func (o *ExtrasReportsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras reports read o k response
func (o *ExtrasReportsReadOK) Code() int {
	return 200
}

func (o *ExtrasReportsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadOK", 200)
}

func (o *ExtrasReportsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadOK", 200)
}

func (o *ExtrasReportsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}