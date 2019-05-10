// Code generated by go-swagger; DO NOT EDIT.

package billing_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeRechargesReader is a Reader for the DescribeRecharges structure.
type DescribeRechargesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeRechargesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeRechargesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeRechargesOK creates a DescribeRechargesOK with default headers values
func NewDescribeRechargesOK() *DescribeRechargesOK {
	return &DescribeRechargesOK{}
}

/*DescribeRechargesOK handles this case with default header values.

A successful response.
*/
type DescribeRechargesOK struct {
	Payload *models.OpenpitrixDescribeRechargesResponse
}

func (o *DescribeRechargesOK) Error() string {
	return fmt.Sprintf("[GET /v1/billing/recharges][%d] describeRechargesOK  %+v", 200, o.Payload)
}

func (o *DescribeRechargesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeRechargesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
