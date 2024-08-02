// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimCableTerminationsDeleteReader is a Reader for the DcimCableTerminationsDelete structure.
type DcimCableTerminationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCableTerminationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/cable-terminations/{id}/] dcim_cable-terminations_delete", response, response.Code())
	}
}

// NewDcimCableTerminationsDeleteNoContent creates a DcimCableTerminationsDeleteNoContent with default headers values
func NewDcimCableTerminationsDeleteNoContent() *DcimCableTerminationsDeleteNoContent {
	return &DcimCableTerminationsDeleteNoContent{}
}

/*
DcimCableTerminationsDeleteNoContent describes a response with status code 204, with default header values.

DcimCableTerminationsDeleteNoContent dcim cable terminations delete no content
*/
type DcimCableTerminationsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim cable terminations delete no content response has a 2xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations delete no content response has a 3xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations delete no content response has a 4xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations delete no content response has a 5xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations delete no content response a status code equal to that given
func (o *DcimCableTerminationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim cable terminations delete no content response
func (o *DcimCableTerminationsDeleteNoContent) Code() int {
	return 204
}

func (o *DcimCableTerminationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteNoContent", 204)
}

func (o *DcimCableTerminationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteNoContent", 204)
}

func (o *DcimCableTerminationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}