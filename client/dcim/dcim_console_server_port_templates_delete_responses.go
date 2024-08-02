// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsoleServerPortTemplatesDeleteReader is a Reader for the DcimConsoleServerPortTemplatesDelete structure.
type DcimConsoleServerPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/console-server-port-templates/{id}/] dcim_console-server-port-templates_delete", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesDeleteNoContent creates a DcimConsoleServerPortTemplatesDeleteNoContent with default headers values
func NewDcimConsoleServerPortTemplatesDeleteNoContent() *DcimConsoleServerPortTemplatesDeleteNoContent {
	return &DcimConsoleServerPortTemplatesDeleteNoContent{}
}

/*
DcimConsoleServerPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortTemplatesDeleteNoContent dcim console server port templates delete no content
*/
type DcimConsoleServerPortTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim console server port templates delete no content response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates delete no content response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates delete no content response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates delete no content response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates delete no content response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim console server port templates delete no content response
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesDeleteNoContent", 204)
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesDeleteNoContent", 204)
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}