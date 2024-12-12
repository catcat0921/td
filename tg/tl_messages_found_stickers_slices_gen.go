//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// MessagesFoundStickersClassArray is adapter for slice of MessagesFoundStickersClass.
type MessagesFoundStickersClassArray []MessagesFoundStickersClass

// Sort sorts slice of MessagesFoundStickersClass.
func (s MessagesFoundStickersClassArray) Sort(less func(a, b MessagesFoundStickersClass) bool) MessagesFoundStickersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickersClass.
func (s MessagesFoundStickersClassArray) SortStable(less func(a, b MessagesFoundStickersClass) bool) MessagesFoundStickersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickersClass.
func (s MessagesFoundStickersClassArray) Retain(keep func(x MessagesFoundStickersClass) bool) MessagesFoundStickersClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFoundStickersClassArray) First() (v MessagesFoundStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickersClassArray) Last() (v MessagesFoundStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickersClassArray) PopFirst() (v MessagesFoundStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickersClassArray) Pop() (v MessagesFoundStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesFoundStickersNotModified returns copy with only MessagesFoundStickersNotModified constructors.
func (s MessagesFoundStickersClassArray) AsMessagesFoundStickersNotModified() (to MessagesFoundStickersNotModifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFoundStickersNotModified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesFoundStickers returns copy with only MessagesFoundStickers constructors.
func (s MessagesFoundStickersClassArray) AsMessagesFoundStickers() (to MessagesFoundStickersArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFoundStickers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesFoundStickersClassArray) AppendOnlyModified(to []*MessagesFoundStickers) []*MessagesFoundStickers {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesFoundStickersClassArray) AsModified() (to []*MessagesFoundStickers) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesFoundStickersClassArray) FirstAsModified() (v *MessagesFoundStickers, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesFoundStickersClassArray) LastAsModified() (v *MessagesFoundStickers, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesFoundStickersClassArray) PopFirstAsModified() (v *MessagesFoundStickers, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesFoundStickersClassArray) PopAsModified() (v *MessagesFoundStickers, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesFoundStickersNotModifiedArray is adapter for slice of MessagesFoundStickersNotModified.
type MessagesFoundStickersNotModifiedArray []MessagesFoundStickersNotModified

// Sort sorts slice of MessagesFoundStickersNotModified.
func (s MessagesFoundStickersNotModifiedArray) Sort(less func(a, b MessagesFoundStickersNotModified) bool) MessagesFoundStickersNotModifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickersNotModified.
func (s MessagesFoundStickersNotModifiedArray) SortStable(less func(a, b MessagesFoundStickersNotModified) bool) MessagesFoundStickersNotModifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickersNotModified.
func (s MessagesFoundStickersNotModifiedArray) Retain(keep func(x MessagesFoundStickersNotModified) bool) MessagesFoundStickersNotModifiedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFoundStickersNotModifiedArray) First() (v MessagesFoundStickersNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickersNotModifiedArray) Last() (v MessagesFoundStickersNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickersNotModifiedArray) PopFirst() (v MessagesFoundStickersNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickersNotModified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickersNotModifiedArray) Pop() (v MessagesFoundStickersNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessagesFoundStickersArray is adapter for slice of MessagesFoundStickers.
type MessagesFoundStickersArray []MessagesFoundStickers

// Sort sorts slice of MessagesFoundStickers.
func (s MessagesFoundStickersArray) Sort(less func(a, b MessagesFoundStickers) bool) MessagesFoundStickersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickers.
func (s MessagesFoundStickersArray) SortStable(less func(a, b MessagesFoundStickers) bool) MessagesFoundStickersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickers.
func (s MessagesFoundStickersArray) Retain(keep func(x MessagesFoundStickers) bool) MessagesFoundStickersArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFoundStickersArray) First() (v MessagesFoundStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickersArray) Last() (v MessagesFoundStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickersArray) PopFirst() (v MessagesFoundStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickersArray) Pop() (v MessagesFoundStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}