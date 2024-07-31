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

// DcimDeviceBaysUpdateReader is a Reader for the DcimDeviceBaysUpdate structure.
type DcimDeviceBaysUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/device-bays/{id}/] dcim_device-bays_update", response, response.Code())
	}
}

// NewDcimDeviceBaysUpdateOK creates a DcimDeviceBaysUpdateOK with default headers values
func NewDcimDeviceBaysUpdateOK() *DcimDeviceBaysUpdateOK {
	return &DcimDeviceBaysUpdateOK{}
}

/*
DcimDeviceBaysUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysUpdateOK dcim device bays update o k
*/
type DcimDeviceBaysUpdateOK struct {
	Payload *models.DeviceBay
}

// IsSuccess returns true when this dcim device bays update o k response has a 2xx status code
func (o *DcimDeviceBaysUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bays update o k response has a 3xx status code
func (o *DcimDeviceBaysUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bays update o k response has a 4xx status code
func (o *DcimDeviceBaysUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bays update o k response has a 5xx status code
func (o *DcimDeviceBaysUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bays update o k response a status code equal to that given
func (o *DcimDeviceBaysUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device bays update o k response
func (o *DcimDeviceBaysUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceBaysUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/device-bays/{id}/][%d] dcimDeviceBaysUpdateOK %s", 200, payload)
}

func (o *DcimDeviceBaysUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/device-bays/{id}/][%d] dcimDeviceBaysUpdateOK %s", 200, payload)
}

func (o *DcimDeviceBaysUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
