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

// DcimDeviceBaysBulkPartialUpdateReader is a Reader for the DcimDeviceBaysBulkPartialUpdate structure.
type DcimDeviceBaysBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/device-bays/] dcim_device-bays_bulk_partial_update", response, response.Code())
	}
}

// NewDcimDeviceBaysBulkPartialUpdateOK creates a DcimDeviceBaysBulkPartialUpdateOK with default headers values
func NewDcimDeviceBaysBulkPartialUpdateOK() *DcimDeviceBaysBulkPartialUpdateOK {
	return &DcimDeviceBaysBulkPartialUpdateOK{}
}

/*
DcimDeviceBaysBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysBulkPartialUpdateOK dcim device bays bulk partial update o k
*/
type DcimDeviceBaysBulkPartialUpdateOK struct {
	Payload *models.DeviceBay
}

// IsSuccess returns true when this dcim device bays bulk partial update o k response has a 2xx status code
func (o *DcimDeviceBaysBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bays bulk partial update o k response has a 3xx status code
func (o *DcimDeviceBaysBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bays bulk partial update o k response has a 4xx status code
func (o *DcimDeviceBaysBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bays bulk partial update o k response has a 5xx status code
func (o *DcimDeviceBaysBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bays bulk partial update o k response a status code equal to that given
func (o *DcimDeviceBaysBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device bays bulk partial update o k response
func (o *DcimDeviceBaysBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceBaysBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-bays/][%d] dcimDeviceBaysBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceBaysBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-bays/][%d] dcimDeviceBaysBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceBaysBulkPartialUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}