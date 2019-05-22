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

// DescribeAccountsReader is a Reader for the DescribeAccounts structure.
type DescribeAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAccountsOK creates a DescribeAccountsOK with default headers values
func NewDescribeAccountsOK() *DescribeAccountsOK {
	return &DescribeAccountsOK{}
}

/*DescribeAccountsOK handles this case with default header values.

A successful response.
*/
type DescribeAccountsOK struct {
	Payload *models.OpenpitrixDescribeAccountsResponse
}

func (o *DescribeAccountsOK) Error() string {
	return fmt.Sprintf("[GET /v1/billing/accounts][%d] describeAccountsOK  %+v", 200, o.Payload)
}

func (o *DescribeAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}