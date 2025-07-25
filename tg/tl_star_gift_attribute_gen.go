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

// StarGiftAttributeModel represents TL type `starGiftAttributeModel#39d99013`.
//
// See https://core.telegram.org/constructor/starGiftAttributeModel for reference.
type StarGiftAttributeModel struct {
	// Name field of StarGiftAttributeModel.
	Name string
	// Document field of StarGiftAttributeModel.
	Document DocumentClass
	// RarityPermille field of StarGiftAttributeModel.
	RarityPermille int
}

// StarGiftAttributeModelTypeID is TL type id of StarGiftAttributeModel.
const StarGiftAttributeModelTypeID = 0x39d99013

// construct implements constructor of StarGiftAttributeClass.
func (s StarGiftAttributeModel) construct() StarGiftAttributeClass { return &s }

// Ensuring interfaces in compile-time for StarGiftAttributeModel.
var (
	_ bin.Encoder     = &StarGiftAttributeModel{}
	_ bin.Decoder     = &StarGiftAttributeModel{}
	_ bin.BareEncoder = &StarGiftAttributeModel{}
	_ bin.BareDecoder = &StarGiftAttributeModel{}

	_ StarGiftAttributeClass = &StarGiftAttributeModel{}
)

func (s *StarGiftAttributeModel) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.Document == nil) {
		return false
	}
	if !(s.RarityPermille == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarGiftAttributeModel) String() string {
	if s == nil {
		return "StarGiftAttributeModel(nil)"
	}
	type Alias StarGiftAttributeModel
	return fmt.Sprintf("StarGiftAttributeModel%+v", Alias(*s))
}

// FillFrom fills StarGiftAttributeModel from given interface.
func (s *StarGiftAttributeModel) FillFrom(from interface {
	GetName() (value string)
	GetDocument() (value DocumentClass)
	GetRarityPermille() (value int)
}) {
	s.Name = from.GetName()
	s.Document = from.GetDocument()
	s.RarityPermille = from.GetRarityPermille()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarGiftAttributeModel) TypeID() uint32 {
	return StarGiftAttributeModelTypeID
}

// TypeName returns name of type in TL schema.
func (*StarGiftAttributeModel) TypeName() string {
	return "starGiftAttributeModel"
}

// TypeInfo returns info about TL type.
func (s *StarGiftAttributeModel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starGiftAttributeModel",
		ID:   StarGiftAttributeModelTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Document",
			SchemaName: "document",
		},
		{
			Name:       "RarityPermille",
			SchemaName: "rarity_permille",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarGiftAttributeModel) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributeModel#39d99013 as nil")
	}
	b.PutID(StarGiftAttributeModelTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarGiftAttributeModel) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributeModel#39d99013 as nil")
	}
	b.PutString(s.Name)
	if s.Document == nil {
		return fmt.Errorf("unable to encode starGiftAttributeModel#39d99013: field document is nil")
	}
	if err := s.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starGiftAttributeModel#39d99013: field document: %w", err)
	}
	b.PutInt(s.RarityPermille)
	return nil
}

// Decode implements bin.Decoder.
func (s *StarGiftAttributeModel) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributeModel#39d99013 to nil")
	}
	if err := b.ConsumeID(StarGiftAttributeModelTypeID); err != nil {
		return fmt.Errorf("unable to decode starGiftAttributeModel#39d99013: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarGiftAttributeModel) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributeModel#39d99013 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeModel#39d99013: field name: %w", err)
		}
		s.Name = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeModel#39d99013: field document: %w", err)
		}
		s.Document = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeModel#39d99013: field rarity_permille: %w", err)
		}
		s.RarityPermille = value
	}
	return nil
}

// GetName returns value of Name field.
func (s *StarGiftAttributeModel) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

// GetDocument returns value of Document field.
func (s *StarGiftAttributeModel) GetDocument() (value DocumentClass) {
	if s == nil {
		return
	}
	return s.Document
}

// GetRarityPermille returns value of RarityPermille field.
func (s *StarGiftAttributeModel) GetRarityPermille() (value int) {
	if s == nil {
		return
	}
	return s.RarityPermille
}

