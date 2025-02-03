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

// HelpPeerColorSet represents TL type `help.peerColorSet#26219a58`.
// Represents a color palette that can be used in message accents »¹.
//
// Links:
//  1. https://core.telegram.org/api/colors
//
// See https://core.telegram.org/constructor/help.peerColorSet for reference.
type HelpPeerColorSet struct {
	// A list of 1-3 colors in RGB format, describing the accent color.
	Colors []int
}

// HelpPeerColorSetTypeID is TL type id of HelpPeerColorSet.
const HelpPeerColorSetTypeID = 0x26219a58

// construct implements constructor of HelpPeerColorSetClass.
func (p HelpPeerColorSet) construct() HelpPeerColorSetClass { return &p }

// Ensuring interfaces in compile-time for HelpPeerColorSet.
var (
	_ bin.Encoder     = &HelpPeerColorSet{}
	_ bin.Decoder     = &HelpPeerColorSet{}
	_ bin.BareEncoder = &HelpPeerColorSet{}
	_ bin.BareDecoder = &HelpPeerColorSet{}

	_ HelpPeerColorSetClass = &HelpPeerColorSet{}
)

func (p *HelpPeerColorSet) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Colors == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPeerColorSet) String() string {
	if p == nil {
		return "HelpPeerColorSet(nil)"
	}
	type Alias HelpPeerColorSet
	return fmt.Sprintf("HelpPeerColorSet%+v", Alias(*p))
}

// FillFrom fills HelpPeerColorSet from given interface.
func (p *HelpPeerColorSet) FillFrom(from interface {
	GetColors() (value []int)
}) {
	p.Colors = from.GetColors()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpPeerColorSet) TypeID() uint32 {
	return HelpPeerColorSetTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpPeerColorSet) TypeName() string {
	return "help.peerColorSet"
}

// TypeInfo returns info about TL type.
func (p *HelpPeerColorSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.peerColorSet",
		ID:   HelpPeerColorSetTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Colors",
			SchemaName: "colors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *HelpPeerColorSet) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.peerColorSet#26219a58 as nil")
	}
	b.PutID(HelpPeerColorSetTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *HelpPeerColorSet) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.peerColorSet#26219a58 as nil")
	}
	b.PutVectorHeader(len(p.Colors))
	for _, v := range p.Colors {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPeerColorSet) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.peerColorSet#26219a58 to nil")
	}
	if err := b.ConsumeID(HelpPeerColorSetTypeID); err != nil {
		return fmt.Errorf("unable to decode help.peerColorSet#26219a58: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *HelpPeerColorSet) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.peerColorSet#26219a58 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.peerColorSet#26219a58: field colors: %w", err)
		}

		if headerLen > 0 {
			p.Colors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode help.peerColorSet#26219a58: field colors: %w", err)
			}
			p.Colors = append(p.Colors, value)
		}
	}
	return nil
}

// GetColors returns value of Colors field.
func (p *HelpPeerColorSet) GetColors() (value []int) {
	if p == nil {
		return
	}
	return p.Colors
}

// HelpPeerColorProfileSet represents TL type `help.peerColorProfileSet#767d61eb`.
// Represents a color palette that can be used in profile pages »¹.
//
// Links:
//  1. https://core.telegram.org/api/colors
//
// See https://core.telegram.org/constructor/help.peerColorProfileSet for reference.
type HelpPeerColorProfileSet struct {
	// A list of 1-2 colors in RGB format, shown in the color palette settings to describe
	// the current palette.
	PaletteColors []int
	// A list of 1-2 colors in RGB format describing the colors used to generate the actual
	// background used in the profile page.
	BgColors []int
	// A list of 2 colors in RGB format describing the colors of the gradient used for the
	// unread active story indicator around the profile photo.
	StoryColors []int
}

// HelpPeerColorProfileSetTypeID is TL type id of HelpPeerColorProfileSet.
const HelpPeerColorProfileSetTypeID = 0x767d61eb

// construct implements constructor of HelpPeerColorSetClass.
func (p HelpPeerColorProfileSet) construct() HelpPeerColorSetClass { return &p }

// Ensuring interfaces in compile-time for HelpPeerColorProfileSet.
var (
	_ bin.Encoder     = &HelpPeerColorProfileSet{}
	_ bin.Decoder     = &HelpPeerColorProfileSet{}
	_ bin.BareEncoder = &HelpPeerColorProfileSet{}
	_ bin.BareDecoder = &HelpPeerColorProfileSet{}

	_ HelpPeerColorSetClass = &HelpPeerColorProfileSet{}
)

func (p *HelpPeerColorProfileSet) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PaletteColors == nil) {
		return false
	}
	if !(p.BgColors == nil) {
		return false
	}
	if !(p.StoryColors == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPeerColorProfileSet) String() string {
	if p == nil {
		return "HelpPeerColorProfileSet(nil)"
	}
	type Alias HelpPeerColorProfileSet
	return fmt.Sprintf("HelpPeerColorProfileSet%+v", Alias(*p))
}

