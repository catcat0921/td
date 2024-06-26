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

// BroadcastRevenueBalances represents TL type `broadcastRevenueBalances#8438f1c6`.
//
// See https://core.telegram.org/constructor/broadcastRevenueBalances for reference.
type BroadcastRevenueBalances struct {
	// CurrentBalance field of BroadcastRevenueBalances.
	CurrentBalance int64
	// AvailableBalance field of BroadcastRevenueBalances.
	AvailableBalance int64
	// OverallRevenue field of BroadcastRevenueBalances.
	OverallRevenue int64
}

// BroadcastRevenueBalancesTypeID is TL type id of BroadcastRevenueBalances.
const BroadcastRevenueBalancesTypeID = 0x8438f1c6

// Ensuring interfaces in compile-time for BroadcastRevenueBalances.
var (
	_ bin.Encoder     = &BroadcastRevenueBalances{}
	_ bin.Decoder     = &BroadcastRevenueBalances{}
	_ bin.BareEncoder = &BroadcastRevenueBalances{}
	_ bin.BareDecoder = &BroadcastRevenueBalances{}
)

func (b *BroadcastRevenueBalances) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.CurrentBalance == 0) {
		return false
	}
	if !(b.AvailableBalance == 0) {
		return false
	}
	if !(b.OverallRevenue == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BroadcastRevenueBalances) String() string {
	if b == nil {
		return "BroadcastRevenueBalances(nil)"
	}
	type Alias BroadcastRevenueBalances
	return fmt.Sprintf("BroadcastRevenueBalances%+v", Alias(*b))
}

// FillFrom fills BroadcastRevenueBalances from given interface.
func (b *BroadcastRevenueBalances) FillFrom(from interface {
	GetCurrentBalance() (value int64)
	GetAvailableBalance() (value int64)
	GetOverallRevenue() (value int64)
}) {
	b.CurrentBalance = from.GetCurrentBalance()
	b.AvailableBalance = from.GetAvailableBalance()
	b.OverallRevenue = from.GetOverallRevenue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BroadcastRevenueBalances) TypeID() uint32 {
	return BroadcastRevenueBalancesTypeID
}

// TypeName returns name of type in TL schema.
func (*BroadcastRevenueBalances) TypeName() string {
	return "broadcastRevenueBalances"
}

// TypeInfo returns info about TL type.
func (b *BroadcastRevenueBalances) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "broadcastRevenueBalances",
		ID:   BroadcastRevenueBalancesTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CurrentBalance",
			SchemaName: "current_balance",
		},
		{
			Name:       "AvailableBalance",
			SchemaName: "available_balance",
		},
		{
			Name:       "OverallRevenue",
			SchemaName: "overall_revenue",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BroadcastRevenueBalances) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode broadcastRevenueBalances#8438f1c6 as nil")
	}
	buf.PutID(BroadcastRevenueBalancesTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BroadcastRevenueBalances) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode broadcastRevenueBalances#8438f1c6 as nil")
	}
	buf.PutLong(b.CurrentBalance)
	buf.PutLong(b.AvailableBalance)
	buf.PutLong(b.OverallRevenue)
	return nil
}

// Decode implements bin.Decoder.
func (b *BroadcastRevenueBalances) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode broadcastRevenueBalances#8438f1c6 to nil")
	}
	if err := buf.ConsumeID(BroadcastRevenueBalancesTypeID); err != nil {
		return fmt.Errorf("unable to decode broadcastRevenueBalances#8438f1c6: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BroadcastRevenueBalances) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode broadcastRevenueBalances#8438f1c6 to nil")
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode broadcastRevenueBalances#8438f1c6: field current_balance: %w", err)
		}
		b.CurrentBalance = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode broadcastRevenueBalances#8438f1c6: field available_balance: %w", err)
		}
		b.AvailableBalance = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode broadcastRevenueBalances#8438f1c6: field overall_revenue: %w", err)
		}
		b.OverallRevenue = value
	}
	return nil
}

// GetCurrentBalance returns value of CurrentBalance field.
func (b *BroadcastRevenueBalances) GetCurrentBalance() (value int64) {
	if b == nil {
		return
	}
	return b.CurrentBalance
}

// GetAvailableBalance returns value of AvailableBalance field.
func (b *BroadcastRevenueBalances) GetAvailableBalance() (value int64) {
	if b == nil {
		return
	}
	return b.AvailableBalance
}

// GetOverallRevenue returns value of OverallRevenue field.
func (b *BroadcastRevenueBalances) GetOverallRevenue() (value int64) {
	if b == nil {
		return
	}
	return b.OverallRevenue
}
