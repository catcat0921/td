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

// InputPeerNotifySettings represents TL type `inputPeerNotifySettings#9c3d198e`.
// Notification settings.
//
// See https://core.telegram.org/constructor/inputPeerNotifySettings for reference.
type InputPeerNotifySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If the text of the message shall be displayed in notification
	//
	// Use SetShowPreviews and GetShowPreviews helpers.
	ShowPreviews bool
	// Peer was muted?
	//
	// Use SetSilent and GetSilent helpers.
	Silent bool
	// Date until which all notifications shall be switched off
	//
	// Use SetMuteUntil and GetMuteUntil helpers.
	MuteUntil int
	// Name of an audio file for notification
	//
	// Use SetSound and GetSound helpers.
	Sound string
}

// InputPeerNotifySettingsTypeID is TL type id of InputPeerNotifySettings.
const InputPeerNotifySettingsTypeID = 0x9c3d198e

func (i *InputPeerNotifySettings) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.ShowPreviews == false) {
		return false
	}
	if !(i.Silent == false) {
		return false
	}
	if !(i.MuteUntil == 0) {
		return false
	}
	if !(i.Sound == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerNotifySettings) String() string {
	if i == nil {
		return "InputPeerNotifySettings(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPeerNotifySettings")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(i.Flags))
	sb.WriteString(",\n")
	if i.Flags.Has(0) {
		sb.WriteString("\tShowPreviews: ")
		sb.WriteString(fmt.Sprint(i.ShowPreviews))
		sb.WriteString(",\n")
	}
	if i.Flags.Has(1) {
		sb.WriteString("\tSilent: ")
		sb.WriteString(fmt.Sprint(i.Silent))
		sb.WriteString(",\n")
	}
	if i.Flags.Has(2) {
		sb.WriteString("\tMuteUntil: ")
		sb.WriteString(fmt.Sprint(i.MuteUntil))
		sb.WriteString(",\n")
	}
	if i.Flags.Has(3) {
		sb.WriteString("\tSound: ")
		sb.WriteString(fmt.Sprint(i.Sound))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerNotifySettings) TypeID() uint32 {
	return InputPeerNotifySettingsTypeID
}

// Encode implements bin.Encoder.
func (i *InputPeerNotifySettings) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerNotifySettings#9c3d198e as nil")
	}
	b.PutID(InputPeerNotifySettingsTypeID)
	if !(i.ShowPreviews == false) {
		i.Flags.Set(0)
	}
	if !(i.Silent == false) {
		i.Flags.Set(1)
	}
	if !(i.MuteUntil == 0) {
		i.Flags.Set(2)
	}
	if !(i.Sound == "") {
		i.Flags.Set(3)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerNotifySettings#9c3d198e: field flags: %w", err)
	}
	if i.Flags.Has(0) {
		b.PutBool(i.ShowPreviews)
	}
	if i.Flags.Has(1) {
		b.PutBool(i.Silent)
	}
	if i.Flags.Has(2) {
		b.PutInt(i.MuteUntil)
	}
	if i.Flags.Has(3) {
		b.PutString(i.Sound)
	}
	return nil
}

// SetShowPreviews sets value of ShowPreviews conditional field.
func (i *InputPeerNotifySettings) SetShowPreviews(value bool) {
	i.Flags.Set(0)
	i.ShowPreviews = value
}

// GetShowPreviews returns value of ShowPreviews conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetShowPreviews() (value bool, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.ShowPreviews, true
}

// SetSilent sets value of Silent conditional field.
func (i *InputPeerNotifySettings) SetSilent(value bool) {
	i.Flags.Set(1)
	i.Silent = value
}

// GetSilent returns value of Silent conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetSilent() (value bool, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Silent, true
}

// SetMuteUntil sets value of MuteUntil conditional field.
func (i *InputPeerNotifySettings) SetMuteUntil(value int) {
	i.Flags.Set(2)
	i.MuteUntil = value
}

// GetMuteUntil returns value of MuteUntil conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetMuteUntil() (value int, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.MuteUntil, true
}

// SetSound sets value of Sound conditional field.
func (i *InputPeerNotifySettings) SetSound(value string) {
	i.Flags.Set(3)
	i.Sound = value
}

// GetSound returns value of Sound conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetSound() (value string, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.Sound, true
}

// Decode implements bin.Decoder.
func (i *InputPeerNotifySettings) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerNotifySettings#9c3d198e to nil")
	}
	if err := b.ConsumeID(InputPeerNotifySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerNotifySettings#9c3d198e: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#9c3d198e: field flags: %w", err)
		}
	}
	if i.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#9c3d198e: field show_previews: %w", err)
		}
		i.ShowPreviews = value
	}
	if i.Flags.Has(1) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#9c3d198e: field silent: %w", err)
		}
		i.Silent = value
	}
	if i.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#9c3d198e: field mute_until: %w", err)
		}
		i.MuteUntil = value
	}
	if i.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#9c3d198e: field sound: %w", err)
		}
		i.Sound = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputPeerNotifySettings.
var (
	_ bin.Encoder = &InputPeerNotifySettings{}
	_ bin.Decoder = &InputPeerNotifySettings{}
)
