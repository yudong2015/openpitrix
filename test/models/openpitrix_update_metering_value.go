// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixUpdateMeteringValue openpitrix update metering value
// swagger:model openpitrixUpdateMeteringValue
type OpenpitrixUpdateMeteringValue struct {

	// attribute id
	AttributeID string `json:"attribute_id,omitempty"`

	// type
	Type UpdateMeteringValueUpdateValueTypes `json:"type,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this openpitrix update metering value
func (m *OpenpitrixUpdateMeteringValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixUpdateMeteringValue) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixUpdateMeteringValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixUpdateMeteringValue) UnmarshalBinary(b []byte) error {
	var res OpenpitrixUpdateMeteringValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
