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

// AuthImportBotAuthorizationRequest represents TL type `auth.importBotAuthorization#67a3ff2c`.
// Login as a bot
//
// See https://core.telegram.org/method/auth.importBotAuthorization for reference.
type AuthImportBotAuthorizationRequest struct {
	// Reserved for future use
	Flags int
	// Application identifier (see. App configuration¹)
	//
	// Links:
	//  1) https://core.telegram.org/myapp
	APIID int
	// Application identifier hash (see. App configuration¹)
	//
	// Links:
	//  1) https://core.telegram.org/myapp
	APIHash string
	// Bot token (see bots¹)
	//
	// Links:
	//  1) https://core.telegram.org/bots
	BotAuthToken string
}

// AuthImportBotAuthorizationRequestTypeID is TL type id of AuthImportBotAuthorizationRequest.
const AuthImportBotAuthorizationRequestTypeID = 0x67a3ff2c

func (i *AuthImportBotAuthorizationRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags == 0) {
		return false
	}
	if !(i.APIID == 0) {
		return false
	}
	if !(i.APIHash == "") {
		return false
	}
	if !(i.BotAuthToken == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AuthImportBotAuthorizationRequest) String() string {
	if i == nil {
		return "AuthImportBotAuthorizationRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthImportBotAuthorizationRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(i.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tAPIID: ")
	sb.WriteString(fmt.Sprint(i.APIID))
	sb.WriteString(",\n")
	sb.WriteString("\tAPIHash: ")
	sb.WriteString(fmt.Sprint(i.APIHash))
	sb.WriteString(",\n")
	sb.WriteString("\tBotAuthToken: ")
	sb.WriteString(fmt.Sprint(i.BotAuthToken))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *AuthImportBotAuthorizationRequest) TypeID() uint32 {
	return AuthImportBotAuthorizationRequestTypeID
}

// Encode implements bin.Encoder.
func (i *AuthImportBotAuthorizationRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importBotAuthorization#67a3ff2c as nil")
	}
	b.PutID(AuthImportBotAuthorizationRequestTypeID)
	b.PutInt(i.Flags)
	b.PutInt(i.APIID)
	b.PutString(i.APIHash)
	b.PutString(i.BotAuthToken)
	return nil
}

// GetFlags returns value of Flags field.
func (i *AuthImportBotAuthorizationRequest) GetFlags() (value int) {
	return i.Flags
}

// GetAPIID returns value of APIID field.
func (i *AuthImportBotAuthorizationRequest) GetAPIID() (value int) {
	return i.APIID
}

// GetAPIHash returns value of APIHash field.
func (i *AuthImportBotAuthorizationRequest) GetAPIHash() (value string) {
	return i.APIHash
}

// GetBotAuthToken returns value of BotAuthToken field.
func (i *AuthImportBotAuthorizationRequest) GetBotAuthToken() (value string) {
	return i.BotAuthToken
}

// Decode implements bin.Decoder.
func (i *AuthImportBotAuthorizationRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importBotAuthorization#67a3ff2c to nil")
	}
	if err := b.ConsumeID(AuthImportBotAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field flags: %w", err)
		}
		i.Flags = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field api_id: %w", err)
		}
		i.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field api_hash: %w", err)
		}
		i.APIHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field bot_auth_token: %w", err)
		}
		i.BotAuthToken = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthImportBotAuthorizationRequest.
var (
	_ bin.Encoder = &AuthImportBotAuthorizationRequest{}
	_ bin.Decoder = &AuthImportBotAuthorizationRequest{}
)

// AuthImportBotAuthorization invokes method auth.importBotAuthorization#67a3ff2c returning error if any.
// Login as a bot
//
// Possible errors:
//  400 ACCESS_TOKEN_EXPIRED: Bot token expired
//  400 ACCESS_TOKEN_INVALID: The provided token is not valid
//  400 API_ID_INVALID: The api_id/api_hash combination is invalid
//  401 AUTH_KEY_INVALID: Auth key invalid
//
// See https://core.telegram.org/method/auth.importBotAuthorization for reference.
// Can be used by bots.
func (c *Client) AuthImportBotAuthorization(ctx context.Context, request *AuthImportBotAuthorizationRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
