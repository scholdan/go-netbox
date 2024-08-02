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

	"github.com/scholdan/go-netbox/models"
)

// ExtrasJournalEntriesPartialUpdateReader is a Reader for the ExtrasJournalEntriesPartialUpdate structure.
type ExtrasJournalEntriesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/journal-entries/{id}/] extras_journal-entries_partial_update", response, response.Code())
	}
}

// NewExtrasJournalEntriesPartialUpdateOK creates a ExtrasJournalEntriesPartialUpdateOK with default headers values
func NewExtrasJournalEntriesPartialUpdateOK() *ExtrasJournalEntriesPartialUpdateOK {
	return &ExtrasJournalEntriesPartialUpdateOK{}
}

/*
ExtrasJournalEntriesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesPartialUpdateOK extras journal entries partial update o k
*/
type ExtrasJournalEntriesPartialUpdateOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries partial update o k response has a 2xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries partial update o k response has a 3xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries partial update o k response has a 4xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries partial update o k response has a 5xx status code
func (o *ExtrasJournalEntriesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries partial update o k response a status code equal to that given
func (o *ExtrasJournalEntriesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries partial update o k response
func (o *ExtrasJournalEntriesPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/journal-entries/{id}/][%d] extrasJournalEntriesPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/journal-entries/{id}/][%d] extrasJournalEntriesPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesPartialUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}