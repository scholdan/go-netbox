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

// WirelessWirelessLinksReadReader is a Reader for the WirelessWirelessLinksRead structure.
type WirelessWirelessLinksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /wireless/wireless-links/{id}/] wireless_wireless-links_read", response, response.Code())
	}
}

// NewWirelessWirelessLinksReadOK creates a WirelessWirelessLinksReadOK with default headers values
func NewWirelessWirelessLinksReadOK() *WirelessWirelessLinksReadOK {
	return &WirelessWirelessLinksReadOK{}
}

/*
WirelessWirelessLinksReadOK describes a response with status code 200, with default header values.

WirelessWirelessLinksReadOK wireless wireless links read o k
*/
type WirelessWirelessLinksReadOK struct {
	Payload *models.WirelessLink
}

// IsSuccess returns true when this wireless wireless links read o k response has a 2xx status code
func (o *WirelessWirelessLinksReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links read o k response has a 3xx status code
func (o *WirelessWirelessLinksReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links read o k response has a 4xx status code
func (o *WirelessWirelessLinksReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links read o k response has a 5xx status code
func (o *WirelessWirelessLinksReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links read o k response a status code equal to that given
func (o *WirelessWirelessLinksReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links read o k response
func (o *WirelessWirelessLinksReadOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksReadOK %s", 200, payload)
}

func (o *WirelessWirelessLinksReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksReadOK %s", 200, payload)
}

func (o *WirelessWirelessLinksReadOK) GetPayload() *models.WirelessLink {
	return o.Payload
}

func (o *WirelessWirelessLinksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}