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

// NewModifyAccountParams creates a new ModifyAccountParams object
// with the default values initialized.
func NewModifyAccountParams() *ModifyAccountParams {
	var ()
	return &ModifyAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyAccountParamsWithTimeout creates a new ModifyAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyAccountParamsWithTimeout(timeout time.Duration) *ModifyAccountParams {
	var ()
	return &ModifyAccountParams{

		timeout: timeout,
	}
}

// NewModifyAccountParamsWithContext creates a new ModifyAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyAccountParamsWithContext(ctx context.Context) *ModifyAccountParams {
	var ()
	return &ModifyAccountParams{

		Context: ctx,
	}
}

// NewModifyAccountParamsWithHTTPClient creates a new ModifyAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyAccountParamsWithHTTPClient(client *http.Client) *ModifyAccountParams {
	var ()
	return &ModifyAccountParams{
		HTTPClient: client,
	}
}

/*ModifyAccountParams contains all the parameters to send to the API endpoint
for the modify account operation typically these are written to a http.Request
*/
type ModifyAccountParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyAccountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify account params
func (o *ModifyAccountParams) WithTimeout(timeout time.Duration) *ModifyAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify account params
func (o *ModifyAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify account params
func (o *ModifyAccountParams) WithContext(ctx context.Context) *ModifyAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify account params
func (o *ModifyAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify account params
func (o *ModifyAccountParams) WithHTTPClient(client *http.Client) *ModifyAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify account params
func (o *ModifyAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify account params
func (o *ModifyAccountParams) WithBody(body *models.OpenpitrixModifyAccountRequest) *ModifyAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify account params
func (o *ModifyAccountParams) SetBody(body *models.OpenpitrixModifyAccountRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