// FillFrom fills HelpPeerColorProfileSet from given interface.
func (p *HelpPeerColorProfileSet) FillFrom(from interface {
	GetPaletteColors() (value []int)
	GetBgColors() (value []int)
	GetStoryColors() (value []int)
}) {
	p.PaletteColors = from.GetPaletteColors()
	p.BgColors = from.GetBgColors()
	p.StoryColors = from.GetStoryColors()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpPeerColorProfileSet) TypeID() uint32 {
	return HelpPeerColorProfileSetTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpPeerColorProfileSet) TypeName() string {
	return "help.peerColorProfileSet"
}

// TypeInfo returns info about TL type.
func (p *HelpPeerColorProfileSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.peerColorProfileSet",
		ID:   HelpPeerColorProfileSetTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PaletteColors",
			SchemaName: "palette_colors",
		},
		{
			Name:       "BgColors",
			SchemaName: "bg_colors",
		},
		{
			Name:       "StoryColors",
			SchemaName: "story_colors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *HelpPeerColorProfileSet) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.peerColorProfileSet#767d61eb as nil")
	}
	b.PutID(HelpPeerColorProfileSetTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *HelpPeerColorProfileSet) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.peerColorProfileSet#767d61eb as nil")
	}
	b.PutVectorHeader(len(p.PaletteColors))
	for _, v := range p.PaletteColors {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(p.BgColors))
	for _, v := range p.BgColors {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(p.StoryColors))
	for _, v := range p.StoryColors {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPeerColorProfileSet) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.peerColorProfileSet#767d61eb to nil")
	}
	if err := b.ConsumeID(HelpPeerColorProfileSetTypeID); err != nil {
		return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *HelpPeerColorProfileSet) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.peerColorProfileSet#767d61eb to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: field palette_colors: %w", err)
		}

		if headerLen > 0 {
			p.PaletteColors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: field palette_colors: %w", err)
			}
			p.PaletteColors = append(p.PaletteColors, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: field bg_colors: %w", err)
		}

		if headerLen > 0 {
			p.BgColors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: field bg_colors: %w", err)
			}
			p.BgColors = append(p.BgColors, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: field story_colors: %w", err)
		}

		if headerLen > 0 {
			p.StoryColors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode help.peerColorProfileSet#767d61eb: field story_colors: %w", err)
			}
			p.StoryColors = append(p.StoryColors, value)
		}
	}
	return nil
}

// GetPaletteColors returns value of PaletteColors field.
func (p *HelpPeerColorProfileSet) GetPaletteColors() (value []int) {
	if p == nil {
		return
	}
	return p.PaletteColors
}

// GetBgColors returns value of BgColors field.
func (p *HelpPeerColorProfileSet) GetBgColors() (value []int) {
	if p == nil {
		return
	}
	return p.BgColors
}

// GetStoryColors returns value of StoryColors field.
func (p *HelpPeerColorProfileSet) GetStoryColors() (value []int) {
	if p == nil {
		return
	}
	return p.StoryColors
}

// HelpPeerColorSetClassName is schema name of HelpPeerColorSetClass.
const HelpPeerColorSetClassName = "help.PeerColorSet"

// HelpPeerColorSetClass represents help.PeerColorSet generic type.
//
// See https://core.telegram.org/type/help.PeerColorSet for reference.
//
// Constructors:
//   - [HelpPeerColorSet]
//   - [HelpPeerColorProfileSet]
//
// Example:
//
//	g, err := tg.DecodeHelpPeerColorSet(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.HelpPeerColorSet: // help.peerColorSet#26219a58
//	case *tg.HelpPeerColorProfileSet: // help.peerColorProfileSet#767d61eb
//	default: panic(v)
//	}
type HelpPeerColorSetClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpPeerColorSetClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeHelpPeerColorSet implements binary de-serialization for HelpPeerColorSetClass.
func DecodeHelpPeerColorSet(buf *bin.Buffer) (HelpPeerColorSetClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpPeerColorSetTypeID:
		// Decoding help.peerColorSet#26219a58.
		v := HelpPeerColorSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPeerColorSetClass: %w", err)
		}
		return &v, nil
	case HelpPeerColorProfileSetTypeID:
		// Decoding help.peerColorProfileSet#767d61eb.
		v := HelpPeerColorProfileSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPeerColorSetClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpPeerColorSetClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpPeerColorSet boxes the HelpPeerColorSetClass providing a helper.
type HelpPeerColorSetBox struct {
	PeerColorSet HelpPeerColorSetClass
}

// Decode implements bin.Decoder for HelpPeerColorSetBox.
func (b *HelpPeerColorSetBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpPeerColorSetBox to nil")
	}
	v, err := DecodeHelpPeerColorSet(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PeerColorSet = v
	return nil
}

// Encode implements bin.Encode for HelpPeerColorSetBox.
func (b *HelpPeerColorSetBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PeerColorSet == nil {
		return fmt.Errorf("unable to encode HelpPeerColorSetClass as nil")
	}
	return b.PeerColorSet.Encode(buf)
}
