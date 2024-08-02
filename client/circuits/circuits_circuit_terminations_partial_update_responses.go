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

// CircuitsCircuitTerminationsPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsPartialUpdate structure.
type CircuitsCircuitTerminationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /circuits/circuit-terminations/{id}/] circuits_circuit-terminations_partial_update", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsPartialUpdateOK creates a CircuitsCircuitTerminationsPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsPartialUpdateOK() *CircuitsCircuitTerminationsPartialUpdateOK {
	return &CircuitsCircuitTerminationsPartialUpdateOK{}
}

/*
CircuitsCircuitTerminationsPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsPartialUpdateOK circuits circuit terminations partial update o k
*/
type CircuitsCircuitTerminationsPartialUpdateOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations partial update o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations partial update o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations partial update o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations partial update o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations partial update o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit terminations partial update o k response
func (o *CircuitsCircuitTerminationsPartialUpdateOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}