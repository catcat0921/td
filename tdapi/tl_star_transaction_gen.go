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

// StarTransaction represents TL type `starTransaction#ce1084a8`.
type StarTransaction struct {
	// Unique identifier of the transaction
	ID string
	// The amount of added owned Telegram Stars; negative for outgoing transactions
	StarCount int64
	// True, if the transaction is a refund of a previous transaction
	IsRefund bool
	// Point in time (Unix timestamp) when the transaction was completed
	Date int32
	// Source of the incoming transaction, or its recipient for outgoing transactions
	Partner StarTransactionPartnerClass
}

// StarTransactionTypeID is TL type id of StarTransaction.
const StarTransactionTypeID = 0xce1084a8

// Ensuring interfaces in compile-time for StarTransaction.
var (
	_ bin.Encoder     = &StarTransaction{}
	_ bin.Decoder     = &StarTransaction{}
	_ bin.BareEncoder = &StarTransaction{}
	_ bin.BareDecoder = &StarTransaction{}
)

func (s *StarTransaction) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.StarCount == 0) {
		return false
	}
	if !(s.IsRefund == false) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.Partner == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarTransaction) String() string {
	if s == nil {
		return "StarTransaction(nil)"
	}
	type Alias StarTransaction
	return fmt.Sprintf("StarTransaction%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarTransaction) TypeID() uint32 {
	return StarTransactionTypeID
}

// TypeName returns name of type in TL schema.
func (*StarTransaction) TypeName() string {
	return "starTransaction"
}

// TypeInfo returns info about TL type.
func (s *StarTransaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starTransaction",
		ID:   StarTransactionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "StarCount",
			SchemaName: "star_count",
		},
		{
			Name:       "IsRefund",
			SchemaName: "is_refund",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Partner",
			SchemaName: "partner",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarTransaction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransaction#ce1084a8 as nil")
	}
	b.PutID(StarTransactionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarTransaction) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransaction#ce1084a8 as nil")
	}
	b.PutString(s.ID)
	b.PutInt53(s.StarCount)
	b.PutBool(s.IsRefund)
	b.PutInt32(s.Date)
	if s.Partner == nil {
		return fmt.Errorf("unable to encode starTransaction#ce1084a8: field partner is nil")
	}
	if err := s.Partner.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starTransaction#ce1084a8: field partner: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarTransaction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransaction#ce1084a8 to nil")
	}
	if err := b.ConsumeID(StarTransactionTypeID); err != nil {
		return fmt.Errorf("unable to decode starTransaction#ce1084a8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarTransaction) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransaction#ce1084a8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starTransaction#ce1084a8: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode starTransaction#ce1084a8: field star_count: %w", err)
		}
		s.StarCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode starTransaction#ce1084a8: field is_refund: %w", err)
		}
		s.IsRefund = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode starTransaction#ce1084a8: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := DecodeStarTransactionPartner(b)
		if err != nil {
			return fmt.Errorf("unable to decode starTransaction#ce1084a8: field partner: %w", err)
		}
		s.Partner = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StarTransaction) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode starTransaction#ce1084a8 as nil")
	}
	b.ObjStart()
	b.PutID("starTransaction")
	b.Comma()
	b.FieldStart("id")
	b.PutString(s.ID)
	b.Comma()
	b.FieldStart("star_count")
	b.PutInt53(s.StarCount)
	b.Comma()
	b.FieldStart("is_refund")
	b.PutBool(s.IsRefund)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(s.Date)
	b.Comma()
	b.FieldStart("partner")
	if s.Partner == nil {
		return fmt.Errorf("unable to encode starTransaction#ce1084a8: field partner is nil")
	}
	if err := s.Partner.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode starTransaction#ce1084a8: field partner: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StarTransaction) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode starTransaction#ce1084a8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("starTransaction"); err != nil {
				return fmt.Errorf("unable to decode starTransaction#ce1084a8: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode starTransaction#ce1084a8: field id: %w", err)
			}
			s.ID = value
		case "star_count":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode starTransaction#ce1084a8: field star_count: %w", err)
			}
			s.StarCount = value
		case "is_refund":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode starTransaction#ce1084a8: field is_refund: %w", err)
			}
			s.IsRefund = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode starTransaction#ce1084a8: field date: %w", err)
			}
			s.Date = value
		case "partner":
			value, err := DecodeTDLibJSONStarTransactionPartner(b)
			if err != nil {
				return fmt.Errorf("unable to decode starTransaction#ce1084a8: field partner: %w", err)
			}
			s.Partner = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *StarTransaction) GetID() (value string) {
	if s == nil {
		return
	}
	return s.ID
}

// GetStarCount returns value of StarCount field.
func (s *StarTransaction) GetStarCount() (value int64) {
	if s == nil {
		return
	}
	return s.StarCount
}

// GetIsRefund returns value of IsRefund field.
func (s *StarTransaction) GetIsRefund() (value bool) {
	if s == nil {
		return
	}
	return s.IsRefund
}

// GetDate returns value of Date field.
func (s *StarTransaction) GetDate() (value int32) {
	if s == nil {
		return
	}
	return s.Date
}

// GetPartner returns value of Partner field.
func (s *StarTransaction) GetPartner() (value StarTransactionPartnerClass) {
	if s == nil {
		return
	}
	return s.Partner
}