// StarGiftAttributePattern represents TL type `starGiftAttributePattern#13acff19`.
//
// See https://core.telegram.org/constructor/starGiftAttributePattern for reference.
type StarGiftAttributePattern struct {
	// Name field of StarGiftAttributePattern.
	Name string
	// Document field of StarGiftAttributePattern.
	Document DocumentClass
	// RarityPermille field of StarGiftAttributePattern.
	RarityPermille int
}

// StarGiftAttributePatternTypeID is TL type id of StarGiftAttributePattern.
const StarGiftAttributePatternTypeID = 0x13acff19

// construct implements constructor of StarGiftAttributeClass.
func (s StarGiftAttributePattern) construct() StarGiftAttributeClass { return &s }

// Ensuring interfaces in compile-time for StarGiftAttributePattern.
var (
	_ bin.Encoder     = &StarGiftAttributePattern{}
	_ bin.Decoder     = &StarGiftAttributePattern{}
	_ bin.BareEncoder = &StarGiftAttributePattern{}
	_ bin.BareDecoder = &StarGiftAttributePattern{}

	_ StarGiftAttributeClass = &StarGiftAttributePattern{}
)

func (s *StarGiftAttributePattern) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.Document == nil) {
		return false
	}
	if !(s.RarityPermille == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarGiftAttributePattern) String() string {
	if s == nil {
		return "StarGiftAttributePattern(nil)"
	}
	type Alias StarGiftAttributePattern
	return fmt.Sprintf("StarGiftAttributePattern%+v", Alias(*s))
}

// FillFrom fills StarGiftAttributePattern from given interface.
func (s *StarGiftAttributePattern) FillFrom(from interface {
	GetName() (value string)
	GetDocument() (value DocumentClass)
	GetRarityPermille() (value int)
}) {
	s.Name = from.GetName()
	s.Document = from.GetDocument()
	s.RarityPermille = from.GetRarityPermille()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarGiftAttributePattern) TypeID() uint32 {
	return StarGiftAttributePatternTypeID
}

// TypeName returns name of type in TL schema.
func (*StarGiftAttributePattern) TypeName() string {
	return "starGiftAttributePattern"
}

// TypeInfo returns info about TL type.
func (s *StarGiftAttributePattern) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starGiftAttributePattern",
		ID:   StarGiftAttributePatternTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Document",
			SchemaName: "document",
		},
		{
			Name:       "RarityPermille",
			SchemaName: "rarity_permille",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarGiftAttributePattern) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributePattern#13acff19 as nil")
	}
	b.PutID(StarGiftAttributePatternTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarGiftAttributePattern) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributePattern#13acff19 as nil")
	}
	b.PutString(s.Name)
	if s.Document == nil {
		return fmt.Errorf("unable to encode starGiftAttributePattern#13acff19: field document is nil")
	}
	if err := s.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starGiftAttributePattern#13acff19: field document: %w", err)
	}
	b.PutInt(s.RarityPermille)
	return nil
}

// Decode implements bin.Decoder.
func (s *StarGiftAttributePattern) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributePattern#13acff19 to nil")
	}
	if err := b.ConsumeID(StarGiftAttributePatternTypeID); err != nil {
		return fmt.Errorf("unable to decode starGiftAttributePattern#13acff19: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarGiftAttributePattern) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributePattern#13acff19 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributePattern#13acff19: field name: %w", err)
		}
		s.Name = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributePattern#13acff19: field document: %w", err)
		}
		s.Document = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributePattern#13acff19: field rarity_permille: %w", err)
		}
		s.RarityPermille = value
	}
	return nil
}

// GetName returns value of Name field.
func (s *StarGiftAttributePattern) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

// GetDocument returns value of Document field.
func (s *StarGiftAttributePattern) GetDocument() (value DocumentClass) {
	if s == nil {
		return
	}
	return s.Document
}

