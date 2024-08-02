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

// DcimModuleBaysBulkUpdateReader is a Reader for the DcimModuleBaysBulkUpdate structure.
type DcimModuleBaysBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/module-bays/] dcim_module-bays_bulk_update", response, response.Code())
	}
}

// NewDcimModuleBaysBulkUpdateOK creates a DcimModuleBaysBulkUpdateOK with default headers values
func NewDcimModuleBaysBulkUpdateOK() *DcimModuleBaysBulkUpdateOK {
	return &DcimModuleBaysBulkUpdateOK{}
}

/*
DcimModuleBaysBulkUpdateOK describes a response with status code 200, with default header values.

DcimModuleBaysBulkUpdateOK dcim module bays bulk update o k
*/
type DcimModuleBaysBulkUpdateOK struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays bulk update o k response has a 2xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays bulk update o k response has a 3xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays bulk update o k response has a 4xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays bulk update o k response has a 5xx status code
func (o *DcimModuleBaysBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays bulk update o k response a status code equal to that given
func (o *DcimModuleBaysBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim module bays bulk update o k response
func (o *DcimModuleBaysBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimModuleBaysBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcimModuleBaysBulkUpdateOK %s", 200, payload)
}

func (o *DcimModuleBaysBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-bays/][%d] dcimModuleBaysBulkUpdateOK %s", 200, payload)
}

func (o *DcimModuleBaysBulkUpdateOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}