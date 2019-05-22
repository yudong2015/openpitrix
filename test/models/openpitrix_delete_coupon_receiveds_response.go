// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteCouponReceivedsResponse openpitrix delete coupon receiveds response
// swagger:model openpitrixDeleteCouponReceivedsResponse
type OpenpitrixDeleteCouponReceivedsResponse struct {

	// coupon received ids
	CouponReceivedIds []string `json:"coupon_received_ids"`
}

// Validate validates this openpitrix delete coupon receiveds response
func (m *OpenpitrixDeleteCouponReceivedsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCouponReceivedIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteCouponReceivedsResponse) validateCouponReceivedIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CouponReceivedIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteCouponReceivedsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteCouponReceivedsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteCouponReceivedsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}