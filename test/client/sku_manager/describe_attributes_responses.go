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

// DescribeAttributesReader is a Reader for the DescribeAttributes structure.
type DescribeAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAttributesOK creates a DescribeAttributesOK with default headers values
func NewDescribeAttributesOK() *DescribeAttributesOK {
	return &DescribeAttributesOK{}
}

/*DescribeAttributesOK handles this case with default header values.

A successful response.
*/
type DescribeAttributesOK struct {
	Payload *models.OpenpitrixDescribeAttributesResponse
}

func (o *DescribeAttributesOK) Error() string {
	return fmt.Sprintf("[GET /v1/metering/attributes][%d] describeAttributesOK  %+v", 200, o.Payload)
}

func (o *DescribeAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAttributesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}