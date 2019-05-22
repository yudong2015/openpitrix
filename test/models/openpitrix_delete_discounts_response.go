// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteDiscountsResponse openpitrix delete discounts response
// swagger:model openpitrixDeleteDiscountsResponse
type OpenpitrixDeleteDiscountsResponse struct {

	// discount ids
	DiscountIds []string `json:"discount_ids"`
}

// Validate validates this openpitrix delete discounts response
func (m *OpenpitrixDeleteDiscountsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteDiscountsResponse) validateDiscountIds(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteDiscountsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteDiscountsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteDiscountsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}