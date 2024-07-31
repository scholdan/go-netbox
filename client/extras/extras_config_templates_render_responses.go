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

// ExtrasConfigTemplatesRenderReader is a Reader for the ExtrasConfigTemplatesRender structure.
type ExtrasConfigTemplatesRenderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigTemplatesRenderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasConfigTemplatesRenderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /extras/config-templates/{id}/render/] extras_config-templates_render", response, response.Code())
	}
}

// NewExtrasConfigTemplatesRenderCreated creates a ExtrasConfigTemplatesRenderCreated with default headers values
func NewExtrasConfigTemplatesRenderCreated() *ExtrasConfigTemplatesRenderCreated {
	return &ExtrasConfigTemplatesRenderCreated{}
}

/*
ExtrasConfigTemplatesRenderCreated describes a response with status code 201, with default header values.

ExtrasConfigTemplatesRenderCreated extras config templates render created
*/
type ExtrasConfigTemplatesRenderCreated struct {
	Payload *models.ConfigTemplate
}

// IsSuccess returns true when this extras config templates render created response has a 2xx status code
func (o *ExtrasConfigTemplatesRenderCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config templates render created response has a 3xx status code
func (o *ExtrasConfigTemplatesRenderCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config templates render created response has a 4xx status code
func (o *ExtrasConfigTemplatesRenderCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config templates render created response has a 5xx status code
func (o *ExtrasConfigTemplatesRenderCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config templates render created response a status code equal to that given
func (o *ExtrasConfigTemplatesRenderCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the extras config templates render created response
func (o *ExtrasConfigTemplatesRenderCreated) Code() int {
	return 201
}

func (o *ExtrasConfigTemplatesRenderCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/config-templates/{id}/render/][%d] extrasConfigTemplatesRenderCreated %s", 201, payload)
}

func (o *ExtrasConfigTemplatesRenderCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/config-templates/{id}/render/][%d] extrasConfigTemplatesRenderCreated %s", 201, payload)
}

func (o *ExtrasConfigTemplatesRenderCreated) GetPayload() *models.ConfigTemplate {
	return o.Payload
}

func (o *ExtrasConfigTemplatesRenderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
