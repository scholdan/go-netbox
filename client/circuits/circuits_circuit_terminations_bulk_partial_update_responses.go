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

// CircuitsCircuitTerminationsBulkPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsBulkPartialUpdate structure.
type CircuitsCircuitTerminationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /circuits/circuit-terminations/] circuits_circuit-terminations_bulk_partial_update", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsBulkPartialUpdateOK creates a CircuitsCircuitTerminationsBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsBulkPartialUpdateOK() *CircuitsCircuitTerminationsBulkPartialUpdateOK {
	return &CircuitsCircuitTerminationsBulkPartialUpdateOK{}
}

/*
CircuitsCircuitTerminationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsBulkPartialUpdateOK circuits circuit terminations bulk partial update o k
*/
type CircuitsCircuitTerminationsBulkPartialUpdateOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations bulk partial update o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations bulk partial update o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations bulk partial update o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations bulk partial update o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations bulk partial update o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit terminations bulk partial update o k response
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
