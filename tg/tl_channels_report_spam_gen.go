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

// ChannelsReportSpamRequest represents TL type `channels.reportSpam#fe087810`.
// Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup
//
// See https://core.telegram.org/method/channels.reportSpam for reference.
type ChannelsReportSpamRequest struct {
	// Supergroup
	Channel InputChannelClass
	// ID of the user that sent the spam messages
	UserID InputUserClass
	// IDs of spam messages
	ID []int
}

// ChannelsReportSpamRequestTypeID is TL type id of ChannelsReportSpamRequest.
const ChannelsReportSpamRequestTypeID = 0xfe087810

func (r *ChannelsReportSpamRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Channel == nil) {
		return false
	}
	if !(r.UserID == nil) {
		return false
	}
	if !(r.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ChannelsReportSpamRequest) String() string {
	if r == nil {
		return "ChannelsReportSpamRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsReportSpamRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(r.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(r.UserID))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range r.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ChannelsReportSpamRequest) TypeID() uint32 {
	return ChannelsReportSpamRequestTypeID
}

// Encode implements bin.Encoder.
func (r *ChannelsReportSpamRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.reportSpam#fe087810 as nil")
	}
	b.PutID(ChannelsReportSpamRequestTypeID)
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field channel: %w", err)
	}
	if r.UserID == nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.reportSpam#fe087810: field user_id: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (r *ChannelsReportSpamRequest) GetChannel() (value InputChannelClass) {
	return r.Channel
}

// GetUserID returns value of UserID field.
func (r *ChannelsReportSpamRequest) GetUserID() (value InputUserClass) {
	return r.UserID
}

// GetID returns value of ID field.
func (r *ChannelsReportSpamRequest) GetID() (value []int) {
	return r.ID
}

// Decode implements bin.Decoder.
func (r *ChannelsReportSpamRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.reportSpam#fe087810 to nil")
	}
	if err := b.ConsumeID(ChannelsReportSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.reportSpam#fe087810: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.reportSpam#fe087810: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsReportSpamRequest.
var (
	_ bin.Encoder = &ChannelsReportSpamRequest{}
	_ bin.Decoder = &ChannelsReportSpamRequest{}
)

// ChannelsReportSpam invokes method channels.reportSpam#fe087810 returning error if any.
// Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/channels.reportSpam for reference.
func (c *Client) ChannelsReportSpam(ctx context.Context, request *ChannelsReportSpamRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
