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

// DcimRackRolesCreateReader is a Reader for the DcimRackRolesCreate structure.
type DcimRackRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRackRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/rack-roles/] dcim_rack-roles_create", response, response.Code())
	}
}

// NewDcimRackRolesCreateCreated creates a DcimRackRolesCreateCreated with default headers values
func NewDcimRackRolesCreateCreated() *DcimRackRolesCreateCreated {
	return &DcimRackRolesCreateCreated{}
}

/*
DcimRackRolesCreateCreated describes a response with status code 201, with default header values.

DcimRackRolesCreateCreated dcim rack roles create created
*/
type DcimRackRolesCreateCreated struct {
	Payload *models.RackRole
}

// IsSuccess returns true when this dcim rack roles create created response has a 2xx status code
func (o *DcimRackRolesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles create created response has a 3xx status code
func (o *DcimRackRolesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles create created response has a 4xx status code
func (o *DcimRackRolesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles create created response has a 5xx status code
func (o *DcimRackRolesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles create created response a status code equal to that given
func (o *DcimRackRolesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim rack roles create created response
func (o *DcimRackRolesCreateCreated) Code() int {
	return 201
}

func (o *DcimRackRolesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateCreated %s", 201, payload)
}

func (o *DcimRackRolesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateCreated %s", 201, payload)
}

func (o *DcimRackRolesCreateCreated) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
