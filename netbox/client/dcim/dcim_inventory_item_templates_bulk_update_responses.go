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

// DcimInventoryItemTemplatesBulkUpdateReader is a Reader for the DcimInventoryItemTemplatesBulkUpdate structure.
type DcimInventoryItemTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/inventory-item-templates/] dcim_inventory-item-templates_bulk_update", response, response.Code())
	}
}

// NewDcimInventoryItemTemplatesBulkUpdateOK creates a DcimInventoryItemTemplatesBulkUpdateOK with default headers values
func NewDcimInventoryItemTemplatesBulkUpdateOK() *DcimInventoryItemTemplatesBulkUpdateOK {
	return &DcimInventoryItemTemplatesBulkUpdateOK{}
}

/*
DcimInventoryItemTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesBulkUpdateOK dcim inventory item templates bulk update o k
*/
type DcimInventoryItemTemplatesBulkUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates bulk update o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates bulk update o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates bulk update o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates bulk update o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates bulk update o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim inventory item templates bulk update o k response
func (o *DcimInventoryItemTemplatesBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateOK %s", 200, payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateOK %s", 200, payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}