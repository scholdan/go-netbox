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

// DcimInventoryItemsCreateReader is a Reader for the DcimInventoryItemsCreate structure.
type DcimInventoryItemsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInventoryItemsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/inventory-items/] dcim_inventory-items_create", response, response.Code())
	}
}

// NewDcimInventoryItemsCreateCreated creates a DcimInventoryItemsCreateCreated with default headers values
func NewDcimInventoryItemsCreateCreated() *DcimInventoryItemsCreateCreated {
	return &DcimInventoryItemsCreateCreated{}
}

/*
DcimInventoryItemsCreateCreated describes a response with status code 201, with default header values.

DcimInventoryItemsCreateCreated dcim inventory items create created
*/
type DcimInventoryItemsCreateCreated struct {
	Payload *models.InventoryItem
}

// IsSuccess returns true when this dcim inventory items create created response has a 2xx status code
func (o *DcimInventoryItemsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items create created response has a 3xx status code
func (o *DcimInventoryItemsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items create created response has a 4xx status code
func (o *DcimInventoryItemsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items create created response has a 5xx status code
func (o *DcimInventoryItemsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items create created response a status code equal to that given
func (o *DcimInventoryItemsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim inventory items create created response
func (o *DcimInventoryItemsCreateCreated) Code() int {
	return 201
}

func (o *DcimInventoryItemsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated %s", 201, payload)
}

func (o *DcimInventoryItemsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated %s", 201, payload)
}

func (o *DcimInventoryItemsCreateCreated) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
