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

// UsersUsersReadReader is a Reader for the UsersUsersRead structure.
type UsersUsersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /users/users/{id}/] users_users_read", response, response.Code())
	}
}

// NewUsersUsersReadOK creates a UsersUsersReadOK with default headers values
func NewUsersUsersReadOK() *UsersUsersReadOK {
	return &UsersUsersReadOK{}
}

/*
UsersUsersReadOK describes a response with status code 200, with default header values.

UsersUsersReadOK users users read o k
*/
type UsersUsersReadOK struct {
	Payload *models.User
}

// IsSuccess returns true when this users users read o k response has a 2xx status code
func (o *UsersUsersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users read o k response has a 3xx status code
func (o *UsersUsersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users read o k response has a 4xx status code
func (o *UsersUsersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users read o k response has a 5xx status code
func (o *UsersUsersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users read o k response a status code equal to that given
func (o *UsersUsersReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users users read o k response
func (o *UsersUsersReadOK) Code() int {
	return 200
}

func (o *UsersUsersReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/users/{id}/][%d] usersUsersReadOK %s", 200, payload)
}

func (o *UsersUsersReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/users/{id}/][%d] usersUsersReadOK %s", 200, payload)
}

func (o *UsersUsersReadOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
