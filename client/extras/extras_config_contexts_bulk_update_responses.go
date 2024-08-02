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

// ExtrasConfigContextsBulkUpdateReader is a Reader for the ExtrasConfigContextsBulkUpdate structure.
type ExtrasConfigContextsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/config-contexts/] extras_config-contexts_bulk_update", response, response.Code())
	}
}

// NewExtrasConfigContextsBulkUpdateOK creates a ExtrasConfigContextsBulkUpdateOK with default headers values
func NewExtrasConfigContextsBulkUpdateOK() *ExtrasConfigContextsBulkUpdateOK {
	return &ExtrasConfigContextsBulkUpdateOK{}
}

/*
ExtrasConfigContextsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsBulkUpdateOK extras config contexts bulk update o k
*/
type ExtrasConfigContextsBulkUpdateOK struct {
	Payload *models.ConfigContext
}

// IsSuccess returns true when this extras config contexts bulk update o k response has a 2xx status code
func (o *ExtrasConfigContextsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts bulk update o k response has a 3xx status code
func (o *ExtrasConfigContextsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts bulk update o k response has a 4xx status code
func (o *ExtrasConfigContextsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts bulk update o k response has a 5xx status code
func (o *ExtrasConfigContextsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts bulk update o k response a status code equal to that given
func (o *ExtrasConfigContextsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config contexts bulk update o k response
func (o *ExtrasConfigContextsBulkUpdateOK) Code() int {
	return 200
}

func (o *ExtrasConfigContextsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-contexts/][%d] extrasConfigContextsBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigContextsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/config-contexts/][%d] extrasConfigContextsBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigContextsBulkUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
