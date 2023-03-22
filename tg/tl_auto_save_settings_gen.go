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

// AutoSaveSettings represents TL type `autoSaveSettings#c84834ce`.
//
// See https://core.telegram.org/constructor/autoSaveSettings for reference.
type AutoSaveSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Photos field of AutoSaveSettings.
	Photos bool
	// Videos field of AutoSaveSettings.
	Videos bool
	// VideoMaxSize field of AutoSaveSettings.
	//
	// Use SetVideoMaxSize and GetVideoMaxSize helpers.
	VideoMaxSize int64
}

// AutoSaveSettingsTypeID is TL type id of AutoSaveSettings.
const AutoSaveSettingsTypeID = 0xc84834ce

// Ensuring interfaces in compile-time for AutoSaveSettings.
var (
	_ bin.Encoder     = &AutoSaveSettings{}
	_ bin.Decoder     = &AutoSaveSettings{}
	_ bin.BareEncoder = &AutoSaveSettings{}
	_ bin.BareDecoder = &AutoSaveSettings{}
)

func (a *AutoSaveSettings) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Photos == false) {
		return false
	}
	if !(a.Videos == false) {
		return false
	}
	if !(a.VideoMaxSize == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutoSaveSettings) String() string {
	if a == nil {
		return "AutoSaveSettings(nil)"
	}
	type Alias AutoSaveSettings
	return fmt.Sprintf("AutoSaveSettings%+v", Alias(*a))
}

// FillFrom fills AutoSaveSettings from given interface.
func (a *AutoSaveSettings) FillFrom(from interface {
	GetPhotos() (value bool)
	GetVideos() (value bool)
	GetVideoMaxSize() (value int64, ok bool)
}) {
	a.Photos = from.GetPhotos()
	a.Videos = from.GetVideos()
	if val, ok := from.GetVideoMaxSize(); ok {
		a.VideoMaxSize = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutoSaveSettings) TypeID() uint32 {
	return AutoSaveSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AutoSaveSettings) TypeName() string {
	return "autoSaveSettings"
}

// TypeInfo returns info about TL type.
func (a *AutoSaveSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autoSaveSettings",
		ID:   AutoSaveSettingsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photos",
			SchemaName: "photos",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "Videos",
			SchemaName: "videos",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "VideoMaxSize",
			SchemaName: "video_max_size",
			Null:       !a.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AutoSaveSettings) SetFlags() {
	if !(a.Photos == false) {
		a.Flags.Set(0)
	}
	if !(a.Videos == false) {
		a.Flags.Set(1)
	}
	if !(a.VideoMaxSize == 0) {
		a.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (a *AutoSaveSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoSaveSettings#c84834ce as nil")
	}
	b.PutID(AutoSaveSettingsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AutoSaveSettings) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoSaveSettings#c84834ce as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoSaveSettings#c84834ce: field flags: %w", err)
	}
	if a.Flags.Has(2) {
		b.PutLong(a.VideoMaxSize)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AutoSaveSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoSaveSettings#c84834ce to nil")
	}
	if err := b.ConsumeID(AutoSaveSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode autoSaveSettings#c84834ce: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AutoSaveSettings) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoSaveSettings#c84834ce to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autoSaveSettings#c84834ce: field flags: %w", err)
		}
	}
	a.Photos = a.Flags.Has(0)
	a.Videos = a.Flags.Has(1)
	if a.Flags.Has(2) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode autoSaveSettings#c84834ce: field video_max_size: %w", err)
		}
		a.VideoMaxSize = value
	}
	return nil
}

// SetPhotos sets value of Photos conditional field.
func (a *AutoSaveSettings) SetPhotos(value bool) {
	if value {
		a.Flags.Set(0)
		a.Photos = true
	} else {
		a.Flags.Unset(0)
		a.Photos = false
	}
}

// GetPhotos returns value of Photos conditional field.
func (a *AutoSaveSettings) GetPhotos() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetVideos sets value of Videos conditional field.
func (a *AutoSaveSettings) SetVideos(value bool) {
	if value {
		a.Flags.Set(1)
		a.Videos = true
	} else {
		a.Flags.Unset(1)
		a.Videos = false
	}
}

// GetVideos returns value of Videos conditional field.
func (a *AutoSaveSettings) GetVideos() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// SetVideoMaxSize sets value of VideoMaxSize conditional field.
func (a *AutoSaveSettings) SetVideoMaxSize(value int64) {
	a.Flags.Set(2)
	a.VideoMaxSize = value
}

// GetVideoMaxSize returns value of VideoMaxSize conditional field and
// boolean which is true if field was set.
func (a *AutoSaveSettings) GetVideoMaxSize() (value int64, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(2) {
		return value, false
	}
	return a.VideoMaxSize, true
}
