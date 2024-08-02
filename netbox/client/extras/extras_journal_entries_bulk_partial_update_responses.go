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

// ExtrasJournalEntriesBulkPartialUpdateReader is a Reader for the ExtrasJournalEntriesBulkPartialUpdate structure.
type ExtrasJournalEntriesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/journal-entries/] extras_journal-entries_bulk_partial_update", response, response.Code())
	}
}

// NewExtrasJournalEntriesBulkPartialUpdateOK creates a ExtrasJournalEntriesBulkPartialUpdateOK with default headers values
func NewExtrasJournalEntriesBulkPartialUpdateOK() *ExtrasJournalEntriesBulkPartialUpdateOK {
	return &ExtrasJournalEntriesBulkPartialUpdateOK{}
}

/*
ExtrasJournalEntriesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesBulkPartialUpdateOK extras journal entries bulk partial update o k
*/
type ExtrasJournalEntriesBulkPartialUpdateOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries bulk partial update o k response has a 2xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries bulk partial update o k response has a 3xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries bulk partial update o k response has a 4xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries bulk partial update o k response has a 5xx status code
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries bulk partial update o k response a status code equal to that given
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries bulk partial update o k response
func (o *ExtrasJournalEntriesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/journal-entries/][%d] extrasJournalEntriesBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/journal-entries/][%d] extrasJournalEntriesBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}