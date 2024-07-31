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

// VirtualMachineWithConfigContext virtual machine with config context
//
// swagger:model VirtualMachineWithConfigContext
type VirtualMachineWithConfigContext struct {

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext interface{} `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	Device *NestedDevice `json:"device,omitempty"`

	// Disk (GB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Disk *int64 `json:"disk,omitempty"`

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

	// Local context data
	LocalContextData interface{} `json:"local_context_data,omitempty"`

	// Memory (MB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Memory *int64 `json:"memory,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// platform
	Platform *NestedPlatform `json:"platform,omitempty"`

	// primary ip
	PrimaryIP *NestedIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`

	// role
	Role *NestedDeviceRole `json:"role,omitempty"`

	// site
	Site *NestedSite `json:"site,omitempty"`

	// status
	Status *VirtualMachineWithConfigContextStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// VCPUs
	// Minimum: 0.01
	Vcpus *float64 `json:"vcpus,omitempty"`
}

// Validate validates this virtual machine with config context
func (m *VirtualMachineWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
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

	if err := m.validateVcpus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineWithConfigContext) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if err := validate.MinimumInt("disk", "body", *m.Disk, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("disk", "body", *m.Disk, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if err := validate.MinimumInt("memory", "body", *m.Memory, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("memory", "body", *m.Memory, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateName(formats strfmt.Registry) error {

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

func (m *VirtualMachineWithConfigContext) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePrimaryIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePrimaryIp4(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePrimaryIp6(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateStatus(formats strfmt.Registry) error {
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

func (m *VirtualMachineWithConfigContext) validateTags(formats strfmt.Registry) error {
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

func (m *VirtualMachineWithConfigContext) validateTenant(formats strfmt.Registry) error {
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

func (m *VirtualMachineWithConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateVcpus(formats strfmt.Registry) error {
	if swag.IsZero(m.Vcpus) { // not required
		return nil
	}

	if err := validate.Minimum("vcpus", "body", *m.Vcpus, 0.01, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this virtual machine with config context based on the context it is used
func (m *VirtualMachineWithConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
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

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIp6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {

		if swag.IsZero(m.Cluster) { // not required
			return nil
		}

		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {

		if swag.IsZero(m.Device) { // not required
			return nil
		}

		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {

		if swag.IsZero(m.Platform) { // not required
			return nil
		}

		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIP != nil {

		if swag.IsZero(m.PrimaryIP) { // not required
			return nil
		}

		if err := m.PrimaryIP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidatePrimaryIp4(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp4 != nil {

		if swag.IsZero(m.PrimaryIp4) { // not required
			return nil
		}

		if err := m.PrimaryIp4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidatePrimaryIp6(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryIp6 != nil {

		if swag.IsZero(m.PrimaryIp6) { // not required
			return nil
		}

		if err := m.PrimaryIp6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {

		if swag.IsZero(m.Role) { // not required
			return nil
		}

		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {

		if swag.IsZero(m.Site) { // not required
			return nil
		}

		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualMachineWithConfigContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualMachineWithConfigContext) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualMachineWithConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineWithConfigContext) UnmarshalBinary(b []byte) error {
	var res VirtualMachineWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualMachineWithConfigContextStatus Status
//
// swagger:model VirtualMachineWithConfigContextStatus
type VirtualMachineWithConfigContextStatus struct {

	// label
	// Required: true
	// Enum: ["Offline","Active","Planned","Staged","Failed","Decommissioning"]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: ["offline","active","planned","staged","failed","decommissioning"]
	Value *string `json:"value"`
}

// Validate validates this virtual machine with config context status
func (m *VirtualMachineWithConfigContextStatus) Validate(formats strfmt.Registry) error {
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

var virtualMachineWithConfigContextStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Active","Planned","Staged","Failed","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualMachineWithConfigContextStatusTypeLabelPropEnum = append(virtualMachineWithConfigContextStatusTypeLabelPropEnum, v)
	}
}

const (

	// VirtualMachineWithConfigContextStatusLabelOffline captures enum value "Offline"
	VirtualMachineWithConfigContextStatusLabelOffline string = "Offline"

	// VirtualMachineWithConfigContextStatusLabelActive captures enum value "Active"
	VirtualMachineWithConfigContextStatusLabelActive string = "Active"

	// VirtualMachineWithConfigContextStatusLabelPlanned captures enum value "Planned"
	VirtualMachineWithConfigContextStatusLabelPlanned string = "Planned"

	// VirtualMachineWithConfigContextStatusLabelStaged captures enum value "Staged"
	VirtualMachineWithConfigContextStatusLabelStaged string = "Staged"

	// VirtualMachineWithConfigContextStatusLabelFailed captures enum value "Failed"
	VirtualMachineWithConfigContextStatusLabelFailed string = "Failed"

	// VirtualMachineWithConfigContextStatusLabelDecommissioning captures enum value "Decommissioning"
	VirtualMachineWithConfigContextStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *VirtualMachineWithConfigContextStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, virtualMachineWithConfigContextStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMachineWithConfigContextStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var virtualMachineWithConfigContextStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","staged","failed","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualMachineWithConfigContextStatusTypeValuePropEnum = append(virtualMachineWithConfigContextStatusTypeValuePropEnum, v)
	}
}

const (

	// VirtualMachineWithConfigContextStatusValueOffline captures enum value "offline"
	VirtualMachineWithConfigContextStatusValueOffline string = "offline"

	// VirtualMachineWithConfigContextStatusValueActive captures enum value "active"
	VirtualMachineWithConfigContextStatusValueActive string = "active"

	// VirtualMachineWithConfigContextStatusValuePlanned captures enum value "planned"
	VirtualMachineWithConfigContextStatusValuePlanned string = "planned"

	// VirtualMachineWithConfigContextStatusValueStaged captures enum value "staged"
	VirtualMachineWithConfigContextStatusValueStaged string = "staged"

	// VirtualMachineWithConfigContextStatusValueFailed captures enum value "failed"
	VirtualMachineWithConfigContextStatusValueFailed string = "failed"

	// VirtualMachineWithConfigContextStatusValueDecommissioning captures enum value "decommissioning"
	VirtualMachineWithConfigContextStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *VirtualMachineWithConfigContextStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, virtualMachineWithConfigContextStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VirtualMachineWithConfigContextStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this virtual machine with config context status based on context it is used
func (m *VirtualMachineWithConfigContextStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineWithConfigContextStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineWithConfigContextStatus) UnmarshalBinary(b []byte) error {
	var res VirtualMachineWithConfigContextStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
