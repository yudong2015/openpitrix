// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixUpdateSkuMetering openpitrix update sku metering
// swagger:model openpitrixUpdateSkuMetering
type OpenpitrixUpdateSkuMetering struct {

	// metering values
	MeteringValues OpenpitrixUpdateSkuMeteringMeteringValues `json:"metering_values"`

	// sku id
	SkuID string `json:"sku_id,omitempty"`
}

// Validate validates this openpitrix update sku metering
func (m *OpenpitrixUpdateSkuMetering) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixUpdateSkuMetering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixUpdateSkuMetering) UnmarshalBinary(b []byte) error {
	var res OpenpitrixUpdateSkuMetering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}