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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritablePowerOutletTemplate writable power outlet template
//
// swagger:model WritablePowerOutletTemplate
type WritablePowerOutletTemplate struct {

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

	// Feed leg
	//
	// Phase (for three-phase feeds)
	// Enum: ["A","B","C"]
	FeedLeg string `json:"feed_leg,omitempty"`

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

	// Power port
	PowerPort *int64 `json:"power_port,omitempty"`

	// Type
	// Enum: ["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60320-c21","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","nema-l22-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","ita-multistandard","usb-a","usb-micro-b","usb-c","dc-terminal","hdot-cx","saf-d-grid","neutrik-powercon-20a","neutrik-powercon-32a","neutrik-powercon-true1","neutrik-powercon-true1-top","ubiquiti-smartpower","hardwired","other"]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable power outlet template
func (m *WritablePowerOutletTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedLeg(formats); err != nil {
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

func (m *WritablePowerOutletTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

var writablePowerOutletTemplateTypeFeedLegPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTemplateTypeFeedLegPropEnum = append(writablePowerOutletTemplateTypeFeedLegPropEnum, v)
	}
}

const (

	// WritablePowerOutletTemplateFeedLegA captures enum value "A"
	WritablePowerOutletTemplateFeedLegA string = "A"

	// WritablePowerOutletTemplateFeedLegB captures enum value "B"
	WritablePowerOutletTemplateFeedLegB string = "B"

	// WritablePowerOutletTemplateFeedLegC captures enum value "C"
	WritablePowerOutletTemplateFeedLegC string = "C"
)

// prop value enum
func (m *WritablePowerOutletTemplate) validateFeedLegEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerOutletTemplateTypeFeedLegPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutletTemplate) validateFeedLeg(formats strfmt.Registry) error {
	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeedLegEnum("feed_leg", "body", m.FeedLeg); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validateName(formats strfmt.Registry) error {

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

var writablePowerOutletTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60320-c21","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","nema-l22-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","ita-multistandard","usb-a","usb-micro-b","usb-c","dc-terminal","hdot-cx","saf-d-grid","neutrik-powercon-20a","neutrik-powercon-32a","neutrik-powercon-true1","neutrik-powercon-true1-top","ubiquiti-smartpower","hardwired","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerOutletTemplateTypeTypePropEnum = append(writablePowerOutletTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerOutletTemplateTypeIecDash60320DashC5 captures enum value "iec-60320-c5"
	WritablePowerOutletTemplateTypeIecDash60320DashC5 string = "iec-60320-c5"

	// WritablePowerOutletTemplateTypeIecDash60320DashC7 captures enum value "iec-60320-c7"
	WritablePowerOutletTemplateTypeIecDash60320DashC7 string = "iec-60320-c7"

	// WritablePowerOutletTemplateTypeIecDash60320DashC13 captures enum value "iec-60320-c13"
	WritablePowerOutletTemplateTypeIecDash60320DashC13 string = "iec-60320-c13"

	// WritablePowerOutletTemplateTypeIecDash60320DashC15 captures enum value "iec-60320-c15"
	WritablePowerOutletTemplateTypeIecDash60320DashC15 string = "iec-60320-c15"

	// WritablePowerOutletTemplateTypeIecDash60320DashC19 captures enum value "iec-60320-c19"
	WritablePowerOutletTemplateTypeIecDash60320DashC19 string = "iec-60320-c19"

	// WritablePowerOutletTemplateTypeIecDash60320DashC21 captures enum value "iec-60320-c21"
	WritablePowerOutletTemplateTypeIecDash60320DashC21 string = "iec-60320-c21"

	// WritablePowerOutletTemplateTypeIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerOutletTemplateTypeIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// WritablePowerOutletTemplateTypeIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerOutletTemplateTypeIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// WritablePowerOutletTemplateTypeIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerOutletTemplateTypeIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerOutletTemplateTypeIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerOutletTemplateTypeIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerOutletTemplateTypeIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerOutletTemplateTypeIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerOutletTemplateTypeIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerOutletTemplateTypeIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerOutletTemplateTypeIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerOutletTemplateTypeIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerOutletTemplateTypeIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerOutletTemplateTypeIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerOutletTemplateTypeNemaDash1Dash15r captures enum value "nema-1-15r"
	WritablePowerOutletTemplateTypeNemaDash1Dash15r string = "nema-1-15r"

	// WritablePowerOutletTemplateTypeNemaDash5Dash15r captures enum value "nema-5-15r"
	WritablePowerOutletTemplateTypeNemaDash5Dash15r string = "nema-5-15r"

	// WritablePowerOutletTemplateTypeNemaDash5Dash20r captures enum value "nema-5-20r"
	WritablePowerOutletTemplateTypeNemaDash5Dash20r string = "nema-5-20r"

	// WritablePowerOutletTemplateTypeNemaDash5Dash30r captures enum value "nema-5-30r"
	WritablePowerOutletTemplateTypeNemaDash5Dash30r string = "nema-5-30r"

	// WritablePowerOutletTemplateTypeNemaDash5Dash50r captures enum value "nema-5-50r"
	WritablePowerOutletTemplateTypeNemaDash5Dash50r string = "nema-5-50r"

	// WritablePowerOutletTemplateTypeNemaDash6Dash15r captures enum value "nema-6-15r"
	WritablePowerOutletTemplateTypeNemaDash6Dash15r string = "nema-6-15r"

	// WritablePowerOutletTemplateTypeNemaDash6Dash20r captures enum value "nema-6-20r"
	WritablePowerOutletTemplateTypeNemaDash6Dash20r string = "nema-6-20r"

	// WritablePowerOutletTemplateTypeNemaDash6Dash30r captures enum value "nema-6-30r"
	WritablePowerOutletTemplateTypeNemaDash6Dash30r string = "nema-6-30r"

	// WritablePowerOutletTemplateTypeNemaDash6Dash50r captures enum value "nema-6-50r"
	WritablePowerOutletTemplateTypeNemaDash6Dash50r string = "nema-6-50r"

	// WritablePowerOutletTemplateTypeNemaDash10Dash30r captures enum value "nema-10-30r"
	WritablePowerOutletTemplateTypeNemaDash10Dash30r string = "nema-10-30r"

	// WritablePowerOutletTemplateTypeNemaDash10Dash50r captures enum value "nema-10-50r"
	WritablePowerOutletTemplateTypeNemaDash10Dash50r string = "nema-10-50r"

	// WritablePowerOutletTemplateTypeNemaDash14Dash20r captures enum value "nema-14-20r"
	WritablePowerOutletTemplateTypeNemaDash14Dash20r string = "nema-14-20r"

	// WritablePowerOutletTemplateTypeNemaDash14Dash30r captures enum value "nema-14-30r"
	WritablePowerOutletTemplateTypeNemaDash14Dash30r string = "nema-14-30r"

	// WritablePowerOutletTemplateTypeNemaDash14Dash50r captures enum value "nema-14-50r"
	WritablePowerOutletTemplateTypeNemaDash14Dash50r string = "nema-14-50r"

	// WritablePowerOutletTemplateTypeNemaDash14Dash60r captures enum value "nema-14-60r"
	WritablePowerOutletTemplateTypeNemaDash14Dash60r string = "nema-14-60r"

	// WritablePowerOutletTemplateTypeNemaDash15Dash15r captures enum value "nema-15-15r"
	WritablePowerOutletTemplateTypeNemaDash15Dash15r string = "nema-15-15r"

	// WritablePowerOutletTemplateTypeNemaDash15Dash20r captures enum value "nema-15-20r"
	WritablePowerOutletTemplateTypeNemaDash15Dash20r string = "nema-15-20r"

	// WritablePowerOutletTemplateTypeNemaDash15Dash30r captures enum value "nema-15-30r"
	WritablePowerOutletTemplateTypeNemaDash15Dash30r string = "nema-15-30r"

	// WritablePowerOutletTemplateTypeNemaDash15Dash50r captures enum value "nema-15-50r"
	WritablePowerOutletTemplateTypeNemaDash15Dash50r string = "nema-15-50r"

	// WritablePowerOutletTemplateTypeNemaDash15Dash60r captures enum value "nema-15-60r"
	WritablePowerOutletTemplateTypeNemaDash15Dash60r string = "nema-15-60r"

	// WritablePowerOutletTemplateTypeNemaDashL1Dash15r captures enum value "nema-l1-15r"
	WritablePowerOutletTemplateTypeNemaDashL1Dash15r string = "nema-l1-15r"

	// WritablePowerOutletTemplateTypeNemaDashL5Dash15r captures enum value "nema-l5-15r"
	WritablePowerOutletTemplateTypeNemaDashL5Dash15r string = "nema-l5-15r"

	// WritablePowerOutletTemplateTypeNemaDashL5Dash20r captures enum value "nema-l5-20r"
	WritablePowerOutletTemplateTypeNemaDashL5Dash20r string = "nema-l5-20r"

	// WritablePowerOutletTemplateTypeNemaDashL5Dash30r captures enum value "nema-l5-30r"
	WritablePowerOutletTemplateTypeNemaDashL5Dash30r string = "nema-l5-30r"

	// WritablePowerOutletTemplateTypeNemaDashL5Dash50r captures enum value "nema-l5-50r"
	WritablePowerOutletTemplateTypeNemaDashL5Dash50r string = "nema-l5-50r"

	// WritablePowerOutletTemplateTypeNemaDashL6Dash15r captures enum value "nema-l6-15r"
	WritablePowerOutletTemplateTypeNemaDashL6Dash15r string = "nema-l6-15r"

	// WritablePowerOutletTemplateTypeNemaDashL6Dash20r captures enum value "nema-l6-20r"
	WritablePowerOutletTemplateTypeNemaDashL6Dash20r string = "nema-l6-20r"

	// WritablePowerOutletTemplateTypeNemaDashL6Dash30r captures enum value "nema-l6-30r"
	WritablePowerOutletTemplateTypeNemaDashL6Dash30r string = "nema-l6-30r"

	// WritablePowerOutletTemplateTypeNemaDashL6Dash50r captures enum value "nema-l6-50r"
	WritablePowerOutletTemplateTypeNemaDashL6Dash50r string = "nema-l6-50r"

	// WritablePowerOutletTemplateTypeNemaDashL10Dash30r captures enum value "nema-l10-30r"
	WritablePowerOutletTemplateTypeNemaDashL10Dash30r string = "nema-l10-30r"

	// WritablePowerOutletTemplateTypeNemaDashL14Dash20r captures enum value "nema-l14-20r"
	WritablePowerOutletTemplateTypeNemaDashL14Dash20r string = "nema-l14-20r"

	// WritablePowerOutletTemplateTypeNemaDashL14Dash30r captures enum value "nema-l14-30r"
	WritablePowerOutletTemplateTypeNemaDashL14Dash30r string = "nema-l14-30r"

	// WritablePowerOutletTemplateTypeNemaDashL14Dash50r captures enum value "nema-l14-50r"
	WritablePowerOutletTemplateTypeNemaDashL14Dash50r string = "nema-l14-50r"

	// WritablePowerOutletTemplateTypeNemaDashL14Dash60r captures enum value "nema-l14-60r"
	WritablePowerOutletTemplateTypeNemaDashL14Dash60r string = "nema-l14-60r"

	// WritablePowerOutletTemplateTypeNemaDashL15Dash20r captures enum value "nema-l15-20r"
	WritablePowerOutletTemplateTypeNemaDashL15Dash20r string = "nema-l15-20r"

	// WritablePowerOutletTemplateTypeNemaDashL15Dash30r captures enum value "nema-l15-30r"
	WritablePowerOutletTemplateTypeNemaDashL15Dash30r string = "nema-l15-30r"

	// WritablePowerOutletTemplateTypeNemaDashL15Dash50r captures enum value "nema-l15-50r"
	WritablePowerOutletTemplateTypeNemaDashL15Dash50r string = "nema-l15-50r"

	// WritablePowerOutletTemplateTypeNemaDashL15Dash60r captures enum value "nema-l15-60r"
	WritablePowerOutletTemplateTypeNemaDashL15Dash60r string = "nema-l15-60r"

	// WritablePowerOutletTemplateTypeNemaDashL21Dash20r captures enum value "nema-l21-20r"
	WritablePowerOutletTemplateTypeNemaDashL21Dash20r string = "nema-l21-20r"

	// WritablePowerOutletTemplateTypeNemaDashL21Dash30r captures enum value "nema-l21-30r"
	WritablePowerOutletTemplateTypeNemaDashL21Dash30r string = "nema-l21-30r"

	// WritablePowerOutletTemplateTypeNemaDashL22Dash30r captures enum value "nema-l22-30r"
	WritablePowerOutletTemplateTypeNemaDashL22Dash30r string = "nema-l22-30r"

	// WritablePowerOutletTemplateTypeCS6360C captures enum value "CS6360C"
	WritablePowerOutletTemplateTypeCS6360C string = "CS6360C"

	// WritablePowerOutletTemplateTypeCS6364C captures enum value "CS6364C"
	WritablePowerOutletTemplateTypeCS6364C string = "CS6364C"

	// WritablePowerOutletTemplateTypeCS8164C captures enum value "CS8164C"
	WritablePowerOutletTemplateTypeCS8164C string = "CS8164C"

	// WritablePowerOutletTemplateTypeCS8264C captures enum value "CS8264C"
	WritablePowerOutletTemplateTypeCS8264C string = "CS8264C"

	// WritablePowerOutletTemplateTypeCS8364C captures enum value "CS8364C"
	WritablePowerOutletTemplateTypeCS8364C string = "CS8364C"

	// WritablePowerOutletTemplateTypeCS8464C captures enum value "CS8464C"
	WritablePowerOutletTemplateTypeCS8464C string = "CS8464C"

	// WritablePowerOutletTemplateTypeItaDashe captures enum value "ita-e"
	WritablePowerOutletTemplateTypeItaDashe string = "ita-e"

	// WritablePowerOutletTemplateTypeItaDashf captures enum value "ita-f"
	WritablePowerOutletTemplateTypeItaDashf string = "ita-f"

	// WritablePowerOutletTemplateTypeItaDashg captures enum value "ita-g"
	WritablePowerOutletTemplateTypeItaDashg string = "ita-g"

	// WritablePowerOutletTemplateTypeItaDashh captures enum value "ita-h"
	WritablePowerOutletTemplateTypeItaDashh string = "ita-h"

	// WritablePowerOutletTemplateTypeItaDashi captures enum value "ita-i"
	WritablePowerOutletTemplateTypeItaDashi string = "ita-i"

	// WritablePowerOutletTemplateTypeItaDashj captures enum value "ita-j"
	WritablePowerOutletTemplateTypeItaDashj string = "ita-j"

	// WritablePowerOutletTemplateTypeItaDashk captures enum value "ita-k"
	WritablePowerOutletTemplateTypeItaDashk string = "ita-k"

	// WritablePowerOutletTemplateTypeItaDashl captures enum value "ita-l"
	WritablePowerOutletTemplateTypeItaDashl string = "ita-l"

	// WritablePowerOutletTemplateTypeItaDashm captures enum value "ita-m"
	WritablePowerOutletTemplateTypeItaDashm string = "ita-m"

	// WritablePowerOutletTemplateTypeItaDashn captures enum value "ita-n"
	WritablePowerOutletTemplateTypeItaDashn string = "ita-n"

	// WritablePowerOutletTemplateTypeItaDasho captures enum value "ita-o"
	WritablePowerOutletTemplateTypeItaDasho string = "ita-o"

	// WritablePowerOutletTemplateTypeItaDashMultistandard captures enum value "ita-multistandard"
	WritablePowerOutletTemplateTypeItaDashMultistandard string = "ita-multistandard"

	// WritablePowerOutletTemplateTypeUsbDasha captures enum value "usb-a"
	WritablePowerOutletTemplateTypeUsbDasha string = "usb-a"

	// WritablePowerOutletTemplateTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritablePowerOutletTemplateTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritablePowerOutletTemplateTypeUsbDashc captures enum value "usb-c"
	WritablePowerOutletTemplateTypeUsbDashc string = "usb-c"

	// WritablePowerOutletTemplateTypeDcDashTerminal captures enum value "dc-terminal"
	WritablePowerOutletTemplateTypeDcDashTerminal string = "dc-terminal"

	// WritablePowerOutletTemplateTypeHdotDashCx captures enum value "hdot-cx"
	WritablePowerOutletTemplateTypeHdotDashCx string = "hdot-cx"

	// WritablePowerOutletTemplateTypeSafDashdDashGrid captures enum value "saf-d-grid"
	WritablePowerOutletTemplateTypeSafDashdDashGrid string = "saf-d-grid"

	// WritablePowerOutletTemplateTypeNeutrikDashPowerconDash20a captures enum value "neutrik-powercon-20a"
	WritablePowerOutletTemplateTypeNeutrikDashPowerconDash20a string = "neutrik-powercon-20a"

	// WritablePowerOutletTemplateTypeNeutrikDashPowerconDash32a captures enum value "neutrik-powercon-32a"
	WritablePowerOutletTemplateTypeNeutrikDashPowerconDash32a string = "neutrik-powercon-32a"

	// WritablePowerOutletTemplateTypeNeutrikDashPowerconDashTrue1 captures enum value "neutrik-powercon-true1"
	WritablePowerOutletTemplateTypeNeutrikDashPowerconDashTrue1 string = "neutrik-powercon-true1"

	// WritablePowerOutletTemplateTypeNeutrikDashPowerconDashTrue1DashTop captures enum value "neutrik-powercon-true1-top"
	WritablePowerOutletTemplateTypeNeutrikDashPowerconDashTrue1DashTop string = "neutrik-powercon-true1-top"

	// WritablePowerOutletTemplateTypeUbiquitiDashSmartpower captures enum value "ubiquiti-smartpower"
	WritablePowerOutletTemplateTypeUbiquitiDashSmartpower string = "ubiquiti-smartpower"

	// WritablePowerOutletTemplateTypeHardwired captures enum value "hardwired"
	WritablePowerOutletTemplateTypeHardwired string = "hardwired"

	// WritablePowerOutletTemplateTypeOther captures enum value "other"
	WritablePowerOutletTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritablePowerOutletTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerOutletTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerOutletTemplate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power outlet template based on the context it is used
func (m *WritablePowerOutletTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *WritablePowerOutletTemplate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerOutletTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerOutletTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerOutletTemplate) UnmarshalBinary(b []byte) error {
	var res WritablePowerOutletTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
