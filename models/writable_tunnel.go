// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableTunnel Adds support for custom fields and tags.
//
// swagger:model WritableTunnel
type WritableTunnel struct {

	// comments
	Comments string `json:"comments,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// * `ipsec-transport` - IPsec - Transport
	// * `ipsec-tunnel` - IPsec - Tunnel
	// * `ip-ip` - IP-in-IP
	// * `gre` - GRE
	// Required: true
	// Enum: ["ipsec-transport","ipsec-tunnel","ip-ip","gre"]
	Encapsulation *string `json:"encapsulation"`

	// group
	// Required: true
	Group *int64 `json:"group"`

	// ipsec profile
	IpsecProfile *int64 `json:"ipsec_profile,omitempty"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// * `planned` - Planned
	// * `active` - Active
	// * `disabled` - Disabled
	// Required: true
	// Enum: ["planned","active","disabled"]
	Status *string `json:"status"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// tunnel id
	// Maximum: 9.223372036854776e+18
	// Minimum: 0
	TunnelID *int64 `json:"tunnel_id,omitempty"`
}

// Validate validates this writable tunnel
func (m *WritableTunnel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncapsulation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnelID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableTunnel) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

var writableTunnelTypeEncapsulationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipsec-transport","ipsec-tunnel","ip-ip","gre"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableTunnelTypeEncapsulationPropEnum = append(writableTunnelTypeEncapsulationPropEnum, v)
	}
}

const (

	// WritableTunnelEncapsulationIpsecDashTransport captures enum value "ipsec-transport"
	WritableTunnelEncapsulationIpsecDashTransport string = "ipsec-transport"

	// WritableTunnelEncapsulationIpsecDashTunnel captures enum value "ipsec-tunnel"
	WritableTunnelEncapsulationIpsecDashTunnel string = "ipsec-tunnel"

	// WritableTunnelEncapsulationIPDashIP captures enum value "ip-ip"
	WritableTunnelEncapsulationIPDashIP string = "ip-ip"

	// WritableTunnelEncapsulationGre captures enum value "gre"
	WritableTunnelEncapsulationGre string = "gre"
)

// prop value enum
func (m *WritableTunnel) validateEncapsulationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableTunnelTypeEncapsulationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableTunnel) validateEncapsulation(formats strfmt.Registry) error {

	if err := validate.Required("encapsulation", "body", m.Encapsulation); err != nil {
		return err
	}

	// value enum
	if err := m.validateEncapsulationEnum("encapsulation", "body", *m.Encapsulation); err != nil {
		return err
	}

	return nil
}

func (m *WritableTunnel) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *WritableTunnel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

var writableTunnelTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["planned","active","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableTunnelTypeStatusPropEnum = append(writableTunnelTypeStatusPropEnum, v)
	}
}

const (

	// WritableTunnelStatusPlanned captures enum value "planned"
	WritableTunnelStatusPlanned string = "planned"

	// WritableTunnelStatusActive captures enum value "active"
	WritableTunnelStatusActive string = "active"

	// WritableTunnelStatusDisabled captures enum value "disabled"
	WritableTunnelStatusDisabled string = "disabled"
)

// prop value enum
func (m *WritableTunnel) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableTunnelTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableTunnel) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableTunnel) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableTunnel) validateTunnelID(formats strfmt.Registry) error {
	if swag.IsZero(m.TunnelID) { // not required
		return nil
	}

	if err := validate.MinimumInt("tunnel_id", "body", *m.TunnelID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tunnel_id", "body", *m.TunnelID, 9.223372036854776e+18, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable tunnel based on the context it is used
func (m *WritableTunnel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableTunnel) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableTunnel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableTunnel) UnmarshalBinary(b []byte) error {
	var res WritableTunnel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
