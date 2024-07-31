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

// DcimLocationsBulkPartialUpdateReader is a Reader for the DcimLocationsBulkPartialUpdate structure.
type DcimLocationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/locations/] dcim_locations_bulk_partial_update", response, response.Code())
	}
}

// NewDcimLocationsBulkPartialUpdateOK creates a DcimLocationsBulkPartialUpdateOK with default headers values
func NewDcimLocationsBulkPartialUpdateOK() *DcimLocationsBulkPartialUpdateOK {
	return &DcimLocationsBulkPartialUpdateOK{}
}

/*
DcimLocationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimLocationsBulkPartialUpdateOK dcim locations bulk partial update o k
*/
type DcimLocationsBulkPartialUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations bulk partial update o k response has a 2xx status code
func (o *DcimLocationsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations bulk partial update o k response has a 3xx status code
func (o *DcimLocationsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations bulk partial update o k response has a 4xx status code
func (o *DcimLocationsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations bulk partial update o k response has a 5xx status code
func (o *DcimLocationsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations bulk partial update o k response a status code equal to that given
func (o *DcimLocationsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim locations bulk partial update o k response
func (o *DcimLocationsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimLocationsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/locations/][%d] dcimLocationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimLocationsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/locations/][%d] dcimLocationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimLocationsBulkPartialUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
