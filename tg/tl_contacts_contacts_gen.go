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

// ContactsContactsNotModified represents TL type `contacts.contactsNotModified#b74ba9d2`.
// Contact list on the server is the same as the list on the client.
//
// See https://core.telegram.org/constructor/contacts.contactsNotModified for reference.
type ContactsContactsNotModified struct {
}

// ContactsContactsNotModifiedTypeID is TL type id of ContactsContactsNotModified.
const ContactsContactsNotModifiedTypeID = 0xb74ba9d2

// construct implements constructor of ContactsContactsClass.
func (c ContactsContactsNotModified) construct() ContactsContactsClass { return &c }

// Ensuring interfaces in compile-time for ContactsContactsNotModified.
var (
	_ bin.Encoder     = &ContactsContactsNotModified{}
	_ bin.Decoder     = &ContactsContactsNotModified{}
	_ bin.BareEncoder = &ContactsContactsNotModified{}
	_ bin.BareDecoder = &ContactsContactsNotModified{}

	_ ContactsContactsClass = &ContactsContactsNotModified{}
)

func (c *ContactsContactsNotModified) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ContactsContactsNotModified) String() string {
	if c == nil {
		return "ContactsContactsNotModified(nil)"
	}
	type Alias ContactsContactsNotModified
	return fmt.Sprintf("ContactsContactsNotModified%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsContactsNotModified) TypeID() uint32 {
	return ContactsContactsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsContactsNotModified) TypeName() string {
	return "contacts.contactsNotModified"
}

// TypeInfo returns info about TL type.
func (c *ContactsContactsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.contactsNotModified",
		ID:   ContactsContactsNotModifiedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ContactsContactsNotModified) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contacts.contactsNotModified#b74ba9d2 as nil")
	}
	b.PutID(ContactsContactsNotModifiedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ContactsContactsNotModified) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contacts.contactsNotModified#b74ba9d2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ContactsContactsNotModified) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contacts.contactsNotModified#b74ba9d2 to nil")
	}
	if err := b.ConsumeID(ContactsContactsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.contactsNotModified#b74ba9d2: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ContactsContactsNotModified) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contacts.contactsNotModified#b74ba9d2 to nil")
	}
	return nil
}

// ContactsContacts represents TL type `contacts.contacts#eae87e42`.
// The current user's contact list and info on users.
//
// See https://core.telegram.org/constructor/contacts.contacts for reference.
type ContactsContacts struct {
	// Contact list
	Contacts []Contact
	// Number of contacts that were saved successfully
	SavedCount int
	// User list
	Users []UserClass
}

// ContactsContactsTypeID is TL type id of ContactsContacts.
const ContactsContactsTypeID = 0xeae87e42

// construct implements constructor of ContactsContactsClass.
func (c ContactsContacts) construct() ContactsContactsClass { return &c }

// Ensuring interfaces in compile-time for ContactsContacts.
var (
	_ bin.Encoder     = &ContactsContacts{}
	_ bin.Decoder     = &ContactsContacts{}
	_ bin.BareEncoder = &ContactsContacts{}
	_ bin.BareDecoder = &ContactsContacts{}

	_ ContactsContactsClass = &ContactsContacts{}
)

func (c *ContactsContacts) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Contacts == nil) {
		return false
	}
	if !(c.SavedCount == 0) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ContactsContacts) String() string {
	if c == nil {
		return "ContactsContacts(nil)"
	}
	type Alias ContactsContacts
	return fmt.Sprintf("ContactsContacts%+v", Alias(*c))
}

// FillFrom fills ContactsContacts from given interface.
func (c *ContactsContacts) FillFrom(from interface {
	GetContacts() (value []Contact)
	GetSavedCount() (value int)
	GetUsers() (value []UserClass)
}) {
	c.Contacts = from.GetContacts()
	c.SavedCount = from.GetSavedCount()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsContacts) TypeID() uint32 {
	return ContactsContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsContacts) TypeName() string {
	return "contacts.contacts"
}

