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

// ChannelsGetFullChannelRequest represents TL type `channels.getFullChannel#8736a09`.
// Get full info about a channel
//
// See https://core.telegram.org/method/channels.getFullChannel for reference.
type ChannelsGetFullChannelRequest struct {
	// The channel to get info about
	Channel InputChannelClass
}

// ChannelsGetFullChannelRequestTypeID is TL type id of ChannelsGetFullChannelRequest.
const ChannelsGetFullChannelRequestTypeID = 0x8736a09

func (g *ChannelsGetFullChannelRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetFullChannelRequest) String() string {
	if g == nil {
		return "ChannelsGetFullChannelRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetFullChannelRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(g.Channel))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *ChannelsGetFullChannelRequest) TypeID() uint32 {
	return ChannelsGetFullChannelRequestTypeID
}

// Encode implements bin.Encoder.
func (g *ChannelsGetFullChannelRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getFullChannel#8736a09 as nil")
	}
	b.PutID(ChannelsGetFullChannelRequestTypeID)
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getFullChannel#8736a09: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getFullChannel#8736a09: field channel: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *ChannelsGetFullChannelRequest) GetChannel() (value InputChannelClass) {
	return g.Channel
}

// Decode implements bin.Decoder.
func (g *ChannelsGetFullChannelRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getFullChannel#8736a09 to nil")
	}
	if err := b.ConsumeID(ChannelsGetFullChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getFullChannel#8736a09: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getFullChannel#8736a09: field channel: %w", err)
		}
		g.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetFullChannelRequest.
var (
	_ bin.Encoder = &ChannelsGetFullChannelRequest{}
	_ bin.Decoder = &ChannelsGetFullChannelRequest{}
)

// ChannelsGetFullChannel invokes method channels.getFullChannel#8736a09 returning error if any.
// Get full info about a channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  403 CHANNEL_PUBLIC_GROUP_NA: channel/supergroup not available
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.getFullChannel for reference.
// Can be used by bots.
func (c *Client) ChannelsGetFullChannel(ctx context.Context, channel InputChannelClass) (*MessagesChatFull, error) {
	var result MessagesChatFull

	request := &ChannelsGetFullChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
