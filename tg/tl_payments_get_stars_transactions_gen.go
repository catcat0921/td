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

// PaymentsGetStarsTransactionsRequest represents TL type `payments.getStarsTransactions#69da4557`.
// Fetch Telegram Stars transactions¹.
// The inbound and outbound flags are mutually exclusive: if none of the two are set,
// both incoming and outgoing transactions are fetched.
//
// Links:
//  1. https://core.telegram.org/api/stars#balance-and-transaction-history
//
// See https://core.telegram.org/method/payments.getStarsTransactions for reference.
type PaymentsGetStarsTransactionsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, fetches only incoming transactions.
	Inbound bool
	// If set, fetches only outgoing transactions.
	Outbound bool
	// Return transactions in ascending order by date (instead of descending order by date).
	Ascending bool
	// Ton field of PaymentsGetStarsTransactionsRequest.
	Ton bool
	// If set, fetches only transactions for the specified Telegram Star subscription »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/stars#star-subscriptions
	//
	// Use SetSubscriptionID and GetSubscriptionID helpers.
	SubscriptionID string
	// Fetch the transaction history of the peer (inputPeerSelf¹ or a bot we own).
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputPeerSelf
	Peer InputPeerClass
	// Offset for pagination, obtained from the returned next_offset, initially an empty
	// string »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset string
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// PaymentsGetStarsTransactionsRequestTypeID is TL type id of PaymentsGetStarsTransactionsRequest.
const PaymentsGetStarsTransactionsRequestTypeID = 0x69da4557

// Ensuring interfaces in compile-time for PaymentsGetStarsTransactionsRequest.
var (
	_ bin.Encoder     = &PaymentsGetStarsTransactionsRequest{}
	_ bin.Decoder     = &PaymentsGetStarsTransactionsRequest{}
	_ bin.BareEncoder = &PaymentsGetStarsTransactionsRequest{}
	_ bin.BareDecoder = &PaymentsGetStarsTransactionsRequest{}
)

func (g *PaymentsGetStarsTransactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Inbound == false) {
		return false
	}
	if !(g.Outbound == false) {
		return false
	}
	if !(g.Ascending == false) {
		return false
	}
	if !(g.Ton == false) {
		return false
	}
	if !(g.SubscriptionID == "") {
		return false
	}
	if !(g.Peer == nil) {
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
func (g *PaymentsGetStarsTransactionsRequest) String() string {
	if g == nil {
		return "PaymentsGetStarsTransactionsRequest(nil)"
	}
	type Alias PaymentsGetStarsTransactionsRequest
	return fmt.Sprintf("PaymentsGetStarsTransactionsRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetStarsTransactionsRequest from given interface.
func (g *PaymentsGetStarsTransactionsRequest) FillFrom(from interface {
	GetInbound() (value bool)
	GetOutbound() (value bool)
	GetAscending() (value bool)
	GetTon() (value bool)
	GetSubscriptionID() (value string, ok bool)
	GetPeer() (value InputPeerClass)
	GetOffset() (value string)
	GetLimit() (value int)
}) {
	g.Inbound = from.GetInbound()
	g.Outbound = from.GetOutbound()
	g.Ascending = from.GetAscending()
	g.Ton = from.GetTon()
	if val, ok := from.GetSubscriptionID(); ok {
		g.SubscriptionID = val
	}

	g.Peer = from.GetPeer()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetStarsTransactionsRequest) TypeID() uint32 {
	return PaymentsGetStarsTransactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetStarsTransactionsRequest) TypeName() string {
	return "payments.getStarsTransactions"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetStarsTransactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getStarsTransactions",
		ID:   PaymentsGetStarsTransactionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inbound",
			SchemaName: "inbound",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Outbound",
			SchemaName: "outbound",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Ascending",
			SchemaName: "ascending",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "Ton",
			SchemaName: "ton",
			Null:       !g.Flags.Has(4),
		},
		{
			Name:       "SubscriptionID",
			SchemaName: "subscription_id",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
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

// SetFlags sets flags for non-zero fields.
func (g *PaymentsGetStarsTransactionsRequest) SetFlags() {
	if !(g.Inbound == false) {
		g.Flags.Set(0)
	}
	if !(g.Outbound == false) {
		g.Flags.Set(1)
	}
	if !(g.Ascending == false) {
		g.Flags.Set(2)
	}
	if !(g.Ton == false) {
		g.Flags.Set(4)
	}
	if !(g.SubscriptionID == "") {
		g.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (g *PaymentsGetStarsTransactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarsTransactions#69da4557 as nil")
	}
	b.PutID(PaymentsGetStarsTransactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetStarsTransactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getStarsTransactions#69da4557 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getStarsTransactions#69da4557: field flags: %w", err)
	}
	if g.Flags.Has(3) {
		b.PutString(g.SubscriptionID)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode payments.getStarsTransactions#69da4557: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getStarsTransactions#69da4557: field peer: %w", err)
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetStarsTransactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarsTransactions#69da4557 to nil")
	}
	if err := b.ConsumeID(PaymentsGetStarsTransactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getStarsTransactions#69da4557: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetStarsTransactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getStarsTransactions#69da4557 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.getStarsTransactions#69da4557: field flags: %w", err)
		}
	}
	g.Inbound = g.Flags.Has(0)
	g.Outbound = g.Flags.Has(1)
	g.Ascending = g.Flags.Has(2)
	g.Ton = g.Flags.Has(4)
	if g.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarsTransactions#69da4557: field subscription_id: %w", err)
		}
		g.SubscriptionID = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarsTransactions#69da4557: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarsTransactions#69da4557: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getStarsTransactions#69da4557: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetInbound sets value of Inbound conditional field.
func (g *PaymentsGetStarsTransactionsRequest) SetInbound(value bool) {
	if value {
		g.Flags.Set(0)
		g.Inbound = true
	} else {
		g.Flags.Unset(0)
		g.Inbound = false
	}
}

// GetInbound returns value of Inbound conditional field.
func (g *PaymentsGetStarsTransactionsRequest) GetInbound() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// SetOutbound sets value of Outbound conditional field.
func (g *PaymentsGetStarsTransactionsRequest) SetOutbound(value bool) {
	if value {
		g.Flags.Set(1)
		g.Outbound = true
	} else {
		g.Flags.Unset(1)
		g.Outbound = false
	}
}

// GetOutbound returns value of Outbound conditional field.
func (g *PaymentsGetStarsTransactionsRequest) GetOutbound() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// SetAscending sets value of Ascending conditional field.
func (g *PaymentsGetStarsTransactionsRequest) SetAscending(value bool) {
	if value {
		g.Flags.Set(2)
		g.Ascending = true
	} else {
		g.Flags.Unset(2)
		g.Ascending = false
	}
}

// GetAscending returns value of Ascending conditional field.
func (g *PaymentsGetStarsTransactionsRequest) GetAscending() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// SetTon sets value of Ton conditional field.
func (g *PaymentsGetStarsTransactionsRequest) SetTon(value bool) {
	if value {
		g.Flags.Set(4)
		g.Ton = true
	} else {
		g.Flags.Unset(4)
		g.Ton = false
	}
}

// GetTon returns value of Ton conditional field.
func (g *PaymentsGetStarsTransactionsRequest) GetTon() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(4)
}

// SetSubscriptionID sets value of SubscriptionID conditional field.
func (g *PaymentsGetStarsTransactionsRequest) SetSubscriptionID(value string) {
	g.Flags.Set(3)
	g.SubscriptionID = value
}

// GetSubscriptionID returns value of SubscriptionID conditional field and
// boolean which is true if field was set.
func (g *PaymentsGetStarsTransactionsRequest) GetSubscriptionID() (value string, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.SubscriptionID, true
}

// GetPeer returns value of Peer field.
func (g *PaymentsGetStarsTransactionsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetOffset returns value of Offset field.
func (g *PaymentsGetStarsTransactionsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *PaymentsGetStarsTransactionsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// PaymentsGetStarsTransactions invokes method payments.getStarsTransactions#69da4557 returning error if any.
// Fetch Telegram Stars transactions¹.
// The inbound and outbound flags are mutually exclusive: if none of the two are set,
// both incoming and outgoing transactions are fetched.
//
// Links:
//  1. https://core.telegram.org/api/stars#balance-and-transaction-history
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/payments.getStarsTransactions for reference.
// Can be used by bots.
func (c *Client) PaymentsGetStarsTransactions(ctx context.Context, request *PaymentsGetStarsTransactionsRequest) (*PaymentsStarsStatus, error) {
	var result PaymentsStarsStatus

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
