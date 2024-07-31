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

// WirelessWirelessLanGroupsBulkPartialUpdateReader is a Reader for the WirelessWirelessLanGroupsBulkPartialUpdate structure.
type WirelessWirelessLanGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLanGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /wireless/wireless-lan-groups/] wireless_wireless-lan-groups_bulk_partial_update", response, response.Code())
	}
}

// NewWirelessWirelessLanGroupsBulkPartialUpdateOK creates a WirelessWirelessLanGroupsBulkPartialUpdateOK with default headers values
func NewWirelessWirelessLanGroupsBulkPartialUpdateOK() *WirelessWirelessLanGroupsBulkPartialUpdateOK {
	return &WirelessWirelessLanGroupsBulkPartialUpdateOK{}
}

/*
WirelessWirelessLanGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLanGroupsBulkPartialUpdateOK wireless wireless lan groups bulk partial update o k
*/
type WirelessWirelessLanGroupsBulkPartialUpdateOK struct {
	Payload *models.WirelessLANGroup
}

// IsSuccess returns true when this wireless wireless lan groups bulk partial update o k response has a 2xx status code
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lan groups bulk partial update o k response has a 3xx status code
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups bulk partial update o k response has a 4xx status code
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lan groups bulk partial update o k response has a 5xx status code
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups bulk partial update o k response a status code equal to that given
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless lan groups bulk partial update o k response
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsBulkPartialUpdateOK %s", 200, payload)
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsBulkPartialUpdateOK %s", 200, payload)
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) GetPayload() *models.WirelessLANGroup {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
