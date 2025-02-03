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

// ChannelsSponsoredMessageReportResultChooseOption represents TL type `channels.sponsoredMessageReportResultChooseOption#846f9e42`.
// The user must choose a report option from the localized options available in options,
// and after selection, channels.reportSponsoredMessage¹ must be invoked again, passing
// the option's option field to the option param of the method.
//
// Links:
//  1. https://core.telegram.org/method/channels.reportSponsoredMessage
//
// See https://core.telegram.org/constructor/channels.sponsoredMessageReportResultChooseOption for reference.
type ChannelsSponsoredMessageReportResultChooseOption struct {
	// Title of the option selection popup.
	Title string
	// Localized list of options.
	Options []SponsoredMessageReportOption
}

// ChannelsSponsoredMessageReportResultChooseOptionTypeID is TL type id of ChannelsSponsoredMessageReportResultChooseOption.
const ChannelsSponsoredMessageReportResultChooseOptionTypeID = 0x846f9e42

// construct implements constructor of ChannelsSponsoredMessageReportResultClass.
func (s ChannelsSponsoredMessageReportResultChooseOption) construct() ChannelsSponsoredMessageReportResultClass {
	return &s
}

// Ensuring interfaces in compile-time for ChannelsSponsoredMessageReportResultChooseOption.
var (
	_ bin.Encoder     = &ChannelsSponsoredMessageReportResultChooseOption{}
	_ bin.Decoder     = &ChannelsSponsoredMessageReportResultChooseOption{}
	_ bin.BareEncoder = &ChannelsSponsoredMessageReportResultChooseOption{}
	_ bin.BareDecoder = &ChannelsSponsoredMessageReportResultChooseOption{}

	_ ChannelsSponsoredMessageReportResultClass = &ChannelsSponsoredMessageReportResultChooseOption{}
)

func (s *ChannelsSponsoredMessageReportResultChooseOption) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Options == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ChannelsSponsoredMessageReportResultChooseOption) String() string {
	if s == nil {
		return "ChannelsSponsoredMessageReportResultChooseOption(nil)"
	}
	type Alias ChannelsSponsoredMessageReportResultChooseOption
	return fmt.Sprintf("ChannelsSponsoredMessageReportResultChooseOption%+v", Alias(*s))
}

// FillFrom fills ChannelsSponsoredMessageReportResultChooseOption from given interface.
func (s *ChannelsSponsoredMessageReportResultChooseOption) FillFrom(from interface {
	GetTitle() (value string)
	GetOptions() (value []SponsoredMessageReportOption)
}) {
	s.Title = from.GetTitle()
	s.Options = from.GetOptions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsSponsoredMessageReportResultChooseOption) TypeID() uint32 {
	return ChannelsSponsoredMessageReportResultChooseOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsSponsoredMessageReportResultChooseOption) TypeName() string {
	return "channels.sponsoredMessageReportResultChooseOption"
}

