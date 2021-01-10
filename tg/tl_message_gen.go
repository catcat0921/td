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

// MessageEmpty represents TL type `messageEmpty#83e5de54`.
// Empty constructor, non-existent message.
//
// See https://core.telegram.org/constructor/messageEmpty for reference.
type MessageEmpty struct {
	// Message identifier
	ID int
}

// MessageEmptyTypeID is TL type id of MessageEmpty.
const MessageEmptyTypeID = 0x83e5de54

func (m *MessageEmpty) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageEmpty) String() string {
	if m == nil {
		return "MessageEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessageEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(m.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageEmpty) TypeID() uint32 {
	return MessageEmptyTypeID
}

// Encode implements bin.Encoder.
func (m *MessageEmpty) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEmpty#83e5de54 as nil")
	}
	b.PutID(MessageEmptyTypeID)
	b.PutInt(m.ID)
	return nil
}

// GetID returns value of ID field.
func (m *MessageEmpty) GetID() (value int) {
	return m.ID
}

// Decode implements bin.Decoder.
func (m *MessageEmpty) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEmpty#83e5de54 to nil")
	}
	if err := b.ConsumeID(MessageEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEmpty#83e5de54: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEmpty#83e5de54: field id: %w", err)
		}
		m.ID = value
	}
	return nil
}

// construct implements constructor of MessageClass.
func (m MessageEmpty) construct() MessageClass { return &m }

// Ensuring interfaces in compile-time for MessageEmpty.
var (
	_ bin.Encoder = &MessageEmpty{}
	_ bin.Decoder = &MessageEmpty{}

	_ MessageClass = &MessageEmpty{}
)

// Message represents TL type `message#58ae39c9`.
// A message
//
// See https://core.telegram.org/constructor/message for reference.
type Message struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Is this an outgoing message
	Out bool
	// Whether we were mentioned¹ in this message
	//
	// Links:
	//  1) https://core.telegram.org/api/mentions
	Mentioned bool
	// Whether there are unread media attachments in this message
	MediaUnread bool
	// Whether this is a silent message (no notification triggered)
	Silent bool
	// Whether this is a channel post
	Post bool
	// Whether this is a scheduled message¹
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	FromScheduled bool
	// This is a legacy message: it has to be refetched with the new layer
	Legacy bool
	// Whether the message should be shown as not modified to the user, even if an edit date is present
	EditHide bool
	// Whether this message is pinned¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pin
	Pinned bool
	// ID of the message
	ID int
	// ID of the sender of the message
	//
	// Use SetFromID and GetFromID helpers.
	FromID PeerClass
	// Peer ID, the chat where this message was sent
	PeerID PeerClass
	// Info about forwarded messages
	//
	// Use SetFwdFrom and GetFwdFrom helpers.
	FwdFrom MessageFwdHeader
	// ID of the inline bot that generated the message
	//
	// Use SetViaBotID and GetViaBotID helpers.
	ViaBotID int
	// Reply information
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo MessageReplyHeader
	// Date of the message
	Date int
	// The message
	Message string
	// Media attachment
	//
	// Use SetMedia and GetMedia helpers.
	Media MessageMediaClass
	// Reply markup (bot/inline keyboards)
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
	// Message entities¹ for styled text
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// View count for channel posts
	//
	// Use SetViews and GetViews helpers.
	Views int
	// Forward counter
	//
	// Use SetForwards and GetForwards helpers.
	Forwards int
	// Info about post comments (for channels) or message replies (for groups)¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReplies and GetReplies helpers.
	Replies MessageReplies
	// Last edit date of this message
	//
	// Use SetEditDate and GetEditDate helpers.
	EditDate int
	// Name of the author of this message for channel posts (with signatures enabled)
	//
	// Use SetPostAuthor and GetPostAuthor helpers.
	PostAuthor string
	// Multiple media messages sent using messages.sendMultiMedia¹ with the same grouped ID indicate an album or media group²
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendMultiMedia
	//  2) https://core.telegram.org/api/files#albums-grouped-media
	//
	// Use SetGroupedID and GetGroupedID helpers.
	GroupedID int64
	// Contains the reason why access to this message must be restricted.
	//
	// Use SetRestrictionReason and GetRestrictionReason helpers.
	RestrictionReason []RestrictionReason
}

