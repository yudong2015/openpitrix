// Code generated by go-swagger; DO NOT EDIT.

package promotion_sku_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// CreateCombinationSkuReader is a Reader for the CreateCombinationSku structure.
type CreateCombinationSkuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCombinationSkuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCombinationSkuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCombinationSkuOK creates a CreateCombinationSkuOK with default headers values
func NewCreateCombinationSkuOK() *CreateCombinationSkuOK {
	return &CreateCombinationSkuOK{}
}

/*CreateCombinationSkuOK handles this case with default header values.

A successful response.
*/
type CreateCombinationSkuOK struct {
	Payload *models.OpenpitrixCreateCombinationSkuResponse
}

func (o *CreateCombinationSkuOK) Error() string {
	return fmt.Sprintf("[POST /v1/metering/combination/sku][%d] createCombinationSkuOK  %+v", 200, o.Payload)
}

func (o *CreateCombinationSkuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixCreateCombinationSkuResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
