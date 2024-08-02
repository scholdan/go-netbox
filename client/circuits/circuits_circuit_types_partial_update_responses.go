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

// CircuitsCircuitTypesPartialUpdateReader is a Reader for the CircuitsCircuitTypesPartialUpdate structure.
type CircuitsCircuitTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /circuits/circuit-types/{id}/] circuits_circuit-types_partial_update", response, response.Code())
	}
}

// NewCircuitsCircuitTypesPartialUpdateOK creates a CircuitsCircuitTypesPartialUpdateOK with default headers values
func NewCircuitsCircuitTypesPartialUpdateOK() *CircuitsCircuitTypesPartialUpdateOK {
	return &CircuitsCircuitTypesPartialUpdateOK{}
}

/*
CircuitsCircuitTypesPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesPartialUpdateOK circuits circuit types partial update o k
*/
type CircuitsCircuitTypesPartialUpdateOK struct {
	Payload *models.CircuitType
}

// IsSuccess returns true when this circuits circuit types partial update o k response has a 2xx status code
func (o *CircuitsCircuitTypesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types partial update o k response has a 3xx status code
func (o *CircuitsCircuitTypesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types partial update o k response has a 4xx status code
func (o *CircuitsCircuitTypesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types partial update o k response has a 5xx status code
func (o *CircuitsCircuitTypesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types partial update o k response a status code equal to that given
func (o *CircuitsCircuitTypesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit types partial update o k response
func (o *CircuitsCircuitTypesPartialUpdateOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTypesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitTypesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesPartialUpdateOK %s", 200, payload)
}

func (o *CircuitsCircuitTypesPartialUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}