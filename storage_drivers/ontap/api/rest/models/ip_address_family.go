// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IPAddressFamily IPv4 or IPv6
//
// swagger:model ip_address_family
type IPAddressFamily string

const (

	// IPAddressFamilyIPV4 captures enum value "ipv4"
	IPAddressFamilyIPV4 IPAddressFamily = "ipv4"

	// IPAddressFamilyIPV6 captures enum value "ipv6"
	IPAddressFamilyIPV6 IPAddressFamily = "ipv6"
)

// for schema
var ipAddressFamilyEnum []interface{}

func init() {
	var res []IPAddressFamily
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressFamilyEnum = append(ipAddressFamilyEnum, v)
	}
}

func (m IPAddressFamily) validateIPAddressFamilyEnum(path, location string, value IPAddressFamily) error {
	if err := validate.EnumCase(path, location, value, ipAddressFamilyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ip address family
func (m IPAddressFamily) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIPAddressFamilyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this ip address family based on the context it is used
func (m IPAddressFamily) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", IPAddressFamily(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// HELLO RIPPY
