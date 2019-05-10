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

// DescribeRefundsReader is a Reader for the DescribeRefunds structure.
type DescribeRefundsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeRefundsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeRefundsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeRefundsOK creates a DescribeRefundsOK with default headers values
func NewDescribeRefundsOK() *DescribeRefundsOK {
	return &DescribeRefundsOK{}
}

/*DescribeRefundsOK handles this case with default header values.

A successful response.
*/
type DescribeRefundsOK struct {
	Payload *models.OpenpitrixDescribeRefundsResponse
}

func (o *DescribeRefundsOK) Error() string {
	return fmt.Sprintf("[GET /v1/billing/refunds][%d] describeRefundsOK  %+v", 200, o.Payload)
}

func (o *DescribeRefundsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeRefundsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
