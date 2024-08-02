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

// ExtrasContentTypesReadReader is a Reader for the ExtrasContentTypesRead structure.
type ExtrasContentTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasContentTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasContentTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/content-types/{id}/] extras_content-types_read", response, response.Code())
	}
}

// NewExtrasContentTypesReadOK creates a ExtrasContentTypesReadOK with default headers values
func NewExtrasContentTypesReadOK() *ExtrasContentTypesReadOK {
	return &ExtrasContentTypesReadOK{}
}

/*
ExtrasContentTypesReadOK describes a response with status code 200, with default header values.

ExtrasContentTypesReadOK extras content types read o k
*/
type ExtrasContentTypesReadOK struct {
	Payload *models.ContentType
}

// IsSuccess returns true when this extras content types read o k response has a 2xx status code
func (o *ExtrasContentTypesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras content types read o k response has a 3xx status code
func (o *ExtrasContentTypesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras content types read o k response has a 4xx status code
func (o *ExtrasContentTypesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras content types read o k response has a 5xx status code
func (o *ExtrasContentTypesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras content types read o k response a status code equal to that given
func (o *ExtrasContentTypesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras content types read o k response
func (o *ExtrasContentTypesReadOK) Code() int {
	return 200
}

func (o *ExtrasContentTypesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadOK %s", 200, payload)
}

func (o *ExtrasContentTypesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadOK %s", 200, payload)
}

func (o *ExtrasContentTypesReadOK) GetPayload() *models.ContentType {
	return o.Payload
}

func (o *ExtrasContentTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}