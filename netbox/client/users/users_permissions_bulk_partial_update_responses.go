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

// UsersPermissionsBulkPartialUpdateReader is a Reader for the UsersPermissionsBulkPartialUpdate structure.
type UsersPermissionsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /users/permissions/] users_permissions_bulk_partial_update", response, response.Code())
	}
}

// NewUsersPermissionsBulkPartialUpdateOK creates a UsersPermissionsBulkPartialUpdateOK with default headers values
func NewUsersPermissionsBulkPartialUpdateOK() *UsersPermissionsBulkPartialUpdateOK {
	return &UsersPermissionsBulkPartialUpdateOK{}
}

/*
UsersPermissionsBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsBulkPartialUpdateOK users permissions bulk partial update o k
*/
type UsersPermissionsBulkPartialUpdateOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions bulk partial update o k response has a 2xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions bulk partial update o k response has a 3xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions bulk partial update o k response has a 4xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions bulk partial update o k response has a 5xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions bulk partial update o k response a status code equal to that given
func (o *UsersPermissionsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users permissions bulk partial update o k response
func (o *UsersPermissionsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *UsersPermissionsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/permissions/][%d] usersPermissionsBulkPartialUpdateOK %s", 200, payload)
}

func (o *UsersPermissionsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/permissions/][%d] usersPermissionsBulkPartialUpdateOK %s", 200, payload)
}

func (o *UsersPermissionsBulkPartialUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}