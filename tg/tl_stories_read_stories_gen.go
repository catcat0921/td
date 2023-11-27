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

// StoriesReadStoriesRequest represents TL type `stories.readStories#a556dac8`.
// Mark all stories up to a certain ID as read, for a given peer; will emit an
// updateReadStories¹ update to all logged-in sessions.
//
// Links:
//  1. https://core.telegram.org/constructor/updateReadStories
//
// See https://core.telegram.org/method/stories.readStories for reference.
type StoriesReadStoriesRequest struct {
	// The peer whose stories should be marked as read.
	Peer InputPeerClass
	// Mark all stories up to and including this ID as read
	MaxID int
}

// StoriesReadStoriesRequestTypeID is TL type id of StoriesReadStoriesRequest.
const StoriesReadStoriesRequestTypeID = 0xa556dac8

// Ensuring interfaces in compile-time for StoriesReadStoriesRequest.
var (
	_ bin.Encoder     = &StoriesReadStoriesRequest{}
	_ bin.Decoder     = &StoriesReadStoriesRequest{}
	_ bin.BareEncoder = &StoriesReadStoriesRequest{}
	_ bin.BareDecoder = &StoriesReadStoriesRequest{}
)

func (r *StoriesReadStoriesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *StoriesReadStoriesRequest) String() string {
	if r == nil {
		return "StoriesReadStoriesRequest(nil)"
	}
	type Alias StoriesReadStoriesRequest
	return fmt.Sprintf("StoriesReadStoriesRequest%+v", Alias(*r))
}

// FillFrom fills StoriesReadStoriesRequest from given interface.
func (r *StoriesReadStoriesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMaxID() (value int)
}) {
	r.Peer = from.GetPeer()
	r.MaxID = from.GetMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesReadStoriesRequest) TypeID() uint32 {
	return StoriesReadStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesReadStoriesRequest) TypeName() string {
	return "stories.readStories"
}

// TypeInfo returns info about TL type.
func (r *StoriesReadStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.readStories",
		ID:   StoriesReadStoriesRequestTypeID,
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
			Name:       "MaxID",
			SchemaName: "max_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *StoriesReadStoriesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stories.readStories#a556dac8 as nil")
	}
	b.PutID(StoriesReadStoriesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *StoriesReadStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode stories.readStories#a556dac8 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode stories.readStories#a556dac8: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.readStories#a556dac8: field peer: %w", err)
	}
	b.PutInt(r.MaxID)
	return nil
}

// Decode implements bin.Decoder.
func (r *StoriesReadStoriesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stories.readStories#a556dac8 to nil")
	}
	if err := b.ConsumeID(StoriesReadStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.readStories#a556dac8: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *StoriesReadStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode stories.readStories#a556dac8 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.readStories#a556dac8: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.readStories#a556dac8: field max_id: %w", err)
		}
		r.MaxID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *StoriesReadStoriesRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetMaxID returns value of MaxID field.
func (r *StoriesReadStoriesRequest) GetMaxID() (value int) {
	if r == nil {
		return
	}
	return r.MaxID
}

// StoriesReadStories invokes method stories.readStories#a556dac8 returning error if any.
// Mark all stories up to a certain ID as read, for a given peer; will emit an
// updateReadStories¹ update to all logged-in sessions.
//
// Links:
//  1. https://core.telegram.org/constructor/updateReadStories
//
// Possible errors:
//
//	400 MAX_ID_INVALID: The provided max ID is invalid.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/stories.readStories for reference.
func (c *Client) StoriesReadStories(ctx context.Context, request *StoriesReadStoriesRequest) ([]int, error) {
	var result IntVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int(result.Elems), nil
}
