// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudCloudconnectionsNetworksDeleteParams creates a new PcloudCloudconnectionsNetworksDeleteParams object
// with the default values initialized.
func NewPcloudCloudconnectionsNetworksDeleteParams() *PcloudCloudconnectionsNetworksDeleteParams {
	var ()
	return &PcloudCloudconnectionsNetworksDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudconnectionsNetworksDeleteParamsWithTimeout creates a new PcloudCloudconnectionsNetworksDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudconnectionsNetworksDeleteParamsWithTimeout(timeout time.Duration) *PcloudCloudconnectionsNetworksDeleteParams {
	var ()
	return &PcloudCloudconnectionsNetworksDeleteParams{

		timeout: timeout,
	}
}

// NewPcloudCloudconnectionsNetworksDeleteParamsWithContext creates a new PcloudCloudconnectionsNetworksDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudconnectionsNetworksDeleteParamsWithContext(ctx context.Context) *PcloudCloudconnectionsNetworksDeleteParams {
	var ()
	return &PcloudCloudconnectionsNetworksDeleteParams{

		Context: ctx,
	}
}

// NewPcloudCloudconnectionsNetworksDeleteParamsWithHTTPClient creates a new PcloudCloudconnectionsNetworksDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudconnectionsNetworksDeleteParamsWithHTTPClient(client *http.Client) *PcloudCloudconnectionsNetworksDeleteParams {
	var ()
	return &PcloudCloudconnectionsNetworksDeleteParams{
		HTTPClient: client,
	}
}

/*PcloudCloudconnectionsNetworksDeleteParams contains all the parameters to send to the API endpoint
for the pcloud cloudconnections networks delete operation typically these are written to a http.Request
*/
type PcloudCloudconnectionsNetworksDeleteParams struct {

	/*CloudConnectionID
	  Cloud Connection ID

	*/
	CloudConnectionID string
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) WithTimeout(timeout time.Duration) *PcloudCloudconnectionsNetworksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) WithContext(ctx context.Context) *PcloudCloudconnectionsNetworksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) WithHTTPClient(client *http.Client) *PcloudCloudconnectionsNetworksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudConnectionID adds the cloudConnectionID to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) WithCloudConnectionID(cloudConnectionID string) *PcloudCloudconnectionsNetworksDeleteParams {
	o.SetCloudConnectionID(cloudConnectionID)
	return o
}

// SetCloudConnectionID adds the cloudConnectionId to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) SetCloudConnectionID(cloudConnectionID string) {
	o.CloudConnectionID = cloudConnectionID
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudconnectionsNetworksDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) WithNetworkID(networkID string) *PcloudCloudconnectionsNetworksDeleteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud cloudconnections networks delete params
func (o *PcloudCloudconnectionsNetworksDeleteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudconnectionsNetworksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_connection_id
	if err := r.SetPathParam("cloud_connection_id", o.CloudConnectionID); err != nil {
		return err
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}