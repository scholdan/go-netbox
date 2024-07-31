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

	"github.com/fbreckle/go-netbox/models"
)

// UsersUsersCreateReader is a Reader for the UsersUsersCreate structure.
type UsersUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /users/users/] users_users_create", response, response.Code())
	}
}

// NewUsersUsersCreateCreated creates a UsersUsersCreateCreated with default headers values
func NewUsersUsersCreateCreated() *UsersUsersCreateCreated {
	return &UsersUsersCreateCreated{}
}

/*
UsersUsersCreateCreated describes a response with status code 201, with default header values.

UsersUsersCreateCreated users users create created
*/
type UsersUsersCreateCreated struct {
	Payload *models.User
}

// IsSuccess returns true when this users users create created response has a 2xx status code
func (o *UsersUsersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users create created response has a 3xx status code
func (o *UsersUsersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users create created response has a 4xx status code
func (o *UsersUsersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users create created response has a 5xx status code
func (o *UsersUsersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users users create created response a status code equal to that given
func (o *UsersUsersCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the users users create created response
func (o *UsersUsersCreateCreated) Code() int {
	return 201
}

func (o *UsersUsersCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/users/][%d] usersUsersCreateCreated %s", 201, payload)
}

func (o *UsersUsersCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/users/][%d] usersUsersCreateCreated %s", 201, payload)
}

func (o *UsersUsersCreateCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
