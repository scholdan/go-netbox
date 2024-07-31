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

// DcimRearPortsBulkPartialUpdateReader is a Reader for the DcimRearPortsBulkPartialUpdate structure.
type DcimRearPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/rear-ports/] dcim_rear-ports_bulk_partial_update", response, response.Code())
	}
}

// NewDcimRearPortsBulkPartialUpdateOK creates a DcimRearPortsBulkPartialUpdateOK with default headers values
func NewDcimRearPortsBulkPartialUpdateOK() *DcimRearPortsBulkPartialUpdateOK {
	return &DcimRearPortsBulkPartialUpdateOK{}
}

/*
DcimRearPortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRearPortsBulkPartialUpdateOK dcim rear ports bulk partial update o k
*/
type DcimRearPortsBulkPartialUpdateOK struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports bulk partial update o k response has a 2xx status code
func (o *DcimRearPortsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports bulk partial update o k response has a 3xx status code
func (o *DcimRearPortsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports bulk partial update o k response has a 4xx status code
func (o *DcimRearPortsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports bulk partial update o k response has a 5xx status code
func (o *DcimRearPortsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports bulk partial update o k response a status code equal to that given
func (o *DcimRearPortsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rear ports bulk partial update o k response
func (o *DcimRearPortsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimRearPortsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/rear-ports/][%d] dcimRearPortsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimRearPortsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/rear-ports/][%d] dcimRearPortsBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimRearPortsBulkPartialUpdateOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
