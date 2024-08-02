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

// DcimRearPortsUpdateReader is a Reader for the DcimRearPortsUpdate structure.
type DcimRearPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/rear-ports/{id}/] dcim_rear-ports_update", response, response.Code())
	}
}

// NewDcimRearPortsUpdateOK creates a DcimRearPortsUpdateOK with default headers values
func NewDcimRearPortsUpdateOK() *DcimRearPortsUpdateOK {
	return &DcimRearPortsUpdateOK{}
}

/*
DcimRearPortsUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsUpdateOK dcim rear ports update o k
*/
type DcimRearPortsUpdateOK struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports update o k response has a 2xx status code
func (o *DcimRearPortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports update o k response has a 3xx status code
func (o *DcimRearPortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports update o k response has a 4xx status code
func (o *DcimRearPortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports update o k response has a 5xx status code
func (o *DcimRearPortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports update o k response a status code equal to that given
func (o *DcimRearPortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rear ports update o k response
func (o *DcimRearPortsUpdateOK) Code() int {
	return 200
}

func (o *DcimRearPortsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcimRearPortsUpdateOK %s", 200, payload)
}

func (o *DcimRearPortsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/rear-ports/{id}/][%d] dcimRearPortsUpdateOK %s", 200, payload)
}

func (o *DcimRearPortsUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}