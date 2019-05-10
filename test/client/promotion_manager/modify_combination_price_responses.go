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

// ModifyCombinationPriceReader is a Reader for the ModifyCombinationPrice structure.
type ModifyCombinationPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyCombinationPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyCombinationPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyCombinationPriceOK creates a ModifyCombinationPriceOK with default headers values
func NewModifyCombinationPriceOK() *ModifyCombinationPriceOK {
	return &ModifyCombinationPriceOK{}
}

/*ModifyCombinationPriceOK handles this case with default header values.

A successful response.
*/
type ModifyCombinationPriceOK struct {
	Payload *models.OpenpitrixModifyCombinationPriceResponse
}

func (o *ModifyCombinationPriceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/billing/combination/price][%d] modifyCombinationPriceOK  %+v", 200, o.Payload)
}

func (o *ModifyCombinationPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyCombinationPriceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
