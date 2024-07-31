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

// ExtrasSavedFiltersReadReader is a Reader for the ExtrasSavedFiltersRead structure.
type ExtrasSavedFiltersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasSavedFiltersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasSavedFiltersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/saved-filters/{id}/] extras_saved-filters_read", response, response.Code())
	}
}

// NewExtrasSavedFiltersReadOK creates a ExtrasSavedFiltersReadOK with default headers values
func NewExtrasSavedFiltersReadOK() *ExtrasSavedFiltersReadOK {
	return &ExtrasSavedFiltersReadOK{}
}

/*
ExtrasSavedFiltersReadOK describes a response with status code 200, with default header values.

ExtrasSavedFiltersReadOK extras saved filters read o k
*/
type ExtrasSavedFiltersReadOK struct {
	Payload *models.SavedFilter
}

// IsSuccess returns true when this extras saved filters read o k response has a 2xx status code
func (o *ExtrasSavedFiltersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras saved filters read o k response has a 3xx status code
func (o *ExtrasSavedFiltersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras saved filters read o k response has a 4xx status code
func (o *ExtrasSavedFiltersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras saved filters read o k response has a 5xx status code
func (o *ExtrasSavedFiltersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras saved filters read o k response a status code equal to that given
func (o *ExtrasSavedFiltersReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras saved filters read o k response
func (o *ExtrasSavedFiltersReadOK) Code() int {
	return 200
}

func (o *ExtrasSavedFiltersReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/saved-filters/{id}/][%d] extrasSavedFiltersReadOK %s", 200, payload)
}

func (o *ExtrasSavedFiltersReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/saved-filters/{id}/][%d] extrasSavedFiltersReadOK %s", 200, payload)
}

func (o *ExtrasSavedFiltersReadOK) GetPayload() *models.SavedFilter {
	return o.Payload
}

func (o *ExtrasSavedFiltersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
