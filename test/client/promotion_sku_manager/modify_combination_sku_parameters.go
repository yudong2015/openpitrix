// Code generated by go-swagger; DO NOT EDIT.

package promotion_sku_manager

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

// NewModifyCombinationSkuParams creates a new ModifyCombinationSkuParams object
// with the default values initialized.
func NewModifyCombinationSkuParams() *ModifyCombinationSkuParams {
	var ()
	return &ModifyCombinationSkuParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyCombinationSkuParamsWithTimeout creates a new ModifyCombinationSkuParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyCombinationSkuParamsWithTimeout(timeout time.Duration) *ModifyCombinationSkuParams {
	var ()
	return &ModifyCombinationSkuParams{

		timeout: timeout,
	}
}

// NewModifyCombinationSkuParamsWithContext creates a new ModifyCombinationSkuParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyCombinationSkuParamsWithContext(ctx context.Context) *ModifyCombinationSkuParams {
	var ()
	return &ModifyCombinationSkuParams{

		Context: ctx,
	}
}

// NewModifyCombinationSkuParamsWithHTTPClient creates a new ModifyCombinationSkuParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyCombinationSkuParamsWithHTTPClient(client *http.Client) *ModifyCombinationSkuParams {
	var ()
	return &ModifyCombinationSkuParams{
		HTTPClient: client,
	}
}

/*ModifyCombinationSkuParams contains all the parameters to send to the API endpoint
for the modify combination sku operation typically these are written to a http.Request
*/
type ModifyCombinationSkuParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyCombinationSkuRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify combination sku params
func (o *ModifyCombinationSkuParams) WithTimeout(timeout time.Duration) *ModifyCombinationSkuParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify combination sku params
func (o *ModifyCombinationSkuParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify combination sku params
func (o *ModifyCombinationSkuParams) WithContext(ctx context.Context) *ModifyCombinationSkuParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify combination sku params
func (o *ModifyCombinationSkuParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify combination sku params
func (o *ModifyCombinationSkuParams) WithHTTPClient(client *http.Client) *ModifyCombinationSkuParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify combination sku params
func (o *ModifyCombinationSkuParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify combination sku params
func (o *ModifyCombinationSkuParams) WithBody(body *models.OpenpitrixModifyCombinationSkuRequest) *ModifyCombinationSkuParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify combination sku params
func (o *ModifyCombinationSkuParams) SetBody(body *models.OpenpitrixModifyCombinationSkuRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyCombinationSkuParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
