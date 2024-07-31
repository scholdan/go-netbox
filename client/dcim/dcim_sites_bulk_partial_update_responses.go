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

// DcimSitesBulkPartialUpdateReader is a Reader for the DcimSitesBulkPartialUpdate structure.
type DcimSitesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/sites/] dcim_sites_bulk_partial_update", response, response.Code())
	}
}

// NewDcimSitesBulkPartialUpdateOK creates a DcimSitesBulkPartialUpdateOK with default headers values
func NewDcimSitesBulkPartialUpdateOK() *DcimSitesBulkPartialUpdateOK {
	return &DcimSitesBulkPartialUpdateOK{}
}

/*
DcimSitesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimSitesBulkPartialUpdateOK dcim sites bulk partial update o k
*/
type DcimSitesBulkPartialUpdateOK struct {
	Payload *models.Site
}

// IsSuccess returns true when this dcim sites bulk partial update o k response has a 2xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites bulk partial update o k response has a 3xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites bulk partial update o k response has a 4xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites bulk partial update o k response has a 5xx status code
func (o *DcimSitesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites bulk partial update o k response a status code equal to that given
func (o *DcimSitesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim sites bulk partial update o k response
func (o *DcimSitesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimSitesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcimSitesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimSitesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/sites/][%d] dcimSitesBulkPartialUpdateOK %s", 200, payload)
}

func (o *DcimSitesBulkPartialUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
