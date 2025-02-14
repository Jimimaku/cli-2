// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddOrganization(params *AddOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationOK, error)

	AddOrganizationAutoInvite(params *AddOrganizationAutoInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationAutoInviteOK, error)

	BulkInviteOrganization(params *BulkInviteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BulkInviteOrganizationOK, error)

	DeleteInvite(params *DeleteInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInviteOK, error)

	DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationOK, error)

	DeleteOrganizationAutoInvite(params *DeleteOrganizationAutoInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationAutoInviteOK, error)

	EditBilling(params *EditBillingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditBillingOK, error)

	EditMember(params *EditMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditMemberOK, error)

	EditOrganization(params *EditOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditOrganizationOK, error)

	GetBilling(params *GetBillingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBillingOK, error)

	GetNextMutationID(params *GetNextMutationIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNextMutationIDOK, error)

	GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationOK, error)

	GetOrganizationAutoInvite(params *GetOrganizationAutoInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationAutoInviteOK, error)

	GetOrganizationInvitations(params *GetOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationInvitationsOK, error)

	GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationMembersOK, error)

	GetOrganizationMutations(params *GetOrganizationMutationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationMutationsOK, error)

	GetOrganizationTier(params *GetOrganizationTierParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationTierOK, error)

	InviteOrganization(params *InviteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InviteOrganizationOK, error)

	JoinOrganization(params *JoinOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*JoinOrganizationOK, error)

	KomodoAuthorized(params *KomodoAuthorizedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KomodoAuthorizedOK, error)

	ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationsOK, error)

	MutateOrganization(params *MutateOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MutateOrganizationOK, error)

	QuitOrganization(params *QuitOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QuitOrganizationOK, error)

	UpdateBillingDate(params *UpdateBillingDateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBillingDateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddOrganization creates a new organization

  Create a new organization
*/
func (a *Client) AddOrganization(params *AddOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addOrganization",
		Method:             "POST",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOrganizationReader{formats: a.formats},
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
	success, ok := result.(*AddOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddOrganizationAutoInvite adds a domain to organization auto invite

  Add a domain to an organization's auto-invite settings
*/
func (a *Client) AddOrganizationAutoInvite(params *AddOrganizationAutoInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddOrganizationAutoInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationAutoInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addOrganizationAutoInvite",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/autoinvite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddOrganizationAutoInviteReader{formats: a.formats},
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
	success, ok := result.(*AddOrganizationAutoInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addOrganizationAutoInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BulkInviteOrganization bulks organization invitations

  Invite many users to an organization at once. Note that while the content type
must be sent as `text/plain`, the body is actually two column CSV data. First
column is the role to assign them. It should be one of `admin`, `editor`, or
`reader`. Second is email address.

Example:
  ```
  editor, person1@email.com
  editor, person2@email.com
  ```

  Note that quoted strings are not supported.

*/
func (a *Client) BulkInviteOrganization(params *BulkInviteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BulkInviteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBulkInviteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "bulkInviteOrganization",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/invitations/bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BulkInviteOrganizationReader{formats: a.formats},
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
	success, ok := result.(*BulkInviteOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bulkInviteOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteInvite invites a user to an organization

  Revoke a user's invitation
*/
func (a *Client) DeleteInvite(params *DeleteInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInvite",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationName}/invitations/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteInviteReader{formats: a.formats},
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
	success, ok := result.(*DeleteInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOrganization deletes an organization

  Delete an organization
*/
func (a *Client) DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrganization",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationIdentifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrganizationReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteOrganizationAutoInvite removes a domain from an organization s auto invite settings

  Remove a domain from an organization's auto-invite settings
*/
func (a *Client) DeleteOrganizationAutoInvite(params *DeleteOrganizationAutoInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationAutoInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationAutoInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteOrganizationAutoInvite",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationName}/autoinvite/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrganizationAutoInviteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrganizationAutoInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteOrganizationAutoInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EditBilling updates an orgs billing information

  Update an orgs billing information
*/
func (a *Client) EditBilling(params *EditBillingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditBillingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditBillingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editBilling",
		Method:             "PUT",
		PathPattern:        "/organizations/{organizationIdentifier}/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditBillingReader{formats: a.formats},
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
	success, ok := result.(*EditBillingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for editBilling: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EditMember edits a member of an organization

  Edit a member of an organization
*/
func (a *Client) EditMember(params *EditMemberParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditMemberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditMemberParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editMember",
		Method:             "PATCH",
		PathPattern:        "/organizations/{organizationName}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditMemberReader{formats: a.formats},
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
	success, ok := result.(*EditMemberOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for editMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EditOrganization edits an organization

  Edit an organization
*/
func (a *Client) EditOrganization(params *EditOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EditOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "editOrganization",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationIdentifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditOrganizationReader{formats: a.formats},
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
	success, ok := result.(*EditOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for editOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBilling retrieves an orgs billing information

  Retrieve an orgs billing information
*/
func (a *Client) GetBilling(params *GetBillingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBillingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBilling",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationIdentifier}/billing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBillingReader{formats: a.formats},
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
	success, ok := result.(*GetBillingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBilling: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNextMutationID nexts mutation ID

  Get the id that the next mutation of this org should use
*/
func (a *Client) GetNextMutationID(params *GetNextMutationIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNextMutationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNextMutationIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNextMutationID",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationIdentifier}/nextMutationID",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNextMutationIDReader{formats: a.formats},
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
	success, ok := result.(*GetNextMutationIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNextMutationID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganization retrieves an organization

  Return a specific organization
*/
func (a *Client) GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganization",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationIdentifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationAutoInvite organizations auto invite settings

  Return organization auto-invite settings
*/
func (a *Client) GetOrganizationAutoInvite(params *GetOrganizationAutoInviteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationAutoInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationAutoInviteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationAutoInvite",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/autoinvite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationAutoInviteReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationAutoInviteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationAutoInvite: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationInvitations organizations invitations

  Return a list of pending invitations
*/
func (a *Client) GetOrganizationInvitations(params *GetOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationInvitationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationInvitationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationInvitations",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationInvitationsReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationInvitationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationInvitations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationMembers organizations membership

  Return a list of users who are members of the organization
*/
func (a *Client) GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationMembers",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationMembersReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationMutations gets history of mutations applied to an organization

  Query mutation records for the org
*/
func (a *Client) GetOrganizationMutations(params *GetOrganizationMutationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationMutationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMutationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationMutations",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationIdentifier}/mutations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationMutationsReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationMutationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationMutations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrganizationTier gets information about an organization s tier

  Get information about an organization's tier
*/
func (a *Client) GetOrganizationTier(params *GetOrganizationTierParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationTierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationTierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrganizationTier",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationIdentifier}/tier",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationTierReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationTierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrganizationTier: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InviteOrganization invites a user to an organization

  Invite a user to an organization's roster
*/
func (a *Client) InviteOrganization(params *InviteOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InviteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInviteOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "inviteOrganization",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/invitations/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InviteOrganizationReader{formats: a.formats},
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
	success, ok := result.(*InviteOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inviteOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  JoinOrganization joins a user to an organization

  Add a user to an organization's roster
*/
func (a *Client) JoinOrganization(params *JoinOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*JoinOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJoinOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "joinOrganization",
		Method:             "PUT",
		PathPattern:        "/organizations/{organizationName}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JoinOrganizationReader{formats: a.formats},
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
	success, ok := result.(*JoinOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for joinOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KomodoAuthorized is user authorized to use komodo ID e

  Check that the authenticated user is permitted to use Komodo IDE
*/
func (a *Client) KomodoAuthorized(params *KomodoAuthorizedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KomodoAuthorizedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKomodoAuthorizedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "komodoAuthorized",
		Method:             "GET",
		PathPattern:        "/status/komodo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KomodoAuthorizedReader{formats: a.formats},
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
	success, ok := result.(*KomodoAuthorizedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for komodoAuthorized: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListOrganizations lists of visible organizations

  Retrieve all organizations from the system that the user has access to
*/
func (a *Client) ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOrganizations",
		Method:             "GET",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOrganizationsReader{formats: a.formats},
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
	success, ok := result.(*ListOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MutateOrganization mutates organization

  Perform an atomic mutation on the org
*/
func (a *Client) MutateOrganization(params *MutateOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MutateOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMutateOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "mutateOrganization",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationIdentifier}/mutations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MutateOrganizationReader{formats: a.formats},
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
	success, ok := result.(*MutateOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mutateOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QuitOrganization drops a user from an organization

  Remove a user from an organization's roster
*/
func (a *Client) QuitOrganization(params *QuitOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*QuitOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuitOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quitOrganization",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationName}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QuitOrganizationReader{formats: a.formats},
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
	success, ok := result.(*QuitOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for quitOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBillingDate changes billing date for organization

  Set a new billing date
*/
func (a *Client) UpdateBillingDate(params *UpdateBillingDateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBillingDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBillingDateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBillingDate",
		Method:             "PUT",
		PathPattern:        "/admin/organizations/{organizationName}/updateBillingDate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateBillingDateReader{formats: a.formats},
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
	success, ok := result.(*UpdateBillingDateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBillingDate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
