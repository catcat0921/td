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

// PollAnswerVoters represents TL type `pollAnswerVoters#3b6ddad2`.
// A poll answer, and how users voted on it
//
// See https://core.telegram.org/constructor/pollAnswerVoters for reference.
type PollAnswerVoters struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether we have chosen this answer
	Chosen bool
	// For quizes, whether the option we have chosen is correct
	Correct bool
	// The param that has to be passed to messages.sendVote¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendVote
	Option []byte
	// How many users voted for this option
	Voters int
}

// PollAnswerVotersTypeID is TL type id of PollAnswerVoters.
const PollAnswerVotersTypeID = 0x3b6ddad2

func (p *PollAnswerVoters) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Chosen == false) {
		return false
	}
	if !(p.Correct == false) {
		return false
	}
	if !(p.Option == nil) {
		return false
	}
	if !(p.Voters == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollAnswerVoters) String() string {
	if p == nil {
		return "PollAnswerVoters(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PollAnswerVoters")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(p.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tOption: ")
	sb.WriteString(fmt.Sprint(p.Option))
	sb.WriteString(",\n")
	sb.WriteString("\tVoters: ")
	sb.WriteString(fmt.Sprint(p.Voters))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PollAnswerVoters) TypeID() uint32 {
	return PollAnswerVotersTypeID
}

// Encode implements bin.Encoder.
func (p *PollAnswerVoters) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollAnswerVoters#3b6ddad2 as nil")
	}
	b.PutID(PollAnswerVotersTypeID)
	if !(p.Chosen == false) {
		p.Flags.Set(0)
	}
	if !(p.Correct == false) {
		p.Flags.Set(1)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pollAnswerVoters#3b6ddad2: field flags: %w", err)
	}
	b.PutBytes(p.Option)
	b.PutInt(p.Voters)
	return nil
}

// SetChosen sets value of Chosen conditional field.
func (p *PollAnswerVoters) SetChosen(value bool) {
	if value {
		p.Flags.Set(0)
		p.Chosen = true
	} else {
		p.Flags.Unset(0)
		p.Chosen = false
	}
}

// GetChosen returns value of Chosen conditional field.
func (p *PollAnswerVoters) GetChosen() (value bool) {
	return p.Flags.Has(0)
}

// SetCorrect sets value of Correct conditional field.
func (p *PollAnswerVoters) SetCorrect(value bool) {
	if value {
		p.Flags.Set(1)
		p.Correct = true
	} else {
		p.Flags.Unset(1)
		p.Correct = false
	}
}

// GetCorrect returns value of Correct conditional field.
func (p *PollAnswerVoters) GetCorrect() (value bool) {
	return p.Flags.Has(1)
}

// GetOption returns value of Option field.
func (p *PollAnswerVoters) GetOption() (value []byte) {
	return p.Option
}

// GetVoters returns value of Voters field.
func (p *PollAnswerVoters) GetVoters() (value int) {
	return p.Voters
}

// Decode implements bin.Decoder.
func (p *PollAnswerVoters) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollAnswerVoters#3b6ddad2 to nil")
	}
	if err := b.ConsumeID(PollAnswerVotersTypeID); err != nil {
		return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: field flags: %w", err)
		}
	}
	p.Chosen = p.Flags.Has(0)
	p.Correct = p.Flags.Has(1)
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: field option: %w", err)
		}
		p.Option = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: field voters: %w", err)
		}
		p.Voters = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PollAnswerVoters.
var (
	_ bin.Encoder = &PollAnswerVoters{}
	_ bin.Decoder = &PollAnswerVoters{}
)
