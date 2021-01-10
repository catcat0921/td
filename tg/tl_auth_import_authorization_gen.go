// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// AuthImportAuthorizationRequest represents TL type `auth.importAuthorization#e3ef9613`.
// Logs in a user using a key transmitted from his native data-centre.
//
// See https://core.telegram.org/method/auth.importAuthorization for reference.
type AuthImportAuthorizationRequest struct {
	// User ID
	ID int
	// Authorization key
	Bytes []byte
}

// AuthImportAuthorizationRequestTypeID is TL type id of AuthImportAuthorizationRequest.
const AuthImportAuthorizationRequestTypeID = 0xe3ef9613

func (i *AuthImportAuthorizationRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AuthImportAuthorizationRequest) String() string {
	if i == nil {
		return "AuthImportAuthorizationRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthImportAuthorizationRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tBytes: ")
	sb.WriteString(fmt.Sprint(i.Bytes))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *AuthImportAuthorizationRequest) TypeID() uint32 {
	return AuthImportAuthorizationRequestTypeID
}

// Encode implements bin.Encoder.
func (i *AuthImportAuthorizationRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importAuthorization#e3ef9613 as nil")
	}
	b.PutID(AuthImportAuthorizationRequestTypeID)
	b.PutInt(i.ID)
	b.PutBytes(i.Bytes)
	return nil
}

// GetID returns value of ID field.
func (i *AuthImportAuthorizationRequest) GetID() (value int) {
	return i.ID
}

// GetBytes returns value of Bytes field.
func (i *AuthImportAuthorizationRequest) GetBytes() (value []byte) {
	return i.Bytes
}

// Decode implements bin.Decoder.
func (i *AuthImportAuthorizationRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importAuthorization#e3ef9613 to nil")
	}
	if err := b.ConsumeID(AuthImportAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importAuthorization#e3ef9613: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importAuthorization#e3ef9613: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importAuthorization#e3ef9613: field bytes: %w", err)
		}
		i.Bytes = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthImportAuthorizationRequest.
var (
	_ bin.Encoder = &AuthImportAuthorizationRequest{}
	_ bin.Decoder = &AuthImportAuthorizationRequest{}
)

// AuthImportAuthorization invokes method auth.importAuthorization#e3ef9613 returning error if any.
// Logs in a user using a key transmitted from his native data-centre.
//
// Possible errors:
//  400 AUTH_BYTES_INVALID: The provided authorization is invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/auth.importAuthorization for reference.
// Can be used by bots.
func (c *Client) AuthImportAuthorization(ctx context.Context, request *AuthImportAuthorizationRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
