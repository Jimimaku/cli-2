// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CurrentUser current user
//
// swagger:model current_user
type CurrentUser struct {

	// email
	Email string `json:"email,omitempty"`

	// is activator
	IsActivator bool `json:"is_activator,omitempty"`

	// is anonymous
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// is superuser
	IsSuperuser bool `json:"is_superuser,omitempty"`

	// is superuser readonly
	IsSuperuserReadonly bool `json:"is_superuser_readonly,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organizations
	Organizations []*Organization `json:"organizations"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this current user
func (m *CurrentUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentUser) validateOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {
		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this current user based on the context it is used
func (m *CurrentUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentUser) contextValidateOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Organizations); i++ {

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrentUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentUser) UnmarshalBinary(b []byte) error {
	var res CurrentUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
