// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// ExtrasObjectChangesReadReader is a Reader for the ExtrasObjectChangesRead structure.
type ExtrasObjectChangesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasObjectChangesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasObjectChangesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/object-changes/{id}/] extras_object-changes_read", response, response.Code())
	}
}

// NewExtrasObjectChangesReadOK creates a ExtrasObjectChangesReadOK with default headers values
func NewExtrasObjectChangesReadOK() *ExtrasObjectChangesReadOK {
	return &ExtrasObjectChangesReadOK{}
}

/*
ExtrasObjectChangesReadOK describes a response with status code 200, with default header values.

ExtrasObjectChangesReadOK extras object changes read o k
*/
type ExtrasObjectChangesReadOK struct {
	Payload *models.ObjectChange
}

// IsSuccess returns true when this extras object changes read o k response has a 2xx status code
func (o *ExtrasObjectChangesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras object changes read o k response has a 3xx status code
func (o *ExtrasObjectChangesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras object changes read o k response has a 4xx status code
func (o *ExtrasObjectChangesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras object changes read o k response has a 5xx status code
func (o *ExtrasObjectChangesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras object changes read o k response a status code equal to that given
func (o *ExtrasObjectChangesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras object changes read o k response
func (o *ExtrasObjectChangesReadOK) Code() int {
	return 200
}

func (o *ExtrasObjectChangesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK %s", 200, payload)
}

func (o *ExtrasObjectChangesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK %s", 200, payload)
}

func (o *ExtrasObjectChangesReadOK) GetPayload() *models.ObjectChange {
	return o.Payload
}

func (o *ExtrasObjectChangesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectChange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
