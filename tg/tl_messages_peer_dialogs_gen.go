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

// MessagesPeerDialogs represents TL type `messages.peerDialogs#3371c354`.
// Dialog info of multiple peers
//
// See https://core.telegram.org/constructor/messages.peerDialogs for reference.
type MessagesPeerDialogs struct {
	// Dialog info
	Dialogs []DialogClass
	// Messages mentioned in dialog info
	Messages []MessageClass
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
	// Current update state of dialog¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	State UpdatesState
}

// MessagesPeerDialogsTypeID is TL type id of MessagesPeerDialogs.
const MessagesPeerDialogsTypeID = 0x3371c354

func (p *MessagesPeerDialogs) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Dialogs == nil) {
		return false
	}
	if !(p.Messages == nil) {
		return false
	}
	if !(p.Chats == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}
	if !(p.State.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *MessagesPeerDialogs) String() string {
	if p == nil {
		return "MessagesPeerDialogs(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesPeerDialogs")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range p.Dialogs {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range p.Messages {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range p.Chats {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range p.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tState: ")
	sb.WriteString(fmt.Sprint(p.State))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *MessagesPeerDialogs) TypeID() uint32 {
	return MessagesPeerDialogsTypeID
}

// Encode implements bin.Encoder.
func (p *MessagesPeerDialogs) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode messages.peerDialogs#3371c354 as nil")
	}
	b.PutID(MessagesPeerDialogsTypeID)
	b.PutVectorHeader(len(p.Dialogs))
	for idx, v := range p.Dialogs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field dialogs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field dialogs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Messages))
	for idx, v := range p.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field users element with index %d: %w", idx, err)
		}
	}
	if err := p.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.peerDialogs#3371c354: field state: %w", err)
	}
	return nil
}

// GetDialogs returns value of Dialogs field.
func (p *MessagesPeerDialogs) GetDialogs() (value []DialogClass) {
	return p.Dialogs
}

// GetMessages returns value of Messages field.
func (p *MessagesPeerDialogs) GetMessages() (value []MessageClass) {
	return p.Messages
}

// GetChats returns value of Chats field.
func (p *MessagesPeerDialogs) GetChats() (value []ChatClass) {
	return p.Chats
}

// GetUsers returns value of Users field.
func (p *MessagesPeerDialogs) GetUsers() (value []UserClass) {
	return p.Users
}

// GetState returns value of State field.
func (p *MessagesPeerDialogs) GetState() (value UpdatesState) {
	return p.State
}

// Decode implements bin.Decoder.
func (p *MessagesPeerDialogs) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode messages.peerDialogs#3371c354 to nil")
	}
	if err := b.ConsumeID(MessagesPeerDialogsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field dialogs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialog(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field dialogs: %w", err)
			}
			p.Dialogs = append(p.Dialogs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field messages: %w", err)
			}
			p.Messages = append(p.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	{
		if err := p.State.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.peerDialogs#3371c354: field state: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesPeerDialogs.
var (
	_ bin.Encoder = &MessagesPeerDialogs{}
	_ bin.Decoder = &MessagesPeerDialogs{}
)
