package models

import (
    "github.com/go-openapi/errors"
    "github.com/go-openapi/strfmt"
    "github.com/go-openapi/validate"
)

// Gateway represents a single gateway object in the API
type Gateway struct {
    ID          int64       `json:"id"`
    URL         string      `json:"url"`
    Display     string      `json:"display"`
    VRF         *int64      `json:"vrf,omitempty"`
    GatewayIP   IPAddress   `json:"gateway_ip"`
    Prefix      Prefix      `json:"prefix"`
    Tags        []string    `json:"tags"`
    CustomFields map[string]interface{} `json:"custom_fields"`
    Created     string      `json:"created"`
    LastUpdated string      `json:"last_updated"`
}

// Validate validates this gateway
func (g *Gateway) Validate(formats strfmt.Registry) error {
    var res []error

    if err := g.validateID(); err != nil {
        res = append(res, err)
    }
    if err := g.GatewayIP.Validate(formats); err != nil {
        res = append(res, err)
    }
    if err := g.Prefix.Validate(formats); err != nil {
        res = append(res, err)
    }
    if err := g.validateCreated(); err != nil {
        res = append(res, err)
    }
    if err := g.validateLastUpdated(); err != nil {
        res = append(res, err)
    }

    if len(res) > 0 {
        return errors.CompositeValidationError(res...)
    }
    return nil
}

func (g *Gateway) validateID() error {
    if g.ID <= 0 {
        return errors.Required("id", "body", "ID is required and must be greater than 0")
    }
    return nil
}

func (g *Gateway) validateCreated() error {
    if g.Created == "" {
        return errors.Required("created", "body", "Created is required")
    }
    return nil
}

func (g *Gateway) validateLastUpdated() error {
    if g.LastUpdated == "" {
        return errors.Required("last_updated", "body", "LastUpdated is required")
    }
    return nil
}

// IPAddress represents the IP address structure
type IPAddress struct {
    ID      int64  `json:"id"`
    URL     string `json:"url"`
    Display string `json:"display"`
    Family  int    `json:"family"`
    Address string `json:"address"`
}

// Validate validates this IP address
func (ip *IPAddress) Validate(formats strfmt.Registry) error {
    var res []error

    if err := ip.validateID(); err != nil {
        res = append(res, err)
    }
    if err := ip.validateAddress(); err != nil {
        res = append(res, err)
    }

    if len(res) > 0 {
        return errors.CompositeValidationError(res...)
    }
    return nil
}

func (ip *IPAddress) validateID() error {
    if ip.ID <= 0 {
        return errors.Required("id", "body", "ID is required and must be greater than 0")
    }
    return nil
}

func (ip *IPAddress) validateAddress() error {
    if ip.Address == "" {
        return errors.Required("address", "body", "Address is required")
    }
    return nil
}

// Prefix represents the prefix structure
type Prefix struct {
    ID     int64  `json:"id"`
    URL    string `json:"url"`
    Display string `json:"display"`
    Family int    `json:"family"`
    Prefix string `json:"prefix"`
    Depth  int    `json:"_depth"`
}

// Validate validates this prefix
func (p *Prefix) Validate(formats strfmt.Registry) error {
    var res []error

    if err := p.validateID(); err != nil {
        res = append(res, err)
    }
    if err := p.validatePrefix(); err != nil {
        res = append(res, err)
    }

    if len(res) > 0 {
        return errors.CompositeValidationError(res...)
    }
    return nil
}

func (p *Prefix) validateID() error {
    if p.ID <= 0 {
        return errors.Required("id", "body", "ID is required and must be greater than 0")
    }
    return nil
}

func (p *Prefix) validatePrefix() error {
    if p.Prefix == "" {
        return errors.Required("prefix", "body", "Prefix is required")
    }
    return nil
}
