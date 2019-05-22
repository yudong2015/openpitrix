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

// NewModifyCombinationParams creates a new ModifyCombinationParams object
// with the default values initialized.
func NewModifyCombinationParams() *ModifyCombinationParams {
	var ()
	return &ModifyCombinationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyCombinationParamsWithTimeout creates a new ModifyCombinationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyCombinationParamsWithTimeout(timeout time.Duration) *ModifyCombinationParams {
	var ()
	return &ModifyCombinationParams{

		timeout: timeout,
	}
}

// NewModifyCombinationParamsWithContext creates a new ModifyCombinationParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyCombinationParamsWithContext(ctx context.Context) *ModifyCombinationParams {
	var ()
	return &ModifyCombinationParams{

		Context: ctx,
	}
}

// NewModifyCombinationParamsWithHTTPClient creates a new ModifyCombinationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyCombinationParamsWithHTTPClient(client *http.Client) *ModifyCombinationParams {
	var ()
	return &ModifyCombinationParams{
		HTTPClient: client,
	}
}

/*ModifyCombinationParams contains all the parameters to send to the API endpoint
for the modify combination operation typically these are written to a http.Request
*/
type ModifyCombinationParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyCombinationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify combination params
func (o *ModifyCombinationParams) WithTimeout(timeout time.Duration) *ModifyCombinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify combination params
func (o *ModifyCombinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify combination params
func (o *ModifyCombinationParams) WithContext(ctx context.Context) *ModifyCombinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify combination params
func (o *ModifyCombinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify combination params
func (o *ModifyCombinationParams) WithHTTPClient(client *http.Client) *ModifyCombinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify combination params
func (o *ModifyCombinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify combination params
func (o *ModifyCombinationParams) WithBody(body *models.OpenpitrixModifyCombinationRequest) *ModifyCombinationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify combination params
func (o *ModifyCombinationParams) SetBody(body *models.OpenpitrixModifyCombinationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyCombinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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