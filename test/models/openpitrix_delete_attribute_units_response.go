// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteAttributeUnitsResponse openpitrix delete attribute units response
// swagger:model openpitrixDeleteAttributeUnitsResponse
type OpenpitrixDeleteAttributeUnitsResponse struct {

	// attribute unit ids
	AttributeUnitIds []string `json:"attribute_unit_ids"`
}

// Validate validates this openpitrix delete attribute units response
func (m *OpenpitrixDeleteAttributeUnitsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeUnitIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteAttributeUnitsResponse) validateAttributeUnitIds(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeUnitIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteAttributeUnitsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteAttributeUnitsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteAttributeUnitsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}