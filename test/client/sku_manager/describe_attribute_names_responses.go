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

// DescribeAttributeNamesReader is a Reader for the DescribeAttributeNames structure.
type DescribeAttributeNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAttributeNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAttributeNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAttributeNamesOK creates a DescribeAttributeNamesOK with default headers values
func NewDescribeAttributeNamesOK() *DescribeAttributeNamesOK {
	return &DescribeAttributeNamesOK{}
}

/*DescribeAttributeNamesOK handles this case with default header values.

A successful response.
*/
type DescribeAttributeNamesOK struct {
	Payload *models.OpenpitrixDescribeAttributeNamesResponse
}

func (o *DescribeAttributeNamesOK) Error() string {
	return fmt.Sprintf("[GET /v1/metering/attribute/names][%d] describeAttributeNamesOK  %+v", 200, o.Payload)
}

func (o *DescribeAttributeNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAttributeNamesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}