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

// DcimCablesBulkPartialUpdateReader is a Reader for the DcimCablesBulkPartialUpdate structure.
type DcimCablesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/cables/] dcim_cables_bulk_partial_update", response, response.Code())
	}
}

// NewDcimCablesBulkPartialUpdateOK creates a DcimCablesBulkPartialUpdateOK with default headers values
func NewDcimCablesBulkPartialUpdateOK() *DcimCablesBulkPartialUpdateOK {
	return &DcimCablesBulkPartialUpdateOK{}
}

/*
DcimCablesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimCablesBulkPartialUpdateOK dcim cables bulk partial update o k
*/
type DcimCablesBulkPartialUpdateOK struct {
	Payload *models.Cable
}

// IsSuccess returns true when this dcim cables bulk partial update o k response has a 2xx status code
func (o *DcimCablesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables bulk partial update o k response has a 3xx status code
func (o *DcimCablesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables bulk partial update o k response has a 4xx status code
func (o *DcimCablesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables bulk partial update o k response has a 5xx status code
func (o *DcimCablesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables bulk partial update o k response a status code equal to that given
func (o *DcimCablesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim cables bulk partial update o k response
func (o *DcimCablesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimCablesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/cables/][%d] dcimCablesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimCablesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/cables/][%d] dcimCablesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimCablesBulkPartialUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
