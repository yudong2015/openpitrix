// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeZonesResponse openpitrix describe zones response
// swagger:model openpitrixDescribeZonesResponse
type OpenpitrixDescribeZonesResponse struct {

	// list of zones in runtime provider
	Zones []string `json:"zones"`
}

// Validate validates this openpitrix describe zones response
func (m *OpenpitrixDescribeZonesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZones(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDescribeZonesResponse) validateZones(formats strfmt.Registry) error {

	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeZonesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeZonesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeZonesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
