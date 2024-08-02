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

// CircuitsProvidersUpdateReader is a Reader for the CircuitsProvidersUpdate structure.
type CircuitsProvidersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /circuits/providers/{id}/] circuits_providers_update", response, response.Code())
	}
}

// NewCircuitsProvidersUpdateOK creates a CircuitsProvidersUpdateOK with default headers values
func NewCircuitsProvidersUpdateOK() *CircuitsProvidersUpdateOK {
	return &CircuitsProvidersUpdateOK{}
}

/*
CircuitsProvidersUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersUpdateOK circuits providers update o k
*/
type CircuitsProvidersUpdateOK struct {
	Payload *models.Provider
}

// IsSuccess returns true when this circuits providers update o k response has a 2xx status code
func (o *CircuitsProvidersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers update o k response has a 3xx status code
func (o *CircuitsProvidersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers update o k response has a 4xx status code
func (o *CircuitsProvidersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers update o k response has a 5xx status code
func (o *CircuitsProvidersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers update o k response a status code equal to that given
func (o *CircuitsProvidersUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits providers update o k response
func (o *CircuitsProvidersUpdateOK) Code() int {
	return 200
}

func (o *CircuitsProvidersUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /circuits/providers/{id}/][%d] circuitsProvidersUpdateOK %s", 200, payload)
}

func (o *CircuitsProvidersUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /circuits/providers/{id}/][%d] circuitsProvidersUpdateOK %s", 200, payload)
}

func (o *CircuitsProvidersUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}