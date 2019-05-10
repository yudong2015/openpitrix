// Code generated by go-swagger; DO NOT EDIT.

package billing_manager

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

// NewModifyPriceParams creates a new ModifyPriceParams object
// with the default values initialized.
func NewModifyPriceParams() *ModifyPriceParams {
	var ()
	return &ModifyPriceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyPriceParamsWithTimeout creates a new ModifyPriceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyPriceParamsWithTimeout(timeout time.Duration) *ModifyPriceParams {
	var ()
	return &ModifyPriceParams{

		timeout: timeout,
	}
}

// NewModifyPriceParamsWithContext creates a new ModifyPriceParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyPriceParamsWithContext(ctx context.Context) *ModifyPriceParams {
	var ()
	return &ModifyPriceParams{

		Context: ctx,
	}
}

// NewModifyPriceParamsWithHTTPClient creates a new ModifyPriceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyPriceParamsWithHTTPClient(client *http.Client) *ModifyPriceParams {
	var ()
	return &ModifyPriceParams{
		HTTPClient: client,
	}
}

/*ModifyPriceParams contains all the parameters to send to the API endpoint
for the modify price operation typically these are written to a http.Request
*/
type ModifyPriceParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyPriceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify price params
func (o *ModifyPriceParams) WithTimeout(timeout time.Duration) *ModifyPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify price params
func (o *ModifyPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify price params
func (o *ModifyPriceParams) WithContext(ctx context.Context) *ModifyPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify price params
func (o *ModifyPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify price params
func (o *ModifyPriceParams) WithHTTPClient(client *http.Client) *ModifyPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify price params
func (o *ModifyPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify price params
func (o *ModifyPriceParams) WithBody(body *models.OpenpitrixModifyPriceRequest) *ModifyPriceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify price params
func (o *ModifyPriceParams) SetBody(body *models.OpenpitrixModifyPriceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
