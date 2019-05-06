// Code generated by go-swagger; DO NOT EDIT.

package service_config

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

// NewValidateEmailServiceParams creates a new ValidateEmailServiceParams object
// with the default values initialized.
func NewValidateEmailServiceParams() *ValidateEmailServiceParams {
	var ()
	return &ValidateEmailServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateEmailServiceParamsWithTimeout creates a new ValidateEmailServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateEmailServiceParamsWithTimeout(timeout time.Duration) *ValidateEmailServiceParams {
	var ()
	return &ValidateEmailServiceParams{

		timeout: timeout,
	}
}

// NewValidateEmailServiceParamsWithContext creates a new ValidateEmailServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateEmailServiceParamsWithContext(ctx context.Context) *ValidateEmailServiceParams {
	var ()
	return &ValidateEmailServiceParams{

		Context: ctx,
	}
}

// NewValidateEmailServiceParamsWithHTTPClient creates a new ValidateEmailServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateEmailServiceParamsWithHTTPClient(client *http.Client) *ValidateEmailServiceParams {
	var ()
	return &ValidateEmailServiceParams{
		HTTPClient: client,
	}
}

/*ValidateEmailServiceParams contains all the parameters to send to the API endpoint
for the validate email service operation typically these are written to a http.Request
*/
type ValidateEmailServiceParams struct {

	/*Body*/
	Body *models.OpenpitrixValidateEmailServiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate email service params
func (o *ValidateEmailServiceParams) WithTimeout(timeout time.Duration) *ValidateEmailServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate email service params
func (o *ValidateEmailServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate email service params
func (o *ValidateEmailServiceParams) WithContext(ctx context.Context) *ValidateEmailServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate email service params
func (o *ValidateEmailServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate email service params
func (o *ValidateEmailServiceParams) WithHTTPClient(client *http.Client) *ValidateEmailServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate email service params
func (o *ValidateEmailServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate email service params
func (o *ValidateEmailServiceParams) WithBody(body *models.OpenpitrixValidateEmailServiceRequest) *ValidateEmailServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate email service params
func (o *ValidateEmailServiceParams) SetBody(body *models.OpenpitrixValidateEmailServiceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateEmailServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
