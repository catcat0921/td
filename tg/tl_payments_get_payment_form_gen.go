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

// PaymentsGetPaymentFormRequest represents TL type `payments.getPaymentForm#99f09745`.
// Get a payment form
//
// See https://core.telegram.org/method/payments.getPaymentForm for reference.
type PaymentsGetPaymentFormRequest struct {
	// Message ID of payment form
	MsgID int
}

// PaymentsGetPaymentFormRequestTypeID is TL type id of PaymentsGetPaymentFormRequest.
const PaymentsGetPaymentFormRequestTypeID = 0x99f09745

func (g *PaymentsGetPaymentFormRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetPaymentFormRequest) String() string {
	if g == nil {
		return "PaymentsGetPaymentFormRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PaymentsGetPaymentFormRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(g.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *PaymentsGetPaymentFormRequest) TypeID() uint32 {
	return PaymentsGetPaymentFormRequestTypeID
}

// Encode implements bin.Encoder.
func (g *PaymentsGetPaymentFormRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getPaymentForm#99f09745 as nil")
	}
	b.PutID(PaymentsGetPaymentFormRequestTypeID)
	b.PutInt(g.MsgID)
	return nil
}

// GetMsgID returns value of MsgID field.
func (g *PaymentsGetPaymentFormRequest) GetMsgID() (value int) {
	return g.MsgID
}

// Decode implements bin.Decoder.
func (g *PaymentsGetPaymentFormRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getPaymentForm#99f09745 to nil")
	}
	if err := b.ConsumeID(PaymentsGetPaymentFormRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getPaymentForm#99f09745: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getPaymentForm#99f09745: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsGetPaymentFormRequest.
var (
	_ bin.Encoder = &PaymentsGetPaymentFormRequest{}
	_ bin.Decoder = &PaymentsGetPaymentFormRequest{}
)

// PaymentsGetPaymentForm invokes method payments.getPaymentForm#99f09745 returning error if any.
// Get a payment form
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/payments.getPaymentForm for reference.
func (c *Client) PaymentsGetPaymentForm(ctx context.Context, msgid int) (*PaymentsPaymentForm, error) {
	var result PaymentsPaymentForm

	request := &PaymentsGetPaymentFormRequest{
		MsgID: msgid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
