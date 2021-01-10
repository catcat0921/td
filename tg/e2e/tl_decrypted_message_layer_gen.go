// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// DecryptedMessageLayer represents TL type `decryptedMessageLayer#1be31789`.
//
// See https://core.telegram.org/constructor/decryptedMessageLayer for reference.
type DecryptedMessageLayer struct {
	// RandomBytes field of DecryptedMessageLayer.
	RandomBytes []byte
	// Layer field of DecryptedMessageLayer.
	Layer int
	// InSeqNo field of DecryptedMessageLayer.
	InSeqNo int
	// OutSeqNo field of DecryptedMessageLayer.
	OutSeqNo int
	// Message field of DecryptedMessageLayer.
	Message DecryptedMessageClass
}

// DecryptedMessageLayerTypeID is TL type id of DecryptedMessageLayer.
const DecryptedMessageLayerTypeID = 0x1be31789

func (d *DecryptedMessageLayer) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.RandomBytes == nil) {
		return false
	}
	if !(d.Layer == 0) {
		return false
	}
	if !(d.InSeqNo == 0) {
		return false
	}
	if !(d.OutSeqNo == 0) {
		return false
	}
	if !(d.Message == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DecryptedMessageLayer) String() string {
	if d == nil {
		return "DecryptedMessageLayer(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DecryptedMessageLayer")
	sb.WriteString("{\n")
	sb.WriteString("\tRandomBytes: ")
	sb.WriteString(fmt.Sprint(d.RandomBytes))
	sb.WriteString(",\n")
	sb.WriteString("\tLayer: ")
	sb.WriteString(fmt.Sprint(d.Layer))
	sb.WriteString(",\n")
	sb.WriteString("\tInSeqNo: ")
	sb.WriteString(fmt.Sprint(d.InSeqNo))
	sb.WriteString(",\n")
	sb.WriteString("\tOutSeqNo: ")
	sb.WriteString(fmt.Sprint(d.OutSeqNo))
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(d.Message))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DecryptedMessageLayer) TypeID() uint32 {
	return DecryptedMessageLayerTypeID
}

// Encode implements bin.Encoder.
func (d *DecryptedMessageLayer) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode decryptedMessageLayer#1be31789 as nil")
	}
	b.PutID(DecryptedMessageLayerTypeID)
	b.PutBytes(d.RandomBytes)
	b.PutInt(d.Layer)
	b.PutInt(d.InSeqNo)
	b.PutInt(d.OutSeqNo)
	if d.Message == nil {
		return fmt.Errorf("unable to encode decryptedMessageLayer#1be31789: field message is nil")
	}
	if err := d.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode decryptedMessageLayer#1be31789: field message: %w", err)
	}
	return nil
}

// GetRandomBytes returns value of RandomBytes field.
func (d *DecryptedMessageLayer) GetRandomBytes() (value []byte) {
	return d.RandomBytes
}

// GetLayer returns value of Layer field.
func (d *DecryptedMessageLayer) GetLayer() (value int) {
	return d.Layer
}

// GetInSeqNo returns value of InSeqNo field.
func (d *DecryptedMessageLayer) GetInSeqNo() (value int) {
	return d.InSeqNo
}

// GetOutSeqNo returns value of OutSeqNo field.
func (d *DecryptedMessageLayer) GetOutSeqNo() (value int) {
	return d.OutSeqNo
}

// GetMessage returns value of Message field.
func (d *DecryptedMessageLayer) GetMessage() (value DecryptedMessageClass) {
	return d.Message
}

// Decode implements bin.Decoder.
func (d *DecryptedMessageLayer) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode decryptedMessageLayer#1be31789 to nil")
	}
	if err := b.ConsumeID(DecryptedMessageLayerTypeID); err != nil {
		return fmt.Errorf("unable to decode decryptedMessageLayer#1be31789: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode decryptedMessageLayer#1be31789: field random_bytes: %w", err)
		}
		d.RandomBytes = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode decryptedMessageLayer#1be31789: field layer: %w", err)
		}
		d.Layer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode decryptedMessageLayer#1be31789: field in_seq_no: %w", err)
		}
		d.InSeqNo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode decryptedMessageLayer#1be31789: field out_seq_no: %w", err)
		}
		d.OutSeqNo = value
	}
	{
		value, err := DecodeDecryptedMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode decryptedMessageLayer#1be31789: field message: %w", err)
		}
		d.Message = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DecryptedMessageLayer.
var (
	_ bin.Encoder = &DecryptedMessageLayer{}
	_ bin.Decoder = &DecryptedMessageLayer{}
)
