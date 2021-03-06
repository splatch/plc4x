//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package utils

import (
	"container/list"
	"fmt"
	"math/big"
)

type WriteBufferBoxBased interface {
	WriteBuffer
	GetBox() AsciiBox
}

func NewBoxedWriteBuffer() WriteBufferBoxBased {
	return &boxedWriteBuffer{
		List:         list.New(),
		desiredWidth: 120,
		currentWidth: 118,
	}
}

// WithAdditionalStringRepresentation can be used by e.g. enums to supply an additional string representation
func WithAdditionalStringRepresentation(stringRepresentation string) WithWriterArgs {
	return withAdditionalStringRepresentation{stringRepresentation: stringRepresentation}
}

///////////////////////////////////////
///////////////////////////////////////
//
// Internal section
//

type withAdditionalStringRepresentation struct {
	writerArg
	stringRepresentation string
}

type boxedWriteBuffer struct {
	*list.List
	desiredWidth int
	currentWidth int
}

//
// Internal section
//
///////////////////////////////////////
///////////////////////////////////////

func (b *boxedWriteBuffer) GetBox() AsciiBox {
	back := b.Back()
	return back.Value.(AsciiBox)
}

func (b *boxedWriteBuffer) PushContext(_ string, _ ...WithWriterArgs) error {
	b.currentWidth -= boxLineOverheat
	b.PushBack(make([]AsciiBox, 0))
	return nil
}

func (b *boxedWriteBuffer) WriteBit(logicalName string, value bool, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	asInt := 0
	if value {
		asInt = 1
	}
	b.PushBack(BoxString(logicalName, fmt.Sprintf("b%d %t%s", asInt, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteUint8(logicalName string, bitLength uint8, value uint8, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteUint16(logicalName string, bitLength uint8, value uint16, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteUint32(logicalName string, bitLength uint8, value uint32, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteUint64(logicalName string, bitLength uint8, value uint64, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteInt8(logicalName string, bitLength uint8, value int8, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteInt16(logicalName string, bitLength uint8, value int16, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteInt32(logicalName string, bitLength uint8, value int32, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteInt64(logicalName string, bitLength uint8, value int64, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteBigInt(logicalName string, bitLength uint8, value *big.Int, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %d%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteFloat32(logicalName string, bitLength uint8, value float32, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %f%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteFloat64(logicalName string, bitLength uint8, value float64, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %f%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteBigFloat(logicalName string, bitLength uint8, value *big.Float, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %f%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) WriteString(logicalName string, bitLength uint8, _ string, value string, writerArgs ...WithWriterArgs) error {
	additionalStringRepresentation := extractAdditionalStringRepresentation(writerArgs...)
	b.PushBack(BoxString(logicalName, fmt.Sprintf("%#0*x %s%s", bitLength/4, value, value, additionalStringRepresentation), 0))
	return nil
}

func (b *boxedWriteBuffer) PopContext(logicalName string, _ ...WithWriterArgs) error {
	b.currentWidth += boxLineOverheat
	finalBoxes := make([]AsciiBox, 0)
findTheBox:
	for back := b.Back(); back != nil; back = b.Back() {
		switch back.Value.(type) {
		case AsciiBox:
			asciiBox := b.Remove(back).(AsciiBox)
			finalBoxes = append([]AsciiBox{asciiBox}, finalBoxes...)
		case []AsciiBox:
			b.Remove(back)
			asciiBoxes := b.Remove(back).([]AsciiBox)
			finalBoxes = append(asciiBoxes, finalBoxes...)
			break findTheBox
		default:
			panic("We should never reach this point")
		}
	}
	asciiBox := BoxBox(logicalName, AlignBoxes(finalBoxes, b.currentWidth), 0)
	b.PushBack(asciiBox)
	return nil
}

func extractAdditionalStringRepresentation(writerArgs ...WithWriterArgs) string {
	for _, arg := range writerArgs {
		if !arg.isWriterArgs() {
			panic("not a writer arg")
		}
		switch arg.(type) {
		case withAdditionalStringRepresentation:
			return " " + arg.(withAdditionalStringRepresentation).stringRepresentation
		}
	}
	return ""
}
