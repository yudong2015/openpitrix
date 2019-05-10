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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeCombinationSkusParams creates a new DescribeCombinationSkusParams object
// with the default values initialized.
func NewDescribeCombinationSkusParams() *DescribeCombinationSkusParams {
	var ()
	return &DescribeCombinationSkusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeCombinationSkusParamsWithTimeout creates a new DescribeCombinationSkusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeCombinationSkusParamsWithTimeout(timeout time.Duration) *DescribeCombinationSkusParams {
	var ()
	return &DescribeCombinationSkusParams{

		timeout: timeout,
	}
}

// NewDescribeCombinationSkusParamsWithContext creates a new DescribeCombinationSkusParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeCombinationSkusParamsWithContext(ctx context.Context) *DescribeCombinationSkusParams {
	var ()
	return &DescribeCombinationSkusParams{

		Context: ctx,
	}
}

// NewDescribeCombinationSkusParamsWithHTTPClient creates a new DescribeCombinationSkusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeCombinationSkusParamsWithHTTPClient(client *http.Client) *DescribeCombinationSkusParams {
	var ()
	return &DescribeCombinationSkusParams{
		HTTPClient: client,
	}
}

/*DescribeCombinationSkusParams contains all the parameters to send to the API endpoint
for the describe combination skus operation typically these are written to a http.Request
*/
type DescribeCombinationSkusParams struct {

	/*CombinationID*/
	CombinationID *string
	/*CombinationSkuID*/
	CombinationSkuID *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Reverse*/
	Reverse *bool
	/*SkuID*/
	SkuID *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithTimeout(timeout time.Duration) *DescribeCombinationSkusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithContext(ctx context.Context) *DescribeCombinationSkusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithHTTPClient(client *http.Client) *DescribeCombinationSkusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCombinationID adds the combinationID to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithCombinationID(combinationID *string) *DescribeCombinationSkusParams {
	o.SetCombinationID(combinationID)
	return o
}

// SetCombinationID adds the combinationId to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetCombinationID(combinationID *string) {
	o.CombinationID = combinationID
}

// WithCombinationSkuID adds the combinationSkuID to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithCombinationSkuID(combinationSkuID *string) *DescribeCombinationSkusParams {
	o.SetCombinationSkuID(combinationSkuID)
	return o
}

// SetCombinationSkuID adds the combinationSkuId to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetCombinationSkuID(combinationSkuID *string) {
	o.CombinationSkuID = combinationSkuID
}

// WithLimit adds the limit to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithLimit(limit *int64) *DescribeCombinationSkusParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithOffset(offset *int64) *DescribeCombinationSkusParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithReverse adds the reverse to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithReverse(reverse *bool) *DescribeCombinationSkusParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSkuID adds the skuID to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithSkuID(skuID *string) *DescribeCombinationSkusParams {
	o.SetSkuID(skuID)
	return o
}

// SetSkuID adds the skuId to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetSkuID(skuID *string) {
	o.SkuID = skuID
}

// WithSortKey adds the sortKey to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithSortKey(sortKey *string) *DescribeCombinationSkusParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe combination skus params
func (o *DescribeCombinationSkusParams) WithStatus(status *string) *DescribeCombinationSkusParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe combination skus params
func (o *DescribeCombinationSkusParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeCombinationSkusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CombinationID != nil {

		// query param combination_id
		var qrCombinationID string
		if o.CombinationID != nil {
			qrCombinationID = *o.CombinationID
		}
		qCombinationID := qrCombinationID
		if qCombinationID != "" {
			if err := r.SetQueryParam("combination_id", qCombinationID); err != nil {
				return err
			}
		}

	}

	if o.CombinationSkuID != nil {

		// query param combination_sku_id
		var qrCombinationSkuID string
		if o.CombinationSkuID != nil {
			qrCombinationSkuID = *o.CombinationSkuID
		}
		qCombinationSkuID := qrCombinationSkuID
		if qCombinationSkuID != "" {
			if err := r.SetQueryParam("combination_sku_id", qCombinationSkuID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SkuID != nil {

		// query param sku_id
		var qrSkuID string
		if o.SkuID != nil {
			qrSkuID = *o.SkuID
		}
		qSkuID := qrSkuID
		if qSkuID != "" {
			if err := r.SetQueryParam("sku_id", qSkuID); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
