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

	"github.com/fbreckle/go-netbox/models"
)

// DcimDeviceTypesBulkPartialUpdateReader is a Reader for the DcimDeviceTypesBulkPartialUpdate structure.
type DcimDeviceTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/device-types/] dcim_device-types_bulk_partial_update", response, response.Code())
	}
}

// NewDcimDeviceTypesBulkPartialUpdateOK creates a DcimDeviceTypesBulkPartialUpdateOK with default headers values
func NewDcimDeviceTypesBulkPartialUpdateOK() *DcimDeviceTypesBulkPartialUpdateOK {
	return &DcimDeviceTypesBulkPartialUpdateOK{}
}

/*
DcimDeviceTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesBulkPartialUpdateOK dcim device types bulk partial update o k
*/
type DcimDeviceTypesBulkPartialUpdateOK struct {
	Payload *models.DeviceType
}

// IsSuccess returns true when this dcim device types bulk partial update o k response has a 2xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types bulk partial update o k response has a 3xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types bulk partial update o k response has a 4xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types bulk partial update o k response has a 5xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types bulk partial update o k response a status code equal to that given
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device types bulk partial update o k response
func (o *DcimDeviceTypesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
