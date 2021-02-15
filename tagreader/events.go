// This file is modified from the original at https://github.com/gvalkov/golang-evdev/blob/master/events.go
// License for this file specifically:
//
// Copyright (c) 2016 Georgi Valkov. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//  1. Redistributions of source code must retain the above copyright
//     notice, this list of conditions and the following disclaimer.
//
//  2. Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in
//     the documentation and/or other materials provided with the
//     distribution.
//
//  3. Neither the name of author nor the names of its contributors may
//     be used to endorse or promote products derived from this software
//     without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL GEORGI VALKOV BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package tagreader

import (
	"fmt"
	"syscall"
	"unsafe"
)

type InputEvent struct {
	Time  syscall.Timeval // time in seconds since epoch at which event occurred
	Type  uint16          // event type - one of ecodes.EV_*
	Code  uint16          // event code related to the event type
	Value int32           // event value related to the event type
}

// Get a useful description for an input event. Example:
//   event at 1347905437.435795, code 01, type 02, val 02
func (ev *InputEvent) String() string {
	return fmt.Sprintf("event at %d.%d, code %02d, type %02d, val %02d",
		ev.Time.Sec, ev.Time.Usec, ev.Code, ev.Type, ev.Value)
}

var eventSize = int(unsafe.Sizeof(InputEvent{}))

type KeyEventState = int32

const (
	KeyUp KeyEventState = iota
	KeyDown
	KeyHold
)
