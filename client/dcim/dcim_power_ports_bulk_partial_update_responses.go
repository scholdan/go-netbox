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

// DcimPowerPortsBulkPartialUpdateReader is a Reader for the DcimPowerPortsBulkPartialUpdate structure.
type DcimPowerPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/power-ports/] dcim_power-ports_bulk_partial_update", response, response.Code())
	}
}

// NewDcimPowerPortsBulkPartialUpdateOK creates a DcimPowerPortsBulkPartialUpdateOK with default headers values
func NewDcimPowerPortsBulkPartialUpdateOK() *DcimPowerPortsBulkPartialUpdateOK {
	return &DcimPowerPortsBulkPartialUpdateOK{}
}

/*
DcimPowerPortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortsBulkPartialUpdateOK dcim power ports bulk partial update o k
*/
type DcimPowerPortsBulkPartialUpdateOK struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports bulk partial update o k response has a 2xx status code
func (o *DcimPowerPortsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports bulk partial update o k response has a 3xx status code
func (o *DcimPowerPortsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports bulk partial update o k response has a 4xx status code
func (o *DcimPowerPortsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports bulk partial update o k response has a 5xx status code
func (o *DcimPowerPortsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports bulk partial update o k response a status code equal to that given
func (o *DcimPowerPortsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power ports bulk partial update o k response
func (o *DcimPowerPortsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPortsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-ports/][%d] dcimPowerPortsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerPortsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-ports/][%d] dcimPowerPortsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerPortsBulkPartialUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
