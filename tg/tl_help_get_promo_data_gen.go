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

// HelpGetPromoDataRequest represents TL type `help.getPromoData#c0977421`.
// Get MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/method/help.getPromoData for reference.
type HelpGetPromoDataRequest struct {
}

// HelpGetPromoDataRequestTypeID is TL type id of HelpGetPromoDataRequest.
const HelpGetPromoDataRequestTypeID = 0xc0977421

func (g *HelpGetPromoDataRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetPromoDataRequest) String() string {
	if g == nil {
		return "HelpGetPromoDataRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpGetPromoDataRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetPromoDataRequest) TypeID() uint32 {
	return HelpGetPromoDataRequestTypeID
}

// Encode implements bin.Encoder.
func (g *HelpGetPromoDataRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPromoData#c0977421 as nil")
	}
	b.PutID(HelpGetPromoDataRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetPromoDataRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPromoData#c0977421 to nil")
	}
	if err := b.ConsumeID(HelpGetPromoDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getPromoData#c0977421: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetPromoDataRequest.
var (
	_ bin.Encoder = &HelpGetPromoDataRequest{}
	_ bin.Decoder = &HelpGetPromoDataRequest{}
)

// HelpGetPromoData invokes method help.getPromoData#c0977421 returning error if any.
// Get MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/method/help.getPromoData for reference.
// Can be used by bots.
func (c *Client) HelpGetPromoData(ctx context.Context) (HelpPromoDataClass, error) {
	var result HelpPromoDataBox

	request := &HelpGetPromoDataRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PromoData, nil
}