// TypeInfo returns info about TL type.
func (s *ChannelsSponsoredMessageReportResultChooseOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.sponsoredMessageReportResultChooseOption",
		ID:   ChannelsSponsoredMessageReportResultChooseOptionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ChannelsSponsoredMessageReportResultChooseOption) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.sponsoredMessageReportResultChooseOption#846f9e42 as nil")
	}
	b.PutID(ChannelsSponsoredMessageReportResultChooseOptionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ChannelsSponsoredMessageReportResultChooseOption) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.sponsoredMessageReportResultChooseOption#846f9e42 as nil")
	}
	b.PutString(s.Title)
	b.PutVectorHeader(len(s.Options))
	for idx, v := range s.Options {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.sponsoredMessageReportResultChooseOption#846f9e42: field options element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ChannelsSponsoredMessageReportResultChooseOption) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.sponsoredMessageReportResultChooseOption#846f9e42 to nil")
	}
	if err := b.ConsumeID(ChannelsSponsoredMessageReportResultChooseOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.sponsoredMessageReportResultChooseOption#846f9e42: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ChannelsSponsoredMessageReportResultChooseOption) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.sponsoredMessageReportResultChooseOption#846f9e42 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.sponsoredMessageReportResultChooseOption#846f9e42: field title: %w", err)
		}
		s.Title = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.sponsoredMessageReportResultChooseOption#846f9e42: field options: %w", err)
		}

		if headerLen > 0 {
			s.Options = make([]SponsoredMessageReportOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SponsoredMessageReportOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channels.sponsoredMessageReportResultChooseOption#846f9e42: field options: %w", err)
			}
			s.Options = append(s.Options, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (s *ChannelsSponsoredMessageReportResultChooseOption) GetTitle() (value string) {
	if s == nil {
		return
	}
	return s.Title
}

// GetOptions returns value of Options field.
func (s *ChannelsSponsoredMessageReportResultChooseOption) GetOptions() (value []SponsoredMessageReportOption) {
	if s == nil {
		return
	}
	return s.Options
}

// ChannelsSponsoredMessageReportResultAdsHidden represents TL type `channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f`.
// Sponsored messages were hidden for the user in all chats.
//
// See https://core.telegram.org/constructor/channels.sponsoredMessageReportResultAdsHidden for reference.
type ChannelsSponsoredMessageReportResultAdsHidden struct {
}

// ChannelsSponsoredMessageReportResultAdsHiddenTypeID is TL type id of ChannelsSponsoredMessageReportResultAdsHidden.
const ChannelsSponsoredMessageReportResultAdsHiddenTypeID = 0x3e3bcf2f

// construct implements constructor of ChannelsSponsoredMessageReportResultClass.
func (s ChannelsSponsoredMessageReportResultAdsHidden) construct() ChannelsSponsoredMessageReportResultClass {
	return &s
}

// Ensuring interfaces in compile-time for ChannelsSponsoredMessageReportResultAdsHidden.
var (
	_ bin.Encoder     = &ChannelsSponsoredMessageReportResultAdsHidden{}
	_ bin.Decoder     = &ChannelsSponsoredMessageReportResultAdsHidden{}
	_ bin.BareEncoder = &ChannelsSponsoredMessageReportResultAdsHidden{}
	_ bin.BareDecoder = &ChannelsSponsoredMessageReportResultAdsHidden{}

	_ ChannelsSponsoredMessageReportResultClass = &ChannelsSponsoredMessageReportResultAdsHidden{}
)

func (s *ChannelsSponsoredMessageReportResultAdsHidden) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *ChannelsSponsoredMessageReportResultAdsHidden) String() string {
	if s == nil {
		return "ChannelsSponsoredMessageReportResultAdsHidden(nil)"
	}
	type Alias ChannelsSponsoredMessageReportResultAdsHidden
	return fmt.Sprintf("ChannelsSponsoredMessageReportResultAdsHidden%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsSponsoredMessageReportResultAdsHidden) TypeID() uint32 {
	return ChannelsSponsoredMessageReportResultAdsHiddenTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsSponsoredMessageReportResultAdsHidden) TypeName() string {
	return "channels.sponsoredMessageReportResultAdsHidden"
}

// TypeInfo returns info about TL type.
func (s *ChannelsSponsoredMessageReportResultAdsHidden) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.sponsoredMessageReportResultAdsHidden",
		ID:   ChannelsSponsoredMessageReportResultAdsHiddenTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *ChannelsSponsoredMessageReportResultAdsHidden) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f as nil")
	}
	b.PutID(ChannelsSponsoredMessageReportResultAdsHiddenTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ChannelsSponsoredMessageReportResultAdsHidden) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ChannelsSponsoredMessageReportResultAdsHidden) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f to nil")
	}
	if err := b.ConsumeID(ChannelsSponsoredMessageReportResultAdsHiddenTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ChannelsSponsoredMessageReportResultAdsHidden) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f to nil")
	}
	return nil
}

// ChannelsSponsoredMessageReportResultReported represents TL type `channels.sponsoredMessageReportResultReported#ad798849`.
// The sponsored message was reported successfully.
//
// See https://core.telegram.org/constructor/channels.sponsoredMessageReportResultReported for reference.
type ChannelsSponsoredMessageReportResultReported struct {
}

// ChannelsSponsoredMessageReportResultReportedTypeID is TL type id of ChannelsSponsoredMessageReportResultReported.
const ChannelsSponsoredMessageReportResultReportedTypeID = 0xad798849

// construct implements constructor of ChannelsSponsoredMessageReportResultClass.
func (s ChannelsSponsoredMessageReportResultReported) construct() ChannelsSponsoredMessageReportResultClass {
	return &s
}

// Ensuring interfaces in compile-time for ChannelsSponsoredMessageReportResultReported.
var (
	_ bin.Encoder     = &ChannelsSponsoredMessageReportResultReported{}
	_ bin.Decoder     = &ChannelsSponsoredMessageReportResultReported{}
	_ bin.BareEncoder = &ChannelsSponsoredMessageReportResultReported{}
	_ bin.BareDecoder = &ChannelsSponsoredMessageReportResultReported{}

	_ ChannelsSponsoredMessageReportResultClass = &ChannelsSponsoredMessageReportResultReported{}
)

