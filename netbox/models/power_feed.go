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

// PowerFeed power feed
//
// swagger:model PowerFeed
type PowerFeed struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable end
	// Read Only: true
	// Min Length: 1
	CableEnd string `json:"cable_end,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoints []interface{} `json:"connected_endpoints"`

	// Connected endpoints reachable
	// Read Only: true
	ConnectedEndpointsReachable *bool `json:"connected_endpoints_reachable,omitempty"`

	// Connected endpoints type
	// Read Only: true
	ConnectedEndpointsType string `json:"connected_endpoints_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	//
	// Return the appropriate serializer for the link termination model.
	//
	// Read Only: true
	LinkPeers []interface{} `json:"link_peers"`

	// Link peers type
	// Read Only: true
	LinkPeersType string `json:"link_peers_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// phase
	Phase *PowerFeedPhase `json:"phase,omitempty"`

	// power panel
	// Required: true
	PowerPanel *NestedPowerPanel `json:"power_panel"`

	// rack
	Rack *NestedRack `json:"rack,omitempty"`

	// status
	Status *PowerFeedStatus `json:"status,omitempty"`

	// supply
	Supply *PowerFeedSupply `json:"supply,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *PowerFeedType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: -32768
	Voltage *int64 `json:"voltage,omitempty"`
}

