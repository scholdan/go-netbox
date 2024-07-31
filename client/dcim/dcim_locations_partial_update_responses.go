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

// DcimLocationsPartialUpdateReader is a Reader for the DcimLocationsPartialUpdate structure.
type DcimLocationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/locations/{id}/] dcim_locations_partial_update", response, response.Code())
	}
}

// NewDcimLocationsPartialUpdateOK creates a DcimLocationsPartialUpdateOK with default headers values
func NewDcimLocationsPartialUpdateOK() *DcimLocationsPartialUpdateOK {
	return &DcimLocationsPartialUpdateOK{}
}

/*
DcimLocationsPartialUpdateOK describes a response with status code 200, with default header values.

DcimLocationsPartialUpdateOK dcim locations partial update o k
*/
type DcimLocationsPartialUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations partial update o k response has a 2xx status code
func (o *DcimLocationsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations partial update o k response has a 3xx status code
func (o *DcimLocationsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations partial update o k response has a 4xx status code
func (o *DcimLocationsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations partial update o k response has a 5xx status code
func (o *DcimLocationsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations partial update o k response a status code equal to that given
func (o *DcimLocationsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim locations partial update o k response
func (o *DcimLocationsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimLocationsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/locations/{id}/][%d] dcimLocationsPartialUpdateOK %s", 200, payload)
}

func (o *DcimLocationsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/locations/{id}/][%d] dcimLocationsPartialUpdateOK %s", 200, payload)
}

func (o *DcimLocationsPartialUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
