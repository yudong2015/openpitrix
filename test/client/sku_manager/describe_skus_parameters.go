// Code generated by go-swagger; DO NOT EDIT.

package sku_manager

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

// NewDescribeSkusParams creates a new DescribeSkusParams object
// with the default values initialized.
func NewDescribeSkusParams() *DescribeSkusParams {
	var ()
	return &DescribeSkusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeSkusParamsWithTimeout creates a new DescribeSkusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeSkusParamsWithTimeout(timeout time.Duration) *DescribeSkusParams {
	var ()
	return &DescribeSkusParams{

		timeout: timeout,
	}
}

// NewDescribeSkusParamsWithContext creates a new DescribeSkusParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeSkusParamsWithContext(ctx context.Context) *DescribeSkusParams {
	var ()
	return &DescribeSkusParams{

		Context: ctx,
	}
}

// NewDescribeSkusParamsWithHTTPClient creates a new DescribeSkusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeSkusParamsWithHTTPClient(client *http.Client) *DescribeSkusParams {
	var ()
	return &DescribeSkusParams{
		HTTPClient: client,
	}
}

/*DescribeSkusParams contains all the parameters to send to the API endpoint
for the describe skus operation typically these are written to a http.Request
*/
type DescribeSkusParams struct {

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
	/*SpuID*/
	SpuID *string
	/*Status*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe skus params
func (o *DescribeSkusParams) WithTimeout(timeout time.Duration) *DescribeSkusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe skus params
func (o *DescribeSkusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe skus params
func (o *DescribeSkusParams) WithContext(ctx context.Context) *DescribeSkusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe skus params
func (o *DescribeSkusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe skus params
func (o *DescribeSkusParams) WithHTTPClient(client *http.Client) *DescribeSkusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe skus params
func (o *DescribeSkusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe skus params
func (o *DescribeSkusParams) WithLimit(limit *int64) *DescribeSkusParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe skus params
func (o *DescribeSkusParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe skus params
func (o *DescribeSkusParams) WithOffset(offset *int64) *DescribeSkusParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe skus params
func (o *DescribeSkusParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithReverse adds the reverse to the describe skus params
func (o *DescribeSkusParams) WithReverse(reverse *bool) *DescribeSkusParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe skus params
func (o *DescribeSkusParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSkuID adds the skuID to the describe skus params
func (o *DescribeSkusParams) WithSkuID(skuID *string) *DescribeSkusParams {
	o.SetSkuID(skuID)
	return o
}

// SetSkuID adds the skuId to the describe skus params
func (o *DescribeSkusParams) SetSkuID(skuID *string) {
	o.SkuID = skuID
}

// WithSortKey adds the sortKey to the describe skus params
func (o *DescribeSkusParams) WithSortKey(sortKey *string) *DescribeSkusParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe skus params
func (o *DescribeSkusParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithSpuID adds the spuID to the describe skus params
func (o *DescribeSkusParams) WithSpuID(spuID *string) *DescribeSkusParams {
	o.SetSpuID(spuID)
	return o
}

// SetSpuID adds the spuId to the describe skus params
func (o *DescribeSkusParams) SetSpuID(spuID *string) {
	o.SpuID = spuID
}

// WithStatus adds the status to the describe skus params
func (o *DescribeSkusParams) WithStatus(status *string) *DescribeSkusParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe skus params
func (o *DescribeSkusParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeSkusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.SpuID != nil {

		// query param spu_id
		var qrSpuID string
		if o.SpuID != nil {
			qrSpuID = *o.SpuID
		}
		qSpuID := qrSpuID
		if qSpuID != "" {
			if err := r.SetQueryParam("spu_id", qSpuID); err != nil {
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