// MessageTypeID is TL type id of Message.
const MessageTypeID = 0x58ae39c9

func (m *Message) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Out == false) {
		return false
	}
	if !(m.Mentioned == false) {
		return false
	}
	if !(m.MediaUnread == false) {
		return false
	}
	if !(m.Silent == false) {
		return false
	}
	if !(m.Post == false) {
		return false
	}
	if !(m.FromScheduled == false) {
		return false
	}
	if !(m.Legacy == false) {
		return false
	}
	if !(m.EditHide == false) {
		return false
	}
	if !(m.Pinned == false) {
		return false
	}
	if !(m.ID == 0) {
		return false
	}
	if !(m.FromID == nil) {
		return false
	}
	if !(m.PeerID == nil) {
		return false
	}
	if !(m.FwdFrom.Zero()) {
		return false
	}
	if !(m.ViaBotID == 0) {
		return false
	}
	if !(m.ReplyTo.Zero()) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}
	if !(m.Message == "") {
		return false
	}
	if !(m.Media == nil) {
		return false
	}
	if !(m.ReplyMarkup == nil) {
		return false
	}
	if !(m.Entities == nil) {
		return false
	}
	if !(m.Views == 0) {
		return false
	}
	if !(m.Forwards == 0) {
		return false
	}
	if !(m.Replies.Zero()) {
		return false
	}
	if !(m.EditDate == 0) {
		return false
	}
	if !(m.PostAuthor == "") {
		return false
	}
	if !(m.GroupedID == 0) {
		return false
	}
	if !(m.RestrictionReason == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *Message) String() string {
	if m == nil {
		return "Message(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Message")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(m.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(m.ID))
	sb.WriteString(",\n")
	if m.Flags.Has(8) {
		sb.WriteString("\tFromID: ")
		sb.WriteString(fmt.Sprint(m.FromID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tPeerID: ")
	sb.WriteString(fmt.Sprint(m.PeerID))
	sb.WriteString(",\n")
	if m.Flags.Has(2) {
		sb.WriteString("\tFwdFrom: ")
		sb.WriteString(fmt.Sprint(m.FwdFrom))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(11) {
		sb.WriteString("\tViaBotID: ")
		sb.WriteString(fmt.Sprint(m.ViaBotID))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(3) {
		sb.WriteString("\tReplyTo: ")
		sb.WriteString(fmt.Sprint(m.ReplyTo))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(m.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(m.Message))
	sb.WriteString(",\n")
	if m.Flags.Has(9) {
		sb.WriteString("\tMedia: ")
		sb.WriteString(fmt.Sprint(m.Media))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(6) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(fmt.Sprint(m.ReplyMarkup))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(7) {
		sb.WriteByte('[')
		for _, v := range m.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if m.Flags.Has(10) {
		sb.WriteString("\tViews: ")
		sb.WriteString(fmt.Sprint(m.Views))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(10) {
		sb.WriteString("\tForwards: ")
		sb.WriteString(fmt.Sprint(m.Forwards))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(23) {
		sb.WriteString("\tReplies: ")
		sb.WriteString(fmt.Sprint(m.Replies))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(15) {
		sb.WriteString("\tEditDate: ")
		sb.WriteString(fmt.Sprint(m.EditDate))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(16) {
		sb.WriteString("\tPostAuthor: ")
		sb.WriteString(fmt.Sprint(m.PostAuthor))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(17) {
		sb.WriteString("\tGroupedID: ")
		sb.WriteString(fmt.Sprint(m.GroupedID))
		sb.WriteString(",\n")
	}
	if m.Flags.Has(22) {
		sb.WriteByte('[')
		for _, v := range m.RestrictionReason {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *Message) TypeID() uint32 {
	return MessageTypeID
}

// Encode implements bin.Encoder.
func (m *Message) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#58ae39c9 as nil")
	}
	b.PutID(MessageTypeID)
	if !(m.Out == false) {
		m.Flags.Set(1)
	}
	if !(m.Mentioned == false) {
		m.Flags.Set(4)
	}
	if !(m.MediaUnread == false) {
		m.Flags.Set(5)
	}
	if !(m.Silent == false) {
		m.Flags.Set(13)
	}
	if !(m.Post == false) {
		m.Flags.Set(14)
	}
	if !(m.FromScheduled == false) {
		m.Flags.Set(18)
	}
	if !(m.Legacy == false) {
		m.Flags.Set(19)
	}
	if !(m.EditHide == false) {
		m.Flags.Set(21)
	}
	if !(m.Pinned == false) {
		m.Flags.Set(24)
	}
	if !(m.FromID == nil) {
		m.Flags.Set(8)
	}
	if !(m.FwdFrom.Zero()) {
		m.Flags.Set(2)
	}
	if !(m.ViaBotID == 0) {
		m.Flags.Set(11)
	}
	if !(m.ReplyTo.Zero()) {
		m.Flags.Set(3)
	}
	if !(m.Media == nil) {
		m.Flags.Set(9)
	}
	if !(m.ReplyMarkup == nil) {
		m.Flags.Set(6)
	}
	if !(m.Entities == nil) {
		m.Flags.Set(7)
	}
	if !(m.Views == 0) {
		m.Flags.Set(10)
	}
	if !(m.Forwards == 0) {
		m.Flags.Set(10)
	}
	if !(m.Replies.Zero()) {
		m.Flags.Set(23)
	}
	if !(m.EditDate == 0) {
		m.Flags.Set(15)
	}
	if !(m.PostAuthor == "") {
		m.Flags.Set(16)
	}
	if !(m.GroupedID == 0) {
		m.Flags.Set(17)
	}
	if !(m.RestrictionReason == nil) {
		m.Flags.Set(22)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#58ae39c9: field flags: %w", err)
	}
	b.PutInt(m.ID)
	if m.Flags.Has(8) {
		if m.FromID == nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field from_id is nil")
		}
		if err := m.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field from_id: %w", err)
		}
	}
	if m.PeerID == nil {
		return fmt.Errorf("unable to encode message#58ae39c9: field peer_id is nil")
	}
	if err := m.PeerID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#58ae39c9: field peer_id: %w", err)
	}
	if m.Flags.Has(2) {
		if err := m.FwdFrom.Encode(b); err != nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field fwd_from: %w", err)
		}
	}
	if m.Flags.Has(11) {
		b.PutInt(m.ViaBotID)
	}
	if m.Flags.Has(3) {
		if err := m.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field reply_to: %w", err)
		}
	}
	b.PutInt(m.Date)
	b.PutString(m.Message)
	if m.Flags.Has(9) {
		if m.Media == nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field media is nil")
		}
		if err := m.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field media: %w", err)
		}
	}
	if m.Flags.Has(6) {
		if m.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field reply_markup is nil")
		}
		if err := m.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field reply_markup: %w", err)
		}
	}
	if m.Flags.Has(7) {
		b.PutVectorHeader(len(m.Entities))
		for idx, v := range m.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode message#58ae39c9: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode message#58ae39c9: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if m.Flags.Has(10) {
		b.PutInt(m.Views)
	}
	if m.Flags.Has(10) {
		b.PutInt(m.Forwards)
	}
	if m.Flags.Has(23) {
		if err := m.Replies.Encode(b); err != nil {
			return fmt.Errorf("unable to encode message#58ae39c9: field replies: %w", err)
		}
	}
	if m.Flags.Has(15) {
		b.PutInt(m.EditDate)
	}
	if m.Flags.Has(16) {
		b.PutString(m.PostAuthor)
	}
	if m.Flags.Has(17) {
		b.PutLong(m.GroupedID)
	}
	if m.Flags.Has(22) {
		b.PutVectorHeader(len(m.RestrictionReason))
		for idx, v := range m.RestrictionReason {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode message#58ae39c9: field restriction_reason element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetOut sets value of Out conditional field.
func (m *Message) SetOut(value bool) {
	if value {
		m.Flags.Set(1)
		m.Out = true
	} else {
		m.Flags.Unset(1)
		m.Out = false
	}
}

// GetOut returns value of Out conditional field.
func (m *Message) GetOut() (value bool) {
	return m.Flags.Has(1)
}

// SetMentioned sets value of Mentioned conditional field.
func (m *Message) SetMentioned(value bool) {
	if value {
		m.Flags.Set(4)
		m.Mentioned = true
	} else {
		m.Flags.Unset(4)
		m.Mentioned = false
	}
}

// GetMentioned returns value of Mentioned conditional field.
func (m *Message) GetMentioned() (value bool) {
	return m.Flags.Has(4)
}

// SetMediaUnread sets value of MediaUnread conditional field.
func (m *Message) SetMediaUnread(value bool) {
	if value {
		m.Flags.Set(5)
		m.MediaUnread = true
	} else {
		m.Flags.Unset(5)
		m.MediaUnread = false
	}
}

// GetMediaUnread returns value of MediaUnread conditional field.
func (m *Message) GetMediaUnread() (value bool) {
	return m.Flags.Has(5)
}

// SetSilent sets value of Silent conditional field.
func (m *Message) SetSilent(value bool) {
	if value {
		m.Flags.Set(13)
		m.Silent = true
	} else {
		m.Flags.Unset(13)
		m.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (m *Message) GetSilent() (value bool) {
	return m.Flags.Has(13)
}

// SetPost sets value of Post conditional field.
func (m *Message) SetPost(value bool) {
	if value {
		m.Flags.Set(14)
		m.Post = true
	} else {
		m.Flags.Unset(14)
		m.Post = false
	}
}

// GetPost returns value of Post conditional field.
func (m *Message) GetPost() (value bool) {
	return m.Flags.Has(14)
}

// SetFromScheduled sets value of FromScheduled conditional field.
func (m *Message) SetFromScheduled(value bool) {
	if value {
		m.Flags.Set(18)
		m.FromScheduled = true
	} else {
		m.Flags.Unset(18)
		m.FromScheduled = false
	}
}

// GetFromScheduled returns value of FromScheduled conditional field.
func (m *Message) GetFromScheduled() (value bool) {
	return m.Flags.Has(18)
}

// SetLegacy sets value of Legacy conditional field.
func (m *Message) SetLegacy(value bool) {
	if value {
		m.Flags.Set(19)
		m.Legacy = true
	} else {
		m.Flags.Unset(19)
		m.Legacy = false
	}
}

// GetLegacy returns value of Legacy conditional field.
func (m *Message) GetLegacy() (value bool) {
	return m.Flags.Has(19)
}

// SetEditHide sets value of EditHide conditional field.
func (m *Message) SetEditHide(value bool) {
	if value {
		m.Flags.Set(21)
		m.EditHide = true
	} else {
		m.Flags.Unset(21)
		m.EditHide = false
	}
}

// GetEditHide returns value of EditHide conditional field.
func (m *Message) GetEditHide() (value bool) {
	return m.Flags.Has(21)
}

// SetPinned sets value of Pinned conditional field.
func (m *Message) SetPinned(value bool) {
	if value {
		m.Flags.Set(24)
		m.Pinned = true
	} else {
		m.Flags.Unset(24)
		m.Pinned = false
	}
}

// GetPinned returns value of Pinned conditional field.
func (m *Message) GetPinned() (value bool) {
	return m.Flags.Has(24)
}

// GetID returns value of ID field.
func (m *Message) GetID() (value int) {
	return m.ID
}

// SetFromID sets value of FromID conditional field.
func (m *Message) SetFromID(value PeerClass) {
	m.Flags.Set(8)
	m.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (m *Message) GetFromID() (value PeerClass, ok bool) {
	if !m.Flags.Has(8) {
		return value, false
	}
	return m.FromID, true
}

// GetPeerID returns value of PeerID field.
func (m *Message) GetPeerID() (value PeerClass) {
	return m.PeerID
}

// SetFwdFrom sets value of FwdFrom conditional field.
func (m *Message) SetFwdFrom(value MessageFwdHeader) {
	m.Flags.Set(2)
	m.FwdFrom = value
}

// GetFwdFrom returns value of FwdFrom conditional field and
// boolean which is true if field was set.
func (m *Message) GetFwdFrom() (value MessageFwdHeader, ok bool) {
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.FwdFrom, true
}

// SetViaBotID sets value of ViaBotID conditional field.
func (m *Message) SetViaBotID(value int) {
	m.Flags.Set(11)
	m.ViaBotID = value
}

// GetViaBotID returns value of ViaBotID conditional field and
// boolean which is true if field was set.
func (m *Message) GetViaBotID() (value int, ok bool) {
	if !m.Flags.Has(11) {
		return value, false
	}
	return m.ViaBotID, true
}

// SetReplyTo sets value of ReplyTo conditional field.
func (m *Message) SetReplyTo(value MessageReplyHeader) {
	m.Flags.Set(3)
	m.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (m *Message) GetReplyTo() (value MessageReplyHeader, ok bool) {
	if !m.Flags.Has(3) {
		return value, false
	}
	return m.ReplyTo, true
}

// GetDate returns value of Date field.
func (m *Message) GetDate() (value int) {
	return m.Date
}

// GetMessage returns value of Message field.
func (m *Message) GetMessage() (value string) {
	return m.Message
}

// SetMedia sets value of Media conditional field.
func (m *Message) SetMedia(value MessageMediaClass) {
	m.Flags.Set(9)
	m.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (m *Message) GetMedia() (value MessageMediaClass, ok bool) {
	if !m.Flags.Has(9) {
		return value, false
	}
	return m.Media, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (m *Message) SetReplyMarkup(value ReplyMarkupClass) {
	m.Flags.Set(6)
	m.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (m *Message) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !m.Flags.Has(6) {
		return value, false
	}
	return m.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (m *Message) SetEntities(value []MessageEntityClass) {
	m.Flags.Set(7)
	m.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (m *Message) GetEntities() (value []MessageEntityClass, ok bool) {
	if !m.Flags.Has(7) {
		return value, false
	}
	return m.Entities, true
}

// SetViews sets value of Views conditional field.
func (m *Message) SetViews(value int) {
	m.Flags.Set(10)
	m.Views = value
}

// GetViews returns value of Views conditional field and
// boolean which is true if field was set.
func (m *Message) GetViews() (value int, ok bool) {
	if !m.Flags.Has(10) {
		return value, false
	}
	return m.Views, true
}

// SetForwards sets value of Forwards conditional field.
func (m *Message) SetForwards(value int) {
	m.Flags.Set(10)
	m.Forwards = value
}

// GetForwards returns value of Forwards conditional field and
// boolean which is true if field was set.
func (m *Message) GetForwards() (value int, ok bool) {
	if !m.Flags.Has(10) {
		return value, false
	}
	return m.Forwards, true
}

// SetReplies sets value of Replies conditional field.
func (m *Message) SetReplies(value MessageReplies) {
	m.Flags.Set(23)
	m.Replies = value
}

// GetReplies returns value of Replies conditional field and
// boolean which is true if field was set.
func (m *Message) GetReplies() (value MessageReplies, ok bool) {
	if !m.Flags.Has(23) {
		return value, false
	}
	return m.Replies, true
}

// SetEditDate sets value of EditDate conditional field.
func (m *Message) SetEditDate(value int) {
	m.Flags.Set(15)
	m.EditDate = value
}

// GetEditDate returns value of EditDate conditional field and
// boolean which is true if field was set.
func (m *Message) GetEditDate() (value int, ok bool) {
	if !m.Flags.Has(15) {
		return value, false
	}
	return m.EditDate, true
}

// SetPostAuthor sets value of PostAuthor conditional field.
func (m *Message) SetPostAuthor(value string) {
	m.Flags.Set(16)
	m.PostAuthor = value
}

// GetPostAuthor returns value of PostAuthor conditional field and
// boolean which is true if field was set.
func (m *Message) GetPostAuthor() (value string, ok bool) {
	if !m.Flags.Has(16) {
		return value, false
	}
	return m.PostAuthor, true
}

// SetGroupedID sets value of GroupedID conditional field.
func (m *Message) SetGroupedID(value int64) {
	m.Flags.Set(17)
	m.GroupedID = value
}

// GetGroupedID returns value of GroupedID conditional field and
// boolean which is true if field was set.
func (m *Message) GetGroupedID() (value int64, ok bool) {
	if !m.Flags.Has(17) {
		return value, false
	}
	return m.GroupedID, true
}

// SetRestrictionReason sets value of RestrictionReason conditional field.
func (m *Message) SetRestrictionReason(value []RestrictionReason) {
	m.Flags.Set(22)
	m.RestrictionReason = value
}

// GetRestrictionReason returns value of RestrictionReason conditional field and
// boolean which is true if field was set.
func (m *Message) GetRestrictionReason() (value []RestrictionReason, ok bool) {
	if !m.Flags.Has(22) {
		return value, false
	}
	return m.RestrictionReason, true
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#58ae39c9 to nil")
	}
	if err := b.ConsumeID(MessageTypeID); err != nil {
		return fmt.Errorf("unable to decode message#58ae39c9: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field flags: %w", err)
		}
	}
	m.Out = m.Flags.Has(1)
	m.Mentioned = m.Flags.Has(4)
	m.MediaUnread = m.Flags.Has(5)
	m.Silent = m.Flags.Has(13)
	m.Post = m.Flags.Has(14)
	m.FromScheduled = m.Flags.Has(18)
	m.Legacy = m.Flags.Has(19)
	m.EditHide = m.Flags.Has(21)
	m.Pinned = m.Flags.Has(24)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field id: %w", err)
		}
		m.ID = value
	}
	if m.Flags.Has(8) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field from_id: %w", err)
		}
		m.FromID = value
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field peer_id: %w", err)
		}
		m.PeerID = value
	}
	if m.Flags.Has(2) {
		if err := m.FwdFrom.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field fwd_from: %w", err)
		}
	}
	if m.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field via_bot_id: %w", err)
		}
		m.ViaBotID = value
	}
	if m.Flags.Has(3) {
		if err := m.ReplyTo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field reply_to: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field date: %w", err)
		}
		m.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field message: %w", err)
		}
		m.Message = value
	}
	if m.Flags.Has(9) {
		value, err := DecodeMessageMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field media: %w", err)
		}
		m.Media = value
	}
	if m.Flags.Has(6) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field reply_markup: %w", err)
		}
		m.ReplyMarkup = value
	}
	if m.Flags.Has(7) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode message#58ae39c9: field entities: %w", err)
			}
			m.Entities = append(m.Entities, value)
		}
	}
	if m.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field views: %w", err)
		}
		m.Views = value
	}
	if m.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field forwards: %w", err)
		}
		m.Forwards = value
	}
	if m.Flags.Has(23) {
		if err := m.Replies.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field replies: %w", err)
		}
	}
	if m.Flags.Has(15) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field edit_date: %w", err)
		}
		m.EditDate = value
	}
	if m.Flags.Has(16) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field post_author: %w", err)
		}
		m.PostAuthor = value
	}
	if m.Flags.Has(17) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field grouped_id: %w", err)
		}
		m.GroupedID = value
	}
	if m.Flags.Has(22) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode message#58ae39c9: field restriction_reason: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value RestrictionReason
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode message#58ae39c9: field restriction_reason: %w", err)
			}
			m.RestrictionReason = append(m.RestrictionReason, value)
		}
	}
	return nil
}

