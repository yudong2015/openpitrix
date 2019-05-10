// Code generated by go-swagger; DO NOT EDIT.

package promotion_manager

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

// NewDeleteProbationsParams creates a new DeleteProbationsParams object
// with the default values initialized.
func NewDeleteProbationsParams() *DeleteProbationsParams {
	var ()
	return &DeleteProbationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProbationsParamsWithTimeout creates a new DeleteProbationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProbationsParamsWithTimeout(timeout time.Duration) *DeleteProbationsParams {
	var ()
	return &DeleteProbationsParams{

		timeout: timeout,
	}
}

// NewDeleteProbationsParamsWithContext creates a new DeleteProbationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProbationsParamsWithContext(ctx context.Context) *DeleteProbationsParams {
	var ()
	return &DeleteProbationsParams{

		Context: ctx,
	}
}

// NewDeleteProbationsParamsWithHTTPClient creates a new DeleteProbationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProbationsParamsWithHTTPClient(client *http.Client) *DeleteProbationsParams {
	var ()
	return &DeleteProbationsParams{
		HTTPClient: client,
	}
}

/*DeleteProbationsParams contains all the parameters to send to the API endpoint
for the delete probations operation typically these are written to a http.Request
*/
type DeleteProbationsParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteProbationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete probations params
func (o *DeleteProbationsParams) WithTimeout(timeout time.Duration) *DeleteProbationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete probations params
func (o *DeleteProbationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete probations params
func (o *DeleteProbationsParams) WithContext(ctx context.Context) *DeleteProbationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete probations params
func (o *DeleteProbationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete probations params
func (o *DeleteProbationsParams) WithHTTPClient(client *http.Client) *DeleteProbationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete probations params
func (o *DeleteProbationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete probations params
func (o *DeleteProbationsParams) WithBody(body *models.OpenpitrixDeleteProbationsRequest) *DeleteProbationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete probations params
func (o *DeleteProbationsParams) SetBody(body *models.OpenpitrixDeleteProbationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProbationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
