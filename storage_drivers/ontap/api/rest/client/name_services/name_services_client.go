// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new name services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for name services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DNSCollectionGet(params *DNSCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSCollectionGetOK, error)

	DNSCreate(params *DNSCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSCreateCreated, error)

	DNSDelete(params *DNSDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSDeleteOK, error)

	DNSGet(params *DNSGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSGetOK, error)

	DNSModify(params *DNSModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSModifyOK, error)

	LdapCollectionGet(params *LdapCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapCollectionGetOK, error)

	LdapCreate(params *LdapCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapCreateCreated, error)

	LdapDelete(params *LdapDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapDeleteOK, error)

	LdapGet(params *LdapGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapGetOK, error)

	LdapModify(params *LdapModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapModifyOK, error)

	NameMappingCollectionGet(params *NameMappingCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingCollectionGetOK, error)

	NameMappingCreate(params *NameMappingCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingCreateCreated, error)

	NameMappingDelete(params *NameMappingDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingDeleteOK, error)

	NameMappingModify(params *NameMappingModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingModifyOK, error)

	NameMappingPositionGet(params *NameMappingPositionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingPositionGetOK, error)

	NisCollectionGet(params *NisCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisCollectionGetOK, error)

	NisCreate(params *NisCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisCreateCreated, error)

	NisDelete(params *NisDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisDeleteOK, error)

	NisGet(params *NisGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisGetOK, error)

	NisModify(params *NisModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisModifyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DNSCollectionGet Retrieves the DNS configurations of all data SVMs.
DNS configuration for the cluster is retrieved and managed via [`/api/cluster`](#docs-cluster-cluster).
### Related ONTAP commands
* `vserver services name-service dns show`
* `vserver services name-service dns check`
### Learn more
* [`DOC /name-services/dns`](#docs-name-services-name-services_dns)

*/
func (a *Client) DNSCollectionGet(params *DNSCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSCollectionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDNSCollectionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dns_collection_get",
		Method:             "GET",
		PathPattern:        "/name-services/dns",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DNSCollectionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DNSCollectionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DNSCollectionGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DNSCreate Creates DNS domain and server configurations for an SVM.<br/>
### Important notes
- Each SVM can have only one DNS configuration.
- The domain name and the servers fields cannot be empty.
- IPv6 must be enabled if IPv6 family addresses are specified in the `servers` field.
- Configuring more than one DNS server is recommended to avoid a single point of failure.
- The DNS server specified using the `servers` field is validated during this operation.<br/>
</br> The validation fails in the following scenarios:<br/>
1. The server is not a DNS server.
2. The server does not exist.
3. The server is unreachable.<br/>

### Learn more
* [`DOC /name-services/dns`](#docs-name-services-name-services_dns)
*/
func (a *Client) DNSCreate(params *DNSCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDNSCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dns_create",
		Method:             "POST",
		PathPattern:        "/name-services/dns",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DNSCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DNSCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DNSCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DNSDelete Deletes DNS domain configuration of the specified SVM.
### Related ONTAP commands
* `vserver services name-service dns delete`
### Learn more
* [`DOC /name-services/dns`](#docs-name-services-name-services_dns)

*/
func (a *Client) DNSDelete(params *DNSDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDNSDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dns_delete",
		Method:             "DELETE",
		PathPattern:        "/name-services/dns/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DNSDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DNSDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DNSDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DNSGet Retrieves DNS domain and server configuration of an SVM. By default, both DNS domains and servers are displayed.
DNS configuration for the cluster is retrieved and managed via [`/api/cluster`](#docs-cluster-cluster).
### Related ONTAP commands
* `vserver services name-service dns show`
* `vserver services name-service dns check`
### Learn more
* [`DOC /name-services/dns`](#docs-name-services-name-services_dns)

*/
func (a *Client) DNSGet(params *DNSGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDNSGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dns_get",
		Method:             "GET",
		PathPattern:        "/name-services/dns/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DNSGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DNSGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DNSGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DNSModify Updates DNS domain and server configurations of an SVM.
### Important notes
- Both DNS domains and servers can be modified.
- The domains and servers fields cannot be empty.
- IPv6 must be enabled if IPv6 family addresses are specified for the `servers` field.
- The DNS server specified using the `servers` field is validated during this operation.<br/>
The validation fails in the following scenarios:<br/>
1. The server is not a DNS server.
2. The server does not exist.
3. The server is unreachable.<br/>

### Learn more
* [`DOC /name-services/dns`](#docs-name-services-name-services_dns)
*/
func (a *Client) DNSModify(params *DNSModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DNSModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDNSModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dns_modify",
		Method:             "PATCH",
		PathPattern:        "/name-services/dns/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DNSModifyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DNSModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DNSModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LdapCollectionGet Retrieves the LDAP configurations for all SVMs.

### Learn more
* [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)
*/
func (a *Client) LdapCollectionGet(params *LdapCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapCollectionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLdapCollectionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ldap_collection_get",
		Method:             "GET",
		PathPattern:        "/name-services/ldap",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LdapCollectionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LdapCollectionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LdapCollectionGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LdapCreate Creates an LDAP configuration for an SVM.
### Important notes
* Each SVM can have one LDAP configuration.
* The LDAP servers and Active Directory domain are mutually exclusive fields. These fields cannot be empty. At any point in time, either the LDAP servers or Active Directory domain must be populated.
* LDAP configuration with Active Directory domain cannot be created on an admin SVM.
* IPv6 must be enabled if IPv6 family addresses are specified.</br>
#### The following parameters are optional:
- preferred AD servers
- schema
- port
- min_bind_level
- bind_password
- base_scope
- use_start_tls
- session_security</br>
Configuring more than one LDAP server is recommended to avoid a single point of failure.
Both FQDNs and IP addresses are supported for the "servers" field.
The Acitve Directory domain or LDAP servers are validated as part of this operation.</br>
LDAP validation fails in the following scenarios:<br/>
1. The server does not have LDAP installed.
2. The server or Active Directory domain is invalid.
3. The server or Active Directory domain is unreachable.<br/>

### Learn more
* [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)
*/
func (a *Client) LdapCreate(params *LdapCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLdapCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ldap_create",
		Method:             "POST",
		PathPattern:        "/name-services/ldap",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LdapCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LdapCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LdapCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LdapDelete Deletes the LDAP configuration of the specified SVM. LDAP can be removed as a source from the ns-switch if LDAP is not used as a source for lookups.

### Learn more
* [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)
*/
func (a *Client) LdapDelete(params *LdapDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLdapDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ldap_delete",
		Method:             "DELETE",
		PathPattern:        "/name-services/ldap/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LdapDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LdapDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LdapDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LdapGet Retrieves LDAP configuration for an SVM. All parameters for the LDAP configuration are displayed by default.

### Learn more
* [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)
*/
func (a *Client) LdapGet(params *LdapGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLdapGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ldap_get",
		Method:             "GET",
		PathPattern:        "/name-services/ldap/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LdapGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LdapGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LdapGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  LdapModify Updates an LDAP configuration of an SVM.
### Important notes
* Both mandatory and optional parameters of the LDAP configuration can be updated.
* The LDAP servers and Active Directory domain are mutually exclusive fields. These fields cannot be empty. At any point in time, either the LDAP servers or Active Directory domain must be populated.
* IPv6 must be enabled if IPv6 family addresses are specified.<br/>
</br>Configuring more than one LDAP server is recommended to avoid a sinlge point of failure.
Both FQDNs and IP addresses are supported for the "servers" field.
The Active Directory domain or LDAP servers are validated as part of this operation.<br/>
LDAP validation fails in the following scenarios:<br/>
1. The server does not have LDAP installed.
2. The server or Active Directory domain is invalid.
3. The server or Active Directory domain is unreachable<br/>

### Learn more
* [`DOC /name-services/ldap`](#docs-name-services-name-services_ldap)
*/
func (a *Client) LdapModify(params *LdapModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LdapModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLdapModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ldap_modify",
		Method:             "PATCH",
		PathPattern:        "/name-services/ldap/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LdapModifyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LdapModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LdapModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NameMappingCollectionGet Retrieves the name mapping configuration for all SVMs.
### Related ONTAP commands
* `vserver name-mapping show`
### Learn more
* [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings)

*/
func (a *Client) NameMappingCollectionGet(params *NameMappingCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingCollectionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNameMappingCollectionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "name_mapping_collection_get",
		Method:             "GET",
		PathPattern:        "/name-services/name-mappings",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NameMappingCollectionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NameMappingCollectionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NameMappingCollectionGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NameMappingCreate Creates name mappings for an SVM.
### Required properties
* `svm.uuid` or `svm.name` - Existing SVM in which to create the name mapping.
* `index` - Name mapping's position in the priority list.
* `direction` - Direction of the name mapping.
* `pattern` - Pattern to match to. Maximum length is 256 characters.
* `replacement` - Replacement pattern to match to. Maximum length is 256 characters.
### Recommended optional properties
* `client_match` - Hostname or IP address added to match the pattern to the client's workstation IP address.
### Related ONTAP commands
* `vserver name-mapping create`
* `vserver name-mapping insert`
### Learn more
* [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings)

*/
func (a *Client) NameMappingCreate(params *NameMappingCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNameMappingCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "name_mapping_create",
		Method:             "POST",
		PathPattern:        "/name-services/name-mappings",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NameMappingCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NameMappingCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NameMappingCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NameMappingDelete Deletes the name mapping configuration.
### Related ONTAP commands
* `vserver name-mapping delete`
### Learn more
* [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings)

*/
func (a *Client) NameMappingDelete(params *NameMappingDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNameMappingDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "name_mapping_delete",
		Method:             "DELETE",
		PathPattern:        "/name-services/name-mappings/{svm.uuid}/{direction}/{index}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NameMappingDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NameMappingDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NameMappingDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NameMappingModify Updates the name mapping configuration of an SVM. The positions can be swapped by providing the `new_index` property.
Swapping is not allowed for entries that have `client_match` property configured.
### Related ONTAP commands
* `vserver name-mapping insert`
* `vserver name-mapping modify`
* `vserver name-mapping swap`
### Learn more
* [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings)

*/
func (a *Client) NameMappingModify(params *NameMappingModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNameMappingModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "name_mapping_modify",
		Method:             "PATCH",
		PathPattern:        "/name-services/name-mappings/{svm.uuid}/{direction}/{index}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NameMappingModifyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NameMappingModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NameMappingModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NameMappingPositionGet Retrieves the name mapping configuration of an SVM.
### Related ONTAP commands
* `vserver name-mapping show`
### Learn more
* [`DOC /name-services/name-mappings`](#docs-name-services-name-services_name-mappings)

*/
func (a *Client) NameMappingPositionGet(params *NameMappingPositionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NameMappingPositionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNameMappingPositionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "name_mapping_position_get",
		Method:             "GET",
		PathPattern:        "/name-services/name-mappings/{svm.uuid}/{direction}/{index}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NameMappingPositionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NameMappingPositionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NameMappingPositionGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NisCollectionGet Retrieves NIS domain configurations of all the SVMs. The bound_servers field indicates the successfully bound NIS servers. Lookups and authentications fail if there are no bound servers.
### Related ONTAP commands
* `vserver services name-service nis-domain show`
* `vserver services name-service nis-domain show-bound`
### Learn more
* [`DOC /name-services/nis`](#docs-name-services-name-services_nis)

*/
func (a *Client) NisCollectionGet(params *NisCollectionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisCollectionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNisCollectionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nis_collection_get",
		Method:             "GET",
		PathPattern:        "/name-services/nis",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NisCollectionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NisCollectionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NisCollectionGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NisCreate Creates an NIS domain and server confguration for a data SVM.
NIS configuration for the cluster is managed via [`/api/security/authentication/cluster/nis`](#docs-security-security_authentication_cluster_nis).<br/>
### Important notes
  - Each SVM can have one NIS domain configuration.
  - Multiple SVMs can be configured with the same NIS domain. Specify the NIS domain and NIS servers as input.Domain name and servers fields cannot be empty.
  - Both FQDNs and IP addresses are supported for the servers field.
  - IPv6 must be enabled if IPv6 family addresses are specified in the servers field.
  - A maximum of ten NIS servers are supported.
### Required properties
* `svm.uuid` or `svm.name` - Existing SVM in which to create the NIS configuration.
* `domain` - NIS domain to which the configuration belongs.
* `servers` - List of NIS server IP addresses.
### Related ONTAP commands
* `vserver services name-service nis-domain create`
### Learn more
* [`DOC /name-services/nis`](#docs-name-services-name-services_nis)

*/
func (a *Client) NisCreate(params *NisCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNisCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nis_create",
		Method:             "POST",
		PathPattern:        "/name-services/nis",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NisCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NisCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NisCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NisDelete Deletes the NIS domain configuration of an SVM. NIS can be removed as a source from ns-switch if NIS is not used for lookups.
### Related ONTAP commands
* `vserver services name-service nis-domain delete`
### Learn more
* [`DOC /name-services/nis`](#docs-name-services-name-services_nis)

*/
func (a *Client) NisDelete(params *NisDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNisDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nis_delete",
		Method:             "DELETE",
		PathPattern:        "/name-services/nis/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NisDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NisDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NisDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NisGet Retrieves NIS domain and server configurations of an SVM. Both NIS domain and servers are displayed by default. The bound_servers field indicates the successfully bound NIS servers.
### Related ONTAP commands
* `vserver services name-service nis-domain show`
* `vserver services name-service nis-domain show-bound`
### Learn more
* [`DOC /name-services/nis`](#docs-name-services-name-services_nis)

*/
func (a *Client) NisGet(params *NisGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNisGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nis_get",
		Method:             "GET",
		PathPattern:        "/name-services/nis/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NisGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NisGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NisGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NisModify Updates NIS domain and server configuration of an SVM.<br/>
### Important notes
  - Both NIS domain and servers can be modified.
  - Domains and servers cannot be empty.
  - Both FQDNs and IP addresses are supported for the servers field.
  - If the domain is modified, NIS servers must also be specified.
  - IPv6 must be enabled if IPv6 family addresses are specified for the servers field.
### Related ONTAP commands
* `vserver services name-service nis-domain modify`
### Learn more
* [`DOC /name-services/nis`](#docs-name-services-name-services_nis)

*/
func (a *Client) NisModify(params *NisModifyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NisModifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNisModifyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "nis_modify",
		Method:             "PATCH",
		PathPattern:        "/name-services/nis/{svm.uuid}",
		ProducesMediaTypes: []string{"application/hal+json", "application/json"},
		ConsumesMediaTypes: []string{"application/hal+json", "application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NisModifyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NisModifyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NisModifyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}