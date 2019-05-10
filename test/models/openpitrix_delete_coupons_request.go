// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteCouponsRequest openpitrix delete coupons request
// swagger:model openpitrixDeleteCouponsRequest
type OpenpitrixDeleteCouponsRequest struct {

	// coupon ids
	CouponIds []string `json:"coupon_ids"`
}

// Validate validates this openpitrix delete coupons request
func (m *OpenpitrixDeleteCouponsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCouponIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteCouponsRequest) validateCouponIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CouponIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteCouponsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteCouponsRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteCouponsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
