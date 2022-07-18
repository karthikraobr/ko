// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

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

// IntotoV001Schema intoto v0.0.1 Schema
//
// Schema for intoto object
//
// swagger:model intotoV001Schema
type IntotoV001Schema struct {

	// content
	// Required: true
	Content *IntotoV001SchemaContent `json:"content"`

	// The public key that can verify the signature
	// Required: true
	// Format: byte
	PublicKey *strfmt.Base64 `json:"publicKey"`
}

// Validate validates this intoto v001 schema
func (m *IntotoV001Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV001Schema) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV001Schema) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this intoto v001 schema based on the context it is used
func (m *IntotoV001Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV001Schema) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV001Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV001Schema) UnmarshalBinary(b []byte) error {
	var res IntotoV001Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV001SchemaContent intoto v001 schema content
//
// swagger:model IntotoV001SchemaContent
type IntotoV001SchemaContent struct {

	// envelope
	Envelope string `json:"envelope,omitempty"`

	// hash
	Hash *IntotoV001SchemaContentHash `json:"hash,omitempty"`

	// payload hash
	PayloadHash *IntotoV001SchemaContentPayloadHash `json:"payloadHash,omitempty"`
}

// Validate validates this intoto v001 schema content
func (m *IntotoV001SchemaContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayloadHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV001SchemaContent) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV001SchemaContent) validatePayloadHash(formats strfmt.Registry) error {
	if swag.IsZero(m.PayloadHash) { // not required
		return nil
	}

	if m.PayloadHash != nil {
		if err := m.PayloadHash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "payloadHash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "payloadHash")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this intoto v001 schema content based on the context it is used
func (m *IntotoV001SchemaContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayloadHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV001SchemaContent) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {
		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV001SchemaContent) contextValidatePayloadHash(ctx context.Context, formats strfmt.Registry) error {

	if m.PayloadHash != nil {
		if err := m.PayloadHash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "payloadHash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "payloadHash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV001SchemaContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV001SchemaContent) UnmarshalBinary(b []byte) error {
	var res IntotoV001SchemaContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV001SchemaContentHash Specifies the hash algorithm and value encompassing the entire signed envelope
//
// swagger:model IntotoV001SchemaContentHash
type IntotoV001SchemaContentHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value for the archive
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this intoto v001 schema content hash
func (m *IntotoV001SchemaContentHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var intotoV001SchemaContentHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		intotoV001SchemaContentHashTypeAlgorithmPropEnum = append(intotoV001SchemaContentHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// IntotoV001SchemaContentHashAlgorithmSha256 captures enum value "sha256"
	IntotoV001SchemaContentHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *IntotoV001SchemaContentHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, intotoV001SchemaContentHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntotoV001SchemaContentHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("content"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *IntotoV001SchemaContentHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this intoto v001 schema content hash based on the context it is used
func (m *IntotoV001SchemaContentHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV001SchemaContentHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV001SchemaContentHash) UnmarshalBinary(b []byte) error {
	var res IntotoV001SchemaContentHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV001SchemaContentPayloadHash Specifies the hash algorithm and value covering the payload within the DSSE envelope
//
// swagger:model IntotoV001SchemaContentPayloadHash
type IntotoV001SchemaContentPayloadHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value for the envelope's payload
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this intoto v001 schema content payload hash
func (m *IntotoV001SchemaContentPayloadHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var intotoV001SchemaContentPayloadHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		intotoV001SchemaContentPayloadHashTypeAlgorithmPropEnum = append(intotoV001SchemaContentPayloadHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// IntotoV001SchemaContentPayloadHashAlgorithmSha256 captures enum value "sha256"
	IntotoV001SchemaContentPayloadHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *IntotoV001SchemaContentPayloadHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, intotoV001SchemaContentPayloadHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntotoV001SchemaContentPayloadHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"payloadHash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("content"+"."+"payloadHash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *IntotoV001SchemaContentPayloadHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"payloadHash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this intoto v001 schema content payload hash based on the context it is used
func (m *IntotoV001SchemaContentPayloadHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV001SchemaContentPayloadHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV001SchemaContentPayloadHash) UnmarshalBinary(b []byte) error {
	var res IntotoV001SchemaContentPayloadHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}