// construct implements constructor of MessageClass.
func (m Message) construct() MessageClass { return &m }

// Ensuring interfaces in compile-time for Message.
var (
	_ bin.Encoder = &Message{}
	_ bin.Decoder = &Message{}

	_ MessageClass = &Message{}
)

// MessageService represents TL type `messageService#286fa604`.
// Indicates a service message
//
// See https://core.telegram.org/constructor/messageService for reference.
type MessageService struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the message is outgoing
	Out bool
	// Whether we were mentioned in the message
	Mentioned bool
	// Whether the message contains unread media
	MediaUnread bool
	// Whether the message is silent
	Silent bool
	// Whether it's a channel post
	Post bool
	// This is a legacy message: it has to be refetched with the new layer
	Legacy bool
	// Message ID
	ID int
	// ID of the sender of this message
	//
	// Use SetFromID and GetFromID helpers.
	FromID PeerClass
	// Sender of service message
	PeerID PeerClass
	// Reply (thread) information
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo MessageReplyHeader
	// Message date
	Date int
	// Event connected with the service message
	Action MessageActionClass
}

// MessageServiceTypeID is TL type id of MessageService.
const MessageServiceTypeID = 0x286fa604

func (m *MessageService) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Out == false) {
		return false
	}
	if !(m.Mentioned == false) {
		return false
	}
	if !(m.MediaUnread == false) {
		return false
	}
	if !(m.Silent == false) {
		return false
	}
	if !(m.Post == false) {
		return false
	}
	if !(m.Legacy == false) {
		return false
	}
	if !(m.ID == 0) {
		return false
	}
	if !(m.FromID == nil) {
		return false
	}
	if !(m.PeerID == nil) {
		return false
	}
	if !(m.ReplyTo.Zero()) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}
	if !(m.Action == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageService) String() string {
	if m == nil {
		return "MessageService(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessageService")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(m.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(m.ID))
	sb.WriteString(",\n")
	if m.Flags.Has(8) {
		sb.WriteString("\tFromID: ")
		sb.WriteString(fmt.Sprint(m.FromID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tPeerID: ")
	sb.WriteString(fmt.Sprint(m.PeerID))
	sb.WriteString(",\n")
	if m.Flags.Has(3) {
		sb.WriteString("\tReplyTo: ")
		sb.WriteString(fmt.Sprint(m.ReplyTo))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(m.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tAction: ")
	sb.WriteString(fmt.Sprint(m.Action))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageService) TypeID() uint32 {
	return MessageServiceTypeID
}

// Encode implements bin.Encoder.
func (m *MessageService) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageService#286fa604 as nil")
	}
	b.PutID(MessageServiceTypeID)
	if !(m.Out == false) {
		m.Flags.Set(1)
	}
	if !(m.Mentioned == false) {
		m.Flags.Set(4)
	}
	if !(m.MediaUnread == false) {
		m.Flags.Set(5)
	}
	if !(m.Silent == false) {
		m.Flags.Set(13)
	}
	if !(m.Post == false) {
		m.Flags.Set(14)
	}
	if !(m.Legacy == false) {
		m.Flags.Set(19)
	}
	if !(m.FromID == nil) {
		m.Flags.Set(8)
	}
	if !(m.ReplyTo.Zero()) {
		m.Flags.Set(3)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageService#286fa604: field flags: %w", err)
	}
	b.PutInt(m.ID)
	if m.Flags.Has(8) {
		if m.FromID == nil {
			return fmt.Errorf("unable to encode messageService#286fa604: field from_id is nil")
		}
		if err := m.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageService#286fa604: field from_id: %w", err)
		}
	}
	if m.PeerID == nil {
		return fmt.Errorf("unable to encode messageService#286fa604: field peer_id is nil")
	}
	if err := m.PeerID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageService#286fa604: field peer_id: %w", err)
	}
	if m.Flags.Has(3) {
		if err := m.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageService#286fa604: field reply_to: %w", err)
		}
	}
	b.PutInt(m.Date)
	if m.Action == nil {
		return fmt.Errorf("unable to encode messageService#286fa604: field action is nil")
	}
	if err := m.Action.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageService#286fa604: field action: %w", err)
	}
	return nil
}

