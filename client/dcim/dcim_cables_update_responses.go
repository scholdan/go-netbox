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

// DcimCablesUpdateReader is a Reader for the DcimCablesUpdate structure.
type DcimCablesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/cables/{id}/] dcim_cables_update", response, response.Code())
	}
}

// NewDcimCablesUpdateOK creates a DcimCablesUpdateOK with default headers values
func NewDcimCablesUpdateOK() *DcimCablesUpdateOK {
	return &DcimCablesUpdateOK{}
}

/*
DcimCablesUpdateOK describes a response with status code 200, with default header values.

DcimCablesUpdateOK dcim cables update o k
*/
type DcimCablesUpdateOK struct {
	Payload *models.Cable
}

// IsSuccess returns true when this dcim cables update o k response has a 2xx status code
func (o *DcimCablesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables update o k response has a 3xx status code
func (o *DcimCablesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables update o k response has a 4xx status code
func (o *DcimCablesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables update o k response has a 5xx status code
func (o *DcimCablesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables update o k response a status code equal to that given
func (o *DcimCablesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim cables update o k response
func (o *DcimCablesUpdateOK) Code() int {
	return 200
}

func (o *DcimCablesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/cables/{id}/][%d] dcimCablesUpdateOK %s", 200, payload)
}

func (o *DcimCablesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/cables/{id}/][%d] dcimCablesUpdateOK %s", 200, payload)
}

func (o *DcimCablesUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}