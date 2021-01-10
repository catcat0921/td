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

// TopPeerCategoryPeers represents TL type `topPeerCategoryPeers#fb834291`.
// Top peer category
//
// See https://core.telegram.org/constructor/topPeerCategoryPeers for reference.
type TopPeerCategoryPeers struct {
	// Top peer category of peers
	Category TopPeerCategoryClass
	// Count of peers
	Count int
	// Peers
	Peers []TopPeer
}

// TopPeerCategoryPeersTypeID is TL type id of TopPeerCategoryPeers.
const TopPeerCategoryPeersTypeID = 0xfb834291

func (t *TopPeerCategoryPeers) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Category == nil) {
		return false
	}
	if !(t.Count == 0) {
		return false
	}
	if !(t.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryPeers) String() string {
	if t == nil {
		return "TopPeerCategoryPeers(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryPeers")
	sb.WriteString("{\n")
	sb.WriteString("\tCategory: ")
	sb.WriteString(fmt.Sprint(t.Category))
	sb.WriteString(",\n")
	sb.WriteString("\tCount: ")
	sb.WriteString(fmt.Sprint(t.Count))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range t.Peers {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TopPeerCategoryPeers) TypeID() uint32 {
	return TopPeerCategoryPeersTypeID
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryPeers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryPeers#fb834291 as nil")
	}
	b.PutID(TopPeerCategoryPeersTypeID)
	if t.Category == nil {
		return fmt.Errorf("unable to encode topPeerCategoryPeers#fb834291: field category is nil")
	}
	if err := t.Category.Encode(b); err != nil {
		return fmt.Errorf("unable to encode topPeerCategoryPeers#fb834291: field category: %w", err)
	}
	b.PutInt(t.Count)
	b.PutVectorHeader(len(t.Peers))
	for idx, v := range t.Peers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode topPeerCategoryPeers#fb834291: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCategory returns value of Category field.
func (t *TopPeerCategoryPeers) GetCategory() (value TopPeerCategoryClass) {
	return t.Category
}

// GetCount returns value of Count field.
func (t *TopPeerCategoryPeers) GetCount() (value int) {
	return t.Count
}

// GetPeers returns value of Peers field.
func (t *TopPeerCategoryPeers) GetPeers() (value []TopPeer) {
	return t.Peers
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryPeers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryPeers#fb834291 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryPeersTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: %w", err)
	}
	{
		value, err := DecodeTopPeerCategory(b)
		if err != nil {
			return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field category: %w", err)
		}
		t.Category = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field count: %w", err)
		}
		t.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field peers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TopPeer
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field peers: %w", err)
			}
			t.Peers = append(t.Peers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TopPeerCategoryPeers.
var (
	_ bin.Encoder = &TopPeerCategoryPeers{}
	_ bin.Decoder = &TopPeerCategoryPeers{}
)
