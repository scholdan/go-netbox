// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInterfaceTemplatesDeleteReader is a Reader for the DcimInterfaceTemplatesDelete structure.
type DcimInterfaceTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInterfaceTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/interface-templates/{id}/] dcim_interface-templates_delete", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesDeleteNoContent creates a DcimInterfaceTemplatesDeleteNoContent with default headers values
func NewDcimInterfaceTemplatesDeleteNoContent() *DcimInterfaceTemplatesDeleteNoContent {
	return &DcimInterfaceTemplatesDeleteNoContent{}
}

/*
DcimInterfaceTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimInterfaceTemplatesDeleteNoContent dcim interface templates delete no content
*/
type DcimInterfaceTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim interface templates delete no content response has a 2xx status code
func (o *DcimInterfaceTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates delete no content response has a 3xx status code
func (o *DcimInterfaceTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates delete no content response has a 4xx status code
func (o *DcimInterfaceTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates delete no content response has a 5xx status code
func (o *DcimInterfaceTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates delete no content response a status code equal to that given
func (o *DcimInterfaceTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim interface templates delete no content response
func (o *DcimInterfaceTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimInterfaceTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesDeleteNoContent", 204)
}

func (o *DcimInterfaceTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesDeleteNoContent", 204)
}

func (o *DcimInterfaceTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
