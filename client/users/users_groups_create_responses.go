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

// UsersGroupsCreateReader is a Reader for the UsersGroupsCreate structure.
type UsersGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /users/groups/] users_groups_create", response, response.Code())
	}
}

// NewUsersGroupsCreateCreated creates a UsersGroupsCreateCreated with default headers values
func NewUsersGroupsCreateCreated() *UsersGroupsCreateCreated {
	return &UsersGroupsCreateCreated{}
}

/*
UsersGroupsCreateCreated describes a response with status code 201, with default header values.

UsersGroupsCreateCreated users groups create created
*/
type UsersGroupsCreateCreated struct {
	Payload *models.Group
}

// IsSuccess returns true when this users groups create created response has a 2xx status code
func (o *UsersGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups create created response has a 3xx status code
func (o *UsersGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups create created response has a 4xx status code
func (o *UsersGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups create created response has a 5xx status code
func (o *UsersGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups create created response a status code equal to that given
func (o *UsersGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the users groups create created response
func (o *UsersGroupsCreateCreated) Code() int {
	return 201
}

func (o *UsersGroupsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/groups/][%d] usersGroupsCreateCreated %s", 201, payload)
}

func (o *UsersGroupsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/groups/][%d] usersGroupsCreateCreated %s", 201, payload)
}

func (o *UsersGroupsCreateCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}