// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// IpamServiceTemplatesBulkPartialUpdateReader is a Reader for the IpamServiceTemplatesBulkPartialUpdate structure.
type IpamServiceTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServiceTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/service-templates/] ipam_service-templates_bulk_partial_update", response, response.Code())
	}
}

// NewIpamServiceTemplatesBulkPartialUpdateOK creates a IpamServiceTemplatesBulkPartialUpdateOK with default headers values
func NewIpamServiceTemplatesBulkPartialUpdateOK() *IpamServiceTemplatesBulkPartialUpdateOK {
	return &IpamServiceTemplatesBulkPartialUpdateOK{}
}

/*
IpamServiceTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamServiceTemplatesBulkPartialUpdateOK ipam service templates bulk partial update o k
*/
type IpamServiceTemplatesBulkPartialUpdateOK struct {
	Payload *models.ServiceTemplate
}

// IsSuccess returns true when this ipam service templates bulk partial update o k response has a 2xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam service templates bulk partial update o k response has a 3xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam service templates bulk partial update o k response has a 4xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam service templates bulk partial update o k response has a 5xx status code
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam service templates bulk partial update o k response a status code equal to that given
func (o *IpamServiceTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam service templates bulk partial update o k response
func (o *IpamServiceTemplatesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/service-templates/][%d] ipamServiceTemplatesBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/service-templates/][%d] ipamServiceTemplatesBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) GetPayload() *models.ServiceTemplate {
	return o.Payload
}

func (o *IpamServiceTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
