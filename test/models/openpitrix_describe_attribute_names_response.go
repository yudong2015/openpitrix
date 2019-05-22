// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeAttributeNamesResponse openpitrix describe attribute names response
// swagger:model openpitrixDescribeAttributeNamesResponse
type OpenpitrixDescribeAttributeNamesResponse struct {

	// attribute name set
	AttributeNameSet OpenpitrixDescribeAttributeNamesResponseAttributeNameSet `json:"attribute_name_set"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe attribute names response
func (m *OpenpitrixDescribeAttributeNamesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeAttributeNamesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeAttributeNamesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeAttributeNamesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}