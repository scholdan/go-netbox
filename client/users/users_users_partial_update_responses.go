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

// UsersUsersPartialUpdateReader is a Reader for the UsersUsersPartialUpdate structure.
type UsersUsersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /users/users/{id}/] users_users_partial_update", response, response.Code())
	}
}

// NewUsersUsersPartialUpdateOK creates a UsersUsersPartialUpdateOK with default headers values
func NewUsersUsersPartialUpdateOK() *UsersUsersPartialUpdateOK {
	return &UsersUsersPartialUpdateOK{}
}

/*
UsersUsersPartialUpdateOK describes a response with status code 200, with default header values.

UsersUsersPartialUpdateOK users users partial update o k
*/
type UsersUsersPartialUpdateOK struct {
	Payload *models.User
}

// IsSuccess returns true when this users users partial update o k response has a 2xx status code
func (o *UsersUsersPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users partial update o k response has a 3xx status code
func (o *UsersUsersPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users partial update o k response has a 4xx status code
func (o *UsersUsersPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users partial update o k response has a 5xx status code
func (o *UsersUsersPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users partial update o k response a status code equal to that given
func (o *UsersUsersPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users users partial update o k response
func (o *UsersUsersPartialUpdateOK) Code() int {
	return 200
}

func (o *UsersUsersPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/users/{id}/][%d] usersUsersPartialUpdateOK %s", 200, payload)
}

func (o *UsersUsersPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /users/users/{id}/][%d] usersUsersPartialUpdateOK %s", 200, payload)
}

func (o *UsersUsersPartialUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
