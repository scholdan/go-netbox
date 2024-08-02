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

// DcimInterfaceTemplatesBulkPartialUpdateReader is a Reader for the DcimInterfaceTemplatesBulkPartialUpdate structure.
type DcimInterfaceTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/interface-templates/] dcim_interface-templates_bulk_partial_update", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesBulkPartialUpdateOK creates a DcimInterfaceTemplatesBulkPartialUpdateOK with default headers values
func NewDcimInterfaceTemplatesBulkPartialUpdateOK() *DcimInterfaceTemplatesBulkPartialUpdateOK {
	return &DcimInterfaceTemplatesBulkPartialUpdateOK{}
}

/*
DcimInterfaceTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesBulkPartialUpdateOK dcim interface templates bulk partial update o k
*/
type DcimInterfaceTemplatesBulkPartialUpdateOK struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates bulk partial update o k response has a 2xx status code
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates bulk partial update o k response has a 3xx status code
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates bulk partial update o k response has a 4xx status code
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates bulk partial update o k response has a 5xx status code
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates bulk partial update o k response a status code equal to that given
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim interface templates bulk partial update o k response
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}