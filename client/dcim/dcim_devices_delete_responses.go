// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDevicesDeleteReader is a Reader for the DcimDevicesDelete structure.
type DcimDevicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDevicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/devices/{id}/] dcim_devices_delete", response, response.Code())
	}
}

// NewDcimDevicesDeleteNoContent creates a DcimDevicesDeleteNoContent with default headers values
func NewDcimDevicesDeleteNoContent() *DcimDevicesDeleteNoContent {
	return &DcimDevicesDeleteNoContent{}
}

/*
DcimDevicesDeleteNoContent describes a response with status code 204, with default header values.

DcimDevicesDeleteNoContent dcim devices delete no content
*/
type DcimDevicesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim devices delete no content response has a 2xx status code
func (o *DcimDevicesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices delete no content response has a 3xx status code
func (o *DcimDevicesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices delete no content response has a 4xx status code
func (o *DcimDevicesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices delete no content response has a 5xx status code
func (o *DcimDevicesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices delete no content response a status code equal to that given
func (o *DcimDevicesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim devices delete no content response
func (o *DcimDevicesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimDevicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcimDevicesDeleteNoContent", 204)
}

func (o *DcimDevicesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcimDevicesDeleteNoContent", 204)
}

func (o *DcimDevicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}