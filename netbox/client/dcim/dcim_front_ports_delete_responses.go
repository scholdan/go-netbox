// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimFrontPortsDeleteReader is a Reader for the DcimFrontPortsDelete structure.
type DcimFrontPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimFrontPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/front-ports/{id}/] dcim_front-ports_delete", response, response.Code())
	}
}

// NewDcimFrontPortsDeleteNoContent creates a DcimFrontPortsDeleteNoContent with default headers values
func NewDcimFrontPortsDeleteNoContent() *DcimFrontPortsDeleteNoContent {
	return &DcimFrontPortsDeleteNoContent{}
}

/*
DcimFrontPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimFrontPortsDeleteNoContent dcim front ports delete no content
*/
type DcimFrontPortsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim front ports delete no content response has a 2xx status code
func (o *DcimFrontPortsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports delete no content response has a 3xx status code
func (o *DcimFrontPortsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports delete no content response has a 4xx status code
func (o *DcimFrontPortsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports delete no content response has a 5xx status code
func (o *DcimFrontPortsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports delete no content response a status code equal to that given
func (o *DcimFrontPortsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim front ports delete no content response
func (o *DcimFrontPortsDeleteNoContent) Code() int {
	return 204
}

func (o *DcimFrontPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/front-ports/{id}/][%d] dcimFrontPortsDeleteNoContent", 204)
}

func (o *DcimFrontPortsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/front-ports/{id}/][%d] dcimFrontPortsDeleteNoContent", 204)
}

func (o *DcimFrontPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
