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

	"github.com/scholdan/go-netbox/models"
)

// WirelessWirelessLinksUpdateReader is a Reader for the WirelessWirelessLinksUpdate structure.
type WirelessWirelessLinksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /wireless/wireless-links/{id}/] wireless_wireless-links_update", response, response.Code())
	}
}

// NewWirelessWirelessLinksUpdateOK creates a WirelessWirelessLinksUpdateOK with default headers values
func NewWirelessWirelessLinksUpdateOK() *WirelessWirelessLinksUpdateOK {
	return &WirelessWirelessLinksUpdateOK{}
}

/*
WirelessWirelessLinksUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLinksUpdateOK wireless wireless links update o k
*/
type WirelessWirelessLinksUpdateOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links update o k response has a 2xx status code
func (o *WirelessWirelessLinksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links update o k response has a 3xx status code
func (o *WirelessWirelessLinksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links update o k response has a 4xx status code
func (o *WirelessWirelessLinksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links update o k response has a 5xx status code
func (o *WirelessWirelessLinksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links update o k response a status code equal to that given
func (o *WirelessWirelessLinksUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links update o k response
func (o *WirelessWirelessLinksUpdateOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksUpdateOK %s", 200, payload)
}

func (o *WirelessWirelessLinksUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksUpdateOK %s", 200, payload)
}

func (o *WirelessWirelessLinksUpdateOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}