// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInventoryItemRolesBulkDeleteReader is a Reader for the DcimInventoryItemRolesBulkDelete structure.
type DcimInventoryItemRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/inventory-item-roles/] dcim_inventory-item-roles_bulk_delete", response, response.Code())
	}
}

// NewDcimInventoryItemRolesBulkDeleteNoContent creates a DcimInventoryItemRolesBulkDeleteNoContent with default headers values
func NewDcimInventoryItemRolesBulkDeleteNoContent() *DcimInventoryItemRolesBulkDeleteNoContent {
	return &DcimInventoryItemRolesBulkDeleteNoContent{}
}

/*
DcimInventoryItemRolesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemRolesBulkDeleteNoContent dcim inventory item roles bulk delete no content
*/
type DcimInventoryItemRolesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim inventory item roles bulk delete no content response has a 2xx status code
func (o *DcimInventoryItemRolesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item roles bulk delete no content response has a 3xx status code
func (o *DcimInventoryItemRolesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item roles bulk delete no content response has a 4xx status code
func (o *DcimInventoryItemRolesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item roles bulk delete no content response has a 5xx status code
func (o *DcimInventoryItemRolesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item roles bulk delete no content response a status code equal to that given
func (o *DcimInventoryItemRolesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim inventory item roles bulk delete no content response
func (o *DcimInventoryItemRolesBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimInventoryItemRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkDeleteNoContent", 204)
}

func (o *DcimInventoryItemRolesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkDeleteNoContent", 204)
}

func (o *DcimInventoryItemRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
