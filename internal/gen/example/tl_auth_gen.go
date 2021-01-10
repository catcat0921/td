// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Auth represents TL type `auth#f8bb4a38`.
//
// See https://localhost:80/doc/constructor/auth for reference.
type Auth struct {
	// Name field of Auth.
	Name string
}

// AuthTypeID is TL type id of Auth.
const AuthTypeID = 0xf8bb4a38

func (a *Auth) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *Auth) String() string {
	if a == nil {
		return "Auth(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Auth")
	sb.WriteString("{\n")
	sb.WriteString("\tName: ")
	sb.WriteString(fmt.Sprint(a.Name))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *Auth) TypeID() uint32 {
	return AuthTypeID
}

// Encode implements bin.Encoder.
func (a *Auth) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth#f8bb4a38 as nil")
	}
	b.PutID(AuthTypeID)
	b.PutString(a.Name)
	return nil
}

// GetName returns value of Name field.
func (a *Auth) GetName() (value string) {
	return a.Name
}

// Decode implements bin.Decoder.
func (a *Auth) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth#f8bb4a38 to nil")
	}
	if err := b.ConsumeID(AuthTypeID); err != nil {
		return fmt.Errorf("unable to decode auth#f8bb4a38: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth#f8bb4a38: field name: %w", err)
		}
		a.Name = value
	}
	return nil
}

// construct implements constructor of AuthClass.
func (a Auth) construct() AuthClass { return &a }

// Ensuring interfaces in compile-time for Auth.
var (
	_ bin.Encoder = &Auth{}
	_ bin.Decoder = &Auth{}

	_ AuthClass = &Auth{}
)

// AuthPassword represents TL type `authPassword#29bacabb`.
//
// See https://localhost:80/doc/constructor/authPassword for reference.
type AuthPassword struct {
	// Name field of AuthPassword.
	Name string
	// Password field of AuthPassword.
	Password string
}

// AuthPasswordTypeID is TL type id of AuthPassword.
const AuthPasswordTypeID = 0x29bacabb

func (a *AuthPassword) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Name == "") {
		return false
	}
	if !(a.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthPassword) String() string {
	if a == nil {
		return "AuthPassword(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthPassword")
	sb.WriteString("{\n")
	sb.WriteString("\tName: ")
	sb.WriteString(fmt.Sprint(a.Name))
	sb.WriteString(",\n")
	sb.WriteString("\tPassword: ")
	sb.WriteString(fmt.Sprint(a.Password))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *AuthPassword) TypeID() uint32 {
	return AuthPasswordTypeID
}

// Encode implements bin.Encoder.
func (a *AuthPassword) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authPassword#29bacabb as nil")
	}
	b.PutID(AuthPasswordTypeID)
	b.PutString(a.Name)
	b.PutString(a.Password)
	return nil
}

// GetName returns value of Name field.
func (a *AuthPassword) GetName() (value string) {
	return a.Name
}

// GetPassword returns value of Password field.
func (a *AuthPassword) GetPassword() (value string) {
	return a.Password
}

// Decode implements bin.Decoder.
func (a *AuthPassword) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authPassword#29bacabb to nil")
	}
	if err := b.ConsumeID(AuthPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode authPassword#29bacabb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authPassword#29bacabb: field name: %w", err)
		}
		a.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authPassword#29bacabb: field password: %w", err)
		}
		a.Password = value
	}
	return nil
}

// construct implements constructor of AuthClass.
func (a AuthPassword) construct() AuthClass { return &a }

// Ensuring interfaces in compile-time for AuthPassword.
var (
	_ bin.Encoder = &AuthPassword{}
	_ bin.Decoder = &AuthPassword{}

	_ AuthClass = &AuthPassword{}
)

// AuthClass represents Auth generic type.
//
// See https://localhost:80/doc/type/Auth for reference.
//
// Example:
//  g, err := DecodeAuth(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *Auth: // auth#f8bb4a38
//  case *AuthPassword: // authPassword#29bacabb
//  default: panic(v)
//  }
type AuthClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthClass

	// Name field of Auth.
	GetName() (value string)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeAuth implements binary de-serialization for AuthClass.
func DecodeAuth(buf *bin.Buffer) (AuthClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthTypeID:
		// Decoding auth#f8bb4a38.
		v := Auth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthClass: %w", err)
		}
		return &v, nil
	case AuthPasswordTypeID:
		// Decoding authPassword#29bacabb.
		v := AuthPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthClass: %w", bin.NewUnexpectedID(id))
	}
}

// Auth boxes the AuthClass providing a helper.
type AuthBox struct {
	Auth AuthClass
}

// Decode implements bin.Decoder for AuthBox.
func (b *AuthBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthBox to nil")
	}
	v, err := DecodeAuth(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Auth = v
	return nil
}

// Encode implements bin.Encode for AuthBox.
func (b *AuthBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Auth == nil {
		return fmt.Errorf("unable to encode AuthClass as nil")
	}
	return b.Auth.Encode(buf)
}
