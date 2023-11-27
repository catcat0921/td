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

// MessagesSendMessageRequest represents TL type `messages.sendMessage#280d096f`.
// Sends a message to a chat
//
// See https://core.telegram.org/method/messages.sendMessage for reference.
type MessagesSendMessageRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag to disable generation of the webpage preview
	NoWebpage bool
	// Send this message silently (no notifications for the receivers)
	Silent bool
	// Send this message as background message
	Background bool
	// Clear the draft field
	ClearDraft bool
	// Only for bots, disallows forwarding and saving of the messages, even if the
	// destination chat doesn't have content protection¹ enabled
	//
	// Links:
	//  1) https://telegram.org/blog/protected-content-delete-by-date-and-more
	Noforwards bool
	// Whether to move used stickersets to top, see here for more info on this flag »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/stickers#recent-stickersets
	UpdateStickersetsOrder bool
	// InvertMedia field of MessagesSendMessageRequest.
	InvertMedia bool
	// The destination where the message will be sent
	Peer InputPeerClass
	// If set, indicates that the message should be sent in reply to the specified message or
	// story.
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo InputReplyToClass
	// The message
	Message string
	// Unique client message ID required to prevent message resending
	RandomID int64
	// Reply markup for sending bot buttons
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
	// Message entities¹ for sending styled text
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Scheduled message date for scheduled messages¹
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// Send this message as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// MessagesSendMessageRequestTypeID is TL type id of MessagesSendMessageRequest.
const MessagesSendMessageRequestTypeID = 0x280d096f

// Ensuring interfaces in compile-time for MessagesSendMessageRequest.
var (
	_ bin.Encoder     = &MessagesSendMessageRequest{}
	_ bin.Decoder     = &MessagesSendMessageRequest{}
	_ bin.BareEncoder = &MessagesSendMessageRequest{}
	_ bin.BareDecoder = &MessagesSendMessageRequest{}
)

func (s *MessagesSendMessageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.NoWebpage == false) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.Noforwards == false) {
		return false
	}
	if !(s.UpdateStickersetsOrder == false) {
		return false
	}
	if !(s.InvertMedia == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.ReplyMarkup == nil) {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}
	if !(s.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendMessageRequest) String() string {
	if s == nil {
		return "MessagesSendMessageRequest(nil)"
	}
	type Alias MessagesSendMessageRequest
	return fmt.Sprintf("MessagesSendMessageRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendMessageRequest from given interface.
func (s *MessagesSendMessageRequest) FillFrom(from interface {
	GetNoWebpage() (value bool)
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetClearDraft() (value bool)
	GetNoforwards() (value bool)
	GetUpdateStickersetsOrder() (value bool)
	GetInvertMedia() (value bool)
	GetPeer() (value InputPeerClass)
	GetReplyTo() (value InputReplyToClass, ok bool)
	GetMessage() (value string)
	GetRandomID() (value int64)
	GetReplyMarkup() (value ReplyMarkupClass, ok bool)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetScheduleDate() (value int, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	s.NoWebpage = from.GetNoWebpage()
	s.Silent = from.GetSilent()
	s.Background = from.GetBackground()
	s.ClearDraft = from.GetClearDraft()
	s.Noforwards = from.GetNoforwards()
	s.UpdateStickersetsOrder = from.GetUpdateStickersetsOrder()
	s.InvertMedia = from.GetInvertMedia()
	s.Peer = from.GetPeer()
	if val, ok := from.GetReplyTo(); ok {
		s.ReplyTo = val
	}

	s.Message = from.GetMessage()
	s.RandomID = from.GetRandomID()
	if val, ok := from.GetReplyMarkup(); ok {
		s.ReplyMarkup = val
	}

	if val, ok := from.GetEntities(); ok {
		s.Entities = val
	}

	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}

	if val, ok := from.GetSendAs(); ok {
		s.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendMessageRequest) TypeID() uint32 {
	return MessagesSendMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendMessageRequest) TypeName() string {
	return "messages.sendMessage"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendMessage",
		ID:   MessagesSendMessageRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NoWebpage",
			SchemaName: "no_webpage",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "ClearDraft",
			SchemaName: "clear_draft",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "UpdateStickersetsOrder",
			SchemaName: "update_stickersets_order",
			Null:       !s.Flags.Has(15),
		},
		{
			Name:       "InvertMedia",
			SchemaName: "invert_media",
			Null:       !s.Flags.Has(16),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !s.Flags.Has(13),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendMessageRequest) SetFlags() {
	if !(s.NoWebpage == false) {
		s.Flags.Set(1)
	}
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.Noforwards == false) {
		s.Flags.Set(14)
	}
	if !(s.UpdateStickersetsOrder == false) {
		s.Flags.Set(15)
	}
	if !(s.InvertMedia == false) {
		s.Flags.Set(16)
	}
	if !(s.ReplyTo == nil) {
		s.Flags.Set(0)
	}
	if !(s.ReplyMarkup == nil) {
		s.Flags.Set(2)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(3)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if !(s.SendAs == nil) {
		s.Flags.Set(13)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendMessageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMessage#280d096f as nil")
	}
	b.PutID(MessagesSendMessageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendMessageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMessage#280d096f as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		if s.ReplyTo == nil {
			return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field reply_to is nil")
		}
		if err := s.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field reply_to: %w", err)
		}
	}
	b.PutString(s.Message)
	b.PutLong(s.RandomID)
	if s.Flags.Has(2) {
		if s.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field reply_markup is nil")
		}
		if err := s.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field reply_markup: %w", err)
		}
	}
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	if s.Flags.Has(13) {
		if s.SendAs == nil {
			return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field send_as is nil")
		}
		if err := s.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMessage#280d096f: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendMessageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMessage#280d096f to nil")
	}
	if err := b.ConsumeID(MessagesSendMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMessage#280d096f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendMessageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMessage#280d096f to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field flags: %w", err)
		}
	}
	s.NoWebpage = s.Flags.Has(1)
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.Noforwards = s.Flags.Has(14)
	s.UpdateStickersetsOrder = s.Flags.Has(15)
	s.InvertMedia = s.Flags.Has(16)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field message: %w", err)
		}
		s.Message = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field reply_markup: %w", err)
		}
		s.ReplyMarkup = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	if s.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMessage#280d096f: field send_as: %w", err)
		}
		s.SendAs = value
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (s *MessagesSendMessageRequest) SetNoWebpage(value bool) {
	if value {
		s.Flags.Set(1)
		s.NoWebpage = true
	} else {
		s.Flags.Unset(1)
		s.NoWebpage = false
	}
}

