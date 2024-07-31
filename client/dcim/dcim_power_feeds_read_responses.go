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

// DcimPowerFeedsReadReader is a Reader for the DcimPowerFeedsRead structure.
type DcimPowerFeedsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/power-feeds/{id}/] dcim_power-feeds_read", response, response.Code())
	}
}

// NewDcimPowerFeedsReadOK creates a DcimPowerFeedsReadOK with default headers values
func NewDcimPowerFeedsReadOK() *DcimPowerFeedsReadOK {
	return &DcimPowerFeedsReadOK{}
}

/*
DcimPowerFeedsReadOK describes a response with status code 200, with default header values.

DcimPowerFeedsReadOK dcim power feeds read o k
*/
type DcimPowerFeedsReadOK struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds read o k response has a 2xx status code
func (o *DcimPowerFeedsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds read o k response has a 3xx status code
func (o *DcimPowerFeedsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds read o k response has a 4xx status code
func (o *DcimPowerFeedsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds read o k response has a 5xx status code
func (o *DcimPowerFeedsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds read o k response a status code equal to that given
func (o *DcimPowerFeedsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power feeds read o k response
func (o *DcimPowerFeedsReadOK) Code() int {
	return 200
}

func (o *DcimPowerFeedsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcimPowerFeedsReadOK %s", 200, payload)
}

func (o *DcimPowerFeedsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcimPowerFeedsReadOK %s", 200, payload)
}

func (o *DcimPowerFeedsReadOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
