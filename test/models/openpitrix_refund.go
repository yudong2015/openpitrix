// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRefund openpitrix refund
// swagger:model openpitrixRefund
type OpenpitrixRefund struct {

	// refund id
	RefundID string `json:"Refund_id,omitempty"`

	// contract id
	ContractID string `json:"contract_id,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// fee
	Fee float64 `json:"fee,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix refund
func (m *OpenpitrixRefund) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRefund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRefund) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRefund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}