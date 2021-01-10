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

// MessagesGetEmojiURLRequest represents TL type `messages.getEmojiURL#d5b10c26`.
// Returns an HTTP URL which can be used to automatically log in into translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
//
// See https://core.telegram.org/method/messages.getEmojiURL for reference.
type MessagesGetEmojiURLRequest struct {
	// Language code for which the emoji replacements will be suggested
	LangCode string
}

// MessagesGetEmojiURLRequestTypeID is TL type id of MessagesGetEmojiURLRequest.
const MessagesGetEmojiURLRequestTypeID = 0xd5b10c26

func (g *MessagesGetEmojiURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiURLRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiURLRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetEmojiURLRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tLangCode: ")
	sb.WriteString(fmt.Sprint(g.LangCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetEmojiURLRequest) TypeID() uint32 {
	return MessagesGetEmojiURLRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiURL#d5b10c26 as nil")
	}
	b.PutID(MessagesGetEmojiURLRequestTypeID)
	b.PutString(g.LangCode)
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *MessagesGetEmojiURLRequest) GetLangCode() (value string) {
	return g.LangCode
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiURL#d5b10c26 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiURL#d5b10c26: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiURL#d5b10c26: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetEmojiURLRequest.
var (
	_ bin.Encoder = &MessagesGetEmojiURLRequest{}
	_ bin.Decoder = &MessagesGetEmojiURLRequest{}
)

// MessagesGetEmojiURL invokes method messages.getEmojiURL#d5b10c26 returning error if any.
// Returns an HTTP URL which can be used to automatically log in into translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
//
// See https://core.telegram.org/method/messages.getEmojiURL for reference.
func (c *Client) MessagesGetEmojiURL(ctx context.Context, langcode string) (*EmojiURL, error) {
	var result EmojiURL

	request := &MessagesGetEmojiURLRequest{
		LangCode: langcode,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
