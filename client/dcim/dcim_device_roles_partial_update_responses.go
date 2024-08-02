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

// DcimDeviceRolesPartialUpdateReader is a Reader for the DcimDeviceRolesPartialUpdate structure.
type DcimDeviceRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/device-roles/{id}/] dcim_device-roles_partial_update", response, response.Code())
	}
}

// NewDcimDeviceRolesPartialUpdateOK creates a DcimDeviceRolesPartialUpdateOK with default headers values
func NewDcimDeviceRolesPartialUpdateOK() *DcimDeviceRolesPartialUpdateOK {
	return &DcimDeviceRolesPartialUpdateOK{}
}

/*
DcimDeviceRolesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesPartialUpdateOK dcim device roles partial update o k
*/
type DcimDeviceRolesPartialUpdateOK struct {
	Payload *models.DeviceRole
}

// IsSuccess returns true when this dcim device roles partial update o k response has a 2xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles partial update o k response has a 3xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles partial update o k response has a 4xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles partial update o k response has a 5xx status code
func (o *DcimDeviceRolesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles partial update o k response a status code equal to that given
func (o *DcimDeviceRolesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device roles partial update o k response
func (o *DcimDeviceRolesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceRolesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcimDeviceRolesPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceRolesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-roles/{id}/][%d] dcimDeviceRolesPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceRolesPartialUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}