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

// DeleteCombinationPricesReader is a Reader for the DeleteCombinationPrices structure.
type DeleteCombinationPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCombinationPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCombinationPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCombinationPricesOK creates a DeleteCombinationPricesOK with default headers values
func NewDeleteCombinationPricesOK() *DeleteCombinationPricesOK {
	return &DeleteCombinationPricesOK{}
}

/*DeleteCombinationPricesOK handles this case with default header values.

A successful response.
*/
type DeleteCombinationPricesOK struct {
	Payload *models.OpenpitrixDeleteCombinationPricesResponse
}

func (o *DeleteCombinationPricesOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/billing/combination/prices][%d] deleteCombinationPricesOK  %+v", 200, o.Payload)
}

func (o *DeleteCombinationPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteCombinationPricesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}