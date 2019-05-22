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

// NewDeleteCombinationsParams creates a new DeleteCombinationsParams object
// with the default values initialized.
func NewDeleteCombinationsParams() *DeleteCombinationsParams {
	var ()
	return &DeleteCombinationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCombinationsParamsWithTimeout creates a new DeleteCombinationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCombinationsParamsWithTimeout(timeout time.Duration) *DeleteCombinationsParams {
	var ()
	return &DeleteCombinationsParams{

		timeout: timeout,
	}
}

// NewDeleteCombinationsParamsWithContext creates a new DeleteCombinationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCombinationsParamsWithContext(ctx context.Context) *DeleteCombinationsParams {
	var ()
	return &DeleteCombinationsParams{

		Context: ctx,
	}
}

// NewDeleteCombinationsParamsWithHTTPClient creates a new DeleteCombinationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCombinationsParamsWithHTTPClient(client *http.Client) *DeleteCombinationsParams {
	var ()
	return &DeleteCombinationsParams{
		HTTPClient: client,
	}
}

/*DeleteCombinationsParams contains all the parameters to send to the API endpoint
for the delete combinations operation typically these are written to a http.Request
*/
type DeleteCombinationsParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteCombinationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete combinations params
func (o *DeleteCombinationsParams) WithTimeout(timeout time.Duration) *DeleteCombinationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete combinations params
func (o *DeleteCombinationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete combinations params
func (o *DeleteCombinationsParams) WithContext(ctx context.Context) *DeleteCombinationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete combinations params
func (o *DeleteCombinationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete combinations params
func (o *DeleteCombinationsParams) WithHTTPClient(client *http.Client) *DeleteCombinationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete combinations params
func (o *DeleteCombinationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete combinations params
func (o *DeleteCombinationsParams) WithBody(body *models.OpenpitrixDeleteCombinationsRequest) *DeleteCombinationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete combinations params
func (o *DeleteCombinationsParams) SetBody(body *models.OpenpitrixDeleteCombinationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCombinationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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