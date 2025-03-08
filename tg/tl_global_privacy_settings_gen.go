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

// GlobalPrivacySettings represents TL type `globalPrivacySettings#c9d8df1c`.
// Global privacy settings
//
// See https://core.telegram.org/constructor/globalPrivacySettings for reference.
type GlobalPrivacySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to archive and mute new chats from non-contacts
	ArchiveAndMuteNewNoncontactPeers bool
	// Whether unmuted chats will be kept in the Archive chat list when they get a new
	// message.
	KeepArchivedUnmuted bool
	// Whether unmuted chats that are always included or pinned in a folder¹, will be kept
	// in the Archive chat list when they get a new message. Ignored if keep_archived_unmuted
	// is set.
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	KeepArchivedFolders bool
	// If this flag is set, the inputPrivacyKeyStatusTimestamp¹ key will also apply to the
	// ability to use messages.getOutboxReadDate² on messages sent to us. Meaning, users
	// that cannot see our exact last online date due to the current value of the
	// inputPrivacyKeyStatusTimestamp³ key will receive a 403 USER_PRIVACY_RESTRICTED error
	// when invoking messages.getOutboxReadDate⁴ to fetch the exact read date of a message
	// they sent to us. The userFull⁵.read_dates_private flag will be set for users that
	// have this flag enabled.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputPrivacyKeyStatusTimestamp
	//  2) https://core.telegram.org/method/messages.getOutboxReadDate
	//  3) https://core.telegram.org/constructor/inputPrivacyKeyStatusTimestamp
	//  4) https://core.telegram.org/method/messages.getOutboxReadDate
	//  5) https://core.telegram.org/constructor/userFull
	HideReadMarks bool
	// If set, only users that have a premium account, are in our contact list, or already
	// have a private chat with us can write to us; a 403 PRIVACY_PREMIUM_REQUIRED error will
	// be emitted otherwise.  The userFull¹.contact_require_premium flag will be set for
	// users that have this flag enabled.  To check whether we can write to a user with this
	// flag enabled, if we haven't yet cached all the required information (for example we
	// don't have the userFull² or history of all users while displaying the chat list in
	// the sharing UI) the users.getIsPremiumRequiredToContact³ method may be invoked,
	// passing the list of users currently visible in the UI, returning a list of booleans
	// that directly specify whether we can or cannot write to each user. This option may be
	// enabled by both non-Premium⁴ and Premium⁵ users only if the
	// new_noncontact_peers_require_premium_without_ownpremium client configuration flag
	// »⁶ is equal to true, otherwise it may be enabled only by Premium⁷ users and
	// non-Premium users will receive a PREMIUM_ACCOUNT_REQUIRED error when trying to enable
	// this flag.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/userFull
	//  2) https://core.telegram.org/constructor/userFull
	//  3) https://core.telegram.org/method/users.getIsPremiumRequiredToContact
	//  4) https://core.telegram.org/api/premium
	//  5) https://core.telegram.org/api/premium
	//  6) https://core.telegram.org/api/config#new-noncontact-peers-require-premium-without-ownpremium
	//  7) https://core.telegram.org/api/premium
	NewNoncontactPeersRequirePremium bool
	// NoncontactPeersPaidStars field of GlobalPrivacySettings.
	//
	// Use SetNoncontactPeersPaidStars and GetNoncontactPeersPaidStars helpers.
	NoncontactPeersPaidStars int64
}

// GlobalPrivacySettingsTypeID is TL type id of GlobalPrivacySettings.
const GlobalPrivacySettingsTypeID = 0xc9d8df1c

// Ensuring interfaces in compile-time for GlobalPrivacySettings.
var (
	_ bin.Encoder     = &GlobalPrivacySettings{}
	_ bin.Decoder     = &GlobalPrivacySettings{}
	_ bin.BareEncoder = &GlobalPrivacySettings{}
	_ bin.BareDecoder = &GlobalPrivacySettings{}
)

