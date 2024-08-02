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

// DcimPowerOutletsPartialUpdateReader is a Reader for the DcimPowerOutletsPartialUpdate structure.
type DcimPowerOutletsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/power-outlets/{id}/] dcim_power-outlets_partial_update", response, response.Code())
	}
}

// NewDcimPowerOutletsPartialUpdateOK creates a DcimPowerOutletsPartialUpdateOK with default headers values
func NewDcimPowerOutletsPartialUpdateOK() *DcimPowerOutletsPartialUpdateOK {
	return &DcimPowerOutletsPartialUpdateOK{}
}

/*
DcimPowerOutletsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsPartialUpdateOK dcim power outlets partial update o k
*/
type DcimPowerOutletsPartialUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets partial update o k response has a 2xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets partial update o k response has a 3xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets partial update o k response has a 4xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets partial update o k response has a 5xx status code
func (o *DcimPowerOutletsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets partial update o k response a status code equal to that given
func (o *DcimPowerOutletsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets partial update o k response
func (o *DcimPowerOutletsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcimPowerOutletsPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcimPowerOutletsPartialUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletsPartialUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