// Validate validates this power feed
func (m *PowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCableEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
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

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerFeed) validateAmperage(formats strfmt.Registry) error {
	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", m.Amperage, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", m.Amperage, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateCableEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.CableEnd) { // not required
		return nil
	}

	if err := validate.MinLength("cable_end", "body", m.CableEnd, 1); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateMaxUtilization(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", m.MaxUtilization, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", m.MaxUtilization, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateName(formats strfmt.Registry) error {

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

func (m *PowerFeed) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	if m.PowerPanel != nil {
		if err := m.PowerPanel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_panel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("power_panel")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateRack(formats strfmt.Registry) error {
	if swag.IsZero(m.Rack) { // not required
		return nil
	}

	if m.Rack != nil {
		if err := m.Rack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateSupply(formats strfmt.Registry) error {
	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	if m.Supply != nil {
		if err := m.Supply.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supply")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supply")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateTags(formats strfmt.Registry) error {
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

func (m *PowerFeed) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) validateVoltage(formats strfmt.Registry) error {
	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", *m.Voltage, -32768, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", *m.Voltage, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this power feed based on the context it is used
func (m *PowerFeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCableEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsType(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateLinkPeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeersType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerPanel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupply(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *PowerFeed) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {

		if swag.IsZero(m.Cable) { // not required
			return nil
		}

		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateCableEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_end", "body", string(m.CableEnd)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateConnectedEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints", "body", []interface{}(m.ConnectedEndpoints)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateConnectedEndpointsReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_reachable", "body", m.ConnectedEndpointsReachable); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateConnectedEndpointsType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_type", "body", string(m.ConnectedEndpointsType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateLinkPeers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers", "body", []interface{}(m.LinkPeers)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidateLinkPeersType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers_type", "body", string(m.LinkPeersType)); err != nil {
		return err
	}

	return nil
}

func (m *PowerFeed) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {

		if swag.IsZero(m.Phase) { // not required
			return nil
		}

		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidatePowerPanel(ctx context.Context, formats strfmt.Registry) error {

	if m.PowerPanel != nil {

		if err := m.PowerPanel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_panel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("power_panel")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateRack(ctx context.Context, formats strfmt.Registry) error {

	if m.Rack != nil {

		if swag.IsZero(m.Rack) { // not required
			return nil
		}

		if err := m.Rack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rack")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateSupply(ctx context.Context, formats strfmt.Registry) error {

	if m.Supply != nil {

		if swag.IsZero(m.Supply) { // not required
			return nil
		}

		if err := m.Supply.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supply")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supply")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PowerFeed) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *PowerFeed) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeed) UnmarshalBinary(b []byte) error {
	var res PowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedPhase Phase
//
// swagger:model PowerFeedPhase
type PowerFeedPhase struct {

	// label
	// Required: true
	// Enum: ["Single phase","Three-phase"]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: ["single-phase","three-phase"]
	Value *string `json:"value"`
}

func (m *PowerFeedPhase) UnmarshalJSON(b []byte) error {
	type PowerFeedPhaseAlias PowerFeedPhase
	var t PowerFeedPhaseAlias
	if err := json.Unmarshal([]byte("{\"label\":\"Single phase\",\"value\":\"single-phase\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedPhase(t)
	return nil
}

// Validate validates this power feed phase
func (m *PowerFeedPhase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerFeedPhaseTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Single phase","Three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedPhaseTypeLabelPropEnum = append(powerFeedPhaseTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedPhaseLabelSinglePhase captures enum value "Single phase"
	PowerFeedPhaseLabelSinglePhase string = "Single phase"

	// PowerFeedPhaseLabelThreeDashPhase captures enum value "Three-phase"
	PowerFeedPhaseLabelThreeDashPhase string = "Three-phase"
)

// prop value enum
func (m *PowerFeedPhase) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedPhaseTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedPhase) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("phase"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("phase"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedPhaseTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single-phase","three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedPhaseTypeValuePropEnum = append(powerFeedPhaseTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedPhaseValueSingleDashPhase captures enum value "single-phase"
	PowerFeedPhaseValueSingleDashPhase string = "single-phase"

	// PowerFeedPhaseValueThreeDashPhase captures enum value "three-phase"
	PowerFeedPhaseValueThreeDashPhase string = "three-phase"
)

// prop value enum
func (m *PowerFeedPhase) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedPhaseTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedPhase) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("phase"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("phase"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed phase based on context it is used
func (m *PowerFeedPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedPhase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedPhase) UnmarshalBinary(b []byte) error {
	var res PowerFeedPhase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedStatus Status
//
// swagger:model PowerFeedStatus
type PowerFeedStatus struct {

	// label
	// Required: true
	// Enum: ["Offline","Active","Planned","Failed"]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: ["offline","active","planned","failed"]
	Value *string `json:"value"`
}

func (m *PowerFeedStatus) UnmarshalJSON(b []byte) error {
	type PowerFeedStatusAlias PowerFeedStatus
	var t PowerFeedStatusAlias
	if err := json.Unmarshal([]byte("{\"label\":\"Active\",\"value\":\"active\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedStatus(t)
	return nil
}

// Validate validates this power feed status
func (m *PowerFeedStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerFeedStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Active","Planned","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedStatusTypeLabelPropEnum = append(powerFeedStatusTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedStatusLabelOffline captures enum value "Offline"
	PowerFeedStatusLabelOffline string = "Offline"

	// PowerFeedStatusLabelActive captures enum value "Active"
	PowerFeedStatusLabelActive string = "Active"

	// PowerFeedStatusLabelPlanned captures enum value "Planned"
	PowerFeedStatusLabelPlanned string = "Planned"

	// PowerFeedStatusLabelFailed captures enum value "Failed"
	PowerFeedStatusLabelFailed string = "Failed"
)

// prop value enum
func (m *PowerFeedStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedStatusTypeValuePropEnum = append(powerFeedStatusTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedStatusValueOffline captures enum value "offline"
	PowerFeedStatusValueOffline string = "offline"

	// PowerFeedStatusValueActive captures enum value "active"
	PowerFeedStatusValueActive string = "active"

	// PowerFeedStatusValuePlanned captures enum value "planned"
	PowerFeedStatusValuePlanned string = "planned"

	// PowerFeedStatusValueFailed captures enum value "failed"
	PowerFeedStatusValueFailed string = "failed"
)

// prop value enum
func (m *PowerFeedStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed status based on context it is used
func (m *PowerFeedStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedStatus) UnmarshalBinary(b []byte) error {
	var res PowerFeedStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedSupply Supply
//
// swagger:model PowerFeedSupply
type PowerFeedSupply struct {

	// label
	// Required: true
	// Enum: ["AC","DC"]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: ["ac","dc"]
	Value *string `json:"value"`
}

func (m *PowerFeedSupply) UnmarshalJSON(b []byte) error {
	type PowerFeedSupplyAlias PowerFeedSupply
	var t PowerFeedSupplyAlias
	if err := json.Unmarshal([]byte("{\"label\":\"AC\",\"value\":\"ac\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedSupply(t)
	return nil
}

// Validate validates this power feed supply
func (m *PowerFeedSupply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerFeedSupplyTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AC","DC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedSupplyTypeLabelPropEnum = append(powerFeedSupplyTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedSupplyLabelAC captures enum value "AC"
	PowerFeedSupplyLabelAC string = "AC"

	// PowerFeedSupplyLabelDC captures enum value "DC"
	PowerFeedSupplyLabelDC string = "DC"
)

// prop value enum
func (m *PowerFeedSupply) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedSupplyTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedSupply) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("supply"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("supply"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedSupplyTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ac","dc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedSupplyTypeValuePropEnum = append(powerFeedSupplyTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedSupplyValueAc captures enum value "ac"
	PowerFeedSupplyValueAc string = "ac"

	// PowerFeedSupplyValueDc captures enum value "dc"
	PowerFeedSupplyValueDc string = "dc"
)

// prop value enum
func (m *PowerFeedSupply) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedSupplyTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedSupply) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("supply"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("supply"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed supply based on context it is used
func (m *PowerFeedSupply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedSupply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedSupply) UnmarshalBinary(b []byte) error {
	var res PowerFeedSupply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerFeedType Type
//
// swagger:model PowerFeedType
type PowerFeedType struct {

	// label
	// Required: true
	// Enum: ["Primary","Redundant"]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: ["primary","redundant"]
	Value *string `json:"value"`
}

func (m *PowerFeedType) UnmarshalJSON(b []byte) error {
	type PowerFeedTypeAlias PowerFeedType
	var t PowerFeedTypeAlias
	if err := json.Unmarshal([]byte("{\"label\":\"Primary\",\"value\":\"primary\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PowerFeedType(t)
	return nil
}

// Validate validates this power feed type
func (m *PowerFeedType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerFeedTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Primary","Redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedTypeTypeLabelPropEnum = append(powerFeedTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerFeedTypeLabelPrimary captures enum value "Primary"
	PowerFeedTypeLabelPrimary string = "Primary"

	// PowerFeedTypeLabelRedundant captures enum value "Redundant"
	PowerFeedTypeLabelRedundant string = "Redundant"
)

// prop value enum
func (m *PowerFeedType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerFeedTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerFeedTypeTypeValuePropEnum = append(powerFeedTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerFeedTypeValuePrimary captures enum value "primary"
	PowerFeedTypeValuePrimary string = "primary"

	// PowerFeedTypeValueRedundant captures enum value "redundant"
	PowerFeedTypeValueRedundant string = "redundant"
)

// prop value enum
func (m *PowerFeedType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerFeedTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerFeedType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this power feed type based on context it is used
func (m *PowerFeedType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PowerFeedType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerFeedType) UnmarshalBinary(b []byte) error {
	var res PowerFeedType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
