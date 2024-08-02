// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// CircuitsCircuitsPartialUpdateReader is a Reader for the CircuitsCircuitsPartialUpdate structure.
type CircuitsCircuitsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /circuits/circuits/{id}/] circuits_circuits_partial_update", response, response.Code())
	}
}

// NewCircuitsCircuitsPartialUpdateOK creates a CircuitsCircuitsPartialUpdateOK with default headers values
func NewCircuitsCircuitsPartialUpdateOK() *CircuitsCircuitsPartialUpdateOK {
	return &CircuitsCircuitsPartialUpdateOK{}
}

/*
CircuitsCircuitsPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsPartialUpdateOK circuits circuits partial update o k
*/
type CircuitsCircuitsPartialUpdateOK struct {
	Payload *models.Circuit
}

// IsSuccess returns true when this circuits circuits partial update o k response has a 2xx status code
func (o *CircuitsCircuitsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuits partial update o k response has a 3xx status code
func (o *CircuitsCircuitsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuits partial update o k response has a 4xx status code
func (o *CircuitsCircuitsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuits partial update o k response has a 5xx status code
func (o *CircuitsCircuitsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuits partial update o k response a status code equal to that given
func (o *CircuitsCircuitsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuits partial update o k response
func (o *CircuitsCircuitsPartialUpdateOK) Code() int {
	return 200
}

func (o *CircuitsCircuitsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuits/{id}/][%d] circuitsCircuitsPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuits/{id}/][%d] circuitsCircuitsPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitsPartialUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}