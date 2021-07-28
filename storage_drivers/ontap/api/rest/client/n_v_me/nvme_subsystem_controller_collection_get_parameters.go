// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewNvmeSubsystemControllerCollectionGetParams creates a new NvmeSubsystemControllerCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemControllerCollectionGetParams() *NvmeSubsystemControllerCollectionGetParams {
	return &NvmeSubsystemControllerCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemControllerCollectionGetParamsWithTimeout creates a new NvmeSubsystemControllerCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemControllerCollectionGetParamsWithTimeout(timeout time.Duration) *NvmeSubsystemControllerCollectionGetParams {
	return &NvmeSubsystemControllerCollectionGetParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemControllerCollectionGetParamsWithContext creates a new NvmeSubsystemControllerCollectionGetParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemControllerCollectionGetParamsWithContext(ctx context.Context) *NvmeSubsystemControllerCollectionGetParams {
	return &NvmeSubsystemControllerCollectionGetParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemControllerCollectionGetParamsWithHTTPClient creates a new NvmeSubsystemControllerCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemControllerCollectionGetParamsWithHTTPClient(client *http.Client) *NvmeSubsystemControllerCollectionGetParams {
	return &NvmeSubsystemControllerCollectionGetParams{
		HTTPClient: client,
	}
}

/* NvmeSubsystemControllerCollectionGetParams contains all the parameters to send to the API endpoint
   for the nvme subsystem controller collection get operation.

   Typically these are written to a http.Request.
*/
type NvmeSubsystemControllerCollectionGetParams struct {

	/* AdminQueueDepth.

	   Filter by admin_queue.depth
	*/
	AdminQueueDepthQueryParameter *int64

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* HostID.

	   Filter by host.id
	*/
	HostIDQueryParameter *string

	/* HostNqn.

	   Filter by host.nqn
	*/
	HostNqnQueryParameter *string

	/* HostTransportAddress.

	   Filter by host.transport_address
	*/
	HostTransportAddressQueryParameter *string

	/* ID.

	   Filter by id
	*/
	IDQueryParameter *string

	/* InterfaceName.

	   Filter by interface.name
	*/
	InterfaceNameQueryParameter *string

	/* InterfaceTransportAddress.

	   Filter by interface.transport_address
	*/
	InterfaceTransportAddressQueryParameter *string

	/* InterfaceUUID.

	   Filter by interface.uuid
	*/
	InterfaceUUIDQueryParameter *string

	/* IoQueueCount.

	   Filter by io_queue.count
	*/
	IoQueueCountQueryParameter *int64

	/* IoQueueDepth.

	   Filter by io_queue.depth
	*/
	IoQueueDepthQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SubsystemName.

	   Filter by subsystem.name
	*/
	SubsystemNameQueryParameter *string

	/* SubsystemUUID.

	   Filter by subsystem.uuid
	*/
	SubsystemUUIDQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem controller collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemControllerCollectionGetParams) WithDefaults() *NvmeSubsystemControllerCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem controller collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemControllerCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NvmeSubsystemControllerCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithTimeout(timeout time.Duration) *NvmeSubsystemControllerCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithContext(ctx context.Context) *NvmeSubsystemControllerCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithHTTPClient(client *http.Client) *NvmeSubsystemControllerCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminQueueDepthQueryParameter adds the adminQueueDepth to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithAdminQueueDepthQueryParameter(adminQueueDepth *int64) *NvmeSubsystemControllerCollectionGetParams {
	o.SetAdminQueueDepthQueryParameter(adminQueueDepth)
	return o
}

// SetAdminQueueDepthQueryParameter adds the adminQueueDepth to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetAdminQueueDepthQueryParameter(adminQueueDepth *int64) {
	o.AdminQueueDepthQueryParameter = adminQueueDepth
}

// WithFields adds the fields to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithFields(fields []string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithHostIDQueryParameter adds the hostID to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithHostIDQueryParameter(hostID *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetHostIDQueryParameter(hostID)
	return o
}

// SetHostIDQueryParameter adds the hostId to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetHostIDQueryParameter(hostID *string) {
	o.HostIDQueryParameter = hostID
}

// WithHostNqnQueryParameter adds the hostNqn to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithHostNqnQueryParameter(hostNqn *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetHostNqnQueryParameter(hostNqn)
	return o
}

// SetHostNqnQueryParameter adds the hostNqn to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetHostNqnQueryParameter(hostNqn *string) {
	o.HostNqnQueryParameter = hostNqn
}

// WithHostTransportAddressQueryParameter adds the hostTransportAddress to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithHostTransportAddressQueryParameter(hostTransportAddress *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetHostTransportAddressQueryParameter(hostTransportAddress)
	return o
}