// TypeInfo returns info about TL type.
func (c *ContactsContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.contacts",
		ID:   ContactsContactsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Contacts",
			SchemaName: "contacts",
		},
		{
			Name:       "SavedCount",
			SchemaName: "saved_count",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ContactsContacts) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contacts.contacts#eae87e42 as nil")
	}
	b.PutID(ContactsContactsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ContactsContacts) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contacts.contacts#eae87e42 as nil")
	}
	b.PutVectorHeader(len(c.Contacts))
	for idx, v := range c.Contacts {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.contacts#eae87e42: field contacts element with index %d: %w", idx, err)
		}
	}
	b.PutInt(c.SavedCount)
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.contacts#eae87e42: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.contacts#eae87e42: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ContactsContacts) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contacts.contacts#eae87e42 to nil")
	}
	if err := b.ConsumeID(ContactsContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.contacts#eae87e42: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ContactsContacts) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contacts.contacts#eae87e42 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field contacts: %w", err)
		}

		if headerLen > 0 {
			c.Contacts = make([]Contact, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Contact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field contacts: %w", err)
			}
			c.Contacts = append(c.Contacts, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field saved_count: %w", err)
		}
		c.SavedCount = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.contacts#eae87e42: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetContacts returns value of Contacts field.
func (c *ContactsContacts) GetContacts() (value []Contact) {
	if c == nil {
		return
	}
	return c.Contacts
}

// GetSavedCount returns value of SavedCount field.
func (c *ContactsContacts) GetSavedCount() (value int) {
	if c == nil {
		return
	}
	return c.SavedCount
}

// GetUsers returns value of Users field.
func (c *ContactsContacts) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *ContactsContacts) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}

// ContactsContactsClassName is schema name of ContactsContactsClass.
const ContactsContactsClassName = "contacts.Contacts"

// ContactsContactsClass represents contacts.Contacts generic type.
//
// See https://core.telegram.org/type/contacts.Contacts for reference.
//
// Constructors:
//   - [ContactsContactsNotModified]
//   - [ContactsContacts]
//
// Example:
//
//	g, err := tg.DecodeContactsContacts(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ContactsContactsNotModified: // contacts.contactsNotModified#b74ba9d2
//	case *tg.ContactsContacts: // contacts.contacts#eae87e42
//	default: panic(v)
//	}
type ContactsContactsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ContactsContactsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map ContactsContactsClass to ContactsContacts.
	AsModified() (*ContactsContacts, bool)
}

// AsModified tries to map ContactsContactsNotModified to ContactsContacts.
func (c *ContactsContactsNotModified) AsModified() (*ContactsContacts, bool) {
	return nil, false
}

// AsModified tries to map ContactsContacts to ContactsContacts.
func (c *ContactsContacts) AsModified() (*ContactsContacts, bool) {
	return c, true
}

// DecodeContactsContacts implements binary de-serialization for ContactsContactsClass.
func DecodeContactsContacts(buf *bin.Buffer) (ContactsContactsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ContactsContactsNotModifiedTypeID:
		// Decoding contacts.contactsNotModified#b74ba9d2.
		v := ContactsContactsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsContactsClass: %w", err)
		}
		return &v, nil
	case ContactsContactsTypeID:
		// Decoding contacts.contacts#eae87e42.
		v := ContactsContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ContactsContactsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ContactsContactsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ContactsContacts boxes the ContactsContactsClass providing a helper.
type ContactsContactsBox struct {
	Contacts ContactsContactsClass
}

// Decode implements bin.Decoder for ContactsContactsBox.
func (b *ContactsContactsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ContactsContactsBox to nil")
	}
	v, err := DecodeContactsContacts(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Contacts = v
	return nil
}

// Encode implements bin.Encode for ContactsContactsBox.
func (b *ContactsContactsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Contacts == nil {
		return fmt.Errorf("unable to encode ContactsContactsClass as nil")
	}
	return b.Contacts.Encode(buf)
}
