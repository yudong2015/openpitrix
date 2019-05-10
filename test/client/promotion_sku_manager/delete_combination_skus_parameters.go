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

// NewDeleteCombinationSkusParams creates a new DeleteCombinationSkusParams object
// with the default values initialized.
func NewDeleteCombinationSkusParams() *DeleteCombinationSkusParams {
	var ()
	return &DeleteCombinationSkusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCombinationSkusParamsWithTimeout creates a new DeleteCombinationSkusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCombinationSkusParamsWithTimeout(timeout time.Duration) *DeleteCombinationSkusParams {
	var ()
	return &DeleteCombinationSkusParams{

		timeout: timeout,
	}
}

// NewDeleteCombinationSkusParamsWithContext creates a new DeleteCombinationSkusParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCombinationSkusParamsWithContext(ctx context.Context) *DeleteCombinationSkusParams {
	var ()
	return &DeleteCombinationSkusParams{

		Context: ctx,
	}
}

// NewDeleteCombinationSkusParamsWithHTTPClient creates a new DeleteCombinationSkusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCombinationSkusParamsWithHTTPClient(client *http.Client) *DeleteCombinationSkusParams {
	var ()
	return &DeleteCombinationSkusParams{
		HTTPClient: client,
	}
}

/*DeleteCombinationSkusParams contains all the parameters to send to the API endpoint
for the delete combination skus operation typically these are written to a http.Request
*/
type DeleteCombinationSkusParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteCombinationSkusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete combination skus params
func (o *DeleteCombinationSkusParams) WithTimeout(timeout time.Duration) *DeleteCombinationSkusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete combination skus params
func (o *DeleteCombinationSkusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete combination skus params
func (o *DeleteCombinationSkusParams) WithContext(ctx context.Context) *DeleteCombinationSkusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete combination skus params
func (o *DeleteCombinationSkusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete combination skus params
func (o *DeleteCombinationSkusParams) WithHTTPClient(client *http.Client) *DeleteCombinationSkusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete combination skus params
func (o *DeleteCombinationSkusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete combination skus params
func (o *DeleteCombinationSkusParams) WithBody(body *models.OpenpitrixDeleteCombinationSkusRequest) *DeleteCombinationSkusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete combination skus params
func (o *DeleteCombinationSkusParams) SetBody(body *models.OpenpitrixDeleteCombinationSkusRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCombinationSkusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
