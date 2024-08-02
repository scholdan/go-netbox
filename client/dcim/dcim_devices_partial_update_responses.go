// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scholdan/go-netbox/models"
)

// DcimDevicesPartialUpdateReader is a Reader for the DcimDevicesPartialUpdate structure.
type DcimDevicesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/devices/{id}/] dcim_devices_partial_update", response, response.Code())
	}
}

// NewDcimDevicesPartialUpdateOK creates a DcimDevicesPartialUpdateOK with default headers values
func NewDcimDevicesPartialUpdateOK() *DcimDevicesPartialUpdateOK {
	return &DcimDevicesPartialUpdateOK{}
}

/*
DcimDevicesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDevicesPartialUpdateOK dcim devices partial update o k
*/
type DcimDevicesPartialUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

// IsSuccess returns true when this dcim devices partial update o k response has a 2xx status code
func (o *DcimDevicesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices partial update o k response has a 3xx status code
func (o *DcimDevicesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices partial update o k response has a 4xx status code
func (o *DcimDevicesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices partial update o k response has a 5xx status code
func (o *DcimDevicesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices partial update o k response a status code equal to that given
func (o *DcimDevicesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim devices partial update o k response
func (o *DcimDevicesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimDevicesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/devices/{id}/][%d] dcimDevicesPartialUpdateOK %s", 200, payload)
}

func (o *DcimDevicesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/devices/{id}/][%d] dcimDevicesPartialUpdateOK %s", 200, payload)
}

func (o *DcimDevicesPartialUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}