// GetRarityPermille returns value of RarityPermille field.
func (s *StarGiftAttributePattern) GetRarityPermille() (value int) {
	if s == nil {
		return
	}
	return s.RarityPermille
}

// StarGiftAttributeBackdrop represents TL type `starGiftAttributeBackdrop#d93d859c`.
//
// See https://core.telegram.org/constructor/starGiftAttributeBackdrop for reference.
type StarGiftAttributeBackdrop struct {
	// Name field of StarGiftAttributeBackdrop.
	Name string
	// BackdropID field of StarGiftAttributeBackdrop.
	BackdropID int
	// CenterColor field of StarGiftAttributeBackdrop.
	CenterColor int
	// EdgeColor field of StarGiftAttributeBackdrop.
	EdgeColor int
	// PatternColor field of StarGiftAttributeBackdrop.
	PatternColor int
	// TextColor field of StarGiftAttributeBackdrop.
	TextColor int
	// RarityPermille field of StarGiftAttributeBackdrop.
	RarityPermille int
}

// StarGiftAttributeBackdropTypeID is TL type id of StarGiftAttributeBackdrop.
const StarGiftAttributeBackdropTypeID = 0xd93d859c

// construct implements constructor of StarGiftAttributeClass.
func (s StarGiftAttributeBackdrop) construct() StarGiftAttributeClass { return &s }

// Ensuring interfaces in compile-time for StarGiftAttributeBackdrop.
var (
	_ bin.Encoder     = &StarGiftAttributeBackdrop{}
	_ bin.Decoder     = &StarGiftAttributeBackdrop{}
	_ bin.BareEncoder = &StarGiftAttributeBackdrop{}
	_ bin.BareDecoder = &StarGiftAttributeBackdrop{}

	_ StarGiftAttributeClass = &StarGiftAttributeBackdrop{}
)

func (s *StarGiftAttributeBackdrop) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.BackdropID == 0) {
		return false
	}
	if !(s.CenterColor == 0) {
		return false
	}
	if !(s.EdgeColor == 0) {
		return false
	}
	if !(s.PatternColor == 0) {
		return false
	}
	if !(s.TextColor == 0) {
		return false
	}
	if !(s.RarityPermille == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarGiftAttributeBackdrop) String() string {
	if s == nil {
		return "StarGiftAttributeBackdrop(nil)"
	}
	type Alias StarGiftAttributeBackdrop
	return fmt.Sprintf("StarGiftAttributeBackdrop%+v", Alias(*s))
}

// FillFrom fills StarGiftAttributeBackdrop from given interface.
func (s *StarGiftAttributeBackdrop) FillFrom(from interface {
	GetName() (value string)
	GetBackdropID() (value int)
	GetCenterColor() (value int)
	GetEdgeColor() (value int)
	GetPatternColor() (value int)
	GetTextColor() (value int)
	GetRarityPermille() (value int)
}) {
	s.Name = from.GetName()
	s.BackdropID = from.GetBackdropID()
	s.CenterColor = from.GetCenterColor()
	s.EdgeColor = from.GetEdgeColor()
	s.PatternColor = from.GetPatternColor()
	s.TextColor = from.GetTextColor()
	s.RarityPermille = from.GetRarityPermille()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarGiftAttributeBackdrop) TypeID() uint32 {
	return StarGiftAttributeBackdropTypeID
}

// TypeName returns name of type in TL schema.
func (*StarGiftAttributeBackdrop) TypeName() string {
	return "starGiftAttributeBackdrop"
}

// TypeInfo returns info about TL type.
func (s *StarGiftAttributeBackdrop) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starGiftAttributeBackdrop",
		ID:   StarGiftAttributeBackdropTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "BackdropID",
			SchemaName: "backdrop_id",
		},
		{
			Name:       "CenterColor",
			SchemaName: "center_color",
		},
		{
			Name:       "EdgeColor",
			SchemaName: "edge_color",
		},
		{
			Name:       "PatternColor",
			SchemaName: "pattern_color",
		},
		{
			Name:       "TextColor",
			SchemaName: "text_color",
		},
		{
			Name:       "RarityPermille",
			SchemaName: "rarity_permille",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StarGiftAttributeBackdrop) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributeBackdrop#d93d859c as nil")
	}
	b.PutID(StarGiftAttributeBackdropTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarGiftAttributeBackdrop) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributeBackdrop#d93d859c as nil")
	}
	b.PutString(s.Name)
	b.PutInt(s.BackdropID)
	b.PutInt(s.CenterColor)
	b.PutInt(s.EdgeColor)
	b.PutInt(s.PatternColor)
	b.PutInt(s.TextColor)
	b.PutInt(s.RarityPermille)
	return nil
}

