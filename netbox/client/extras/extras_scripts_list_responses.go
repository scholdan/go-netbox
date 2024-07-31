// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasScriptsListReader is a Reader for the ExtrasScriptsList structure.
type ExtrasScriptsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasScriptsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasScriptsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/scripts/] extras_scripts_list", response, response.Code())
	}
}

// NewExtrasScriptsListOK creates a ExtrasScriptsListOK with default headers values
func NewExtrasScriptsListOK() *ExtrasScriptsListOK {
	return &ExtrasScriptsListOK{}
}

/*
ExtrasScriptsListOK describes a response with status code 200, with default header values.

ExtrasScriptsListOK extras scripts list o k
*/
type ExtrasScriptsListOK struct {
}

// IsSuccess returns true when this extras scripts list o k response has a 2xx status code
func (o *ExtrasScriptsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras scripts list o k response has a 3xx status code
func (o *ExtrasScriptsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras scripts list o k response has a 4xx status code
func (o *ExtrasScriptsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras scripts list o k response has a 5xx status code
func (o *ExtrasScriptsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras scripts list o k response a status code equal to that given
func (o *ExtrasScriptsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras scripts list o k response
func (o *ExtrasScriptsListOK) Code() int {
	return 200
}

func (o *ExtrasScriptsListOK) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/][%d] extrasScriptsListOK", 200)
}

func (o *ExtrasScriptsListOK) String() string {
	return fmt.Sprintf("[GET /extras/scripts/][%d] extrasScriptsListOK", 200)
}

func (o *ExtrasScriptsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
