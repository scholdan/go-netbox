// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// DcimInterfaceTemplatesReadReader is a Reader for the DcimInterfaceTemplatesRead structure.
type DcimInterfaceTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/interface-templates/{id}/] dcim_interface-templates_read", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesReadOK creates a DcimInterfaceTemplatesReadOK with default headers values
func NewDcimInterfaceTemplatesReadOK() *DcimInterfaceTemplatesReadOK {
	return &DcimInterfaceTemplatesReadOK{}
}

/*
DcimInterfaceTemplatesReadOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesReadOK dcim interface templates read o k
*/
type DcimInterfaceTemplatesReadOK struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates read o k response has a 2xx status code
func (o *DcimInterfaceTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates read o k response has a 3xx status code
func (o *DcimInterfaceTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates read o k response has a 4xx status code
func (o *DcimInterfaceTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates read o k response has a 5xx status code
func (o *DcimInterfaceTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates read o k response a status code equal to that given
func (o *DcimInterfaceTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim interface templates read o k response
func (o *DcimInterfaceTemplatesReadOK) Code() int {
	return 200
}

func (o *DcimInterfaceTemplatesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesReadOK %s", 200, payload)
}

func (o *DcimInterfaceTemplatesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesReadOK %s", 200, payload)
}

func (o *DcimInterfaceTemplatesReadOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}