// Decode implements bin.Decoder.
func (s *StarGiftAttributeBackdrop) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributeBackdrop#d93d859c to nil")
	}
	if err := b.ConsumeID(StarGiftAttributeBackdropTypeID); err != nil {
		return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarGiftAttributeBackdrop) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributeBackdrop#d93d859c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field name: %w", err)
		}
		s.Name = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field backdrop_id: %w", err)
		}
		s.BackdropID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field center_color: %w", err)
		}
		s.CenterColor = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field edge_color: %w", err)
		}
		s.EdgeColor = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field pattern_color: %w", err)
		}
		s.PatternColor = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field text_color: %w", err)
		}
		s.TextColor = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeBackdrop#d93d859c: field rarity_permille: %w", err)
		}
		s.RarityPermille = value
	}
	return nil
}

// GetName returns value of Name field.
func (s *StarGiftAttributeBackdrop) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

// GetBackdropID returns value of BackdropID field.
func (s *StarGiftAttributeBackdrop) GetBackdropID() (value int) {
	if s == nil {
		return
	}
	return s.BackdropID
}

// GetCenterColor returns value of CenterColor field.
func (s *StarGiftAttributeBackdrop) GetCenterColor() (value int) {
	if s == nil {
		return
	}
	return s.CenterColor
}

// GetEdgeColor returns value of EdgeColor field.
func (s *StarGiftAttributeBackdrop) GetEdgeColor() (value int) {
	if s == nil {
		return
	}
	return s.EdgeColor
}

// GetPatternColor returns value of PatternColor field.
func (s *StarGiftAttributeBackdrop) GetPatternColor() (value int) {
	if s == nil {
		return
	}
	return s.PatternColor
}

// GetTextColor returns value of TextColor field.
func (s *StarGiftAttributeBackdrop) GetTextColor() (value int) {
	if s == nil {
		return
	}
	return s.TextColor
}

// GetRarityPermille returns value of RarityPermille field.
func (s *StarGiftAttributeBackdrop) GetRarityPermille() (value int) {
	if s == nil {
		return
	}
	return s.RarityPermille
}

// StarGiftAttributeOriginalDetails represents TL type `starGiftAttributeOriginalDetails#e0bff26c`.
//
// See https://core.telegram.org/constructor/starGiftAttributeOriginalDetails for reference.
type StarGiftAttributeOriginalDetails struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// SenderID field of StarGiftAttributeOriginalDetails.
	//
	// Use SetSenderID and GetSenderID helpers.
	SenderID PeerClass
	// RecipientID field of StarGiftAttributeOriginalDetails.
	RecipientID PeerClass
	// Date field of StarGiftAttributeOriginalDetails.
	Date int
	// Message field of StarGiftAttributeOriginalDetails.
	//
	// Use SetMessage and GetMessage helpers.
	Message TextWithEntities
}

// StarGiftAttributeOriginalDetailsTypeID is TL type id of StarGiftAttributeOriginalDetails.
const StarGiftAttributeOriginalDetailsTypeID = 0xe0bff26c

// construct implements constructor of StarGiftAttributeClass.
func (s StarGiftAttributeOriginalDetails) construct() StarGiftAttributeClass { return &s }

// Ensuring interfaces in compile-time for StarGiftAttributeOriginalDetails.
var (
	_ bin.Encoder     = &StarGiftAttributeOriginalDetails{}
	_ bin.Decoder     = &StarGiftAttributeOriginalDetails{}
	_ bin.BareEncoder = &StarGiftAttributeOriginalDetails{}
	_ bin.BareDecoder = &StarGiftAttributeOriginalDetails{}

	_ StarGiftAttributeClass = &StarGiftAttributeOriginalDetails{}
)

