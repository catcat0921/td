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

// HelpSupportName represents TL type `help.supportName#8c05f1c9`.
// Localized name for telegram support
//
// See https://core.telegram.org/constructor/help.supportName for reference.
type HelpSupportName struct {
	// Localized name
	Name string
}

// HelpSupportNameTypeID is TL type id of HelpSupportName.
const HelpSupportNameTypeID = 0x8c05f1c9

func (s *HelpSupportName) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *HelpSupportName) String() string {
	if s == nil {
		return "HelpSupportName(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpSupportName")
	sb.WriteString("{\n")
	sb.WriteString("\tName: ")
	sb.WriteString(fmt.Sprint(s.Name))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *HelpSupportName) TypeID() uint32 {
	return HelpSupportNameTypeID
}

// Encode implements bin.Encoder.
func (s *HelpSupportName) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.supportName#8c05f1c9 as nil")
	}
	b.PutID(HelpSupportNameTypeID)
	b.PutString(s.Name)
	return nil
}

// GetName returns value of Name field.
func (s *HelpSupportName) GetName() (value string) {
	return s.Name
}

// Decode implements bin.Decoder.
func (s *HelpSupportName) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.supportName#8c05f1c9 to nil")
	}
	if err := b.ConsumeID(HelpSupportNameTypeID); err != nil {
		return fmt.Errorf("unable to decode help.supportName#8c05f1c9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.supportName#8c05f1c9: field name: %w", err)
		}
		s.Name = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpSupportName.
var (
	_ bin.Encoder = &HelpSupportName{}
	_ bin.Decoder = &HelpSupportName{}
)
