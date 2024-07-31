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

// WirelessWirelessLansReadReader is a Reader for the WirelessWirelessLansRead structure.
type WirelessWirelessLansReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /wireless/wireless-lans/{id}/] wireless_wireless-lans_read", response, response.Code())
	}
}

// NewWirelessWirelessLansReadOK creates a WirelessWirelessLansReadOK with default headers values
func NewWirelessWirelessLansReadOK() *WirelessWirelessLansReadOK {
	return &WirelessWirelessLansReadOK{}
}

/*
WirelessWirelessLansReadOK describes a response with status code 200, with default header values.

WirelessWirelessLansReadOK wireless wireless lans read o k
*/
type WirelessWirelessLansReadOK struct {
	Payload *models.WirelessLAN
}

// IsSuccess returns true when this wireless wireless lans read o k response has a 2xx status code
func (o *WirelessWirelessLansReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans read o k response has a 3xx status code
func (o *WirelessWirelessLansReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans read o k response has a 4xx status code
func (o *WirelessWirelessLansReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans read o k response has a 5xx status code
func (o *WirelessWirelessLansReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans read o k response a status code equal to that given
func (o *WirelessWirelessLansReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless lans read o k response
func (o *WirelessWirelessLansReadOK) Code() int {
	return 200
}

func (o *WirelessWirelessLansReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansReadOK %s", 200, payload)
}

func (o *WirelessWirelessLansReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansReadOK %s", 200, payload)
}

func (o *WirelessWirelessLansReadOK) GetPayload() *models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
