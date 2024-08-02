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

// DcimManufacturersPartialUpdateReader is a Reader for the DcimManufacturersPartialUpdate structure.
type DcimManufacturersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/manufacturers/{id}/] dcim_manufacturers_partial_update", response, response.Code())
	}
}

// NewDcimManufacturersPartialUpdateOK creates a DcimManufacturersPartialUpdateOK with default headers values
func NewDcimManufacturersPartialUpdateOK() *DcimManufacturersPartialUpdateOK {
	return &DcimManufacturersPartialUpdateOK{}
}

/*
DcimManufacturersPartialUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersPartialUpdateOK dcim manufacturers partial update o k
*/
type DcimManufacturersPartialUpdateOK struct {
	Payload *models.Manufacturer
}

// IsSuccess returns true when this dcim manufacturers partial update o k response has a 2xx status code
func (o *DcimManufacturersPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim manufacturers partial update o k response has a 3xx status code
func (o *DcimManufacturersPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers partial update o k response has a 4xx status code
func (o *DcimManufacturersPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim manufacturers partial update o k response has a 5xx status code
func (o *DcimManufacturersPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers partial update o k response a status code equal to that given
func (o *DcimManufacturersPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim manufacturers partial update o k response
func (o *DcimManufacturersPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimManufacturersPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcimManufacturersPartialUpdateOK %s", 200, payload)
}

func (o *DcimManufacturersPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcimManufacturersPartialUpdateOK %s", 200, payload)
}

func (o *DcimManufacturersPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}