// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeDebugAppClustersReader is a Reader for the DescribeDebugAppClusters structure.
type DescribeDebugAppClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDebugAppClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeDebugAppClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeDebugAppClustersOK creates a DescribeDebugAppClustersOK with default headers values
func NewDescribeDebugAppClustersOK() *DescribeDebugAppClustersOK {
	return &DescribeDebugAppClustersOK{}
}

/*DescribeDebugAppClustersOK handles this case with default header values.

A successful response.
*/
type DescribeDebugAppClustersOK struct {
	Payload *models.OpenpitrixDescribeAppClustersResponse
}

func (o *DescribeDebugAppClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/debug_clusters/apps][%d] describeDebugAppClustersOK  %+v", 200, o.Payload)
}

func (o *DescribeDebugAppClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAppClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
