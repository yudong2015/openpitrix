// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyProbationResponse openpitrix modify probation response
// swagger:model openpitrixModifyProbationResponse
type OpenpitrixModifyProbationResponse struct {

	// probation id
	ProbationID string `json:"probation_id,omitempty"`
}

// Validate validates this openpitrix modify probation response
func (m *OpenpitrixModifyProbationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyProbationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyProbationResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyProbationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}