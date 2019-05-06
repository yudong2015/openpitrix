// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

// NewDescribeAppsParams creates a new DescribeAppsParams object
// with the default values initialized.
func NewDescribeAppsParams() *DescribeAppsParams {
	var ()
	return &DescribeAppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeAppsParamsWithTimeout creates a new DescribeAppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeAppsParamsWithTimeout(timeout time.Duration) *DescribeAppsParams {
	var ()
	return &DescribeAppsParams{

		timeout: timeout,
	}
}

// NewDescribeAppsParamsWithContext creates a new DescribeAppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeAppsParamsWithContext(ctx context.Context) *DescribeAppsParams {
	var ()
	return &DescribeAppsParams{

		Context: ctx,
	}
}

// NewDescribeAppsParamsWithHTTPClient creates a new DescribeAppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeAppsParamsWithHTTPClient(client *http.Client) *DescribeAppsParams {
	var ()
	return &DescribeAppsParams{
		HTTPClient: client,
	}
}

/*DescribeAppsParams contains all the parameters to send to the API endpoint
for the describe apps operation typically these are written to a http.Request
*/
type DescribeAppsParams struct {

	/*AppID
	  app ids.

	*/
	AppID []string
	/*CategoryID
	  app category ids.

	*/
	CategoryID []string
	/*ChartName
	  app chart name.

	*/
	ChartName []string
	/*DisplayColumns
	  select column to display.

	*/
	DisplayColumns []string
	/*Limit
	  data limit per page, default is 20, max value is 200.

	*/
	Limit *int64
	/*Name
	  app name.

	*/
	Name []string
	/*Offset
	  data offset, default is 0.

	*/
	Offset *int64
	/*Owner
	  app owner.

	*/
	Owner []string
	/*RepoID
	  app repository ids.

	*/
	RepoID []string
	/*Reverse
	  value = 0 sort ASC, value = 1 sort DESC.

	*/
	Reverse *bool
	/*SearchWord
	  query key, support these fields(app_id, name, repo_id, description, status, home, icon, screenshots, maintainers, sources, readme, owner, chart_name).

	*/
	SearchWord *string
	/*SortKey
	  sort key, order by sort_key, default create_time.

	*/
	SortKey *string
	/*Status
	  app status eg.[modify|submit|review|cancel|release|delete|pass|reject|suspend|recover].

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe apps params
func (o *DescribeAppsParams) WithTimeout(timeout time.Duration) *DescribeAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe apps params
func (o *DescribeAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe apps params
func (o *DescribeAppsParams) WithContext(ctx context.Context) *DescribeAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe apps params
func (o *DescribeAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe apps params
func (o *DescribeAppsParams) WithHTTPClient(client *http.Client) *DescribeAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe apps params
func (o *DescribeAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the describe apps params
func (o *DescribeAppsParams) WithAppID(appID []string) *DescribeAppsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the describe apps params
func (o *DescribeAppsParams) SetAppID(appID []string) {
	o.AppID = appID
}

// WithCategoryID adds the categoryID to the describe apps params
func (o *DescribeAppsParams) WithCategoryID(categoryID []string) *DescribeAppsParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the describe apps params
func (o *DescribeAppsParams) SetCategoryID(categoryID []string) {
	o.CategoryID = categoryID
}

// WithChartName adds the chartName to the describe apps params
func (o *DescribeAppsParams) WithChartName(chartName []string) *DescribeAppsParams {
	o.SetChartName(chartName)
	return o
}

// SetChartName adds the chartName to the describe apps params
func (o *DescribeAppsParams) SetChartName(chartName []string) {
	o.ChartName = chartName
}

// WithDisplayColumns adds the displayColumns to the describe apps params
func (o *DescribeAppsParams) WithDisplayColumns(displayColumns []string) *DescribeAppsParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe apps params
func (o *DescribeAppsParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe apps params
func (o *DescribeAppsParams) WithLimit(limit *int64) *DescribeAppsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe apps params
func (o *DescribeAppsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the describe apps params
func (o *DescribeAppsParams) WithName(name []string) *DescribeAppsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe apps params
func (o *DescribeAppsParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the describe apps params
func (o *DescribeAppsParams) WithOffset(offset *int64) *DescribeAppsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe apps params
func (o *DescribeAppsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe apps params
func (o *DescribeAppsParams) WithOwner(owner []string) *DescribeAppsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe apps params
func (o *DescribeAppsParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithRepoID adds the repoID to the describe apps params
func (o *DescribeAppsParams) WithRepoID(repoID []string) *DescribeAppsParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the describe apps params
func (o *DescribeAppsParams) SetRepoID(repoID []string) {
	o.RepoID = repoID
}

// WithReverse adds the reverse to the describe apps params
func (o *DescribeAppsParams) WithReverse(reverse *bool) *DescribeAppsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe apps params
func (o *DescribeAppsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe apps params
func (o *DescribeAppsParams) WithSearchWord(searchWord *string) *DescribeAppsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe apps params
func (o *DescribeAppsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe apps params
func (o *DescribeAppsParams) WithSortKey(sortKey *string) *DescribeAppsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe apps params
func (o *DescribeAppsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe apps params
func (o *DescribeAppsParams) WithStatus(status []string) *DescribeAppsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe apps params
func (o *DescribeAppsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAppID := o.AppID

	joinedAppID := swag.JoinByFormat(valuesAppID, "multi")
	// query array param app_id
	if err := r.SetQueryParam("app_id", joinedAppID...); err != nil {
		return err
	}

	valuesCategoryID := o.CategoryID

	joinedCategoryID := swag.JoinByFormat(valuesCategoryID, "multi")
	// query array param category_id
	if err := r.SetQueryParam("category_id", joinedCategoryID...); err != nil {
		return err
	}

	valuesChartName := o.ChartName

	joinedChartName := swag.JoinByFormat(valuesChartName, "multi")
	// query array param chart_name
	if err := r.SetQueryParam("chart_name", joinedChartName...); err != nil {
		return err
	}

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
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

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "multi")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
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

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	valuesRepoID := o.RepoID

	joinedRepoID := swag.JoinByFormat(valuesRepoID, "multi")
	// query array param repo_id
	if err := r.SetQueryParam("repo_id", joinedRepoID...); err != nil {
		return err
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

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
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

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
