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

// MessagesReadMentionsRequest represents TL type `messages.readMentions#36e5bf4d`.
// Mark mentions as read
//
// See https://core.telegram.org/method/messages.readMentions for reference.
type MessagesReadMentionsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Dialog
	Peer InputPeerClass
	// TopMsgID field of MessagesReadMentionsRequest.
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
}

// MessagesReadMentionsRequestTypeID is TL type id of MessagesReadMentionsRequest.
const MessagesReadMentionsRequestTypeID = 0x36e5bf4d

// Ensuring interfaces in compile-time for MessagesReadMentionsRequest.
var (
	_ bin.Encoder     = &MessagesReadMentionsRequest{}
	_ bin.Decoder     = &MessagesReadMentionsRequest{}
	_ bin.BareEncoder = &MessagesReadMentionsRequest{}
	_ bin.BareDecoder = &MessagesReadMentionsRequest{}
)

func (r *MessagesReadMentionsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.TopMsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadMentionsRequest) String() string {
	if r == nil {
		return "MessagesReadMentionsRequest(nil)"
	}
	type Alias MessagesReadMentionsRequest
	return fmt.Sprintf("MessagesReadMentionsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadMentionsRequest from given interface.
func (r *MessagesReadMentionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetTopMsgID() (value int, ok bool)
}) {
	r.Peer = from.GetPeer()
	if val, ok := from.GetTopMsgID(); ok {
		r.TopMsgID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReadMentionsRequest) TypeID() uint32 {
	return MessagesReadMentionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReadMentionsRequest) TypeName() string {
	return "messages.readMentions"
}

// TypeInfo returns info about TL type.
func (r *MessagesReadMentionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.readMentions",
		ID:   MessagesReadMentionsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !r.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesReadMentionsRequest) SetFlags() {
	if !(r.TopMsgID == 0) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesReadMentionsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMentions#36e5bf4d as nil")
	}
	b.PutID(MessagesReadMentionsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReadMentionsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMentions#36e5bf4d as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readMentions#36e5bf4d: field flags: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.readMentions#36e5bf4d: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readMentions#36e5bf4d: field peer: %w", err)
	}
	if r.Flags.Has(0) {
		b.PutInt(r.TopMsgID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadMentionsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMentions#36e5bf4d to nil")
	}
	if err := b.ConsumeID(MessagesReadMentionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readMentions#36e5bf4d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReadMentionsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMentions#36e5bf4d to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.readMentions#36e5bf4d: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.readMentions#36e5bf4d: field peer: %w", err)
		}
		r.Peer = value
	}
	if r.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readMentions#36e5bf4d: field top_msg_id: %w", err)
		}
		r.TopMsgID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReadMentionsRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (r *MessagesReadMentionsRequest) SetTopMsgID(value int) {
	r.Flags.Set(0)
	r.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (r *MessagesReadMentionsRequest) GetTopMsgID() (value int, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.TopMsgID, true
}

// MessagesReadMentions invokes method messages.readMentions#36e5bf4d returning error if any.
// Mark mentions as read
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.readMentions for reference.
func (c *Client) MessagesReadMentions(ctx context.Context, request *MessagesReadMentionsRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
