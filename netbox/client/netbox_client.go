// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/scholdan/go-netbox/tree/master/netbox/client/circuits"
	"github.com/scholdan/go-netbox/netbox/client/dcim"
	"github.com/scholdan/go-netbox/netbox/client/extras"
	"github.com/scholdan/go-netbox/netbox/client/ipam"
	"github.com/scholdan/go-netbox/netbox/client/operations"
	"github.com/scholdan/go-netbox/netbox/client/status"
	"github.com/scholdan/go-netbox/netbox/client/tenancy"
	"github.com/scholdan/go-netbox/netbox/client/users"
	"github.com/scholdan/go-netbox/netbox/client/virtualization"
	"github.com/scholdan/go-netbox/netbox/client/vpn"
	"github.com/scholdan/go-netbox/netbox/client/wireless"
)

// Default netbox HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8001"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new netbox HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Netbox {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new netbox HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Netbox {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new netbox client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Netbox {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Netbox)
	cli.Transport = transport
	cli.Circuits = circuits.New(transport, formats)
	cli.Dcim = dcim.New(transport, formats)
	cli.Extras = extras.New(transport, formats)
	cli.Ipam = ipam.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Status = status.New(transport, formats)
	cli.Tenancy = tenancy.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Virtualization = virtualization.New(transport, formats)
	cli.Vpn = vpn.New(transport, formats)
	cli.Wireless = wireless.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Netbox is a client for netbox
type Netbox struct {
	Circuits circuits.ClientService

	Dcim dcim.ClientService

	Extras extras.ClientService

	Ipam ipam.ClientService

	Operations operations.ClientService

	Status status.ClientService

	Tenancy tenancy.ClientService

	Users users.ClientService

	Virtualization virtualization.ClientService

	Vpn vpn.ClientService

	Wireless wireless.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Netbox) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Circuits.SetTransport(transport)
	c.Dcim.SetTransport(transport)
	c.Extras.SetTransport(transport)
	c.Ipam.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Status.SetTransport(transport)
	c.Tenancy.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Virtualization.SetTransport(transport)
	c.Vpn.SetTransport(transport)
	c.Wireless.SetTransport(transport)
}
