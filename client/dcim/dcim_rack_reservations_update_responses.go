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

// DcimRackReservationsUpdateReader is a Reader for the DcimRackReservationsUpdate structure.
type DcimRackReservationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/rack-reservations/{id}/] dcim_rack-reservations_update", response, response.Code())
	}
}

// NewDcimRackReservationsUpdateOK creates a DcimRackReservationsUpdateOK with default headers values
func NewDcimRackReservationsUpdateOK() *DcimRackReservationsUpdateOK {
	return &DcimRackReservationsUpdateOK{}
}

/*
DcimRackReservationsUpdateOK describes a response with status code 200, with default header values.

DcimRackReservationsUpdateOK dcim rack reservations update o k
*/
type DcimRackReservationsUpdateOK struct {
	Payload *models.RackReservation
}

// IsSuccess returns true when this dcim rack reservations update o k response has a 2xx status code
func (o *DcimRackReservationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack reservations update o k response has a 3xx status code
func (o *DcimRackReservationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack reservations update o k response has a 4xx status code
func (o *DcimRackReservationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack reservations update o k response has a 5xx status code
func (o *DcimRackReservationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack reservations update o k response a status code equal to that given
func (o *DcimRackReservationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rack reservations update o k response
func (o *DcimRackReservationsUpdateOK) Code() int {
	return 200
}

func (o *DcimRackReservationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/rack-reservations/{id}/][%d] dcimRackReservationsUpdateOK %s", 200, payload)
}

func (o *DcimRackReservationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/rack-reservations/{id}/][%d] dcimRackReservationsUpdateOK %s", 200, payload)
}

func (o *DcimRackReservationsUpdateOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}