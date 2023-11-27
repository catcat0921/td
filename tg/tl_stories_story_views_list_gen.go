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

// StoriesStoryViewsList represents TL type `stories.storyViewsList#46e9b9ec`.
// Reaction and view counters for a story¹
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/constructor/stories.storyViewsList for reference.
type StoriesStoryViewsList struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Total number of results that can be fetched
	Count int
	// ReactionsCount field of StoriesStoryViewsList.
	ReactionsCount int
	// Story view date and reaction information
	Views []StoryView
	// Mentioned users
	Users []UserClass
	// Offset for pagination
	//
	// Use SetNextOffset and GetNextOffset helpers.
	NextOffset string
}

// StoriesStoryViewsListTypeID is TL type id of StoriesStoryViewsList.
const StoriesStoryViewsListTypeID = 0x46e9b9ec

// Ensuring interfaces in compile-time for StoriesStoryViewsList.
var (
	_ bin.Encoder     = &StoriesStoryViewsList{}
	_ bin.Decoder     = &StoriesStoryViewsList{}
	_ bin.BareEncoder = &StoriesStoryViewsList{}
	_ bin.BareDecoder = &StoriesStoryViewsList{}
)

func (s *StoriesStoryViewsList) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.ReactionsCount == 0) {
		return false
	}
	if !(s.Views == nil) {
		return false
	}
	if !(s.Users == nil) {
		return false
	}
	if !(s.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoriesStoryViewsList) String() string {
	if s == nil {
		return "StoriesStoryViewsList(nil)"
	}
	type Alias StoriesStoryViewsList
	return fmt.Sprintf("StoriesStoryViewsList%+v", Alias(*s))
}

// FillFrom fills StoriesStoryViewsList from given interface.
func (s *StoriesStoryViewsList) FillFrom(from interface {
	GetCount() (value int)
	GetReactionsCount() (value int)
	GetViews() (value []StoryView)
	GetUsers() (value []UserClass)
	GetNextOffset() (value string, ok bool)
}) {
	s.Count = from.GetCount()
	s.ReactionsCount = from.GetReactionsCount()
	s.Views = from.GetViews()
	s.Users = from.GetUsers()
	if val, ok := from.GetNextOffset(); ok {
		s.NextOffset = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesStoryViewsList) TypeID() uint32 {
	return StoriesStoryViewsListTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesStoryViewsList) TypeName() string {
	return "stories.storyViewsList"
}

// TypeInfo returns info about TL type.
func (s *StoriesStoryViewsList) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.storyViewsList",
		ID:   StoriesStoryViewsListTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "ReactionsCount",
			SchemaName: "reactions_count",
		},
		{
			Name:       "Views",
			SchemaName: "views",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoriesStoryViewsList) SetFlags() {
	if !(s.NextOffset == "") {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *StoriesStoryViewsList) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.storyViewsList#46e9b9ec as nil")
	}
	b.PutID(StoriesStoryViewsListTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoriesStoryViewsList) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stories.storyViewsList#46e9b9ec as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.storyViewsList#46e9b9ec: field flags: %w", err)
	}
	b.PutInt(s.Count)
	b.PutInt(s.ReactionsCount)
	b.PutVectorHeader(len(s.Views))
	for idx, v := range s.Views {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.storyViewsList#46e9b9ec: field views element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Users))
	for idx, v := range s.Users {
		if v == nil {
			return fmt.Errorf("unable to encode stories.storyViewsList#46e9b9ec: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.storyViewsList#46e9b9ec: field users element with index %d: %w", idx, err)
		}
	}
	if s.Flags.Has(0) {
		b.PutString(s.NextOffset)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoriesStoryViewsList) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.storyViewsList#46e9b9ec to nil")
	}
	if err := b.ConsumeID(StoriesStoryViewsListTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoriesStoryViewsList) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stories.storyViewsList#46e9b9ec to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field count: %w", err)
		}
		s.Count = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field reactions_count: %w", err)
		}
		s.ReactionsCount = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field views: %w", err)
		}

		if headerLen > 0 {
			s.Views = make([]StoryView, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StoryView
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field views: %w", err)
			}
			s.Views = append(s.Views, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field users: %w", err)
		}

		if headerLen > 0 {
			s.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field users: %w", err)
			}
			s.Users = append(s.Users, value)
		}
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.storyViewsList#46e9b9ec: field next_offset: %w", err)
		}
		s.NextOffset = value
	}
	return nil
}

// GetCount returns value of Count field.
func (s *StoriesStoryViewsList) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}

// GetReactionsCount returns value of ReactionsCount field.
func (s *StoriesStoryViewsList) GetReactionsCount() (value int) {
	if s == nil {
		return
	}
	return s.ReactionsCount
}

// GetViews returns value of Views field.
func (s *StoriesStoryViewsList) GetViews() (value []StoryView) {
	if s == nil {
		return
	}
	return s.Views
}

// GetUsers returns value of Users field.
func (s *StoriesStoryViewsList) GetUsers() (value []UserClass) {
	if s == nil {
		return
	}
	return s.Users
}

// SetNextOffset sets value of NextOffset conditional field.
func (s *StoriesStoryViewsList) SetNextOffset(value string) {
	s.Flags.Set(0)
	s.NextOffset = value
}

// GetNextOffset returns value of NextOffset conditional field and
// boolean which is true if field was set.
func (s *StoriesStoryViewsList) GetNextOffset() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.NextOffset, true
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (s *StoriesStoryViewsList) MapUsers() (value UserClassArray) {
	return UserClassArray(s.Users)
}
