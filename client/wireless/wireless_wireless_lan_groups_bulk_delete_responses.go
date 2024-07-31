// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLanGroupsBulkDeleteReader is a Reader for the WirelessWirelessLanGroupsBulkDelete structure.
type WirelessWirelessLanGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLanGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /wireless/wireless-lan-groups/] wireless_wireless-lan-groups_bulk_delete", response, response.Code())
	}
}

// NewWirelessWirelessLanGroupsBulkDeleteNoContent creates a WirelessWirelessLanGroupsBulkDeleteNoContent with default headers values
func NewWirelessWirelessLanGroupsBulkDeleteNoContent() *WirelessWirelessLanGroupsBulkDeleteNoContent {
	return &WirelessWirelessLanGroupsBulkDeleteNoContent{}
}

/*
WirelessWirelessLanGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLanGroupsBulkDeleteNoContent wireless wireless lan groups bulk delete no content
*/
type WirelessWirelessLanGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this wireless wireless lan groups bulk delete no content response has a 2xx status code
func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lan groups bulk delete no content response has a 3xx status code
func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups bulk delete no content response has a 4xx status code
func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lan groups bulk delete no content response has a 5xx status code
func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups bulk delete no content response a status code equal to that given
func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the wireless wireless lan groups bulk delete no content response
func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsBulkDeleteNoContent", 204)
}

func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsBulkDeleteNoContent", 204)
}

func (o *WirelessWirelessLanGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
