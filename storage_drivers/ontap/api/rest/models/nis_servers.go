// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NisServers A list of hostnames or IP addresses of NIS servers used
// by the NIS domain configuration.
//
//
// swagger:model nis_servers
type NisServers []string

// Validate validates this nis servers
func (m NisServers) Validate(formats strfmt.Registry) error {
	var res []error

	iNisServersSize := int64(len(m))

	if err := validate.MaxItems("", "body", iNisServersSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if err := validate.MinLength(strconv.Itoa(i), "body", m[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength(strconv.Itoa(i), "body", m[i], 255); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this nis servers based on context it is used
func (m NisServers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// HELLO RIPPY
