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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeCombinationPricesParams creates a new DescribeCombinationPricesParams object
// with the default values initialized.
func NewDescribeCombinationPricesParams() *DescribeCombinationPricesParams {
	var ()
	return &DescribeCombinationPricesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeCombinationPricesParamsWithTimeout creates a new DescribeCombinationPricesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeCombinationPricesParamsWithTimeout(timeout time.Duration) *DescribeCombinationPricesParams {
	var ()
	return &DescribeCombinationPricesParams{

		timeout: timeout,
	}
}

// NewDescribeCombinationPricesParamsWithContext creates a new DescribeCombinationPricesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeCombinationPricesParamsWithContext(ctx context.Context) *DescribeCombinationPricesParams {
	var ()
	return &DescribeCombinationPricesParams{

		Context: ctx,
	}
}

// NewDescribeCombinationPricesParamsWithHTTPClient creates a new DescribeCombinationPricesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeCombinationPricesParamsWithHTTPClient(client *http.Client) *DescribeCombinationPricesParams {
	var ()
	return &DescribeCombinationPricesParams{
		HTTPClient: client,
	}
}

/*DescribeCombinationPricesParams contains all the parameters to send to the API endpoint
for the describe combination prices operation typically these are written to a http.Request
*/
type DescribeCombinationPricesParams struct {

	/*AttributeID*/
	AttributeID *string
	/*CombinationPriceID*/
	CombinationPriceID *string
	/*CombinationSkuID*/
	CombinationSkuID *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Reverse*/
	Reverse *bool
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithTimeout(timeout time.Duration) *DescribeCombinationPricesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithContext(ctx context.Context) *DescribeCombinationPricesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithHTTPClient(client *http.Client) *DescribeCombinationPricesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeID adds the attributeID to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithAttributeID(attributeID *string) *DescribeCombinationPricesParams {
	o.SetAttributeID(attributeID)
	return o
}

// SetAttributeID adds the attributeId to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetAttributeID(attributeID *string) {
	o.AttributeID = attributeID
}

// WithCombinationPriceID adds the combinationPriceID to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithCombinationPriceID(combinationPriceID *string) *DescribeCombinationPricesParams {
	o.SetCombinationPriceID(combinationPriceID)
	return o
}

// SetCombinationPriceID adds the combinationPriceId to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetCombinationPriceID(combinationPriceID *string) {
	o.CombinationPriceID = combinationPriceID
}

// WithCombinationSkuID adds the combinationSkuID to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithCombinationSkuID(combinationSkuID *string) *DescribeCombinationPricesParams {
	o.SetCombinationSkuID(combinationSkuID)
	return o
}

// SetCombinationSkuID adds the combinationSkuId to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetCombinationSkuID(combinationSkuID *string) {
	o.CombinationSkuID = combinationSkuID
}

// WithLimit adds the limit to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithLimit(limit *int64) *DescribeCombinationPricesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithOffset(offset *int64) *DescribeCombinationPricesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithReverse adds the reverse to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithReverse(reverse *bool) *DescribeCombinationPricesParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSortKey adds the sortKey to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithSortKey(sortKey *string) *DescribeCombinationPricesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe combination prices params
func (o *DescribeCombinationPricesParams) WithStatus(status *string) *DescribeCombinationPricesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe combination prices params
func (o *DescribeCombinationPricesParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeCombinationPricesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AttributeID != nil {

		// query param attribute_id
		var qrAttributeID string
		if o.AttributeID != nil {
			qrAttributeID = *o.AttributeID
		}
		qAttributeID := qrAttributeID
		if qAttributeID != "" {
			if err := r.SetQueryParam("attribute_id", qAttributeID); err != nil {
				return err
			}
		}

	}

	if o.CombinationPriceID != nil {

		// query param combination_price_id
		var qrCombinationPriceID string
		if o.CombinationPriceID != nil {
			qrCombinationPriceID = *o.CombinationPriceID
		}
		qCombinationPriceID := qrCombinationPriceID
		if qCombinationPriceID != "" {
			if err := r.SetQueryParam("combination_price_id", qCombinationPriceID); err != nil {
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