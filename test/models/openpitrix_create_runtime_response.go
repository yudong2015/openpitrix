// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateRuntimeResponse openpitrix create runtime response
// swagger:model openpitrixCreateRuntimeResponse
type OpenpitrixCreateRuntimeResponse struct {

	// id of runtime created
	RuntimeID string `json:"runtime_id,omitempty"`
}

// Validate validates this openpitrix create runtime response
func (m *OpenpitrixCreateRuntimeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateRuntimeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateRuntimeResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateRuntimeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