func (s *ChannelsSponsoredMessageReportResultReported) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *ChannelsSponsoredMessageReportResultReported) String() string {
	if s == nil {
		return "ChannelsSponsoredMessageReportResultReported(nil)"
	}
	type Alias ChannelsSponsoredMessageReportResultReported
	return fmt.Sprintf("ChannelsSponsoredMessageReportResultReported%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsSponsoredMessageReportResultReported) TypeID() uint32 {
	return ChannelsSponsoredMessageReportResultReportedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsSponsoredMessageReportResultReported) TypeName() string {
	return "channels.sponsoredMessageReportResultReported"
}

// TypeInfo returns info about TL type.
func (s *ChannelsSponsoredMessageReportResultReported) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.sponsoredMessageReportResultReported",
		ID:   ChannelsSponsoredMessageReportResultReportedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *ChannelsSponsoredMessageReportResultReported) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.sponsoredMessageReportResultReported#ad798849 as nil")
	}
	b.PutID(ChannelsSponsoredMessageReportResultReportedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ChannelsSponsoredMessageReportResultReported) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.sponsoredMessageReportResultReported#ad798849 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ChannelsSponsoredMessageReportResultReported) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.sponsoredMessageReportResultReported#ad798849 to nil")
	}
	if err := b.ConsumeID(ChannelsSponsoredMessageReportResultReportedTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.sponsoredMessageReportResultReported#ad798849: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ChannelsSponsoredMessageReportResultReported) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.sponsoredMessageReportResultReported#ad798849 to nil")
	}
	return nil
}

// ChannelsSponsoredMessageReportResultClassName is schema name of ChannelsSponsoredMessageReportResultClass.
const ChannelsSponsoredMessageReportResultClassName = "channels.SponsoredMessageReportResult"

// ChannelsSponsoredMessageReportResultClass represents channels.SponsoredMessageReportResult generic type.
//
// See https://core.telegram.org/type/channels.SponsoredMessageReportResult for reference.
//
// Constructors:
//   - [ChannelsSponsoredMessageReportResultChooseOption]
//   - [ChannelsSponsoredMessageReportResultAdsHidden]
//   - [ChannelsSponsoredMessageReportResultReported]
//
// Example:
//
//	g, err := tg.DecodeChannelsSponsoredMessageReportResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChannelsSponsoredMessageReportResultChooseOption: // channels.sponsoredMessageReportResultChooseOption#846f9e42
//	case *tg.ChannelsSponsoredMessageReportResultAdsHidden: // channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f
//	case *tg.ChannelsSponsoredMessageReportResultReported: // channels.sponsoredMessageReportResultReported#ad798849
//	default: panic(v)
//	}
type ChannelsSponsoredMessageReportResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChannelsSponsoredMessageReportResultClass

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

// DecodeChannelsSponsoredMessageReportResult implements binary de-serialization for ChannelsSponsoredMessageReportResultClass.
func DecodeChannelsSponsoredMessageReportResult(buf *bin.Buffer) (ChannelsSponsoredMessageReportResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelsSponsoredMessageReportResultChooseOptionTypeID:
		// Decoding channels.sponsoredMessageReportResultChooseOption#846f9e42.
		v := ChannelsSponsoredMessageReportResultChooseOption{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelsSponsoredMessageReportResultClass: %w", err)
		}
		return &v, nil
	case ChannelsSponsoredMessageReportResultAdsHiddenTypeID:
		// Decoding channels.sponsoredMessageReportResultAdsHidden#3e3bcf2f.
		v := ChannelsSponsoredMessageReportResultAdsHidden{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelsSponsoredMessageReportResultClass: %w", err)
		}
		return &v, nil
	case ChannelsSponsoredMessageReportResultReportedTypeID:
		// Decoding channels.sponsoredMessageReportResultReported#ad798849.
		v := ChannelsSponsoredMessageReportResultReported{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelsSponsoredMessageReportResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelsSponsoredMessageReportResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelsSponsoredMessageReportResult boxes the ChannelsSponsoredMessageReportResultClass providing a helper.
type ChannelsSponsoredMessageReportResultBox struct {
	SponsoredMessageReportResult ChannelsSponsoredMessageReportResultClass
}

// Decode implements bin.Decoder for ChannelsSponsoredMessageReportResultBox.
func (b *ChannelsSponsoredMessageReportResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelsSponsoredMessageReportResultBox to nil")
	}
	v, err := DecodeChannelsSponsoredMessageReportResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SponsoredMessageReportResult = v
	return nil
}

// Encode implements bin.Encode for ChannelsSponsoredMessageReportResultBox.
func (b *ChannelsSponsoredMessageReportResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SponsoredMessageReportResult == nil {
		return fmt.Errorf("unable to encode ChannelsSponsoredMessageReportResultClass as nil")
	}
	return b.SponsoredMessageReportResult.Encode(buf)
}
