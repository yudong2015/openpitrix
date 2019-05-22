// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifySkuRequest openpitrix modify sku request
// swagger:model openpitrixModifySkuRequest
type OpenpitrixModifySkuRequest struct {

	// attribute ids
	AttributeIds []string `json:"attribute_ids"`

	// metering attribute ids
	MeteringAttributeIds []string `json:"metering_attribute_ids"`

	// sku id
	SkuID string `json:"sku_id,omitempty"`

	// spu id
	SpuID string `json:"spu_id,omitempty"`
}

// Validate validates this openpitrix modify sku request
func (m *OpenpitrixModifySkuRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeteringAttributeIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixModifySkuRequest) validateAttributeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeIds) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixModifySkuRequest) validateMeteringAttributeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.MeteringAttributeIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifySkuRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifySkuRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifySkuRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}