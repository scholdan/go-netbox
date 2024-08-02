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

// DcimDeviceBayTemplatesCreateReader is a Reader for the DcimDeviceBayTemplatesCreate structure.
type DcimDeviceBayTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDeviceBayTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/device-bay-templates/] dcim_device-bay-templates_create", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesCreateCreated creates a DcimDeviceBayTemplatesCreateCreated with default headers values
func NewDcimDeviceBayTemplatesCreateCreated() *DcimDeviceBayTemplatesCreateCreated {
	return &DcimDeviceBayTemplatesCreateCreated{}
}

/*
DcimDeviceBayTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimDeviceBayTemplatesCreateCreated dcim device bay templates create created
*/
type DcimDeviceBayTemplatesCreateCreated struct {
	Payload *models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates create created response has a 2xx status code
func (o *DcimDeviceBayTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates create created response has a 3xx status code
func (o *DcimDeviceBayTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates create created response has a 4xx status code
func (o *DcimDeviceBayTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates create created response has a 5xx status code
func (o *DcimDeviceBayTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates create created response a status code equal to that given
func (o *DcimDeviceBayTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim device bay templates create created response
func (o *DcimDeviceBayTemplatesCreateCreated) Code() int {
	return 201
}

func (o *DcimDeviceBayTemplatesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimDeviceBayTemplatesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimDeviceBayTemplatesCreateCreated) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}