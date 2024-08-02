// Code generated by go-swagger; DO NOT EDIT.

package users

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

// UsersGroupsBulkUpdateReader is a Reader for the UsersGroupsBulkUpdate structure.
type UsersGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /users/groups/] users_groups_bulk_update", response, response.Code())
	}
}

// NewUsersGroupsBulkUpdateOK creates a UsersGroupsBulkUpdateOK with default headers values
func NewUsersGroupsBulkUpdateOK() *UsersGroupsBulkUpdateOK {
	return &UsersGroupsBulkUpdateOK{}
}

/*
UsersGroupsBulkUpdateOK describes a response with status code 200, with default header values.

UsersGroupsBulkUpdateOK users groups bulk update o k
*/
type UsersGroupsBulkUpdateOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this users groups bulk update o k response has a 2xx status code
func (o *UsersGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups bulk update o k response has a 3xx status code
func (o *UsersGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups bulk update o k response has a 4xx status code
func (o *UsersGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups bulk update o k response has a 5xx status code
func (o *UsersGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups bulk update o k response a status code equal to that given
func (o *UsersGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users groups bulk update o k response
func (o *UsersGroupsBulkUpdateOK) Code() int {
	return 200
}

func (o *UsersGroupsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/groups/][%d] usersGroupsBulkUpdateOK %s", 200, payload)
}

func (o *UsersGroupsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/groups/][%d] usersGroupsBulkUpdateOK %s", 200, payload)
}

func (o *UsersGroupsBulkUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}