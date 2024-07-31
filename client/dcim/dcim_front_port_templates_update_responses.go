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

// DcimFrontPortTemplatesUpdateReader is a Reader for the DcimFrontPortTemplatesUpdate structure.
type DcimFrontPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/front-port-templates/{id}/] dcim_front-port-templates_update", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesUpdateOK creates a DcimFrontPortTemplatesUpdateOK with default headers values
func NewDcimFrontPortTemplatesUpdateOK() *DcimFrontPortTemplatesUpdateOK {
	return &DcimFrontPortTemplatesUpdateOK{}
}

/*
DcimFrontPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesUpdateOK dcim front port templates update o k
*/
type DcimFrontPortTemplatesUpdateOK struct {
	Payload *models.FrontPortTemplate
}

// IsSuccess returns true when this dcim front port templates update o k response has a 2xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front port templates update o k response has a 3xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates update o k response has a 4xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front port templates update o k response has a 5xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates update o k response a status code equal to that given
func (o *DcimFrontPortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim front port templates update o k response
func (o *DcimFrontPortTemplatesUpdateOK) Code() int {
	return 200
}

func (o *DcimFrontPortTemplatesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateOK %s", 200, payload)
}

func (o *DcimFrontPortTemplatesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateOK %s", 200, payload)
}

func (o *DcimFrontPortTemplatesUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
