// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixProbation probation
// swagger:model openpitrixProbation
type OpenpitrixProbation struct {

	// attribute id
	AttributeID string `json:"attribute_id,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// end time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// probation id
	ProbationID string `json:"probation_id,omitempty"`

	// sku id
	SkuID string `json:"sku_id,omitempty"`

	// start time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`
}

// Validate validates this openpitrix probation
func (m *OpenpitrixProbation) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixProbation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixProbation) UnmarshalBinary(b []byte) error {
	var res OpenpitrixProbation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}