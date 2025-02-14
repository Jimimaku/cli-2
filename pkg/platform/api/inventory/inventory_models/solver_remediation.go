// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SolverRemediation Solver Remediation
//
// An action that a user can take to attempt to resolve a solver error
//
// swagger:model solverRemediation
type SolverRemediation struct {

	// An explanation of this remediation
	// Required: true
	Command *string `json:"command"`

	// parameters
	Parameters *SolverRemediationParameters `json:"parameters,omitempty"`

	// The remediation type. Different types of remediations have different fields.
	// Required: true
	// Enum: [ADVANCE_TIMESTAMP CHANGE_REQUIREMENT_NAME CHANGE_REQUIREMENT_VERSION REMOVE_PLATFORM REMOVE_REQUIREMENT REQUEST_PACKAGE_IMPORT]
	RemediationType *string `json:"remediation_type"`
}

// Validate validates this solver remediation
func (m *SolverRemediation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemediationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SolverRemediation) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

func (m *SolverRemediation) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

var solverRemediationTypeRemediationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADVANCE_TIMESTAMP","CHANGE_REQUIREMENT_NAME","CHANGE_REQUIREMENT_VERSION","REMOVE_PLATFORM","REMOVE_REQUIREMENT","REQUEST_PACKAGE_IMPORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		solverRemediationTypeRemediationTypePropEnum = append(solverRemediationTypeRemediationTypePropEnum, v)
	}
}

const (

	// SolverRemediationRemediationTypeADVANCETIMESTAMP captures enum value "ADVANCE_TIMESTAMP"
	SolverRemediationRemediationTypeADVANCETIMESTAMP string = "ADVANCE_TIMESTAMP"

	// SolverRemediationRemediationTypeCHANGEREQUIREMENTNAME captures enum value "CHANGE_REQUIREMENT_NAME"
	SolverRemediationRemediationTypeCHANGEREQUIREMENTNAME string = "CHANGE_REQUIREMENT_NAME"

	// SolverRemediationRemediationTypeCHANGEREQUIREMENTVERSION captures enum value "CHANGE_REQUIREMENT_VERSION"
	SolverRemediationRemediationTypeCHANGEREQUIREMENTVERSION string = "CHANGE_REQUIREMENT_VERSION"

	// SolverRemediationRemediationTypeREMOVEPLATFORM captures enum value "REMOVE_PLATFORM"
	SolverRemediationRemediationTypeREMOVEPLATFORM string = "REMOVE_PLATFORM"

	// SolverRemediationRemediationTypeREMOVEREQUIREMENT captures enum value "REMOVE_REQUIREMENT"
	SolverRemediationRemediationTypeREMOVEREQUIREMENT string = "REMOVE_REQUIREMENT"

	// SolverRemediationRemediationTypeREQUESTPACKAGEIMPORT captures enum value "REQUEST_PACKAGE_IMPORT"
	SolverRemediationRemediationTypeREQUESTPACKAGEIMPORT string = "REQUEST_PACKAGE_IMPORT"
)

// prop value enum
func (m *SolverRemediation) validateRemediationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, solverRemediationTypeRemediationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SolverRemediation) validateRemediationType(formats strfmt.Registry) error {

	if err := validate.Required("remediation_type", "body", m.RemediationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRemediationTypeEnum("remediation_type", "body", *m.RemediationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this solver remediation based on the context it is used
func (m *SolverRemediation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SolverRemediation) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SolverRemediation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SolverRemediation) UnmarshalBinary(b []byte) error {
	var res SolverRemediation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
