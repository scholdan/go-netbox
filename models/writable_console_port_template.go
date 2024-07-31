// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConsolePortTemplate writable console port template
//
// swagger:model WritableConsolePortTemplate
type WritableConsolePortTemplate struct {

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device type
	DeviceType *int64 `json:"device_type,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Module type
	ModuleType *int64 `json:"module_type,omitempty"`

	// Name
	//
	//
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	//
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Type
	// Enum: ["de-9","db-25","rj-11","rj-12","rj-45","mini-din-8","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-micro-ab","other"]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable console port template
func (m *WritableConsolePortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConsolePortTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

var writableConsolePortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","mini-din-8","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-micro-ab","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableConsolePortTemplateTypeTypePropEnum = append(writableConsolePortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableConsolePortTemplateTypeDeDash9 captures enum value "de-9"
	WritableConsolePortTemplateTypeDeDash9 string = "de-9"

	// WritableConsolePortTemplateTypeDbDash25 captures enum value "db-25"
	WritableConsolePortTemplateTypeDbDash25 string = "db-25"

	// WritableConsolePortTemplateTypeRjDash11 captures enum value "rj-11"
	WritableConsolePortTemplateTypeRjDash11 string = "rj-11"

	// WritableConsolePortTemplateTypeRjDash12 captures enum value "rj-12"
	WritableConsolePortTemplateTypeRjDash12 string = "rj-12"

	// WritableConsolePortTemplateTypeRjDash45 captures enum value "rj-45"
	WritableConsolePortTemplateTypeRjDash45 string = "rj-45"

	// WritableConsolePortTemplateTypeMiniDashDinDash8 captures enum value "mini-din-8"
	WritableConsolePortTemplateTypeMiniDashDinDash8 string = "mini-din-8"

	// WritableConsolePortTemplateTypeUsbDasha captures enum value "usb-a"
	WritableConsolePortTemplateTypeUsbDasha string = "usb-a"

	// WritableConsolePortTemplateTypeUsbDashb captures enum value "usb-b"
	WritableConsolePortTemplateTypeUsbDashb string = "usb-b"

	// WritableConsolePortTemplateTypeUsbDashc captures enum value "usb-c"
	WritableConsolePortTemplateTypeUsbDashc string = "usb-c"

	// WritableConsolePortTemplateTypeUsbDashMiniDasha captures enum value "usb-mini-a"
	WritableConsolePortTemplateTypeUsbDashMiniDasha string = "usb-mini-a"

	// WritableConsolePortTemplateTypeUsbDashMiniDashb captures enum value "usb-mini-b"
	WritableConsolePortTemplateTypeUsbDashMiniDashb string = "usb-mini-b"

	// WritableConsolePortTemplateTypeUsbDashMicroDasha captures enum value "usb-micro-a"
	WritableConsolePortTemplateTypeUsbDashMicroDasha string = "usb-micro-a"

	// WritableConsolePortTemplateTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritableConsolePortTemplateTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritableConsolePortTemplateTypeUsbDashMicroDashAb captures enum value "usb-micro-ab"
	WritableConsolePortTemplateTypeUsbDashMicroDashAb string = "usb-micro-ab"

	// WritableConsolePortTemplateTypeOther captures enum value "other"
	WritableConsolePortTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritableConsolePortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableConsolePortTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableConsolePortTemplate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable console port template based on the context it is used
func (m *WritableConsolePortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConsolePortTemplate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableConsolePortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableConsolePortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConsolePortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableConsolePortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
