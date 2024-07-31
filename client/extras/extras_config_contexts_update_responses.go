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

	"github.com/fbreckle/go-netbox/models"
)

// ExtrasConfigContextsUpdateReader is a Reader for the ExtrasConfigContextsUpdate structure.
type ExtrasConfigContextsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/config-contexts/{id}/] extras_config-contexts_update", response, response.Code())
	}
}

// NewExtrasConfigContextsUpdateOK creates a ExtrasConfigContextsUpdateOK with default headers values
func NewExtrasConfigContextsUpdateOK() *ExtrasConfigContextsUpdateOK {
	return &ExtrasConfigContextsUpdateOK{}
}

/*
ExtrasConfigContextsUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsUpdateOK extras config contexts update o k
*/
type ExtrasConfigContextsUpdateOK struct {
	Payload *models.ConfigContext
}

// IsSuccess returns true when this extras config contexts update o k response has a 2xx status code
func (o *ExtrasConfigContextsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts update o k response has a 3xx status code
func (o *ExtrasConfigContextsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts update o k response has a 4xx status code
func (o *ExtrasConfigContextsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts update o k response has a 5xx status code
func (o *ExtrasConfigContextsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts update o k response a status code equal to that given
func (o *ExtrasConfigContextsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config contexts update o k response
func (o *ExtrasConfigContextsUpdateOK) Code() int {
	return 200
}

func (o *ExtrasConfigContextsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-contexts/{id}/][%d] extrasConfigContextsUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigContextsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-contexts/{id}/][%d] extrasConfigContextsUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigContextsUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
