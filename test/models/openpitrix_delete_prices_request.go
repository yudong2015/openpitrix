// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeletePricesRequest openpitrix delete prices request
// swagger:model openpitrixDeletePricesRequest
type OpenpitrixDeletePricesRequest struct {

	// price ids
	PriceIds []string `json:"price_ids"`

	// sku id
	SkuID string `json:"sku_id,omitempty"`
}

// Validate validates this openpitrix delete prices request
func (m *OpenpitrixDeletePricesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeletePricesRequest) validatePriceIds(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeletePricesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeletePricesRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeletePricesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}