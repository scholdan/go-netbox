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

// DcimRackReservationsReadReader is a Reader for the DcimRackReservationsRead structure.
type DcimRackReservationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/rack-reservations/{id}/] dcim_rack-reservations_read", response, response.Code())
	}
}

// NewDcimRackReservationsReadOK creates a DcimRackReservationsReadOK with default headers values
func NewDcimRackReservationsReadOK() *DcimRackReservationsReadOK {
	return &DcimRackReservationsReadOK{}
}

/*
DcimRackReservationsReadOK describes a response with status code 200, with default header values.

DcimRackReservationsReadOK dcim rack reservations read o k
*/
type DcimRackReservationsReadOK struct {
	Payload *models.RackReservation
}

// IsSuccess returns true when this dcim rack reservations read o k response has a 2xx status code
func (o *DcimRackReservationsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack reservations read o k response has a 3xx status code
func (o *DcimRackReservationsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack reservations read o k response has a 4xx status code
func (o *DcimRackReservationsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack reservations read o k response has a 5xx status code
func (o *DcimRackReservationsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack reservations read o k response a status code equal to that given
func (o *DcimRackReservationsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack reservations read o k response
func (o *DcimRackReservationsReadOK) Code() int {
	return 200
}

func (o *DcimRackReservationsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/rack-reservations/{id}/][%d] dcimRackReservationsReadOK %s", 200, payload)
}

func (o *DcimRackReservationsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/rack-reservations/{id}/][%d] dcimRackReservationsReadOK %s", 200, payload)
}

func (o *DcimRackReservationsReadOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}