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

	"github.com/fbreckle/go-netbox/models"
)

// DcimInterfaceTemplatesPartialUpdateReader is a Reader for the DcimInterfaceTemplatesPartialUpdate structure.
type DcimInterfaceTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/interface-templates/{id}/] dcim_interface-templates_partial_update", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesPartialUpdateOK creates a DcimInterfaceTemplatesPartialUpdateOK with default headers values
func NewDcimInterfaceTemplatesPartialUpdateOK() *DcimInterfaceTemplatesPartialUpdateOK {
	return &DcimInterfaceTemplatesPartialUpdateOK{}
}

/*
DcimInterfaceTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesPartialUpdateOK dcim interface templates partial update o k
*/
type DcimInterfaceTemplatesPartialUpdateOK struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates partial update o k response has a 2xx status code
func (o *DcimInterfaceTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates partial update o k response has a 3xx status code
func (o *DcimInterfaceTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates partial update o k response has a 4xx status code
func (o *DcimInterfaceTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates partial update o k response has a 5xx status code
func (o *DcimInterfaceTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates partial update o k response a status code equal to that given
func (o *DcimInterfaceTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim interface templates partial update o k response
func (o *DcimInterfaceTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimInterfaceTemplatesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimInterfaceTemplatesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimInterfaceTemplatesPartialUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
