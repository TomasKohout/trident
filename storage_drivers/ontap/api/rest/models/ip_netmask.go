// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPNetmask Input as netmask length (16) or IPv4 mask (255.255.0.0). For IPv6, you must set the netmask length. The default value is 64. Output is always netmask length.
// Example: 24
//
// swagger:model ip_netmask
type IPNetmask string

// Validate validates this ip netmask
func (m IPNetmask) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ip netmask based on context it is used
func (m IPNetmask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// HELLO RIPPY
