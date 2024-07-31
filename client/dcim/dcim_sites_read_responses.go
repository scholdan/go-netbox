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

// DcimSitesReadReader is a Reader for the DcimSitesRead structure.
type DcimSitesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/sites/{id}/] dcim_sites_read", response, response.Code())
	}
}

// NewDcimSitesReadOK creates a DcimSitesReadOK with default headers values
func NewDcimSitesReadOK() *DcimSitesReadOK {
	return &DcimSitesReadOK{}
}

/*
DcimSitesReadOK describes a response with status code 200, with default header values.

DcimSitesReadOK dcim sites read o k
*/
type DcimSitesReadOK struct {
	Payload *models.Site
}

// IsSuccess returns true when this dcim sites read o k response has a 2xx status code
func (o *DcimSitesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites read o k response has a 3xx status code
func (o *DcimSitesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites read o k response has a 4xx status code
func (o *DcimSitesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites read o k response has a 5xx status code
func (o *DcimSitesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites read o k response a status code equal to that given
func (o *DcimSitesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim sites read o k response
func (o *DcimSitesReadOK) Code() int {
	return 200
}

func (o *DcimSitesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/sites/{id}/][%d] dcimSitesReadOK %s", 200, payload)
}

func (o *DcimSitesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/sites/{id}/][%d] dcimSitesReadOK %s", 200, payload)
}

func (o *DcimSitesReadOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
