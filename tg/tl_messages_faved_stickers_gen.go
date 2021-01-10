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

// MessagesFavedStickersNotModified represents TL type `messages.favedStickersNotModified#9e8fa6d3`.
// No new favorited stickers were found
//
// See https://core.telegram.org/constructor/messages.favedStickersNotModified for reference.
type MessagesFavedStickersNotModified struct {
}

// MessagesFavedStickersNotModifiedTypeID is TL type id of MessagesFavedStickersNotModified.
const MessagesFavedStickersNotModifiedTypeID = 0x9e8fa6d3

func (f *MessagesFavedStickersNotModified) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFavedStickersNotModified) String() string {
	if f == nil {
		return "MessagesFavedStickersNotModified(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesFavedStickersNotModified")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *MessagesFavedStickersNotModified) TypeID() uint32 {
	return MessagesFavedStickersNotModifiedTypeID
}

// Encode implements bin.Encoder.
func (f *MessagesFavedStickersNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.favedStickersNotModified#9e8fa6d3 as nil")
	}
	b.PutID(MessagesFavedStickersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFavedStickersNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.favedStickersNotModified#9e8fa6d3 to nil")
	}
	if err := b.ConsumeID(MessagesFavedStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.favedStickersNotModified#9e8fa6d3: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesFavedStickersClass.
func (f MessagesFavedStickersNotModified) construct() MessagesFavedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFavedStickersNotModified.
var (
	_ bin.Encoder = &MessagesFavedStickersNotModified{}
	_ bin.Decoder = &MessagesFavedStickersNotModified{}

	_ MessagesFavedStickersClass = &MessagesFavedStickersNotModified{}
)

// MessagesFavedStickers represents TL type `messages.favedStickers#f37f2f16`.
// Favorited stickers
//
// See https://core.telegram.org/constructor/messages.favedStickers for reference.
type MessagesFavedStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Emojis associated to stickers
	Packs []StickerPack
	// Favorited stickers
	Stickers []DocumentClass
}

// MessagesFavedStickersTypeID is TL type id of MessagesFavedStickers.
const MessagesFavedStickersTypeID = 0xf37f2f16

func (f *MessagesFavedStickers) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Packs == nil) {
		return false
	}
	if !(f.Stickers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFavedStickers) String() string {
	if f == nil {
		return "MessagesFavedStickers(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesFavedStickers")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(f.Hash))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range f.Packs {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range f.Stickers {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *MessagesFavedStickers) TypeID() uint32 {
	return MessagesFavedStickersTypeID
}

// Encode implements bin.Encoder.
func (f *MessagesFavedStickers) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.favedStickers#f37f2f16 as nil")
	}
	b.PutID(MessagesFavedStickersTypeID)
	b.PutInt(f.Hash)
	b.PutVectorHeader(len(f.Packs))
	for idx, v := range f.Packs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.favedStickers#f37f2f16: field packs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Stickers))
	for idx, v := range f.Stickers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.favedStickers#f37f2f16: field stickers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.favedStickers#f37f2f16: field stickers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (f *MessagesFavedStickers) GetHash() (value int) {
	return f.Hash
}

// GetPacks returns value of Packs field.
func (f *MessagesFavedStickers) GetPacks() (value []StickerPack) {
	return f.Packs
}

// GetStickers returns value of Stickers field.
func (f *MessagesFavedStickers) GetStickers() (value []DocumentClass) {
	return f.Stickers
}

// Decode implements bin.Decoder.
func (f *MessagesFavedStickers) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.favedStickers#f37f2f16 to nil")
	}
	if err := b.ConsumeID(MessagesFavedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.favedStickers#f37f2f16: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.favedStickers#f37f2f16: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.favedStickers#f37f2f16: field packs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerPack
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.favedStickers#f37f2f16: field packs: %w", err)
			}
			f.Packs = append(f.Packs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.favedStickers#f37f2f16: field stickers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.favedStickers#f37f2f16: field stickers: %w", err)
			}
			f.Stickers = append(f.Stickers, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesFavedStickersClass.
func (f MessagesFavedStickers) construct() MessagesFavedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFavedStickers.
var (
	_ bin.Encoder = &MessagesFavedStickers{}
	_ bin.Decoder = &MessagesFavedStickers{}

	_ MessagesFavedStickersClass = &MessagesFavedStickers{}
)

// MessagesFavedStickersClass represents messages.FavedStickers generic type.
//
// See https://core.telegram.org/type/messages.FavedStickers for reference.
//
// Example:
//  g, err := DecodeMessagesFavedStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesFavedStickersNotModified: // messages.favedStickersNotModified#9e8fa6d3
//  case *MessagesFavedStickers: // messages.favedStickers#f37f2f16
//  default: panic(v)
//  }
type MessagesFavedStickersClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesFavedStickersClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessagesFavedStickers implements binary de-serialization for MessagesFavedStickersClass.
func DecodeMessagesFavedStickers(buf *bin.Buffer) (MessagesFavedStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFavedStickersNotModifiedTypeID:
		// Decoding messages.favedStickersNotModified#9e8fa6d3.
		v := MessagesFavedStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFavedStickersClass: %w", err)
		}
		return &v, nil
	case MessagesFavedStickersTypeID:
		// Decoding messages.favedStickers#f37f2f16.
		v := MessagesFavedStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFavedStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFavedStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFavedStickers boxes the MessagesFavedStickersClass providing a helper.
type MessagesFavedStickersBox struct {
	FavedStickers MessagesFavedStickersClass
}

// Decode implements bin.Decoder for MessagesFavedStickersBox.
func (b *MessagesFavedStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFavedStickersBox to nil")
	}
	v, err := DecodeMessagesFavedStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FavedStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesFavedStickersBox.
func (b *MessagesFavedStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FavedStickers == nil {
		return fmt.Errorf("unable to encode MessagesFavedStickersClass as nil")
	}
	return b.FavedStickers.Encode(buf)
}
