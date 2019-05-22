// Code generated by go-swagger; DO NOT EDIT.

package metering_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewStopMeteringsParams creates a new StopMeteringsParams object
// with the default values initialized.
func NewStopMeteringsParams() *StopMeteringsParams {
	var ()
	return &StopMeteringsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopMeteringsParamsWithTimeout creates a new StopMeteringsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopMeteringsParamsWithTimeout(timeout time.Duration) *StopMeteringsParams {
	var ()
	return &StopMeteringsParams{

		timeout: timeout,
	}
}

// NewStopMeteringsParamsWithContext creates a new StopMeteringsParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopMeteringsParamsWithContext(ctx context.Context) *StopMeteringsParams {
	var ()
	return &StopMeteringsParams{

		Context: ctx,
	}
}

// NewStopMeteringsParamsWithHTTPClient creates a new StopMeteringsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopMeteringsParamsWithHTTPClient(client *http.Client) *StopMeteringsParams {
	var ()
	return &StopMeteringsParams{
		HTTPClient: client,
	}
}

/*StopMeteringsParams contains all the parameters to send to the API endpoint
for the stop meterings operation typically these are written to a http.Request
*/
type StopMeteringsParams struct {

	/*Body*/
	Body *models.OpenpitrixStopMeteringsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop meterings params
func (o *StopMeteringsParams) WithTimeout(timeout time.Duration) *StopMeteringsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop meterings params
func (o *StopMeteringsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop meterings params
func (o *StopMeteringsParams) WithContext(ctx context.Context) *StopMeteringsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop meterings params
func (o *StopMeteringsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop meterings params
func (o *StopMeteringsParams) WithHTTPClient(client *http.Client) *StopMeteringsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop meterings params
func (o *StopMeteringsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stop meterings params
func (o *StopMeteringsParams) WithBody(body *models.OpenpitrixStopMeteringsRequest) *StopMeteringsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stop meterings params
func (o *StopMeteringsParams) SetBody(body *models.OpenpitrixStopMeteringsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StopMeteringsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}