// GetNoWebpage returns value of NoWebpage conditional field.
func (s *MessagesSendMessageRequest) GetNoWebpage() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMessageRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendMessageRequest) GetSilent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMessageRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendMessageRequest) GetBackground() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMessageRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendMessageRequest) GetClearDraft() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(7)
}

// SetNoforwards sets value of Noforwards conditional field.
func (s *MessagesSendMessageRequest) SetNoforwards(value bool) {
	if value {
		s.Flags.Set(14)
		s.Noforwards = true
	} else {
		s.Flags.Unset(14)
		s.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (s *MessagesSendMessageRequest) GetNoforwards() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(14)
}

// SetUpdateStickersetsOrder sets value of UpdateStickersetsOrder conditional field.
func (s *MessagesSendMessageRequest) SetUpdateStickersetsOrder(value bool) {
	if value {
		s.Flags.Set(15)
		s.UpdateStickersetsOrder = true
	} else {
		s.Flags.Unset(15)
		s.UpdateStickersetsOrder = false
	}
}

// GetUpdateStickersetsOrder returns value of UpdateStickersetsOrder conditional field.
func (s *MessagesSendMessageRequest) GetUpdateStickersetsOrder() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(15)
}

// SetInvertMedia sets value of InvertMedia conditional field.
func (s *MessagesSendMessageRequest) SetInvertMedia(value bool) {
	if value {
		s.Flags.Set(16)
		s.InvertMedia = true
	} else {
		s.Flags.Unset(16)
		s.InvertMedia = false
	}
}

// GetInvertMedia returns value of InvertMedia conditional field.
func (s *MessagesSendMessageRequest) GetInvertMedia() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(16)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendMessageRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetReplyTo sets value of ReplyTo conditional field.
func (s *MessagesSendMessageRequest) SetReplyTo(value InputReplyToClass) {
	s.Flags.Set(0)
	s.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMessageRequest) GetReplyTo() (value InputReplyToClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyTo, true
}

