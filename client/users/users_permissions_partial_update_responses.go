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

// UsersPermissionsPartialUpdateReader is a Reader for the UsersPermissionsPartialUpdate structure.
type UsersPermissionsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /users/permissions/{id}/] users_permissions_partial_update", response, response.Code())
	}
}

// NewUsersPermissionsPartialUpdateOK creates a UsersPermissionsPartialUpdateOK with default headers values
func NewUsersPermissionsPartialUpdateOK() *UsersPermissionsPartialUpdateOK {
	return &UsersPermissionsPartialUpdateOK{}
}

/*
UsersPermissionsPartialUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsPartialUpdateOK users permissions partial update o k
*/
type UsersPermissionsPartialUpdateOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions partial update o k response has a 2xx status code
func (o *UsersPermissionsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions partial update o k response has a 3xx status code
func (o *UsersPermissionsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions partial update o k response has a 4xx status code
func (o *UsersPermissionsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions partial update o k response has a 5xx status code
func (o *UsersPermissionsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions partial update o k response a status code equal to that given
func (o *UsersPermissionsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users permissions partial update o k response
func (o *UsersPermissionsPartialUpdateOK) Code() int {
	return 200
}

func (o *UsersPermissionsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/permissions/{id}/][%d] usersPermissionsPartialUpdateOK %s", 200, payload)
}

func (o *UsersPermissionsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/permissions/{id}/][%d] usersPermissionsPartialUpdateOK %s", 200, payload)
}

func (o *UsersPermissionsPartialUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}