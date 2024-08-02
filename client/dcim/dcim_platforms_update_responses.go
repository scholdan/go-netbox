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

// DcimPlatformsUpdateReader is a Reader for the DcimPlatformsUpdate structure.
type DcimPlatformsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/platforms/{id}/] dcim_platforms_update", response, response.Code())
	}
}

// NewDcimPlatformsUpdateOK creates a DcimPlatformsUpdateOK with default headers values
func NewDcimPlatformsUpdateOK() *DcimPlatformsUpdateOK {
	return &DcimPlatformsUpdateOK{}
}

/*
DcimPlatformsUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsUpdateOK dcim platforms update o k
*/
type DcimPlatformsUpdateOK struct {
	Payload *models.Platform
}

// IsSuccess returns true when this dcim platforms update o k response has a 2xx status code
func (o *DcimPlatformsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms update o k response has a 3xx status code
func (o *DcimPlatformsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms update o k response has a 4xx status code
func (o *DcimPlatformsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms update o k response has a 5xx status code
func (o *DcimPlatformsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms update o k response a status code equal to that given
func (o *DcimPlatformsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim platforms update o k response
func (o *DcimPlatformsUpdateOK) Code() int {
	return 200
}

func (o *DcimPlatformsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/platforms/{id}/][%d] dcimPlatformsUpdateOK %s", 200, payload)
}

func (o *DcimPlatformsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/platforms/{id}/][%d] dcimPlatformsUpdateOK %s", 200, payload)
}

func (o *DcimPlatformsUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}