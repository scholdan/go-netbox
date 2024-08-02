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

// UsersPermissionsReadReader is a Reader for the UsersPermissionsRead structure.
type UsersPermissionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /users/permissions/{id}/] users_permissions_read", response, response.Code())
	}
}

// NewUsersPermissionsReadOK creates a UsersPermissionsReadOK with default headers values
func NewUsersPermissionsReadOK() *UsersPermissionsReadOK {
	return &UsersPermissionsReadOK{}
}

/*
UsersPermissionsReadOK describes a response with status code 200, with default header values.

UsersPermissionsReadOK users permissions read o k
*/
type UsersPermissionsReadOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions read o k response has a 2xx status code
func (o *UsersPermissionsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions read o k response has a 3xx status code
func (o *UsersPermissionsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions read o k response has a 4xx status code
func (o *UsersPermissionsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions read o k response has a 5xx status code
func (o *UsersPermissionsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions read o k response a status code equal to that given
func (o *UsersPermissionsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users permissions read o k response
func (o *UsersPermissionsReadOK) Code() int {
	return 200
}

func (o *UsersPermissionsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] usersPermissionsReadOK %s", 200, payload)
}

func (o *UsersPermissionsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] usersPermissionsReadOK %s", 200, payload)
}

func (o *UsersPermissionsReadOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}