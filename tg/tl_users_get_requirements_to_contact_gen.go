// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// UsersGetRequirementsToContactRequest represents TL type `users.getRequirementsToContact#d89a83a3`.
//
// See https://core.telegram.org/method/users.getRequirementsToContact for reference.
type UsersGetRequirementsToContactRequest struct {
	// ID field of UsersGetRequirementsToContactRequest.
	ID []InputUserClass
}

// UsersGetRequirementsToContactRequestTypeID is TL type id of UsersGetRequirementsToContactRequest.
const UsersGetRequirementsToContactRequestTypeID = 0xd89a83a3

// Ensuring interfaces in compile-time for UsersGetRequirementsToContactRequest.
var (
	_ bin.Encoder     = &UsersGetRequirementsToContactRequest{}
	_ bin.Decoder     = &UsersGetRequirementsToContactRequest{}
	_ bin.BareEncoder = &UsersGetRequirementsToContactRequest{}
	_ bin.BareDecoder = &UsersGetRequirementsToContactRequest{}
)

func (g *UsersGetRequirementsToContactRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UsersGetRequirementsToContactRequest) String() string {
	if g == nil {
		return "UsersGetRequirementsToContactRequest(nil)"
	}
	type Alias UsersGetRequirementsToContactRequest
	return fmt.Sprintf("UsersGetRequirementsToContactRequest%+v", Alias(*g))
}

// FillFrom fills UsersGetRequirementsToContactRequest from given interface.
func (g *UsersGetRequirementsToContactRequest) FillFrom(from interface {
	GetID() (value []InputUserClass)
}) {
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UsersGetRequirementsToContactRequest) TypeID() uint32 {
	return UsersGetRequirementsToContactRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UsersGetRequirementsToContactRequest) TypeName() string {
	return "users.getRequirementsToContact"
}

// TypeInfo returns info about TL type.
func (g *UsersGetRequirementsToContactRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "users.getRequirementsToContact",
		ID:   UsersGetRequirementsToContactRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UsersGetRequirementsToContactRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getRequirementsToContact#d89a83a3 as nil")
	}
	b.PutID(UsersGetRequirementsToContactRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UsersGetRequirementsToContactRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getRequirementsToContact#d89a83a3 as nil")
	}
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode users.getRequirementsToContact#d89a83a3: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.getRequirementsToContact#d89a83a3: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *UsersGetRequirementsToContactRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getRequirementsToContact#d89a83a3 to nil")
	}
	if err := b.ConsumeID(UsersGetRequirementsToContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode users.getRequirementsToContact#d89a83a3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UsersGetRequirementsToContactRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getRequirementsToContact#d89a83a3 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.getRequirementsToContact#d89a83a3: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.getRequirementsToContact#d89a83a3: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (g *UsersGetRequirementsToContactRequest) GetID() (value []InputUserClass) {
	if g == nil {
		return
	}
	return g.ID
}

// MapID returns field ID wrapped in InputUserClassArray helper.
func (g *UsersGetRequirementsToContactRequest) MapID() (value InputUserClassArray) {
	return InputUserClassArray(g.ID)
}

// UsersGetRequirementsToContact invokes method users.getRequirementsToContact#d89a83a3 returning error if any.
//
// See https://core.telegram.org/method/users.getRequirementsToContact for reference.
func (c *Client) UsersGetRequirementsToContact(ctx context.Context, id []InputUserClass) ([]RequirementToContactClass, error) {
	var result RequirementToContactClassVector

	request := &UsersGetRequirementsToContactRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []RequirementToContactClass(result.Elems), nil
}
