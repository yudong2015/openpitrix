// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteCombinationSkusRequest openpitrix delete combination skus request
// swagger:model openpitrixDeleteCombinationSkusRequest
type OpenpitrixDeleteCombinationSkusRequest struct {

	// combination sku ids
	CombinationSkuIds []string `json:"combination_sku_ids"`
}

// Validate validates this openpitrix delete combination skus request
func (m *OpenpitrixDeleteCombinationSkusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCombinationSkuIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteCombinationSkusRequest) validateCombinationSkuIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CombinationSkuIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteCombinationSkusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteCombinationSkusRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteCombinationSkusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
