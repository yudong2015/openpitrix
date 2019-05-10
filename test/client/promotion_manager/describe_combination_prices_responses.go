// Code generated by go-swagger; DO NOT EDIT.

package promotion_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeCombinationPricesReader is a Reader for the DescribeCombinationPrices structure.
type DescribeCombinationPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeCombinationPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeCombinationPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeCombinationPricesOK creates a DescribeCombinationPricesOK with default headers values
func NewDescribeCombinationPricesOK() *DescribeCombinationPricesOK {
	return &DescribeCombinationPricesOK{}
}

/*DescribeCombinationPricesOK handles this case with default header values.

A successful response.
*/
type DescribeCombinationPricesOK struct {
	Payload *models.OpenpitrixDescribeCombinationPricesResponse
}

func (o *DescribeCombinationPricesOK) Error() string {
	return fmt.Sprintf("[GET /v1/billing/combination/prices][%d] describeCombinationPricesOK  %+v", 200, o.Payload)
}

func (o *DescribeCombinationPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeCombinationPricesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
