// Code generated by go-swagger; DO NOT EDIT.

package sku_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ModifyAttributeReader is a Reader for the ModifyAttribute structure.
type ModifyAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyAttributeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyAttributeOK creates a ModifyAttributeOK with default headers values
func NewModifyAttributeOK() *ModifyAttributeOK {
	return &ModifyAttributeOK{}
}

/*ModifyAttributeOK handles this case with default header values.

A successful response.
*/
type ModifyAttributeOK struct {
	Payload *models.OpenpitrixModifyAttributeResponse
}

func (o *ModifyAttributeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/metering/attribute][%d] modifyAttributeOK  %+v", 200, o.Payload)
}

func (o *ModifyAttributeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyAttributeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
