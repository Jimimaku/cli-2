// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CPUExtensionAllOf0 cpu extension all of0
//
// swagger:model cpuExtensionAllOf0
type CPUExtensionAllOf0 struct {

	// cpu extension id
	// Required: true
	// Format: uuid
	CPUExtensionID *strfmt.UUID `json:"cpu_extension_id"`

	// links
	// Required: true
	Links *SelfLink `json:"links"`
}

// Validate validates this cpu extension all of0
func (m *CPUExtensionAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUExtensionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUExtensionAllOf0) validateCPUExtensionID(formats strfmt.Registry) error {

	if err := validate.Required("cpu_extension_id", "body", m.CPUExtensionID); err != nil {
		return err
	}

	if err := validate.FormatOf("cpu_extension_id", "body", "uuid", m.CPUExtensionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CPUExtensionAllOf0) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cpu extension all of0 based on the context it is used
func (m *CPUExtensionAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUExtensionAllOf0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPUExtensionAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUExtensionAllOf0) UnmarshalBinary(b []byte) error {
	var res CPUExtensionAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
