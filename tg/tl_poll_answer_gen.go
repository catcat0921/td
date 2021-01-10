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

// PollAnswer represents TL type `pollAnswer#6ca9c2e9`.
// A possible answer of a poll
//
// See https://core.telegram.org/constructor/pollAnswer for reference.
type PollAnswer struct {
	// Textual representation of the answer
	Text string
	// The param that has to be passed to messages.sendVote¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendVote
	Option []byte
}

// PollAnswerTypeID is TL type id of PollAnswer.
const PollAnswerTypeID = 0x6ca9c2e9

func (p *PollAnswer) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == "") {
		return false
	}
	if !(p.Option == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollAnswer) String() string {
	if p == nil {
		return "PollAnswer(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PollAnswer")
	sb.WriteString("{\n")
	sb.WriteString("\tText: ")
	sb.WriteString(fmt.Sprint(p.Text))
	sb.WriteString(",\n")
	sb.WriteString("\tOption: ")
	sb.WriteString(fmt.Sprint(p.Option))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PollAnswer) TypeID() uint32 {
	return PollAnswerTypeID
}

// Encode implements bin.Encoder.
func (p *PollAnswer) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollAnswer#6ca9c2e9 as nil")
	}
	b.PutID(PollAnswerTypeID)
	b.PutString(p.Text)
	b.PutBytes(p.Option)
	return nil
}

// GetText returns value of Text field.
func (p *PollAnswer) GetText() (value string) {
	return p.Text
}

// GetOption returns value of Option field.
func (p *PollAnswer) GetOption() (value []byte) {
	return p.Option
}

// Decode implements bin.Decoder.
func (p *PollAnswer) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollAnswer#6ca9c2e9 to nil")
	}
	if err := b.ConsumeID(PollAnswerTypeID); err != nil {
		return fmt.Errorf("unable to decode pollAnswer#6ca9c2e9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswer#6ca9c2e9: field text: %w", err)
		}
		p.Text = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswer#6ca9c2e9: field option: %w", err)
		}
		p.Option = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PollAnswer.
var (
	_ bin.Encoder = &PollAnswer{}
	_ bin.Decoder = &PollAnswer{}
)