// SetOut sets value of Out conditional field.
func (m *MessageService) SetOut(value bool) {
	if value {
		m.Flags.Set(1)
		m.Out = true
	} else {
		m.Flags.Unset(1)
		m.Out = false
	}
}

// GetOut returns value of Out conditional field.
func (m *MessageService) GetOut() (value bool) {
	return m.Flags.Has(1)
}

// SetMentioned sets value of Mentioned conditional field.
func (m *MessageService) SetMentioned(value bool) {
	if value {
		m.Flags.Set(4)
		m.Mentioned = true
	} else {
		m.Flags.Unset(4)
		m.Mentioned = false
	}
}

// GetMentioned returns value of Mentioned conditional field.
func (m *MessageService) GetMentioned() (value bool) {
	return m.Flags.Has(4)
}

// SetMediaUnread sets value of MediaUnread conditional field.
func (m *MessageService) SetMediaUnread(value bool) {
	if value {
		m.Flags.Set(5)
		m.MediaUnread = true
	} else {
		m.Flags.Unset(5)
		m.MediaUnread = false
	}
}

// GetMediaUnread returns value of MediaUnread conditional field.
func (m *MessageService) GetMediaUnread() (value bool) {
	return m.Flags.Has(5)
}

// SetSilent sets value of Silent conditional field.
func (m *MessageService) SetSilent(value bool) {
	if value {
		m.Flags.Set(13)
		m.Silent = true
	} else {
		m.Flags.Unset(13)
		m.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (m *MessageService) GetSilent() (value bool) {
	return m.Flags.Has(13)
}

// SetPost sets value of Post conditional field.
func (m *MessageService) SetPost(value bool) {
	if value {
		m.Flags.Set(14)
		m.Post = true
	} else {
		m.Flags.Unset(14)
		m.Post = false
	}
}

// GetPost returns value of Post conditional field.
func (m *MessageService) GetPost() (value bool) {
	return m.Flags.Has(14)
}

// SetLegacy sets value of Legacy conditional field.
func (m *MessageService) SetLegacy(value bool) {
	if value {
		m.Flags.Set(19)
		m.Legacy = true
	} else {
		m.Flags.Unset(19)
		m.Legacy = false
	}
}

// GetLegacy returns value of Legacy conditional field.
func (m *MessageService) GetLegacy() (value bool) {
	return m.Flags.Has(19)
}

// GetID returns value of ID field.
func (m *MessageService) GetID() (value int) {
	return m.ID
}

// SetFromID sets value of FromID conditional field.
func (m *MessageService) SetFromID(value PeerClass) {
	m.Flags.Set(8)
	m.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (m *MessageService) GetFromID() (value PeerClass, ok bool) {
	if !m.Flags.Has(8) {
		return value, false
	}
	return m.FromID, true
}

// GetPeerID returns value of PeerID field.
func (m *MessageService) GetPeerID() (value PeerClass) {
	return m.PeerID
}

// SetReplyTo sets value of ReplyTo conditional field.
func (m *MessageService) SetReplyTo(value MessageReplyHeader) {
	m.Flags.Set(3)
	m.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (m *MessageService) GetReplyTo() (value MessageReplyHeader, ok bool) {
	if !m.Flags.Has(3) {
		return value, false
	}
	return m.ReplyTo, true
}

// GetDate returns value of Date field.
func (m *MessageService) GetDate() (value int) {
	return m.Date
}

// GetAction returns value of Action field.
func (m *MessageService) GetAction() (value MessageActionClass) {
	return m.Action
}

// Decode implements bin.Decoder.
func (m *MessageService) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageService#286fa604 to nil")
	}
	if err := b.ConsumeID(MessageServiceTypeID); err != nil {
		return fmt.Errorf("unable to decode messageService#286fa604: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field flags: %w", err)
		}
	}
	m.Out = m.Flags.Has(1)
	m.Mentioned = m.Flags.Has(4)
	m.MediaUnread = m.Flags.Has(5)
	m.Silent = m.Flags.Has(13)
	m.Post = m.Flags.Has(14)
	m.Legacy = m.Flags.Has(19)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field id: %w", err)
		}
		m.ID = value
	}
	if m.Flags.Has(8) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field from_id: %w", err)
		}
		m.FromID = value
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field peer_id: %w", err)
		}
		m.PeerID = value
	}
	if m.Flags.Has(3) {
		if err := m.ReplyTo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field reply_to: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field date: %w", err)
		}
		m.Date = value
	}
	{
		value, err := DecodeMessageAction(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageService#286fa604: field action: %w", err)
		}
		m.Action = value
	}
	return nil
}

