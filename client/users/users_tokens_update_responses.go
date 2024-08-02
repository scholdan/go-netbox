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

// UsersTokensUpdateReader is a Reader for the UsersTokensUpdate structure.
type UsersTokensUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /users/tokens/{id}/] users_tokens_update", response, response.Code())
	}
}

// NewUsersTokensUpdateOK creates a UsersTokensUpdateOK with default headers values
func NewUsersTokensUpdateOK() *UsersTokensUpdateOK {
	return &UsersTokensUpdateOK{}
}

/*
UsersTokensUpdateOK describes a response with status code 200, with default header values.

UsersTokensUpdateOK users tokens update o k
*/
type UsersTokensUpdateOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens update o k response has a 2xx status code
func (o *UsersTokensUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens update o k response has a 3xx status code
func (o *UsersTokensUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens update o k response has a 4xx status code
func (o *UsersTokensUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens update o k response has a 5xx status code
func (o *UsersTokensUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens update o k response a status code equal to that given
func (o *UsersTokensUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users tokens update o k response
func (o *UsersTokensUpdateOK) Code() int {
	return 200
}

func (o *UsersTokensUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/tokens/{id}/][%d] usersTokensUpdateOK %s", 200, payload)
}

func (o *UsersTokensUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/tokens/{id}/][%d] usersTokensUpdateOK %s", 200, payload)
}

func (o *UsersTokensUpdateOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}