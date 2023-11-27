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

// StoriesDeleteStoriesRequest represents TL type `stories.deleteStories#ae59db5f`.
// Deletes some posted stories¹.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/method/stories.deleteStories for reference.
type StoriesDeleteStoriesRequest struct {
	// Channel/user from where to delete stories.
	Peer InputPeerClass
	// IDs of stories to delete.
	ID []int
}

// StoriesDeleteStoriesRequestTypeID is TL type id of StoriesDeleteStoriesRequest.
const StoriesDeleteStoriesRequestTypeID = 0xae59db5f

// Ensuring interfaces in compile-time for StoriesDeleteStoriesRequest.
var (
	_ bin.Encoder     = &StoriesDeleteStoriesRequest{}
	_ bin.Decoder     = &StoriesDeleteStoriesRequest{}
	_ bin.BareEncoder = &StoriesDeleteStoriesRequest{}
	_ bin.BareDecoder = &StoriesDeleteStoriesRequest{}
)

func (d *StoriesDeleteStoriesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *StoriesDeleteStoriesRequest) String() string {
	if d == nil {
		return "StoriesDeleteStoriesRequest(nil)"
	}
	type Alias StoriesDeleteStoriesRequest
	return fmt.Sprintf("StoriesDeleteStoriesRequest%+v", Alias(*d))
}

// FillFrom fills StoriesDeleteStoriesRequest from given interface.
func (d *StoriesDeleteStoriesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	d.Peer = from.GetPeer()
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesDeleteStoriesRequest) TypeID() uint32 {
	return StoriesDeleteStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesDeleteStoriesRequest) TypeName() string {
	return "stories.deleteStories"
}

// TypeInfo returns info about TL type.
func (d *StoriesDeleteStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.deleteStories",
		ID:   StoriesDeleteStoriesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *StoriesDeleteStoriesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode stories.deleteStories#ae59db5f as nil")
	}
	b.PutID(StoriesDeleteStoriesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *StoriesDeleteStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode stories.deleteStories#ae59db5f as nil")
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode stories.deleteStories#ae59db5f: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.deleteStories#ae59db5f: field peer: %w", err)
	}
	b.PutVectorHeader(len(d.ID))
	for _, v := range d.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *StoriesDeleteStoriesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode stories.deleteStories#ae59db5f to nil")
	}
	if err := b.ConsumeID(StoriesDeleteStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.deleteStories#ae59db5f: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *StoriesDeleteStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode stories.deleteStories#ae59db5f to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.deleteStories#ae59db5f: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.deleteStories#ae59db5f: field id: %w", err)
		}

		if headerLen > 0 {
			d.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode stories.deleteStories#ae59db5f: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *StoriesDeleteStoriesRequest) GetPeer() (value InputPeerClass) {
	if d == nil {
		return
	}
	return d.Peer
}

// GetID returns value of ID field.
func (d *StoriesDeleteStoriesRequest) GetID() (value []int) {
	if d == nil {
		return
	}
	return d.ID
}

// StoriesDeleteStories invokes method stories.deleteStories#ae59db5f returning error if any.
// Deletes some posted stories¹.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/stories.deleteStories for reference.
func (c *Client) StoriesDeleteStories(ctx context.Context, request *StoriesDeleteStoriesRequest) ([]int, error) {
	var result IntVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int(result.Elems), nil
}
