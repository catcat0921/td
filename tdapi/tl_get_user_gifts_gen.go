// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// GetUserGiftsRequest represents TL type `getUserGifts#b7cdbc7c`.
type GetUserGiftsRequest struct {
	// Identifier of the user
	UserID int64
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of gifts to be returned; must be positive and can't be greater than
	// 100. For optimal performance, the number of returned objects is chosen by TDLib and
	// can be smaller than the specified limit
	Limit int32
}

// GetUserGiftsRequestTypeID is TL type id of GetUserGiftsRequest.
const GetUserGiftsRequestTypeID = 0xb7cdbc7c

// Ensuring interfaces in compile-time for GetUserGiftsRequest.
var (
	_ bin.Encoder     = &GetUserGiftsRequest{}
	_ bin.Decoder     = &GetUserGiftsRequest{}
	_ bin.BareEncoder = &GetUserGiftsRequest{}
	_ bin.BareDecoder = &GetUserGiftsRequest{}
)

func (g *GetUserGiftsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUserGiftsRequest) String() string {
	if g == nil {
		return "GetUserGiftsRequest(nil)"
	}
	type Alias GetUserGiftsRequest
	return fmt.Sprintf("GetUserGiftsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUserGiftsRequest) TypeID() uint32 {
	return GetUserGiftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUserGiftsRequest) TypeName() string {
	return "getUserGifts"
}

// TypeInfo returns info about TL type.
func (g *GetUserGiftsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUserGifts",
		ID:   GetUserGiftsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetUserGiftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserGifts#b7cdbc7c as nil")
	}
	b.PutID(GetUserGiftsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetUserGiftsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserGifts#b7cdbc7c as nil")
	}
	b.PutInt53(g.UserID)
	b.PutString(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUserGiftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserGifts#b7cdbc7c to nil")
	}
	if err := b.ConsumeID(GetUserGiftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetUserGiftsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserGifts#b7cdbc7c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetUserGiftsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserGifts#b7cdbc7c as nil")
	}
	b.ObjStart()
	b.PutID("getUserGifts")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetUserGiftsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserGifts#b7cdbc7c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getUserGifts"); err != nil {
				return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: field user_id: %w", err)
			}
			g.UserID = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getUserGifts#b7cdbc7c: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (g *GetUserGiftsRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetOffset returns value of Offset field.
func (g *GetUserGiftsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetUserGiftsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetUserGifts invokes method getUserGifts#b7cdbc7c returning error if any.
func (c *Client) GetUserGifts(ctx context.Context, request *GetUserGiftsRequest) (*UserGifts, error) {
	var result UserGifts

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
