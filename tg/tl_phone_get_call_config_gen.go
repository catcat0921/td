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

// PhoneGetCallConfigRequest represents TL type `phone.getCallConfig#55451fa9`.
// Get phone call configuration to be passed to libtgvoip's shared config
//
// See https://core.telegram.org/method/phone.getCallConfig for reference.
type PhoneGetCallConfigRequest struct {
}

// PhoneGetCallConfigRequestTypeID is TL type id of PhoneGetCallConfigRequest.
const PhoneGetCallConfigRequestTypeID = 0x55451fa9

func (g *PhoneGetCallConfigRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGetCallConfigRequest) String() string {
	if g == nil {
		return "PhoneGetCallConfigRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneGetCallConfigRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *PhoneGetCallConfigRequest) TypeID() uint32 {
	return PhoneGetCallConfigRequestTypeID
}

// Encode implements bin.Encoder.
func (g *PhoneGetCallConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getCallConfig#55451fa9 as nil")
	}
	b.PutID(PhoneGetCallConfigRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *PhoneGetCallConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getCallConfig#55451fa9 to nil")
	}
	if err := b.ConsumeID(PhoneGetCallConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getCallConfig#55451fa9: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGetCallConfigRequest.
var (
	_ bin.Encoder = &PhoneGetCallConfigRequest{}
	_ bin.Decoder = &PhoneGetCallConfigRequest{}
)

// PhoneGetCallConfig invokes method phone.getCallConfig#55451fa9 returning error if any.
// Get phone call configuration to be passed to libtgvoip's shared config
//
// See https://core.telegram.org/method/phone.getCallConfig for reference.
func (c *Client) PhoneGetCallConfig(ctx context.Context) (*DataJSON, error) {
	var result DataJSON

	request := &PhoneGetCallConfigRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
