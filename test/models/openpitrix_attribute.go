// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAttribute attribute
// swagger:model openpitrixAttribute
type OpenpitrixAttribute struct {

	// attribute id
	AttributeID string `json:"attribute_id,omitempty"`

	// attribute name id
	AttributeNameID string `json:"attribute_name_id,omitempty"`

	// attribute unit id
	AttributeUnitID string `json:"attribute_unit_id,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// the types of value: single int value, scope of value (min_value, max_value], string value
	Value string `json:"value,omitempty"`
}

// Validate validates this openpitrix attribute
func (m *OpenpitrixAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAttribute) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
