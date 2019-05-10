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

// CreateCombinationReader is a Reader for the CreateCombination structure.
type CreateCombinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCombinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCombinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCombinationOK creates a CreateCombinationOK with default headers values
func NewCreateCombinationOK() *CreateCombinationOK {
	return &CreateCombinationOK{}
}

/*CreateCombinationOK handles this case with default header values.

A successful response.
*/
type CreateCombinationOK struct {
	Payload *models.OpenpitrixCreateCombinationResponse
}

func (o *CreateCombinationOK) Error() string {
	return fmt.Sprintf("[POST /v1/metering/combination][%d] createCombinationOK  %+v", 200, o.Payload)
}

func (o *CreateCombinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixCreateCombinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
