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

// DcimPowerPanelsUpdateReader is a Reader for the DcimPowerPanelsUpdate structure.
type DcimPowerPanelsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/power-panels/{id}/] dcim_power-panels_update", response, response.Code())
	}
}

// NewDcimPowerPanelsUpdateOK creates a DcimPowerPanelsUpdateOK with default headers values
func NewDcimPowerPanelsUpdateOK() *DcimPowerPanelsUpdateOK {
	return &DcimPowerPanelsUpdateOK{}
}

/*
DcimPowerPanelsUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsUpdateOK dcim power panels update o k
*/
type DcimPowerPanelsUpdateOK struct {
	Payload *models.PowerPanel
}

// IsSuccess returns true when this dcim power panels update o k response has a 2xx status code
func (o *DcimPowerPanelsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels update o k response has a 3xx status code
func (o *DcimPowerPanelsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels update o k response has a 4xx status code
func (o *DcimPowerPanelsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels update o k response has a 5xx status code
func (o *DcimPowerPanelsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels update o k response a status code equal to that given
func (o *DcimPowerPanelsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power panels update o k response
func (o *DcimPowerPanelsUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerPanelsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-panels/{id}/][%d] dcimPowerPanelsUpdateOK %s", 200, payload)
}

func (o *DcimPowerPanelsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-panels/{id}/][%d] dcimPowerPanelsUpdateOK %s", 200, payload)
}

func (o *DcimPowerPanelsUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}