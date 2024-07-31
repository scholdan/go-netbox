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

	"github.com/fbreckle/go-netbox/models"
)

// DcimLocationsUpdateReader is a Reader for the DcimLocationsUpdate structure.
type DcimLocationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/locations/{id}/] dcim_locations_update", response, response.Code())
	}
}

// NewDcimLocationsUpdateOK creates a DcimLocationsUpdateOK with default headers values
func NewDcimLocationsUpdateOK() *DcimLocationsUpdateOK {
	return &DcimLocationsUpdateOK{}
}

/*
DcimLocationsUpdateOK describes a response with status code 200, with default header values.

DcimLocationsUpdateOK dcim locations update o k
*/
type DcimLocationsUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations update o k response has a 2xx status code
func (o *DcimLocationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations update o k response has a 3xx status code
func (o *DcimLocationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations update o k response has a 4xx status code
func (o *DcimLocationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations update o k response has a 5xx status code
func (o *DcimLocationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations update o k response a status code equal to that given
func (o *DcimLocationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim locations update o k response
func (o *DcimLocationsUpdateOK) Code() int {
	return 200
}

func (o *DcimLocationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateOK %s", 200, payload)
}

func (o *DcimLocationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateOK %s", 200, payload)
}

func (o *DcimLocationsUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
