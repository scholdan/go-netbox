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

// DcimRackReservationsBulkPartialUpdateReader is a Reader for the DcimRackReservationsBulkPartialUpdate structure.
type DcimRackReservationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/rack-reservations/] dcim_rack-reservations_bulk_partial_update", response, response.Code())
	}
}

// NewDcimRackReservationsBulkPartialUpdateOK creates a DcimRackReservationsBulkPartialUpdateOK with default headers values
func NewDcimRackReservationsBulkPartialUpdateOK() *DcimRackReservationsBulkPartialUpdateOK {
	return &DcimRackReservationsBulkPartialUpdateOK{}
}

/*
DcimRackReservationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackReservationsBulkPartialUpdateOK dcim rack reservations bulk partial update o k
*/
type DcimRackReservationsBulkPartialUpdateOK struct {
	Payload *models.RackReservation
}

// IsSuccess returns true when this dcim rack reservations bulk partial update o k response has a 2xx status code
func (o *DcimRackReservationsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack reservations bulk partial update o k response has a 3xx status code
func (o *DcimRackReservationsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack reservations bulk partial update o k response has a 4xx status code
func (o *DcimRackReservationsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack reservations bulk partial update o k response has a 5xx status code
func (o *DcimRackReservationsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack reservations bulk partial update o k response a status code equal to that given
func (o *DcimRackReservationsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack reservations bulk partial update o k response
func (o *DcimRackReservationsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimRackReservationsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/rack-reservations/][%d] dcimRackReservationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimRackReservationsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/rack-reservations/][%d] dcimRackReservationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimRackReservationsBulkPartialUpdateOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
