// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimSitesDeleteReader is a Reader for the DcimSitesDelete structure.
type DcimSitesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSitesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/sites/{id}/] dcim_sites_delete", response, response.Code())
	}
}

// NewDcimSitesDeleteNoContent creates a DcimSitesDeleteNoContent with default headers values
func NewDcimSitesDeleteNoContent() *DcimSitesDeleteNoContent {
	return &DcimSitesDeleteNoContent{}
}

/*
DcimSitesDeleteNoContent describes a response with status code 204, with default header values.

DcimSitesDeleteNoContent dcim sites delete no content
*/
type DcimSitesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim sites delete no content response has a 2xx status code
func (o *DcimSitesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites delete no content response has a 3xx status code
func (o *DcimSitesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites delete no content response has a 4xx status code
func (o *DcimSitesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites delete no content response has a 5xx status code
func (o *DcimSitesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites delete no content response a status code equal to that given
func (o *DcimSitesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim sites delete no content response
func (o *DcimSitesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimSitesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/{id}/][%d] dcimSitesDeleteNoContent", 204)
}

func (o *DcimSitesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/sites/{id}/][%d] dcimSitesDeleteNoContent", 204)
}

func (o *DcimSitesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}