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

// WritableInterfaceTemplate writable interface template
//
// swagger:model WritableInterfaceTemplate
type WritableInterfaceTemplate struct {

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

	// Management only
	MgmtOnly bool `json:"mgmt_only,omitempty"`

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

	// PoE mode
	// Enum: ["pd","pse"]
	PoeMode string `json:"poe_mode,omitempty"`

	// PoE type
	// Enum: ["type1-ieee802.3af","type2-ieee802.3at","type2-ieee802.3az","type3-ieee802.3bt","type4-ieee802.3bt","passive-24v-2pair","passive-24v-4pair","passive-48v-2pair","passive-48v-4pair"]
	PoeType string `json:"poe_type,omitempty"`

	// Type
	// Required: true
	// Enum: ["virtual","bridge","lag","100base-fx","100base-lfx","100base-tx","100base-t1","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","50gbase-x-sfp56","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","800gbase-x-qsfpdd","800gbase-x-osfp","1000base-kx","10gbase-kr","10gbase-kx4","25gbase-kr","40gbase-kr4","50gbase-kr","100gbase-kp4","100gbase-kr2","100gbase-kr4","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","ieee802.11ay","ieee802.15.1","other-wireless","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","64gfc-qsfpp","128gfc-qsfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","xdsl","docsis","gpon","xg-pon","xgs-pon","ng-pon2","epon","10g-epon","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","cisco-stackwise-80","cisco-stackwise-160","cisco-stackwise-320","cisco-stackwise-480","cisco-stackwise-1t","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable interface template
func (m *WritableInterfaceTemplate) Validate(formats strfmt.Registry) error {
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

	if err := m.validatePoeMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoeType(formats); err != nil {
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

func (m *WritableInterfaceTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateName(formats strfmt.Registry) error {

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

var writableInterfaceTemplateTypePoeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pd","pse"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTemplateTypePoeModePropEnum = append(writableInterfaceTemplateTypePoeModePropEnum, v)
	}
}

const (

	// WritableInterfaceTemplatePoeModePd captures enum value "pd"
	WritableInterfaceTemplatePoeModePd string = "pd"

	// WritableInterfaceTemplatePoeModePse captures enum value "pse"
	WritableInterfaceTemplatePoeModePse string = "pse"
)

// prop value enum
func (m *WritableInterfaceTemplate) validatePoeModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableInterfaceTemplateTypePoeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterfaceTemplate) validatePoeMode(formats strfmt.Registry) error {
	if swag.IsZero(m.PoeMode) { // not required
		return nil
	}

	// value enum
	if err := m.validatePoeModeEnum("poe_mode", "body", m.PoeMode); err != nil {
		return err
	}

	return nil
}

var writableInterfaceTemplateTypePoeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["type1-ieee802.3af","type2-ieee802.3at","type2-ieee802.3az","type3-ieee802.3bt","type4-ieee802.3bt","passive-24v-2pair","passive-24v-4pair","passive-48v-2pair","passive-48v-4pair"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTemplateTypePoeTypePropEnum = append(writableInterfaceTemplateTypePoeTypePropEnum, v)
	}
}

const (

	// WritableInterfaceTemplatePoeTypeType1DashIeee802Dot3af captures enum value "type1-ieee802.3af"
	WritableInterfaceTemplatePoeTypeType1DashIeee802Dot3af string = "type1-ieee802.3af"

	// WritableInterfaceTemplatePoeTypeType2DashIeee802Dot3at captures enum value "type2-ieee802.3at"
	WritableInterfaceTemplatePoeTypeType2DashIeee802Dot3at string = "type2-ieee802.3at"

	// WritableInterfaceTemplatePoeTypeType2DashIeee802Dot3az captures enum value "type2-ieee802.3az"
	WritableInterfaceTemplatePoeTypeType2DashIeee802Dot3az string = "type2-ieee802.3az"

	// WritableInterfaceTemplatePoeTypeType3DashIeee802Dot3bt captures enum value "type3-ieee802.3bt"
	WritableInterfaceTemplatePoeTypeType3DashIeee802Dot3bt string = "type3-ieee802.3bt"

	// WritableInterfaceTemplatePoeTypeType4DashIeee802Dot3bt captures enum value "type4-ieee802.3bt"
	WritableInterfaceTemplatePoeTypeType4DashIeee802Dot3bt string = "type4-ieee802.3bt"

	// WritableInterfaceTemplatePoeTypePassiveDash24vDash2pair captures enum value "passive-24v-2pair"
	WritableInterfaceTemplatePoeTypePassiveDash24vDash2pair string = "passive-24v-2pair"

	// WritableInterfaceTemplatePoeTypePassiveDash24vDash4pair captures enum value "passive-24v-4pair"
	WritableInterfaceTemplatePoeTypePassiveDash24vDash4pair string = "passive-24v-4pair"

	// WritableInterfaceTemplatePoeTypePassiveDash48vDash2pair captures enum value "passive-48v-2pair"
	WritableInterfaceTemplatePoeTypePassiveDash48vDash2pair string = "passive-48v-2pair"

	// WritableInterfaceTemplatePoeTypePassiveDash48vDash4pair captures enum value "passive-48v-4pair"
	WritableInterfaceTemplatePoeTypePassiveDash48vDash4pair string = "passive-48v-4pair"
)

// prop value enum
func (m *WritableInterfaceTemplate) validatePoeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableInterfaceTemplateTypePoeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterfaceTemplate) validatePoeType(formats strfmt.Registry) error {
	if swag.IsZero(m.PoeType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePoeTypeEnum("poe_type", "body", m.PoeType); err != nil {
		return err
	}

	return nil
}

var writableInterfaceTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","bridge","lag","100base-fx","100base-lfx","100base-tx","100base-t1","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","50gbase-x-sfp56","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","800gbase-x-qsfpdd","800gbase-x-osfp","1000base-kx","10gbase-kr","10gbase-kx4","25gbase-kr","40gbase-kr4","50gbase-kr","100gbase-kp4","100gbase-kr2","100gbase-kr4","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","ieee802.11ay","ieee802.15.1","other-wireless","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","64gfc-qsfpp","128gfc-qsfp28","infiniband-sdr","infiniband-ddr","infiniband-qdr","infiniband-fdr10","infiniband-fdr","infiniband-edr","infiniband-hdr","infiniband-ndr","infiniband-xdr","t1","e1","t3","e3","xdsl","docsis","gpon","xg-pon","xgs-pon","ng-pon2","epon","10g-epon","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","cisco-stackwise-80","cisco-stackwise-160","cisco-stackwise-320","cisco-stackwise-480","cisco-stackwise-1t","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTemplateTypeTypePropEnum = append(writableInterfaceTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableInterfaceTemplateTypeVirtual captures enum value "virtual"
	WritableInterfaceTemplateTypeVirtual string = "virtual"

	// WritableInterfaceTemplateTypeBridge captures enum value "bridge"
	WritableInterfaceTemplateTypeBridge string = "bridge"

	// WritableInterfaceTemplateTypeLag captures enum value "lag"
	WritableInterfaceTemplateTypeLag string = "lag"

	// WritableInterfaceTemplateTypeNr100baseDashFx captures enum value "100base-fx"
	WritableInterfaceTemplateTypeNr100baseDashFx string = "100base-fx"

	// WritableInterfaceTemplateTypeNr100baseDashLfx captures enum value "100base-lfx"
	WritableInterfaceTemplateTypeNr100baseDashLfx string = "100base-lfx"

	// WritableInterfaceTemplateTypeNr100baseDashTx captures enum value "100base-tx"
	WritableInterfaceTemplateTypeNr100baseDashTx string = "100base-tx"

	// WritableInterfaceTemplateTypeNr100baseDashT1 captures enum value "100base-t1"
	WritableInterfaceTemplateTypeNr100baseDashT1 string = "100base-t1"

	// WritableInterfaceTemplateTypeNr1000baseDasht captures enum value "1000base-t"
	WritableInterfaceTemplateTypeNr1000baseDasht string = "1000base-t"

	// WritableInterfaceTemplateTypeNr2Dot5gbaseDasht captures enum value "2.5gbase-t"
	WritableInterfaceTemplateTypeNr2Dot5gbaseDasht string = "2.5gbase-t"

	// WritableInterfaceTemplateTypeNr5gbaseDasht captures enum value "5gbase-t"
	WritableInterfaceTemplateTypeNr5gbaseDasht string = "5gbase-t"

	// WritableInterfaceTemplateTypeNr10gbaseDasht captures enum value "10gbase-t"
	WritableInterfaceTemplateTypeNr10gbaseDasht string = "10gbase-t"

	// WritableInterfaceTemplateTypeNr10gbaseDashCx4 captures enum value "10gbase-cx4"
	WritableInterfaceTemplateTypeNr10gbaseDashCx4 string = "10gbase-cx4"

	// WritableInterfaceTemplateTypeNr1000baseDashxDashGbic captures enum value "1000base-x-gbic"
	WritableInterfaceTemplateTypeNr1000baseDashxDashGbic string = "1000base-x-gbic"

	// WritableInterfaceTemplateTypeNr1000baseDashxDashSfp captures enum value "1000base-x-sfp"
	WritableInterfaceTemplateTypeNr1000baseDashxDashSfp string = "1000base-x-sfp"

	// WritableInterfaceTemplateTypeNr10gbaseDashxDashSfpp captures enum value "10gbase-x-sfpp"
	WritableInterfaceTemplateTypeNr10gbaseDashxDashSfpp string = "10gbase-x-sfpp"

	// WritableInterfaceTemplateTypeNr10gbaseDashxDashXfp captures enum value "10gbase-x-xfp"
	WritableInterfaceTemplateTypeNr10gbaseDashxDashXfp string = "10gbase-x-xfp"

	// WritableInterfaceTemplateTypeNr10gbaseDashxDashXenpak captures enum value "10gbase-x-xenpak"
	WritableInterfaceTemplateTypeNr10gbaseDashxDashXenpak string = "10gbase-x-xenpak"

	// WritableInterfaceTemplateTypeNr10gbaseDashxDashX2 captures enum value "10gbase-x-x2"
	WritableInterfaceTemplateTypeNr10gbaseDashxDashX2 string = "10gbase-x-x2"

	// WritableInterfaceTemplateTypeNr25gbaseDashxDashSfp28 captures enum value "25gbase-x-sfp28"
	WritableInterfaceTemplateTypeNr25gbaseDashxDashSfp28 string = "25gbase-x-sfp28"

	// WritableInterfaceTemplateTypeNr50gbaseDashxDashSfp56 captures enum value "50gbase-x-sfp56"
	WritableInterfaceTemplateTypeNr50gbaseDashxDashSfp56 string = "50gbase-x-sfp56"

	// WritableInterfaceTemplateTypeNr40gbaseDashxDashQsfpp captures enum value "40gbase-x-qsfpp"
	WritableInterfaceTemplateTypeNr40gbaseDashxDashQsfpp string = "40gbase-x-qsfpp"

	// WritableInterfaceTemplateTypeNr50gbaseDashxDashSfp28 captures enum value "50gbase-x-sfp28"
	WritableInterfaceTemplateTypeNr50gbaseDashxDashSfp28 string = "50gbase-x-sfp28"

	// WritableInterfaceTemplateTypeNr100gbaseDashxDashCfp captures enum value "100gbase-x-cfp"
	WritableInterfaceTemplateTypeNr100gbaseDashxDashCfp string = "100gbase-x-cfp"

	// WritableInterfaceTemplateTypeNr100gbaseDashxDashCfp2 captures enum value "100gbase-x-cfp2"
	WritableInterfaceTemplateTypeNr100gbaseDashxDashCfp2 string = "100gbase-x-cfp2"

	// WritableInterfaceTemplateTypeNr200gbaseDashxDashCfp2 captures enum value "200gbase-x-cfp2"
	WritableInterfaceTemplateTypeNr200gbaseDashxDashCfp2 string = "200gbase-x-cfp2"

	// WritableInterfaceTemplateTypeNr100gbaseDashxDashCfp4 captures enum value "100gbase-x-cfp4"
	WritableInterfaceTemplateTypeNr100gbaseDashxDashCfp4 string = "100gbase-x-cfp4"

	// WritableInterfaceTemplateTypeNr100gbaseDashxDashCpak captures enum value "100gbase-x-cpak"
	WritableInterfaceTemplateTypeNr100gbaseDashxDashCpak string = "100gbase-x-cpak"

	// WritableInterfaceTemplateTypeNr100gbaseDashxDashQsfp28 captures enum value "100gbase-x-qsfp28"
	WritableInterfaceTemplateTypeNr100gbaseDashxDashQsfp28 string = "100gbase-x-qsfp28"

	// WritableInterfaceTemplateTypeNr200gbaseDashxDashQsfp56 captures enum value "200gbase-x-qsfp56"
	WritableInterfaceTemplateTypeNr200gbaseDashxDashQsfp56 string = "200gbase-x-qsfp56"

	// WritableInterfaceTemplateTypeNr400gbaseDashxDashQsfpdd captures enum value "400gbase-x-qsfpdd"
	WritableInterfaceTemplateTypeNr400gbaseDashxDashQsfpdd string = "400gbase-x-qsfpdd"

	// WritableInterfaceTemplateTypeNr400gbaseDashxDashOsfp captures enum value "400gbase-x-osfp"
	WritableInterfaceTemplateTypeNr400gbaseDashxDashOsfp string = "400gbase-x-osfp"

	// WritableInterfaceTemplateTypeNr800gbaseDashxDashQsfpdd captures enum value "800gbase-x-qsfpdd"
	WritableInterfaceTemplateTypeNr800gbaseDashxDashQsfpdd string = "800gbase-x-qsfpdd"

	// WritableInterfaceTemplateTypeNr800gbaseDashxDashOsfp captures enum value "800gbase-x-osfp"
	WritableInterfaceTemplateTypeNr800gbaseDashxDashOsfp string = "800gbase-x-osfp"

	// WritableInterfaceTemplateTypeNr1000baseDashKx captures enum value "1000base-kx"
	WritableInterfaceTemplateTypeNr1000baseDashKx string = "1000base-kx"

	// WritableInterfaceTemplateTypeNr10gbaseDashKr captures enum value "10gbase-kr"
	WritableInterfaceTemplateTypeNr10gbaseDashKr string = "10gbase-kr"

	// WritableInterfaceTemplateTypeNr10gbaseDashKx4 captures enum value "10gbase-kx4"
	WritableInterfaceTemplateTypeNr10gbaseDashKx4 string = "10gbase-kx4"

	// WritableInterfaceTemplateTypeNr25gbaseDashKr captures enum value "25gbase-kr"
	WritableInterfaceTemplateTypeNr25gbaseDashKr string = "25gbase-kr"

	// WritableInterfaceTemplateTypeNr40gbaseDashKr4 captures enum value "40gbase-kr4"
	WritableInterfaceTemplateTypeNr40gbaseDashKr4 string = "40gbase-kr4"

	// WritableInterfaceTemplateTypeNr50gbaseDashKr captures enum value "50gbase-kr"
	WritableInterfaceTemplateTypeNr50gbaseDashKr string = "50gbase-kr"

	// WritableInterfaceTemplateTypeNr100gbaseDashKp4 captures enum value "100gbase-kp4"
	WritableInterfaceTemplateTypeNr100gbaseDashKp4 string = "100gbase-kp4"

	// WritableInterfaceTemplateTypeNr100gbaseDashKr2 captures enum value "100gbase-kr2"
	WritableInterfaceTemplateTypeNr100gbaseDashKr2 string = "100gbase-kr2"

	// WritableInterfaceTemplateTypeNr100gbaseDashKr4 captures enum value "100gbase-kr4"
	WritableInterfaceTemplateTypeNr100gbaseDashKr4 string = "100gbase-kr4"

	// WritableInterfaceTemplateTypeIeee802Dot11a captures enum value "ieee802.11a"
	WritableInterfaceTemplateTypeIeee802Dot11a string = "ieee802.11a"

	// WritableInterfaceTemplateTypeIeee802Dot11g captures enum value "ieee802.11g"
	WritableInterfaceTemplateTypeIeee802Dot11g string = "ieee802.11g"

	// WritableInterfaceTemplateTypeIeee802Dot11n captures enum value "ieee802.11n"
	WritableInterfaceTemplateTypeIeee802Dot11n string = "ieee802.11n"

	// WritableInterfaceTemplateTypeIeee802Dot11ac captures enum value "ieee802.11ac"
	WritableInterfaceTemplateTypeIeee802Dot11ac string = "ieee802.11ac"

	// WritableInterfaceTemplateTypeIeee802Dot11ad captures enum value "ieee802.11ad"
	WritableInterfaceTemplateTypeIeee802Dot11ad string = "ieee802.11ad"

	// WritableInterfaceTemplateTypeIeee802Dot11ax captures enum value "ieee802.11ax"
	WritableInterfaceTemplateTypeIeee802Dot11ax string = "ieee802.11ax"

	// WritableInterfaceTemplateTypeIeee802Dot11ay captures enum value "ieee802.11ay"
	WritableInterfaceTemplateTypeIeee802Dot11ay string = "ieee802.11ay"

	// WritableInterfaceTemplateTypeIeee802Dot15Dot1 captures enum value "ieee802.15.1"
	WritableInterfaceTemplateTypeIeee802Dot15Dot1 string = "ieee802.15.1"

	// WritableInterfaceTemplateTypeOtherDashWireless captures enum value "other-wireless"
	WritableInterfaceTemplateTypeOtherDashWireless string = "other-wireless"

	// WritableInterfaceTemplateTypeGsm captures enum value "gsm"
	WritableInterfaceTemplateTypeGsm string = "gsm"

	// WritableInterfaceTemplateTypeCdma captures enum value "cdma"
	WritableInterfaceTemplateTypeCdma string = "cdma"

	// WritableInterfaceTemplateTypeLte captures enum value "lte"
	WritableInterfaceTemplateTypeLte string = "lte"

	// WritableInterfaceTemplateTypeSonetDashOc3 captures enum value "sonet-oc3"
	WritableInterfaceTemplateTypeSonetDashOc3 string = "sonet-oc3"

	// WritableInterfaceTemplateTypeSonetDashOc12 captures enum value "sonet-oc12"
	WritableInterfaceTemplateTypeSonetDashOc12 string = "sonet-oc12"

	// WritableInterfaceTemplateTypeSonetDashOc48 captures enum value "sonet-oc48"
	WritableInterfaceTemplateTypeSonetDashOc48 string = "sonet-oc48"

	// WritableInterfaceTemplateTypeSonetDashOc192 captures enum value "sonet-oc192"
	WritableInterfaceTemplateTypeSonetDashOc192 string = "sonet-oc192"

	// WritableInterfaceTemplateTypeSonetDashOc768 captures enum value "sonet-oc768"
	WritableInterfaceTemplateTypeSonetDashOc768 string = "sonet-oc768"

	// WritableInterfaceTemplateTypeSonetDashOc1920 captures enum value "sonet-oc1920"
	WritableInterfaceTemplateTypeSonetDashOc1920 string = "sonet-oc1920"

	// WritableInterfaceTemplateTypeSonetDashOc3840 captures enum value "sonet-oc3840"
	WritableInterfaceTemplateTypeSonetDashOc3840 string = "sonet-oc3840"

	// WritableInterfaceTemplateTypeNr1gfcDashSfp captures enum value "1gfc-sfp"
	WritableInterfaceTemplateTypeNr1gfcDashSfp string = "1gfc-sfp"

	// WritableInterfaceTemplateTypeNr2gfcDashSfp captures enum value "2gfc-sfp"
	WritableInterfaceTemplateTypeNr2gfcDashSfp string = "2gfc-sfp"

	// WritableInterfaceTemplateTypeNr4gfcDashSfp captures enum value "4gfc-sfp"
	WritableInterfaceTemplateTypeNr4gfcDashSfp string = "4gfc-sfp"

	// WritableInterfaceTemplateTypeNr8gfcDashSfpp captures enum value "8gfc-sfpp"
	WritableInterfaceTemplateTypeNr8gfcDashSfpp string = "8gfc-sfpp"

	// WritableInterfaceTemplateTypeNr16gfcDashSfpp captures enum value "16gfc-sfpp"
	WritableInterfaceTemplateTypeNr16gfcDashSfpp string = "16gfc-sfpp"

	// WritableInterfaceTemplateTypeNr32gfcDashSfp28 captures enum value "32gfc-sfp28"
	WritableInterfaceTemplateTypeNr32gfcDashSfp28 string = "32gfc-sfp28"

	// WritableInterfaceTemplateTypeNr64gfcDashQsfpp captures enum value "64gfc-qsfpp"
	WritableInterfaceTemplateTypeNr64gfcDashQsfpp string = "64gfc-qsfpp"

	// WritableInterfaceTemplateTypeNr128gfcDashQsfp28 captures enum value "128gfc-qsfp28"
	WritableInterfaceTemplateTypeNr128gfcDashQsfp28 string = "128gfc-qsfp28"

	// WritableInterfaceTemplateTypeInfinibandDashSdr captures enum value "infiniband-sdr"
	WritableInterfaceTemplateTypeInfinibandDashSdr string = "infiniband-sdr"

	// WritableInterfaceTemplateTypeInfinibandDashDdr captures enum value "infiniband-ddr"
	WritableInterfaceTemplateTypeInfinibandDashDdr string = "infiniband-ddr"

	// WritableInterfaceTemplateTypeInfinibandDashQdr captures enum value "infiniband-qdr"
	WritableInterfaceTemplateTypeInfinibandDashQdr string = "infiniband-qdr"

	// WritableInterfaceTemplateTypeInfinibandDashFdr10 captures enum value "infiniband-fdr10"
	WritableInterfaceTemplateTypeInfinibandDashFdr10 string = "infiniband-fdr10"

	// WritableInterfaceTemplateTypeInfinibandDashFdr captures enum value "infiniband-fdr"
	WritableInterfaceTemplateTypeInfinibandDashFdr string = "infiniband-fdr"

	// WritableInterfaceTemplateTypeInfinibandDashEdr captures enum value "infiniband-edr"
	WritableInterfaceTemplateTypeInfinibandDashEdr string = "infiniband-edr"

	// WritableInterfaceTemplateTypeInfinibandDashHdr captures enum value "infiniband-hdr"
	WritableInterfaceTemplateTypeInfinibandDashHdr string = "infiniband-hdr"

	// WritableInterfaceTemplateTypeInfinibandDashNdr captures enum value "infiniband-ndr"
	WritableInterfaceTemplateTypeInfinibandDashNdr string = "infiniband-ndr"

	// WritableInterfaceTemplateTypeInfinibandDashXdr captures enum value "infiniband-xdr"
	WritableInterfaceTemplateTypeInfinibandDashXdr string = "infiniband-xdr"

	// WritableInterfaceTemplateTypeT1 captures enum value "t1"
	WritableInterfaceTemplateTypeT1 string = "t1"

	// WritableInterfaceTemplateTypeE1 captures enum value "e1"
	WritableInterfaceTemplateTypeE1 string = "e1"

	// WritableInterfaceTemplateTypeT3 captures enum value "t3"
	WritableInterfaceTemplateTypeT3 string = "t3"

	// WritableInterfaceTemplateTypeE3 captures enum value "e3"
	WritableInterfaceTemplateTypeE3 string = "e3"

	// WritableInterfaceTemplateTypeXdsl captures enum value "xdsl"
	WritableInterfaceTemplateTypeXdsl string = "xdsl"

	// WritableInterfaceTemplateTypeDocsis captures enum value "docsis"
	WritableInterfaceTemplateTypeDocsis string = "docsis"

	// WritableInterfaceTemplateTypeGpon captures enum value "gpon"
	WritableInterfaceTemplateTypeGpon string = "gpon"

	// WritableInterfaceTemplateTypeXgDashPon captures enum value "xg-pon"
	WritableInterfaceTemplateTypeXgDashPon string = "xg-pon"

	// WritableInterfaceTemplateTypeXgsDashPon captures enum value "xgs-pon"
	WritableInterfaceTemplateTypeXgsDashPon string = "xgs-pon"

	// WritableInterfaceTemplateTypeNgDashPon2 captures enum value "ng-pon2"
	WritableInterfaceTemplateTypeNgDashPon2 string = "ng-pon2"

	// WritableInterfaceTemplateTypeEpon captures enum value "epon"
	WritableInterfaceTemplateTypeEpon string = "epon"

	// WritableInterfaceTemplateTypeNr10gDashEpon captures enum value "10g-epon"
	WritableInterfaceTemplateTypeNr10gDashEpon string = "10g-epon"

	// WritableInterfaceTemplateTypeCiscoDashStackwise captures enum value "cisco-stackwise"
	WritableInterfaceTemplateTypeCiscoDashStackwise string = "cisco-stackwise"

	// WritableInterfaceTemplateTypeCiscoDashStackwiseDashPlus captures enum value "cisco-stackwise-plus"
	WritableInterfaceTemplateTypeCiscoDashStackwiseDashPlus string = "cisco-stackwise-plus"

	// WritableInterfaceTemplateTypeCiscoDashFlexstack captures enum value "cisco-flexstack"
	WritableInterfaceTemplateTypeCiscoDashFlexstack string = "cisco-flexstack"

	// WritableInterfaceTemplateTypeCiscoDashFlexstackDashPlus captures enum value "cisco-flexstack-plus"
	WritableInterfaceTemplateTypeCiscoDashFlexstackDashPlus string = "cisco-flexstack-plus"

	// WritableInterfaceTemplateTypeCiscoDashStackwiseDash80 captures enum value "cisco-stackwise-80"
	WritableInterfaceTemplateTypeCiscoDashStackwiseDash80 string = "cisco-stackwise-80"

	// WritableInterfaceTemplateTypeCiscoDashStackwiseDash160 captures enum value "cisco-stackwise-160"
	WritableInterfaceTemplateTypeCiscoDashStackwiseDash160 string = "cisco-stackwise-160"

	// WritableInterfaceTemplateTypeCiscoDashStackwiseDash320 captures enum value "cisco-stackwise-320"
	WritableInterfaceTemplateTypeCiscoDashStackwiseDash320 string = "cisco-stackwise-320"

	// WritableInterfaceTemplateTypeCiscoDashStackwiseDash480 captures enum value "cisco-stackwise-480"
	WritableInterfaceTemplateTypeCiscoDashStackwiseDash480 string = "cisco-stackwise-480"

	// WritableInterfaceTemplateTypeCiscoDashStackwiseDash1t captures enum value "cisco-stackwise-1t"
	WritableInterfaceTemplateTypeCiscoDashStackwiseDash1t string = "cisco-stackwise-1t"

	// WritableInterfaceTemplateTypeJuniperDashVcp captures enum value "juniper-vcp"
	WritableInterfaceTemplateTypeJuniperDashVcp string = "juniper-vcp"

	// WritableInterfaceTemplateTypeExtremeDashSummitstack captures enum value "extreme-summitstack"
	WritableInterfaceTemplateTypeExtremeDashSummitstack string = "extreme-summitstack"

	// WritableInterfaceTemplateTypeExtremeDashSummitstackDash128 captures enum value "extreme-summitstack-128"
	WritableInterfaceTemplateTypeExtremeDashSummitstackDash128 string = "extreme-summitstack-128"

	// WritableInterfaceTemplateTypeExtremeDashSummitstackDash256 captures enum value "extreme-summitstack-256"
	WritableInterfaceTemplateTypeExtremeDashSummitstackDash256 string = "extreme-summitstack-256"

	// WritableInterfaceTemplateTypeExtremeDashSummitstackDash512 captures enum value "extreme-summitstack-512"
	WritableInterfaceTemplateTypeExtremeDashSummitstackDash512 string = "extreme-summitstack-512"

	// WritableInterfaceTemplateTypeOther captures enum value "other"
	WritableInterfaceTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritableInterfaceTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableInterfaceTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterfaceTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable interface template based on the context it is used
func (m *WritableInterfaceTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *WritableInterfaceTemplate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableInterfaceTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableInterfaceTemplate) UnmarshalBinary(b []byte) error {
	var res WritableInterfaceTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