// SetHostTransportAddressQueryParameter adds the hostTransportAddress to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetHostTransportAddressQueryParameter(hostTransportAddress *string) {
	o.HostTransportAddressQueryParameter = hostTransportAddress
}

// WithIDQueryParameter adds the id to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithIDQueryParameter(id *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetIDQueryParameter(id)
	return o
}

// SetIDQueryParameter adds the id to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetIDQueryParameter(id *string) {
	o.IDQueryParameter = id
}

// WithInterfaceNameQueryParameter adds the interfaceName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithInterfaceNameQueryParameter(interfaceName *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetInterfaceNameQueryParameter(interfaceName)
	return o
}

// SetInterfaceNameQueryParameter adds the interfaceName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetInterfaceNameQueryParameter(interfaceName *string) {
	o.InterfaceNameQueryParameter = interfaceName
}

// WithInterfaceTransportAddressQueryParameter adds the interfaceTransportAddress to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithInterfaceTransportAddressQueryParameter(interfaceTransportAddress *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetInterfaceTransportAddressQueryParameter(interfaceTransportAddress)
	return o
}

// SetInterfaceTransportAddressQueryParameter adds the interfaceTransportAddress to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetInterfaceTransportAddressQueryParameter(interfaceTransportAddress *string) {
	o.InterfaceTransportAddressQueryParameter = interfaceTransportAddress
}

// WithInterfaceUUIDQueryParameter adds the interfaceUUID to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithInterfaceUUIDQueryParameter(interfaceUUID *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetInterfaceUUIDQueryParameter(interfaceUUID)
	return o
}

// SetInterfaceUUIDQueryParameter adds the interfaceUuid to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetInterfaceUUIDQueryParameter(interfaceUUID *string) {
	o.InterfaceUUIDQueryParameter = interfaceUUID
}

// WithIoQueueCountQueryParameter adds the ioQueueCount to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithIoQueueCountQueryParameter(ioQueueCount *int64) *NvmeSubsystemControllerCollectionGetParams {
	o.SetIoQueueCountQueryParameter(ioQueueCount)
	return o
}

// SetIoQueueCountQueryParameter adds the ioQueueCount to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetIoQueueCountQueryParameter(ioQueueCount *int64) {
	o.IoQueueCountQueryParameter = ioQueueCount
}

// WithIoQueueDepthQueryParameter adds the ioQueueDepth to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithIoQueueDepthQueryParameter(ioQueueDepth *int64) *NvmeSubsystemControllerCollectionGetParams {
	o.SetIoQueueDepthQueryParameter(ioQueueDepth)
	return o
}

// SetIoQueueDepthQueryParameter adds the ioQueueDepth to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetIoQueueDepthQueryParameter(ioQueueDepth *int64) {
	o.IoQueueDepthQueryParameter = ioQueueDepth
}

// WithMaxRecords adds the maxRecords to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithMaxRecords(maxRecords *int64) *NvmeSubsystemControllerCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderBy adds the orderBy to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithOrderBy(orderBy []string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithReturnRecords(returnRecords *bool) *NvmeSubsystemControllerCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *NvmeSubsystemControllerCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSubsystemNameQueryParameter adds the subsystemName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithSubsystemNameQueryParameter(subsystemName *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetSubsystemNameQueryParameter(subsystemName)
	return o
}

// SetSubsystemNameQueryParameter adds the subsystemName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetSubsystemNameQueryParameter(subsystemName *string) {
	o.SubsystemNameQueryParameter = subsystemName
}

// WithSubsystemUUIDQueryParameter adds the subsystemUUID to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithSubsystemUUIDQueryParameter(subsystemUUID *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetSubsystemUUIDQueryParameter(subsystemUUID)
	return o
}

// SetSubsystemUUIDQueryParameter adds the subsystemUuid to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetSubsystemUUIDQueryParameter(subsystemUUID *string) {
	o.SubsystemUUIDQueryParameter = subsystemUUID
}

