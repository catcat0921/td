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

// StoriesToggleAllStoriesHiddenRequest represents TL type `stories.toggleAllStoriesHidden#7c2557c4`.
// Hide the active stories of a specific peer, preventing them from being displayed on
// the action bar on the homescreen.
//
// See https://core.telegram.org/method/stories.toggleAllStoriesHidden for reference.
type StoriesToggleAllStoriesHiddenRequest struct {
	// Whether to hide or unhide all active stories of the peer
	Hidden bool
}

// StoriesToggleAllStoriesHiddenRequestTypeID is TL type id of StoriesToggleAllStoriesHiddenRequest.
const StoriesToggleAllStoriesHiddenRequestTypeID = 0x7c2557c4

// Ensuring interfaces in compile-time for StoriesToggleAllStoriesHiddenRequest.
var (
	_ bin.Encoder     = &StoriesToggleAllStoriesHiddenRequest{}
	_ bin.Decoder     = &StoriesToggleAllStoriesHiddenRequest{}
	_ bin.BareEncoder = &StoriesToggleAllStoriesHiddenRequest{}
	_ bin.BareDecoder = &StoriesToggleAllStoriesHiddenRequest{}
)

func (t *StoriesToggleAllStoriesHiddenRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Hidden == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *StoriesToggleAllStoriesHiddenRequest) String() string {
	if t == nil {
		return "StoriesToggleAllStoriesHiddenRequest(nil)"
	}
	type Alias StoriesToggleAllStoriesHiddenRequest
	return fmt.Sprintf("StoriesToggleAllStoriesHiddenRequest%+v", Alias(*t))
}

// FillFrom fills StoriesToggleAllStoriesHiddenRequest from given interface.
func (t *StoriesToggleAllStoriesHiddenRequest) FillFrom(from interface {
	GetHidden() (value bool)
}) {
	t.Hidden = from.GetHidden()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesToggleAllStoriesHiddenRequest) TypeID() uint32 {
	return StoriesToggleAllStoriesHiddenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesToggleAllStoriesHiddenRequest) TypeName() string {
	return "stories.toggleAllStoriesHidden"
}

// TypeInfo returns info about TL type.
func (t *StoriesToggleAllStoriesHiddenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.toggleAllStoriesHidden",
		ID:   StoriesToggleAllStoriesHiddenRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hidden",
			SchemaName: "hidden",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *StoriesToggleAllStoriesHiddenRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode stories.toggleAllStoriesHidden#7c2557c4 as nil")
	}
	b.PutID(StoriesToggleAllStoriesHiddenRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *StoriesToggleAllStoriesHiddenRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode stories.toggleAllStoriesHidden#7c2557c4 as nil")
	}
	b.PutBool(t.Hidden)
	return nil
}

// Decode implements bin.Decoder.
func (t *StoriesToggleAllStoriesHiddenRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode stories.toggleAllStoriesHidden#7c2557c4 to nil")
	}
	if err := b.ConsumeID(StoriesToggleAllStoriesHiddenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.toggleAllStoriesHidden#7c2557c4: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *StoriesToggleAllStoriesHiddenRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode stories.toggleAllStoriesHidden#7c2557c4 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stories.toggleAllStoriesHidden#7c2557c4: field hidden: %w", err)
		}
		t.Hidden = value
	}
	return nil
}

// GetHidden returns value of Hidden field.
func (t *StoriesToggleAllStoriesHiddenRequest) GetHidden() (value bool) {
	if t == nil {
		return
	}
	return t.Hidden
}

// StoriesToggleAllStoriesHidden invokes method stories.toggleAllStoriesHidden#7c2557c4 returning error if any.
// Hide the active stories of a specific peer, preventing them from being displayed on
// the action bar on the homescreen.
//
// See https://core.telegram.org/method/stories.toggleAllStoriesHidden for reference.
func (c *Client) StoriesToggleAllStoriesHidden(ctx context.Context, hidden bool) (bool, error) {
	var result BoolBox

	request := &StoriesToggleAllStoriesHiddenRequest{
		Hidden: hidden,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
