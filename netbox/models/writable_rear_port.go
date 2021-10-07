// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// WritableRearPort writable rear port
//
// swagger:model WritableRearPort
type WritableRearPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 1024
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Type
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc f mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable rear port
func (m *WritableRearPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *WritableRearPort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableRearPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) validateName(formats strfmt.Registry) error {

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

func (m *WritableRearPort) validatePositions(formats strfmt.Registry) error {
	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if err := validate.MinimumInt("positions", "body", m.Positions, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("positions", "body", m.Positions, 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) validateTags(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

var writableRearPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","f","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRearPortTypeTypePropEnum = append(writableRearPortTypeTypePropEnum, v)
	}
}

const (

	// WritableRearPortTypeNr8p8c captures enum value "8p8c"
	WritableRearPortTypeNr8p8c string = "8p8c"

	// WritableRearPortTypeNr8p6c captures enum value "8p6c"
	WritableRearPortTypeNr8p6c string = "8p6c"

	// WritableRearPortTypeNr8p4c captures enum value "8p4c"
	WritableRearPortTypeNr8p4c string = "8p4c"

	// WritableRearPortTypeNr8p2c captures enum value "8p2c"
	WritableRearPortTypeNr8p2c string = "8p2c"

	// WritableRearPortTypeGg45 captures enum value "gg45"
	WritableRearPortTypeGg45 string = "gg45"

	// WritableRearPortTypeTeraDash4p captures enum value "tera-4p"
	WritableRearPortTypeTeraDash4p string = "tera-4p"

	// WritableRearPortTypeTeraDash2p captures enum value "tera-2p"
	WritableRearPortTypeTeraDash2p string = "tera-2p"

	// WritableRearPortTypeTeraDash1p captures enum value "tera-1p"
	WritableRearPortTypeTeraDash1p string = "tera-1p"

	// WritableRearPortTypeNr110DashPunch captures enum value "110-punch"
	WritableRearPortTypeNr110DashPunch string = "110-punch"

	// WritableRearPortTypeBnc captures enum value "bnc"
	WritableRearPortTypeBnc string = "bnc"

	// WritableRearPortTypeF captures enum value "f"
	WritableRearPortTypeF string = "f"

	// WritableRearPortTypeMrj21 captures enum value "mrj21"
	WritableRearPortTypeMrj21 string = "mrj21"

	// WritableRearPortTypeFc captures enum value "fc"
	WritableRearPortTypeFc string = "fc"

	// WritableRearPortTypeLc captures enum value "lc"
	WritableRearPortTypeLc string = "lc"

	// WritableRearPortTypeLcDashApc captures enum value "lc-apc"
	WritableRearPortTypeLcDashApc string = "lc-apc"

	// WritableRearPortTypeLsh captures enum value "lsh"
	WritableRearPortTypeLsh string = "lsh"

	// WritableRearPortTypeLshDashApc captures enum value "lsh-apc"
	WritableRearPortTypeLshDashApc string = "lsh-apc"

	// WritableRearPortTypeMpo captures enum value "mpo"
	WritableRearPortTypeMpo string = "mpo"

	// WritableRearPortTypeMtrj captures enum value "mtrj"
	WritableRearPortTypeMtrj string = "mtrj"

	// WritableRearPortTypeSc captures enum value "sc"
	WritableRearPortTypeSc string = "sc"

	// WritableRearPortTypeScDashApc captures enum value "sc-apc"
	WritableRearPortTypeScDashApc string = "sc-apc"

	// WritableRearPortTypeSt captures enum value "st"
	WritableRearPortTypeSt string = "st"

	// WritableRearPortTypeCs captures enum value "cs"
	WritableRearPortTypeCs string = "cs"

	// WritableRearPortTypeSn captures enum value "sn"
	WritableRearPortTypeSn string = "sn"

	// WritableRearPortTypeSplice captures enum value "splice"
	WritableRearPortTypeSplice string = "splice"
)

// prop value enum
func (m *WritableRearPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRearPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRearPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable rear port based on the context it is used
func (m *WritableRearPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *WritableRearPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableRearPort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableRearPort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableRearPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRearPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRearPort) UnmarshalBinary(b []byte) error {
	var res WritableRearPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
