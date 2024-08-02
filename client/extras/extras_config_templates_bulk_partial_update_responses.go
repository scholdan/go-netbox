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

// ExtrasConfigTemplatesBulkPartialUpdateReader is a Reader for the ExtrasConfigTemplatesBulkPartialUpdate structure.
type ExtrasConfigTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/config-templates/] extras_config-templates_bulk_partial_update", response, response.Code())
	}
}

// NewExtrasConfigTemplatesBulkPartialUpdateOK creates a ExtrasConfigTemplatesBulkPartialUpdateOK with default headers values
func NewExtrasConfigTemplatesBulkPartialUpdateOK() *ExtrasConfigTemplatesBulkPartialUpdateOK {
	return &ExtrasConfigTemplatesBulkPartialUpdateOK{}
}

/*
ExtrasConfigTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigTemplatesBulkPartialUpdateOK extras config templates bulk partial update o k
*/
type ExtrasConfigTemplatesBulkPartialUpdateOK struct {
	Payload *models.ConfigTemplate
}

// IsSuccess returns true when this extras config templates bulk partial update o k response has a 2xx status code
func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config templates bulk partial update o k response has a 3xx status code
func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config templates bulk partial update o k response has a 4xx status code
func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config templates bulk partial update o k response has a 5xx status code
func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config templates bulk partial update o k response a status code equal to that given
func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras config templates bulk partial update o k response
func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/config-templates/][%d] extrasConfigTemplatesBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/config-templates/][%d] extrasConfigTemplatesBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) GetPayload() *models.ConfigTemplate {
	return o.Payload
}

func (o *ExtrasConfigTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