// construct implements constructor of MessageClass.
func (m MessageService) construct() MessageClass { return &m }

// Ensuring interfaces in compile-time for MessageService.
var (
	_ bin.Encoder = &MessageService{}
	_ bin.Decoder = &MessageService{}

	_ MessageClass = &MessageService{}
)

// MessageClass represents Message generic type.
//
// See https://core.telegram.org/type/Message for reference.
//
// Example:
//  g, err := DecodeMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessageEmpty: // messageEmpty#83e5de54
//  case *Message: // message#58ae39c9
//  case *MessageService: // messageService#286fa604
//  default: panic(v)
//  }
type MessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessageClass

	// Message identifier
	GetID() (value int)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessage implements binary de-serialization for MessageClass.
func DecodeMessage(buf *bin.Buffer) (MessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageEmptyTypeID:
		// Decoding messageEmpty#83e5de54.
		v := MessageEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageClass: %w", err)
		}
		return &v, nil
	case MessageTypeID:
		// Decoding message#58ae39c9.
		v := Message{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageClass: %w", err)
		}
		return &v, nil
	case MessageServiceTypeID:
		// Decoding messageService#286fa604.
		v := MessageService{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// Message boxes the MessageClass providing a helper.
type MessageBox struct {
	Message MessageClass
}

// Decode implements bin.Decoder for MessageBox.
func (b *MessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageBox to nil")
	}
	v, err := DecodeMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Message = v
	return nil
}

// Encode implements bin.Encode for MessageBox.
func (b *MessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Message == nil {
		return fmt.Errorf("unable to encode MessageClass as nil")
	}
	return b.Message.Encode(buf)
}
