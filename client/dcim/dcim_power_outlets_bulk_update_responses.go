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

// DcimPowerOutletsBulkUpdateReader is a Reader for the DcimPowerOutletsBulkUpdate structure.
type DcimPowerOutletsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/power-outlets/] dcim_power-outlets_bulk_update", response, response.Code())
	}
}

// NewDcimPowerOutletsBulkUpdateOK creates a DcimPowerOutletsBulkUpdateOK with default headers values
func NewDcimPowerOutletsBulkUpdateOK() *DcimPowerOutletsBulkUpdateOK {
	return &DcimPowerOutletsBulkUpdateOK{}
}

/*
DcimPowerOutletsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsBulkUpdateOK dcim power outlets bulk update o k
*/
type DcimPowerOutletsBulkUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets bulk update o k response has a 2xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets bulk update o k response has a 3xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets bulk update o k response has a 4xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets bulk update o k response has a 5xx status code
func (o *DcimPowerOutletsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets bulk update o k response a status code equal to that given
func (o *DcimPowerOutletsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets bulk update o k response
func (o *DcimPowerOutletsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcimPowerOutletsBulkUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcimPowerOutletsBulkUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletsBulkUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