func (g *GlobalPrivacySettings) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		return false
	}
	if !(g.KeepArchivedUnmuted == false) {
		return false
	}
	if !(g.KeepArchivedFolders == false) {
		return false
	}
	if !(g.HideReadMarks == false) {
		return false
	}
	if !(g.NewNoncontactPeersRequirePremium == false) {
		return false
	}
	if !(g.NoncontactPeersPaidStars == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GlobalPrivacySettings) String() string {
	if g == nil {
		return "GlobalPrivacySettings(nil)"
	}
	type Alias GlobalPrivacySettings
	return fmt.Sprintf("GlobalPrivacySettings%+v", Alias(*g))
}

// FillFrom fills GlobalPrivacySettings from given interface.
func (g *GlobalPrivacySettings) FillFrom(from interface {
	GetArchiveAndMuteNewNoncontactPeers() (value bool)
	GetKeepArchivedUnmuted() (value bool)
	GetKeepArchivedFolders() (value bool)
	GetHideReadMarks() (value bool)
	GetNewNoncontactPeersRequirePremium() (value bool)
	GetNoncontactPeersPaidStars() (value int64, ok bool)
}) {
	g.ArchiveAndMuteNewNoncontactPeers = from.GetArchiveAndMuteNewNoncontactPeers()
	g.KeepArchivedUnmuted = from.GetKeepArchivedUnmuted()
	g.KeepArchivedFolders = from.GetKeepArchivedFolders()
	g.HideReadMarks = from.GetHideReadMarks()
	g.NewNoncontactPeersRequirePremium = from.GetNewNoncontactPeersRequirePremium()
	if val, ok := from.GetNoncontactPeersPaidStars(); ok {
		g.NoncontactPeersPaidStars = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GlobalPrivacySettings) TypeID() uint32 {
	return GlobalPrivacySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*GlobalPrivacySettings) TypeName() string {
	return "globalPrivacySettings"
}

// TypeInfo returns info about TL type.
func (g *GlobalPrivacySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "globalPrivacySettings",
		ID:   GlobalPrivacySettingsTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ArchiveAndMuteNewNoncontactPeers",
			SchemaName: "archive_and_mute_new_noncontact_peers",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "KeepArchivedUnmuted",
			SchemaName: "keep_archived_unmuted",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "KeepArchivedFolders",
			SchemaName: "keep_archived_folders",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "HideReadMarks",
			SchemaName: "hide_read_marks",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "NewNoncontactPeersRequirePremium",
			SchemaName: "new_noncontact_peers_require_premium",
			Null:       !g.Flags.Has(4),
		},
		{
			Name:       "NoncontactPeersPaidStars",
			SchemaName: "noncontact_peers_paid_stars",
			Null:       !g.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *GlobalPrivacySettings) SetFlags() {
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		g.Flags.Set(0)
	}
	if !(g.KeepArchivedUnmuted == false) {
		g.Flags.Set(1)
	}
	if !(g.KeepArchivedFolders == false) {
		g.Flags.Set(2)
	}
	if !(g.HideReadMarks == false) {
		g.Flags.Set(3)
	}
	if !(g.NewNoncontactPeersRequirePremium == false) {
		g.Flags.Set(4)
	}
	if !(g.NoncontactPeersPaidStars == 0) {
		g.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (g *GlobalPrivacySettings) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#c9d8df1c as nil")
	}
	b.PutID(GlobalPrivacySettingsTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GlobalPrivacySettings) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#c9d8df1c as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode globalPrivacySettings#c9d8df1c: field flags: %w", err)
	}
	if g.Flags.Has(5) {
		b.PutLong(g.NoncontactPeersPaidStars)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GlobalPrivacySettings) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#c9d8df1c to nil")
	}
	if err := b.ConsumeID(GlobalPrivacySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode globalPrivacySettings#c9d8df1c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GlobalPrivacySettings) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#c9d8df1c to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#c9d8df1c: field flags: %w", err)
		}
	}
	g.ArchiveAndMuteNewNoncontactPeers = g.Flags.Has(0)
	g.KeepArchivedUnmuted = g.Flags.Has(1)
	g.KeepArchivedFolders = g.Flags.Has(2)
	g.HideReadMarks = g.Flags.Has(3)
	g.NewNoncontactPeersRequirePremium = g.Flags.Has(4)
	if g.Flags.Has(5) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#c9d8df1c: field noncontact_peers_paid_stars: %w", err)
		}
		g.NoncontactPeersPaidStars = value
	}
	return nil
}