func (s *StarGiftAttributeOriginalDetails) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.SenderID == nil) {
		return false
	}
	if !(s.RecipientID == nil) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.Message.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarGiftAttributeOriginalDetails) String() string {
	if s == nil {
		return "StarGiftAttributeOriginalDetails(nil)"
	}
	type Alias StarGiftAttributeOriginalDetails
	return fmt.Sprintf("StarGiftAttributeOriginalDetails%+v", Alias(*s))
}

// FillFrom fills StarGiftAttributeOriginalDetails from given interface.
func (s *StarGiftAttributeOriginalDetails) FillFrom(from interface {
	GetSenderID() (value PeerClass, ok bool)
	GetRecipientID() (value PeerClass)
	GetDate() (value int)
	GetMessage() (value TextWithEntities, ok bool)
}) {
	if val, ok := from.GetSenderID(); ok {
		s.SenderID = val
	}

	s.RecipientID = from.GetRecipientID()
	s.Date = from.GetDate()
	if val, ok := from.GetMessage(); ok {
		s.Message = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarGiftAttributeOriginalDetails) TypeID() uint32 {
	return StarGiftAttributeOriginalDetailsTypeID
}

// TypeName returns name of type in TL schema.
func (*StarGiftAttributeOriginalDetails) TypeName() string {
	return "starGiftAttributeOriginalDetails"
}

// TypeInfo returns info about TL type.
func (s *StarGiftAttributeOriginalDetails) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starGiftAttributeOriginalDetails",
		ID:   StarGiftAttributeOriginalDetailsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "RecipientID",
			SchemaName: "recipient_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Message",
			SchemaName: "message",
			Null:       !s.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StarGiftAttributeOriginalDetails) SetFlags() {
	if !(s.SenderID == nil) {
		s.Flags.Set(0)
	}
	if !(s.Message.Zero()) {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *StarGiftAttributeOriginalDetails) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributeOriginalDetails#e0bff26c as nil")
	}
	b.PutID(StarGiftAttributeOriginalDetailsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarGiftAttributeOriginalDetails) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starGiftAttributeOriginalDetails#e0bff26c as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starGiftAttributeOriginalDetails#e0bff26c: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		if s.SenderID == nil {
			return fmt.Errorf("unable to encode starGiftAttributeOriginalDetails#e0bff26c: field sender_id is nil")
		}
		if err := s.SenderID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starGiftAttributeOriginalDetails#e0bff26c: field sender_id: %w", err)
		}
	}
	if s.RecipientID == nil {
		return fmt.Errorf("unable to encode starGiftAttributeOriginalDetails#e0bff26c: field recipient_id is nil")
	}
	if err := s.RecipientID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starGiftAttributeOriginalDetails#e0bff26c: field recipient_id: %w", err)
	}
	b.PutInt(s.Date)
	if s.Flags.Has(1) {
		if err := s.Message.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starGiftAttributeOriginalDetails#e0bff26c: field message: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarGiftAttributeOriginalDetails) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributeOriginalDetails#e0bff26c to nil")
	}
	if err := b.ConsumeID(StarGiftAttributeOriginalDetailsTypeID); err != nil {
		return fmt.Errorf("unable to decode starGiftAttributeOriginalDetails#e0bff26c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarGiftAttributeOriginalDetails) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starGiftAttributeOriginalDetails#e0bff26c to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeOriginalDetails#e0bff26c: field flags: %w", err)
		}
	}
	if s.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeOriginalDetails#e0bff26c: field sender_id: %w", err)
		}
		s.SenderID = value
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeOriginalDetails#e0bff26c: field recipient_id: %w", err)
		}
		s.RecipientID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeOriginalDetails#e0bff26c: field date: %w", err)
		}
		s.Date = value
	}
	if s.Flags.Has(1) {
		if err := s.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starGiftAttributeOriginalDetails#e0bff26c: field message: %w", err)
		}
	}
	return nil
}

