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

// DcimPowerOutletsBulkPartialUpdateReader is a Reader for the DcimPowerOutletsBulkPartialUpdate structure.
type DcimPowerOutletsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/power-outlets/] dcim_power-outlets_bulk_partial_update", response, response.Code())
	}
}

// NewDcimPowerOutletsBulkPartialUpdateOK creates a DcimPowerOutletsBulkPartialUpdateOK with default headers values
func NewDcimPowerOutletsBulkPartialUpdateOK() *DcimPowerOutletsBulkPartialUpdateOK {
	return &DcimPowerOutletsBulkPartialUpdateOK{}
}

/*
DcimPowerOutletsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsBulkPartialUpdateOK dcim power outlets bulk partial update o k
*/
type DcimPowerOutletsBulkPartialUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets bulk partial update o k response has a 2xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets bulk partial update o k response has a 3xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets bulk partial update o k response has a 4xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets bulk partial update o k response has a 5xx status code
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets bulk partial update o k response a status code equal to that given
func (o *DcimPowerOutletsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets bulk partial update o k response
func (o *DcimPowerOutletsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcimPowerOutletsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcimPowerOutletsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