// SetArchiveAndMuteNewNoncontactPeers sets value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) SetArchiveAndMuteNewNoncontactPeers(value bool) {
	if value {
		g.Flags.Set(0)
		g.ArchiveAndMuteNewNoncontactPeers = true
	} else {
		g.Flags.Unset(0)
		g.ArchiveAndMuteNewNoncontactPeers = false
	}
}

// GetArchiveAndMuteNewNoncontactPeers returns value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) GetArchiveAndMuteNewNoncontactPeers() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// SetKeepArchivedUnmuted sets value of KeepArchivedUnmuted conditional field.
func (g *GlobalPrivacySettings) SetKeepArchivedUnmuted(value bool) {
	if value {
		g.Flags.Set(1)
		g.KeepArchivedUnmuted = true
	} else {
		g.Flags.Unset(1)
		g.KeepArchivedUnmuted = false
	}
}

// GetKeepArchivedUnmuted returns value of KeepArchivedUnmuted conditional field.
func (g *GlobalPrivacySettings) GetKeepArchivedUnmuted() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// SetKeepArchivedFolders sets value of KeepArchivedFolders conditional field.
func (g *GlobalPrivacySettings) SetKeepArchivedFolders(value bool) {
	if value {
		g.Flags.Set(2)
		g.KeepArchivedFolders = true
	} else {
		g.Flags.Unset(2)
		g.KeepArchivedFolders = false
	}
}

// GetKeepArchivedFolders returns value of KeepArchivedFolders conditional field.
func (g *GlobalPrivacySettings) GetKeepArchivedFolders() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// SetHideReadMarks sets value of HideReadMarks conditional field.
func (g *GlobalPrivacySettings) SetHideReadMarks(value bool) {
	if value {
		g.Flags.Set(3)
		g.HideReadMarks = true
	} else {
		g.Flags.Unset(3)
		g.HideReadMarks = false
	}
}

// GetHideReadMarks returns value of HideReadMarks conditional field.
func (g *GlobalPrivacySettings) GetHideReadMarks() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(3)
}

// SetNewNoncontactPeersRequirePremium sets value of NewNoncontactPeersRequirePremium conditional field.
func (g *GlobalPrivacySettings) SetNewNoncontactPeersRequirePremium(value bool) {
	if value {
		g.Flags.Set(4)
		g.NewNoncontactPeersRequirePremium = true
	} else {
		g.Flags.Unset(4)
		g.NewNoncontactPeersRequirePremium = false
	}
}

// GetNewNoncontactPeersRequirePremium returns value of NewNoncontactPeersRequirePremium conditional field.
func (g *GlobalPrivacySettings) GetNewNoncontactPeersRequirePremium() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(4)
}

// SetNoncontactPeersPaidStars sets value of NoncontactPeersPaidStars conditional field.
func (g *GlobalPrivacySettings) SetNoncontactPeersPaidStars(value int64) {
	g.Flags.Set(5)
	g.NoncontactPeersPaidStars = value
}

// GetNoncontactPeersPaidStars returns value of NoncontactPeersPaidStars conditional field and
// boolean which is true if field was set.
func (g *GlobalPrivacySettings) GetNoncontactPeersPaidStars() (value int64, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(5) {
		return value, false
	}
	return g.NoncontactPeersPaidStars, true
}
