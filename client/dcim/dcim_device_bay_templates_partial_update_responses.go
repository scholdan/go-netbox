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

// DcimDeviceBayTemplatesPartialUpdateReader is a Reader for the DcimDeviceBayTemplatesPartialUpdate structure.
type DcimDeviceBayTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/device-bay-templates/{id}/] dcim_device-bay-templates_partial_update", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateOK creates a DcimDeviceBayTemplatesPartialUpdateOK with default headers values
func NewDcimDeviceBayTemplatesPartialUpdateOK() *DcimDeviceBayTemplatesPartialUpdateOK {
	return &DcimDeviceBayTemplatesPartialUpdateOK{}
}

/*
DcimDeviceBayTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesPartialUpdateOK dcim device bay templates partial update o k
*/
type DcimDeviceBayTemplatesPartialUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates partial update o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates partial update o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates partial update o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates partial update o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates partial update o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device bay templates partial update o k response
func (o *DcimDeviceBayTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
