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

// UsersPermissionsCreateReader is a Reader for the UsersPermissionsCreate structure.
type UsersPermissionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersPermissionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /users/permissions/] users_permissions_create", response, response.Code())
	}
}

// NewUsersPermissionsCreateCreated creates a UsersPermissionsCreateCreated with default headers values
func NewUsersPermissionsCreateCreated() *UsersPermissionsCreateCreated {
	return &UsersPermissionsCreateCreated{}
}

/*
UsersPermissionsCreateCreated describes a response with status code 201, with default header values.

UsersPermissionsCreateCreated users permissions create created
*/
type UsersPermissionsCreateCreated struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions create created response has a 2xx status code
func (o *UsersPermissionsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions create created response has a 3xx status code
func (o *UsersPermissionsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions create created response has a 4xx status code
func (o *UsersPermissionsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions create created response has a 5xx status code
func (o *UsersPermissionsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions create created response a status code equal to that given
func (o *UsersPermissionsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the users permissions create created response
func (o *UsersPermissionsCreateCreated) Code() int {
	return 201
}

func (o *UsersPermissionsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/permissions/][%d] usersPermissionsCreateCreated %s", 201, payload)
}

func (o *UsersPermissionsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/permissions/][%d] usersPermissionsCreateCreated %s", 201, payload)
}

func (o *UsersPermissionsCreateCreated) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
