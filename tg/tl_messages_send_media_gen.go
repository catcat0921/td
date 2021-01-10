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

// MessagesSendMediaRequest represents TL type `messages.sendMedia#3491eba9`.
// Send a media
//
// See https://core.telegram.org/method/messages.sendMedia for reference.
type MessagesSendMediaRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Send message silently (no notification should be triggered)
	Silent bool
	// Send message in background
	Background bool
	// Clear the draft
	ClearDraft bool
	// Destination
	Peer InputPeerClass
	// Message ID to which this message should reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// Attached media
	Media InputMediaClass
	// Caption
	Message string
	// Random ID to avoid resending the same message
	RandomID int64
	// Reply markup for bot keyboards
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
	// Scheduled message date for scheduled messages¹
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// MessagesSendMediaRequestTypeID is TL type id of MessagesSendMediaRequest.
const MessagesSendMediaRequestTypeID = 0x3491eba9

func (s *MessagesSendMediaRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
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
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.Media == nil) {
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

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendMediaRequest) String() string {
	if s == nil {
		return "MessagesSendMediaRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSendMediaRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(s.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tReplyToMsgID: ")
		sb.WriteString(fmt.Sprint(s.ReplyToMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tMedia: ")
	sb.WriteString(fmt.Sprint(s.Media))
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(s.Message))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(s.RandomID))
	sb.WriteString(",\n")
	if s.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(fmt.Sprint(s.ReplyMarkup))
		sb.WriteString(",\n")
	}
	if s.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range s.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if s.Flags.Has(10) {
		sb.WriteString("\tScheduleDate: ")
		sb.WriteString(fmt.Sprint(s.ScheduleDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSendMediaRequest) TypeID() uint32 {
	return MessagesSendMediaRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSendMediaRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMedia#3491eba9 as nil")
	}
	b.PutID(MessagesSendMediaRequestTypeID)
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.ReplyToMsgID == 0) {
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
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	if s.Media == nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field media is nil")
	}
	if err := s.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field media: %w", err)
	}
	b.PutString(s.Message)
	b.PutLong(s.RandomID)
	if s.Flags.Has(2) {
		if s.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field reply_markup is nil")
		}
		if err := s.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field reply_markup: %w", err)
		}
	}
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMediaRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendMediaRequest) GetSilent() (value bool) {
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMediaRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendMediaRequest) GetBackground() (value bool) {
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMediaRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendMediaRequest) GetClearDraft() (value bool) {
	return s.Flags.Has(7)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendMediaRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendMediaRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// GetMedia returns value of Media field.
func (s *MessagesSendMediaRequest) GetMedia() (value InputMediaClass) {
	return s.Media
}

// GetMessage returns value of Message field.
func (s *MessagesSendMediaRequest) GetMessage() (value string) {
	return s.Message
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendMediaRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (s *MessagesSendMediaRequest) SetReplyMarkup(value ReplyMarkupClass) {
	s.Flags.Set(2)
	s.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (s *MessagesSendMediaRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(3)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Entities, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMediaRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetScheduleDate() (value int, ok bool) {
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (s *MessagesSendMediaRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMedia#3491eba9 to nil")
	}
	if err := b.ConsumeID(MessagesSendMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field media: %w", err)
		}
		s.Media = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field message: %w", err)
		}
		s.Message = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field reply_markup: %w", err)
		}
		s.ReplyMarkup = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendMediaRequest.
var (
	_ bin.Encoder = &MessagesSendMediaRequest{}
	_ bin.Decoder = &MessagesSendMediaRequest{}
)

// MessagesSendMedia invokes method messages.sendMedia#3491eba9 returning error if any.
// Send a media
//
// Possible errors:
//  400 BROADCAST_PUBLIC_VOTERS_FORBIDDEN: You can't forward polls with public voters
//  400 BUTTON_DATA_INVALID: The data of one or more of the buttons you provided is invalid
//  400 BUTTON_TYPE_INVALID: The type of one or more of the buttons you provided is invalid
//  400 BUTTON_URL_INVALID: Button URL invalid
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_RESTRICTED: You can't send messages in this chat, you were restricted
//  403 CHAT_SEND_GIFS_FORBIDDEN: You can't send gifs in this chat
//  403 CHAT_SEND_MEDIA_FORBIDDEN: You can't send media in this chat
//  403 CHAT_SEND_STICKERS_FORBIDDEN: You can't send stickers in this chat.
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 EXTERNAL_URL_INVALID: External URL invalid
//  400 FILE_PARTS_INVALID: The number of file parts is invalid
//  400 FILE_PART_LENGTH_INVALID: The length of a file part is invalid
//  400 IMAGE_PROCESS_FAILED: Failure while processing image
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MD5_CHECKSUM_INVALID: The MD5 checksums do not match
//  400 MEDIA_CAPTION_TOO_LONG: The caption is too long
//  400 MEDIA_EMPTY: The provided media object is invalid
//  400 MEDIA_INVALID: Media invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 PHOTO_EXT_INVALID: The extension of the photo is invalid
//  400 PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
//  400 PHOTO_SAVE_FILE_INVALID: Internal issues, try again later
//  400 POLL_ANSWERS_INVALID: Invalid poll answers were provided
//  400 POLL_OPTION_DUPLICATE: Duplicate poll options provided
//  400 POLL_OPTION_INVALID: Invalid poll option provided
//  400 QUIZ_CORRECT_ANSWERS_EMPTY: No correct quiz answer was specified
//  400 QUIZ_CORRECT_ANSWER_INVALID: An invalid value was provided to the correct_answers field
//  400 REPLY_MARKUP_BUY_EMPTY: Reply markup for buy button empty
//  400 REPLY_MARKUP_INVALID: The provided reply markup is invalid
//  400 SCHEDULE_TOO_MUCH: There are too many scheduled messages
//  420 SLOWMODE_WAIT_X: Slowmode is enabled in this chat: you must wait for the specified number of seconds before sending another message to the chat.
//  400 TTL_MEDIA_INVALID: Invalid media Time To Live was provided
//  400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels
//  400 USER_IS_BLOCKED: You were blocked by this user
//  400 USER_IS_BOT: Bots can't send messages to other bots
//  400 WEBPAGE_CURL_FAILED: Failure while fetching the webpage with cURL
//  400 WEBPAGE_MEDIA_EMPTY: Webpage media empty
//  400 YOU_BLOCKED_USER: You blocked this user
//
// See https://core.telegram.org/method/messages.sendMedia for reference.
// Can be used by bots.
func (c *Client) MessagesSendMedia(ctx context.Context, request *MessagesSendMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
