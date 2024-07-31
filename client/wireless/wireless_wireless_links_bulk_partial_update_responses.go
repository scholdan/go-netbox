// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// WirelessWirelessLinksBulkPartialUpdateReader is a Reader for the WirelessWirelessLinksBulkPartialUpdate structure.
type WirelessWirelessLinksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /wireless/wireless-links/] wireless_wireless-links_bulk_partial_update", response, response.Code())
	}
}

// NewWirelessWirelessLinksBulkPartialUpdateOK creates a WirelessWirelessLinksBulkPartialUpdateOK with default headers values
func NewWirelessWirelessLinksBulkPartialUpdateOK() *WirelessWirelessLinksBulkPartialUpdateOK {
	return &WirelessWirelessLinksBulkPartialUpdateOK{}
}

/*
WirelessWirelessLinksBulkPartialUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLinksBulkPartialUpdateOK wireless wireless links bulk partial update o k
*/
type WirelessWirelessLinksBulkPartialUpdateOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links bulk partial update o k response has a 2xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links bulk partial update o k response has a 3xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links bulk partial update o k response has a 4xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links bulk partial update o k response has a 5xx status code
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links bulk partial update o k response a status code equal to that given
func (o *WirelessWirelessLinksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links bulk partial update o k response
func (o *WirelessWirelessLinksBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wireless/wireless-links/][%d] wirelessWirelessLinksBulkPartialUpdateOK %s", 200, payload)
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wireless/wireless-links/][%d] wirelessWirelessLinksBulkPartialUpdateOK %s", 200, payload)
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
