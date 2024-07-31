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

// DcimPlatformsBulkUpdateReader is a Reader for the DcimPlatformsBulkUpdate structure.
type DcimPlatformsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/platforms/] dcim_platforms_bulk_update", response, response.Code())
	}
}

// NewDcimPlatformsBulkUpdateOK creates a DcimPlatformsBulkUpdateOK with default headers values
func NewDcimPlatformsBulkUpdateOK() *DcimPlatformsBulkUpdateOK {
	return &DcimPlatformsBulkUpdateOK{}
}

/*
DcimPlatformsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsBulkUpdateOK dcim platforms bulk update o k
*/
type DcimPlatformsBulkUpdateOK struct {
	Payload *models.Platform
}

// IsSuccess returns true when this dcim platforms bulk update o k response has a 2xx status code
func (o *DcimPlatformsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms bulk update o k response has a 3xx status code
func (o *DcimPlatformsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms bulk update o k response has a 4xx status code
func (o *DcimPlatformsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms bulk update o k response has a 5xx status code
func (o *DcimPlatformsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms bulk update o k response a status code equal to that given
func (o *DcimPlatformsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim platforms bulk update o k response
func (o *DcimPlatformsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimPlatformsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcimPlatformsBulkUpdateOK %s", 200, payload)
}

func (o *DcimPlatformsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/platforms/][%d] dcimPlatformsBulkUpdateOK %s", 200, payload)
}

func (o *DcimPlatformsBulkUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
