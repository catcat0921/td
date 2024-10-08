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

// UserGifts represents TL type `userGifts#e274219f`.
type UserGifts struct {
	// The total number of received gifts
	TotalCount int32
	// The list of gifts
	Gifts []UserGift
	// The offset for the next request. If empty, then there are no more results
	NextOffset string
}

// UserGiftsTypeID is TL type id of UserGifts.
const UserGiftsTypeID = 0xe274219f

// Ensuring interfaces in compile-time for UserGifts.
var (
	_ bin.Encoder     = &UserGifts{}
	_ bin.Decoder     = &UserGifts{}
	_ bin.BareEncoder = &UserGifts{}
	_ bin.BareDecoder = &UserGifts{}
)

func (u *UserGifts) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.TotalCount == 0) {
		return false
	}
	if !(u.Gifts == nil) {
		return false
	}
	if !(u.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserGifts) String() string {
	if u == nil {
		return "UserGifts(nil)"
	}
	type Alias UserGifts
	return fmt.Sprintf("UserGifts%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserGifts) TypeID() uint32 {
	return UserGiftsTypeID
}

// TypeName returns name of type in TL schema.
func (*UserGifts) TypeName() string {
	return "userGifts"
}

// TypeInfo returns info about TL type.
func (u *UserGifts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userGifts",
		ID:   UserGiftsTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Gifts",
			SchemaName: "gifts",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserGifts) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userGifts#e274219f as nil")
	}
	b.PutID(UserGiftsTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserGifts) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userGifts#e274219f as nil")
	}
	b.PutInt32(u.TotalCount)
	b.PutInt(len(u.Gifts))
	for idx, v := range u.Gifts {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare userGifts#e274219f: field gifts element with index %d: %w", idx, err)
		}
	}
	b.PutString(u.NextOffset)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserGifts) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userGifts#e274219f to nil")
	}
	if err := b.ConsumeID(UserGiftsTypeID); err != nil {
		return fmt.Errorf("unable to decode userGifts#e274219f: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserGifts) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userGifts#e274219f to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userGifts#e274219f: field total_count: %w", err)
		}
		u.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userGifts#e274219f: field gifts: %w", err)
		}

		if headerLen > 0 {
			u.Gifts = make([]UserGift, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value UserGift
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare userGifts#e274219f: field gifts: %w", err)
			}
			u.Gifts = append(u.Gifts, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userGifts#e274219f: field next_offset: %w", err)
		}
		u.NextOffset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UserGifts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode userGifts#e274219f as nil")
	}
	b.ObjStart()
	b.PutID("userGifts")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(u.TotalCount)
	b.Comma()
	b.FieldStart("gifts")
	b.ArrStart()
	for idx, v := range u.Gifts {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode userGifts#e274219f: field gifts element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("next_offset")
	b.PutString(u.NextOffset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UserGifts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode userGifts#e274219f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("userGifts"); err != nil {
				return fmt.Errorf("unable to decode userGifts#e274219f: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode userGifts#e274219f: field total_count: %w", err)
			}
			u.TotalCount = value
		case "gifts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value UserGift
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode userGifts#e274219f: field gifts: %w", err)
				}
				u.Gifts = append(u.Gifts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode userGifts#e274219f: field gifts: %w", err)
			}
		case "next_offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode userGifts#e274219f: field next_offset: %w", err)
			}
			u.NextOffset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (u *UserGifts) GetTotalCount() (value int32) {
	if u == nil {
		return
	}
	return u.TotalCount
}

// GetGifts returns value of Gifts field.
func (u *UserGifts) GetGifts() (value []UserGift) {
	if u == nil {
		return
	}
	return u.Gifts
}

// GetNextOffset returns value of NextOffset field.
func (u *UserGifts) GetNextOffset() (value string) {
	if u == nil {
		return
	}
	return u.NextOffset
}
