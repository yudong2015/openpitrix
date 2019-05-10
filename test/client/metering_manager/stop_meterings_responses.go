// Code generated by go-swagger; DO NOT EDIT.

package metering_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// StopMeteringsReader is a Reader for the StopMeterings structure.
type StopMeteringsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopMeteringsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStopMeteringsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopMeteringsOK creates a StopMeteringsOK with default headers values
func NewStopMeteringsOK() *StopMeteringsOK {
	return &StopMeteringsOK{}
}

/*StopMeteringsOK handles this case with default header values.

A successful response.
*/
type StopMeteringsOK struct {
	Payload *models.OpenpitrixCommonMeteringsResponse
}

func (o *StopMeteringsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/metering/stop][%d] stopMeteringsOK  %+v", 200, o.Payload)
}

func (o *StopMeteringsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixCommonMeteringsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
