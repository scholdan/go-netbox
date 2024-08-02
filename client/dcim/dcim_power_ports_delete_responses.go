// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPortsDeleteReader is a Reader for the DcimPowerPortsDelete structure.
type DcimPowerPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/power-ports/{id}/] dcim_power-ports_delete", response, response.Code())
	}
}

// NewDcimPowerPortsDeleteNoContent creates a DcimPowerPortsDeleteNoContent with default headers values
func NewDcimPowerPortsDeleteNoContent() *DcimPowerPortsDeleteNoContent {
	return &DcimPowerPortsDeleteNoContent{}
}

/*
DcimPowerPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortsDeleteNoContent dcim power ports delete no content
*/
type DcimPowerPortsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power ports delete no content response has a 2xx status code
func (o *DcimPowerPortsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports delete no content response has a 3xx status code
func (o *DcimPowerPortsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports delete no content response has a 4xx status code
func (o *DcimPowerPortsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports delete no content response has a 5xx status code
func (o *DcimPowerPortsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports delete no content response a status code equal to that given
func (o *DcimPowerPortsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim power ports delete no content response
func (o *DcimPowerPortsDeleteNoContent) Code() int {
	return 204
}

func (o *DcimPowerPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcimPowerPortsDeleteNoContent", 204)
}

func (o *DcimPowerPortsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcimPowerPortsDeleteNoContent", 204)
}

func (o *DcimPowerPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}