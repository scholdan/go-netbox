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

// DcimModuleBayTemplatesUpdateReader is a Reader for the DcimModuleBayTemplatesUpdate structure.
type DcimModuleBayTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBayTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/module-bay-templates/{id}/] dcim_module-bay-templates_update", response, response.Code())
	}
}

// NewDcimModuleBayTemplatesUpdateOK creates a DcimModuleBayTemplatesUpdateOK with default headers values
func NewDcimModuleBayTemplatesUpdateOK() *DcimModuleBayTemplatesUpdateOK {
	return &DcimModuleBayTemplatesUpdateOK{}
}

/*
DcimModuleBayTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimModuleBayTemplatesUpdateOK dcim module bay templates update o k
*/
type DcimModuleBayTemplatesUpdateOK struct {
	Payload *models.ModuleBayTemplate
}

// IsSuccess returns true when this dcim module bay templates update o k response has a 2xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates update o k response has a 3xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates update o k response has a 4xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates update o k response has a 5xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates update o k response a status code equal to that given
func (o *DcimModuleBayTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim module bay templates update o k response
func (o *DcimModuleBayTemplatesUpdateOK) Code() int {
	return 200
}

func (o *DcimModuleBayTemplatesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesUpdateOK %s", 200, payload)
}

func (o *DcimModuleBayTemplatesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesUpdateOK %s", 200, payload)
}

func (o *DcimModuleBayTemplatesUpdateOK) GetPayload() *models.ModuleBayTemplate {
	return o.Payload
}

func (o *DcimModuleBayTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}