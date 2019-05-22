// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateCombinationResponse openpitrix create combination response
// swagger:model openpitrixCreateCombinationResponse
type OpenpitrixCreateCombinationResponse struct {

	// combination id
	CombinationID string `json:"combination_id,omitempty"`
}

// Validate validates this openpitrix create combination response
func (m *OpenpitrixCreateCombinationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateCombinationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateCombinationResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateCombinationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}