// WithSVMNameQueryParameter adds the svmName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NvmeSubsystemControllerCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the nvme subsystem controller collection get params
func (o *NvmeSubsystemControllerCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemControllerCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdminQueueDepthQueryParameter != nil {

		// query param admin_queue.depth
		var qrAdminQueueDepth int64

		if o.AdminQueueDepthQueryParameter != nil {
			qrAdminQueueDepth = *o.AdminQueueDepthQueryParameter
		}
		qAdminQueueDepth := swag.FormatInt64(qrAdminQueueDepth)
		if qAdminQueueDepth != "" {

			if err := r.SetQueryParam("admin_queue.depth", qAdminQueueDepth); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.HostIDQueryParameter != nil {

		// query param host.id
		var qrHostID string

		if o.HostIDQueryParameter != nil {
			qrHostID = *o.HostIDQueryParameter
		}
		qHostID := qrHostID
		if qHostID != "" {

			if err := r.SetQueryParam("host.id", qHostID); err != nil {
				return err
			}
		}
	}

	if o.HostNqnQueryParameter != nil {

		// query param host.nqn
		var qrHostNqn string

		if o.HostNqnQueryParameter != nil {
			qrHostNqn = *o.HostNqnQueryParameter
		}
		qHostNqn := qrHostNqn
		if qHostNqn != "" {

			if err := r.SetQueryParam("host.nqn", qHostNqn); err != nil {
				return err
			}
		}
	}

	if o.HostTransportAddressQueryParameter != nil {

		// query param host.transport_address
		var qrHostTransportAddress string

		if o.HostTransportAddressQueryParameter != nil {
			qrHostTransportAddress = *o.HostTransportAddressQueryParameter
		}
		qHostTransportAddress := qrHostTransportAddress
		if qHostTransportAddress != "" {

			if err := r.SetQueryParam("host.transport_address", qHostTransportAddress); err != nil {
				return err
			}
		}
	}

	if o.IDQueryParameter != nil {

		// query param id
		var qrID string

		if o.IDQueryParameter != nil {
			qrID = *o.IDQueryParameter
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.InterfaceNameQueryParameter != nil {

		// query param interface.name
		var qrInterfaceName string

		if o.InterfaceNameQueryParameter != nil {
			qrInterfaceName = *o.InterfaceNameQueryParameter
		}
		qInterfaceName := qrInterfaceName
		if qInterfaceName != "" {

			if err := r.SetQueryParam("interface.name", qInterfaceName); err != nil {
				return err
			}
		}
	}

	if o.InterfaceTransportAddressQueryParameter != nil {

		// query param interface.transport_address
		var qrInterfaceTransportAddress string

		if o.InterfaceTransportAddressQueryParameter != nil {
			qrInterfaceTransportAddress = *o.InterfaceTransportAddressQueryParameter
		}
		qInterfaceTransportAddress := qrInterfaceTransportAddress
		if qInterfaceTransportAddress != "" {

			if err := r.SetQueryParam("interface.transport_address", qInterfaceTransportAddress); err != nil {
				return err
			}
		}
	}

	if o.InterfaceUUIDQueryParameter != nil {

		// query param interface.uuid
		var qrInterfaceUUID string

		if o.InterfaceUUIDQueryParameter != nil {
			qrInterfaceUUID = *o.InterfaceUUIDQueryParameter
		}
		qInterfaceUUID := qrInterfaceUUID
		if qInterfaceUUID != "" {

			if err := r.SetQueryParam("interface.uuid", qInterfaceUUID); err != nil {
				return err
			}
		}
	}

	if o.IoQueueCountQueryParameter != nil {

		// query param io_queue.count
		var qrIoQueueCount int64

		if o.IoQueueCountQueryParameter != nil {
			qrIoQueueCount = *o.IoQueueCountQueryParameter
		}
		qIoQueueCount := swag.FormatInt64(qrIoQueueCount)
		if qIoQueueCount != "" {

			if err := r.SetQueryParam("io_queue.count", qIoQueueCount); err != nil {
				return err
			}
		}
	}

	if o.IoQueueDepthQueryParameter != nil {

		// query param io_queue.depth
		var qrIoQueueDepth int64

		if o.IoQueueDepthQueryParameter != nil {
			qrIoQueueDepth = *o.IoQueueDepthQueryParameter
		}
		qIoQueueDepth := swag.FormatInt64(qrIoQueueDepth)
		if qIoQueueDepth != "" {

			if err := r.SetQueryParam("io_queue.depth", qIoQueueDepth); err != nil {
				return err
			}
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SubsystemNameQueryParameter != nil {

		// query param subsystem.name
		var qrSubsystemName string

		if o.SubsystemNameQueryParameter != nil {
			qrSubsystemName = *o.SubsystemNameQueryParameter
		}
		qSubsystemName := qrSubsystemName
		if qSubsystemName != "" {

			if err := r.SetQueryParam("subsystem.name", qSubsystemName); err != nil {
				return err
			}
		}
	}

	if o.SubsystemUUIDQueryParameter != nil {

		// query param subsystem.uuid
		var qrSubsystemUUID string

		if o.SubsystemUUIDQueryParameter != nil {
			qrSubsystemUUID = *o.SubsystemUUIDQueryParameter
		}
		qSubsystemUUID := qrSubsystemUUID
		if qSubsystemUUID != "" {

			if err := r.SetQueryParam("subsystem.uuid", qSubsystemUUID); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNvmeSubsystemControllerCollectionGet binds the parameter fields
func (o *NvmeSubsystemControllerCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamNvmeSubsystemControllerCollectionGet binds the parameter order_by
func (o *NvmeSubsystemControllerCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}