// SetSenderID sets value of SenderID conditional field.
func (s *StarGiftAttributeOriginalDetails) SetSenderID(value PeerClass) {
	s.Flags.Set(0)
	s.SenderID = value
}

// GetSenderID returns value of SenderID conditional field and
// boolean which is true if field was set.
func (s *StarGiftAttributeOriginalDetails) GetSenderID() (value PeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.SenderID, true
}

// GetRecipientID returns value of RecipientID field.
func (s *StarGiftAttributeOriginalDetails) GetRecipientID() (value PeerClass) {
	if s == nil {
		return
	}
	return s.RecipientID
}

// GetDate returns value of Date field.
func (s *StarGiftAttributeOriginalDetails) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// SetMessage sets value of Message conditional field.
func (s *StarGiftAttributeOriginalDetails) SetMessage(value TextWithEntities) {
	s.Flags.Set(1)
	s.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (s *StarGiftAttributeOriginalDetails) GetMessage() (value TextWithEntities, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Message, true
}

// StarGiftAttributeClassName is schema name of StarGiftAttributeClass.
const StarGiftAttributeClassName = "StarGiftAttribute"

// StarGiftAttributeClass represents StarGiftAttribute generic type.
//
// See https://core.telegram.org/type/StarGiftAttribute for reference.
//
// Constructors:
//   - [StarGiftAttributeModel]
//   - [StarGiftAttributePattern]
//   - [StarGiftAttributeBackdrop]
//   - [StarGiftAttributeOriginalDetails]
//
// Example:
//
//	g, err := tg.DecodeStarGiftAttribute(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.StarGiftAttributeModel: // starGiftAttributeModel#39d99013
//	case *tg.StarGiftAttributePattern: // starGiftAttributePattern#13acff19
//	case *tg.StarGiftAttributeBackdrop: // starGiftAttributeBackdrop#d93d859c
//	case *tg.StarGiftAttributeOriginalDetails: // starGiftAttributeOriginalDetails#e0bff26c
//	default: panic(v)
//	}
type StarGiftAttributeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StarGiftAttributeClass

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

// DecodeStarGiftAttribute implements binary de-serialization for StarGiftAttributeClass.
func DecodeStarGiftAttribute(buf *bin.Buffer) (StarGiftAttributeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StarGiftAttributeModelTypeID:
		// Decoding starGiftAttributeModel#39d99013.
		v := StarGiftAttributeModel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarGiftAttributeClass: %w", err)
		}
		return &v, nil
	case StarGiftAttributePatternTypeID:
		// Decoding starGiftAttributePattern#13acff19.
		v := StarGiftAttributePattern{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarGiftAttributeClass: %w", err)
		}
		return &v, nil
	case StarGiftAttributeBackdropTypeID:
		// Decoding starGiftAttributeBackdrop#d93d859c.
		v := StarGiftAttributeBackdrop{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarGiftAttributeClass: %w", err)
		}
		return &v, nil
	case StarGiftAttributeOriginalDetailsTypeID:
		// Decoding starGiftAttributeOriginalDetails#e0bff26c.
		v := StarGiftAttributeOriginalDetails{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StarGiftAttributeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StarGiftAttributeClass: %w", bin.NewUnexpectedID(id))
	}
}

// StarGiftAttribute boxes the StarGiftAttributeClass providing a helper.
type StarGiftAttributeBox struct {
	StarGiftAttribute StarGiftAttributeClass
}

// Decode implements bin.Decoder for StarGiftAttributeBox.
func (b *StarGiftAttributeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StarGiftAttributeBox to nil")
	}
	v, err := DecodeStarGiftAttribute(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StarGiftAttribute = v
	return nil
}

// Encode implements bin.Encode for StarGiftAttributeBox.
func (b *StarGiftAttributeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StarGiftAttribute == nil {
		return fmt.Errorf("unable to encode StarGiftAttributeClass as nil")
	}
	return b.StarGiftAttribute.Encode(buf)
}
