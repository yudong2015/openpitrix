// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixTerminateMeteringsRequest openpitrix terminate meterings request
// swagger:model openpitrixTerminateMeteringsRequest
type OpenpitrixTerminateMeteringsRequest struct {

	// resource id
	ResourceID string `json:"resource_id,omitempty"`

	// sku ids
	SkuIds []string `json:"sku_ids"`
}

// Validate validates this openpitrix terminate meterings request
func (m *OpenpitrixTerminateMeteringsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSkuIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixTerminateMeteringsRequest) validateSkuIds(formats strfmt.Registry) error {

	if swag.IsZero(m.SkuIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixTerminateMeteringsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixTerminateMeteringsRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixTerminateMeteringsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
