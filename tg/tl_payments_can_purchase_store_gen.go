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

// PaymentsCanPurchaseStoreRequest represents TL type `payments.canPurchaseStore#4fdc5ea7`.
//
// See https://core.telegram.org/method/payments.canPurchaseStore for reference.
type PaymentsCanPurchaseStoreRequest struct {
	// Purpose field of PaymentsCanPurchaseStoreRequest.
	Purpose InputStorePaymentPurposeClass
}

// PaymentsCanPurchaseStoreRequestTypeID is TL type id of PaymentsCanPurchaseStoreRequest.
const PaymentsCanPurchaseStoreRequestTypeID = 0x4fdc5ea7

// Ensuring interfaces in compile-time for PaymentsCanPurchaseStoreRequest.
var (
	_ bin.Encoder     = &PaymentsCanPurchaseStoreRequest{}
	_ bin.Decoder     = &PaymentsCanPurchaseStoreRequest{}
	_ bin.BareEncoder = &PaymentsCanPurchaseStoreRequest{}
	_ bin.BareDecoder = &PaymentsCanPurchaseStoreRequest{}
)

func (c *PaymentsCanPurchaseStoreRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PaymentsCanPurchaseStoreRequest) String() string {
	if c == nil {
		return "PaymentsCanPurchaseStoreRequest(nil)"
	}
	type Alias PaymentsCanPurchaseStoreRequest
	return fmt.Sprintf("PaymentsCanPurchaseStoreRequest%+v", Alias(*c))
}

// FillFrom fills PaymentsCanPurchaseStoreRequest from given interface.
func (c *PaymentsCanPurchaseStoreRequest) FillFrom(from interface {
	GetPurpose() (value InputStorePaymentPurposeClass)
}) {
	c.Purpose = from.GetPurpose()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsCanPurchaseStoreRequest) TypeID() uint32 {
	return PaymentsCanPurchaseStoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsCanPurchaseStoreRequest) TypeName() string {
	return "payments.canPurchaseStore"
}

// TypeInfo returns info about TL type.
func (c *PaymentsCanPurchaseStoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.canPurchaseStore",
		ID:   PaymentsCanPurchaseStoreRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *PaymentsCanPurchaseStoreRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.canPurchaseStore#4fdc5ea7 as nil")
	}
	b.PutID(PaymentsCanPurchaseStoreRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *PaymentsCanPurchaseStoreRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.canPurchaseStore#4fdc5ea7 as nil")
	}
	if c.Purpose == nil {
		return fmt.Errorf("unable to encode payments.canPurchaseStore#4fdc5ea7: field purpose is nil")
	}
	if err := c.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.canPurchaseStore#4fdc5ea7: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *PaymentsCanPurchaseStoreRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.canPurchaseStore#4fdc5ea7 to nil")
	}
	if err := b.ConsumeID(PaymentsCanPurchaseStoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.canPurchaseStore#4fdc5ea7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *PaymentsCanPurchaseStoreRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.canPurchaseStore#4fdc5ea7 to nil")
	}
	{
		value, err := DecodeInputStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.canPurchaseStore#4fdc5ea7: field purpose: %w", err)
		}
		c.Purpose = value
	}
	return nil
}

// GetPurpose returns value of Purpose field.
func (c *PaymentsCanPurchaseStoreRequest) GetPurpose() (value InputStorePaymentPurposeClass) {
	if c == nil {
		return
	}
	return c.Purpose
}

// PaymentsCanPurchaseStore invokes method payments.canPurchaseStore#4fdc5ea7 returning error if any.
//
// See https://core.telegram.org/method/payments.canPurchaseStore for reference.
func (c *Client) PaymentsCanPurchaseStore(ctx context.Context, purpose InputStorePaymentPurposeClass) (bool, error) {
	var result BoolBox

	request := &PaymentsCanPurchaseStoreRequest{
		Purpose: purpose,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
