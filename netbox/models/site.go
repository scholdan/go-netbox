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

// Site site
//
// swagger:model Site
type Site struct {

	// asns
	// Unique: true
	Asns []*NestedASN `json:"asns"`

	// Circuit count
	// Read Only: true
	CircuitCount int64 `json:"circuit_count,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Facility
	//
	// Local facility ID or description
	// Max Length: 50
	Facility string `json:"facility,omitempty"`

	// group
	Group *NestedSiteGroup `json:"group,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Latitude
	//
	// GPS coordinate (latitude)
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude
	//
	// GPS coordinate (longitude)
	Longitude *float64 `json:"longitude,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Physical address
	// Max Length: 200
	PhysicalAddress string `json:"physical_address,omitempty"`

	// Prefix count
	// Read Only: true
	PrefixCount int64 `json:"prefix_count,omitempty"`

	// Rack count
	// Read Only: true
	RackCount int64 `json:"rack_count,omitempty"`

	// region
	Region *NestedRegion `json:"region,omitempty"`

	// Shipping address
	// Max Length: 200
	ShippingAddress string `json:"shipping_address,omitempty"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// status
	Status *SiteStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Time zone
	TimeZone *string `json:"time_zone,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Virtualmachine count
	// Read Only: true
	VirtualmachineCount int64 `json:"virtualmachine_count,omitempty"`

	// Vlan count
	// Read Only: true
	VlanCount int64 `json:"vlan_count,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
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

func (m *Site) validateAsns(formats strfmt.Registry) error {
	if swag.IsZero(m.Asns) { // not required
		return nil
	}

	if err := validate.UniqueItems("asns", "body", m.Asns); err != nil {
		return err
	}

	for i := 0; i < len(m.Asns); i++ {
		if swag.IsZero(m.Asns[i]) { // not required
			continue
		}

		if m.Asns[i] != nil {
			if err := m.Asns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("asns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("asns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Site) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateFacility(formats strfmt.Registry) error {
	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if err := validate.MaxLength("facility", "body", m.Facility, 50); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateName(formats strfmt.Registry) error {

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

func (m *Site) validatePhysicalAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicalAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("physical_address", "body", m.PhysicalAddress, 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateShippingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("shipping_address", "body", m.ShippingAddress, 200); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", *m.Slug, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", *m.Slug, 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", *m.Slug, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Site) validateStatus(formats strfmt.Registry) error {
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

func (m *Site) validateTags(formats strfmt.Registry) error {
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

func (m *Site) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this site based on the context it is used
func (m *Site) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCircuitCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefixCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRackCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualmachineCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) contextValidateAsns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Asns); i++ {

		if m.Asns[i] != nil {

			if swag.IsZero(m.Asns[i]) { // not required
				return nil
			}

			if err := m.Asns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("asns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("asns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Site) contextValidateCircuitCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "circuit_count", "body", int64(m.CircuitCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateDeviceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "device_count", "body", int64(m.DeviceCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {

		if swag.IsZero(m.Group) { // not required
			return nil
		}

		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Site) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidatePrefixCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "prefix_count", "body", int64(m.PrefixCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateRackCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rack_count", "body", int64(m.RackCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *Site) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Site) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Site) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {

		if swag.IsZero(m.Tenant) { // not required
			return nil
		}

		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Site) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateVirtualmachineCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "virtualmachine_count", "body", int64(m.VirtualmachineCount)); err != nil {
		return err
	}

	return nil
}

func (m *Site) contextValidateVlanCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vlan_count", "body", int64(m.VlanCount)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SiteStatus Status
//
// swagger:model SiteStatus
type SiteStatus struct {

	// label
	// Required: true
	// Enum: ["Planned","Staging","Active","Decommissioning","Retired"]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: ["planned","staging","active","decommissioning","retired"]
	Value *string `json:"value"`
}

// Validate validates this site status
func (m *SiteStatus) Validate(formats strfmt.Registry) error {
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

var siteStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Planned","Staging","Active","Decommissioning","Retired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteStatusTypeLabelPropEnum = append(siteStatusTypeLabelPropEnum, v)
	}
}

const (

	// SiteStatusLabelPlanned captures enum value "Planned"
	SiteStatusLabelPlanned string = "Planned"

	// SiteStatusLabelStaging captures enum value "Staging"
	SiteStatusLabelStaging string = "Staging"

	// SiteStatusLabelActive captures enum value "Active"
	SiteStatusLabelActive string = "Active"

	// SiteStatusLabelDecommissioning captures enum value "Decommissioning"
	SiteStatusLabelDecommissioning string = "Decommissioning"

	// SiteStatusLabelRetired captures enum value "Retired"
	SiteStatusLabelRetired string = "Retired"
)

// prop value enum
func (m *SiteStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var siteStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["planned","staging","active","decommissioning","retired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteStatusTypeValuePropEnum = append(siteStatusTypeValuePropEnum, v)
	}
}

const (

	// SiteStatusValuePlanned captures enum value "planned"
	SiteStatusValuePlanned string = "planned"

	// SiteStatusValueStaging captures enum value "staging"
	SiteStatusValueStaging string = "staging"

	// SiteStatusValueActive captures enum value "active"
	SiteStatusValueActive string = "active"

	// SiteStatusValueDecommissioning captures enum value "decommissioning"
	SiteStatusValueDecommissioning string = "decommissioning"

	// SiteStatusValueRetired captures enum value "retired"
	SiteStatusValueRetired string = "retired"
)

// prop value enum
func (m *SiteStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, siteStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SiteStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this site status based on context it is used
func (m *SiteStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SiteStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteStatus) UnmarshalBinary(b []byte) error {
	var res SiteStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