// GetMessage returns value of Message field.
func (s *MessagesSendMessageRequest) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendMessageRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (s *MessagesSendMessageRequest) SetReplyMarkup(value ReplyMarkupClass) {
	s.Flags.Set(2)
	s.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMessageRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (s *MessagesSendMessageRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(3)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMessageRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Entities, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMessageRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMessageRequest) GetScheduleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// SetSendAs sets value of SendAs conditional field.
func (s *MessagesSendMessageRequest) SetSendAs(value InputPeerClass) {
	s.Flags.Set(13)
	s.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMessageRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.SendAs, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (s *MessagesSendMessageRequest) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(s.Entities), true
}

// MessagesSendMessage invokes method messages.sendMessage#280d096f returning error if any.
// Sends a message to a chat
//
// Possible errors:
//
//	400 BOT_DOMAIN_INVALID: Bot domain invalid.
//	400 BOT_INVALID: This is not a valid bot.
//	400 BUTTON_DATA_INVALID: The data of one or more of the buttons you provided is invalid.
//	400 BUTTON_TYPE_INVALID: The type of one or more of the buttons you provided is invalid.
//	400 BUTTON_URL_INVALID: Button URL invalid.
//	400 BUTTON_USER_PRIVACY_RESTRICTED: The privacy setting of the user specified in a inputKeyboardButtonUserProfile button do not allow creating such a button.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	403 CHAT_GUEST_SEND_FORBIDDEN: You join the discussion group before commenting, see here » for more info.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 CHAT_RESTRICTED: You can't send messages in this chat, you were restricted.
//	403 CHAT_SEND_PLAIN_FORBIDDEN: You can't send non-media (text) messages in this chat.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 DOCUMENT_INVALID: The specified document is invalid.
//	400 ENCRYPTION_DECLINED: The secret chat was declined.
//	400 ENTITIES_TOO_LONG: You provided too many styled message entities.
//	400 ENTITY_BOUNDS_INVALID: A specified entity offset or length is invalid, see here » for info on how to properly compute the entity offset/length.
//	400 ENTITY_MENTION_USER_INVALID: You mentioned an invalid user.
//	400 FROM_MESSAGE_BOT_DISABLED: Bots can't use fromMessage min constructors.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MESSAGE_EMPTY: The provided message is empty.
//	400 MESSAGE_TOO_LONG: The provided message is too long.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	500 MSG_WAIT_FAILED: A waiting call returned an error.
//	406 PAYMENT_UNSUPPORTED: A detailed description of the error will be received separately as described here ».
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 PINNED_DIALOGS_TOO_MUCH: Too many pinned dialogs.
//	400 POLL_OPTION_INVALID: Invalid poll option provided.
//	500 RANDOM_ID_DUPLICATE: You provided a random ID that was already used.
//	400 REPLY_MARKUP_INVALID: The provided reply markup is invalid.
//	400 REPLY_MARKUP_TOO_LONG: The specified reply_markup is too long.
//	400 REPLY_TO_INVALID: The specified reply_to field is invalid.
//	400 REPLY_TO_USER_INVALID:
//	400 SCHEDULE_BOT_NOT_ALLOWED: Bots cannot schedule messages.
//	400 SCHEDULE_DATE_TOO_LATE: You can't schedule a message this far in the future.
//	400 SCHEDULE_STATUS_PRIVATE: Can't schedule until user is online, if the user's last seen timestamp is hidden by their privacy settings.
//	400 SCHEDULE_TOO_MUCH: There are too many scheduled messages.
//	400 SEND_AS_PEER_INVALID: You can't send messages as the specified peer.
//	420 SLOWMODE_WAIT_%d: Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat.
//	400 STORY_ID_INVALID:
//	406 TOPIC_CLOSED: This topic was closed, you can't send messages to it anymore.
//	406 TOPIC_DELETED: The specified topic was deleted.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//	403 USER_IS_BLOCKED: You were blocked by this user.
//	400 USER_IS_BOT: Bots can't send messages to other bots.
//	400 WC_CONVERT_URL_INVALID: WC convert URL invalid.
//	400 YOU_BLOCKED_USER: You blocked this user.
//
// See https://core.telegram.org/method/messages.sendMessage for reference.
// Can be used by bots.
func (c *Client) MessagesSendMessage(ctx context.Context, request *MessagesSendMessageRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
