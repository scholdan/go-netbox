// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimModulesDeleteReader is a Reader for the DcimModulesDelete structure.
type DcimModulesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModulesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModulesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/modules/{id}/] dcim_modules_delete", response, response.Code())
	}
}

// NewDcimModulesDeleteNoContent creates a DcimModulesDeleteNoContent with default headers values
func NewDcimModulesDeleteNoContent() *DcimModulesDeleteNoContent {
	return &DcimModulesDeleteNoContent{}
}

/*
DcimModulesDeleteNoContent describes a response with status code 204, with default header values.

DcimModulesDeleteNoContent dcim modules delete no content
*/
type DcimModulesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim modules delete no content response has a 2xx status code
func (o *DcimModulesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim modules delete no content response has a 3xx status code
func (o *DcimModulesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim modules delete no content response has a 4xx status code
func (o *DcimModulesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim modules delete no content response has a 5xx status code
func (o *DcimModulesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim modules delete no content response a status code equal to that given
func (o *DcimModulesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim modules delete no content response
func (o *DcimModulesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimModulesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/modules/{id}/][%d] dcimModulesDeleteNoContent", 204)
}

func (o *DcimModulesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/modules/{id}/][%d] dcimModulesDeleteNoContent", 204)
}

func (o *DcimModulesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
