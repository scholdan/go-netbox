// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// ExtrasJournalEntriesReadReader is a Reader for the ExtrasJournalEntriesRead structure.
type ExtrasJournalEntriesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/journal-entries/{id}/] extras_journal-entries_read", response, response.Code())
	}
}

// NewExtrasJournalEntriesReadOK creates a ExtrasJournalEntriesReadOK with default headers values
func NewExtrasJournalEntriesReadOK() *ExtrasJournalEntriesReadOK {
	return &ExtrasJournalEntriesReadOK{}
}

/*
ExtrasJournalEntriesReadOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesReadOK extras journal entries read o k
*/
type ExtrasJournalEntriesReadOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries read o k response has a 2xx status code
func (o *ExtrasJournalEntriesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries read o k response has a 3xx status code
func (o *ExtrasJournalEntriesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries read o k response has a 4xx status code
func (o *ExtrasJournalEntriesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries read o k response has a 5xx status code
func (o *ExtrasJournalEntriesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries read o k response a status code equal to that given
func (o *ExtrasJournalEntriesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries read o k response
func (o *ExtrasJournalEntriesReadOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/journal-entries/{id}/][%d] extrasJournalEntriesReadOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/journal-entries/{id}/][%d] extrasJournalEntriesReadOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesReadOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
