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

package s7

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/s7/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"reflect"
	"strings"
	"testing"
)

func TestS7MessageBytes(t *testing.T) {
	type debuggable interface {
		utils.Serializable
		fmt.Stringer
		utils.AsciiBoxer
	}
	type args struct {
		debuggable debuggable
	}
	tests := []struct {
		name                 string
		args                 args
		wantString           string
		wantStringSerialized string
		wantStringXml        string
		wantDump             string
	}{
		{
			name: "TPKT Packet with Read var response data",
			args: args{
				debuggable: model.NewTPKTPacket(
					model.NewCOTPPacketData(
						false,
						13,
						[]*model.COTPParameter{model.NewCOTPParameterTpduSize(model.COTPTpduSize_SIZE_4096)},
						model.NewS7MessageResponseData(
							0,
							0,
							11,
							model.NewS7ParameterReadVarResponse(1),
							model.NewS7PayloadReadVarResponse(
								[]*model.S7VarPayloadDataItem{
									model.NewS7VarPayloadDataItem(
										model.DataTransportErrorCode_OK,
										model.DataTransportSize_BIT,
										[]int8{1},
									),
								},
							),
						),
					),
				),
			},
			wantString: `
╔═TPKTPacket══════════════════════════════════════════════════════════════════════════════════════════════════════════╗
║╔═ProtocolId╗╔═reserved╗╔═Len═════╗                                                                                  ║
║║  0x03 3   ║║ 0x00 0  ║║0x001d 29║                                                                                  ║
║╚═══════════╝╚═════════╝╚═════════╝                                                                                  ║
║╔═COTPPacket/COTPPacketData/payload═════════════════════════════════════════════════════════════════════════════════╗║
║║╔═HeaderLength╗╔═TpduCode╗╔═Eot════╗╔═TpduRef╗╔═Parameters═════════════════════════════════════════════════╗       ║║
║║║   0x05 5    ║║0xf0 240 ║║b0 false║║0x0d 13 ║║╔═COTPParameter/COTPParameterTpduSize══════════════════════╗║       ║║
║║╚═════════════╝╚═════════╝╚════════╝╚════════╝║║╔═ParameterType╗╔═ParameterLength╗╔═COTPTpduSize/tpduSize╗║║       ║║
║║                                              ║║║   0xc0 192   ║║     0x01 1     ║║    0x0c SIZE_4096    ║║║       ║║
║║                                              ║║╚══════════════╝╚════════════════╝╚══════════════════════╝║║       ║║
║║                                              ║╚══════════════════════════════════════════════════════════╝║       ║║
║║                                              ╚════════════════════════════════════════════════════════════╝       ║║
║║╔═S7Message/S7MessageResponseData/payload═════════════════════════════════════════════════════════════════════════╗║║
║║║╔═ProtocolId╗╔═MessageType╗╔═reserved╗╔═TpduReference╗╔═ParameterLength╗╔═PayloadLength╗╔═ErrorClass╗╔═ErrorCode╗║║║
║║║║  0x32 50  ║║   0x03 3   ║║0x0000 0 ║║  0x000b 11   ║║    0x0002 2    ║║   0x0005 5   ║║  0x00 0   ║║  0x00 0  ║║║║
║║║╚═══════════╝╚════════════╝╚═════════╝╚══════════════╝╚════════════════╝╚══════════════╝╚═══════════╝╚══════════╝║║║
║║║╔═S7Parameter/S7ParameterReadVarResponse/parameter╗                                                              ║║║
║║║║           ╔═ParameterType╗╔═NumItems╗           ║                                                              ║║║
║║║║           ║    0x04 4    ║║ 0x01 1  ║           ║                                                              ║║║
║║║║           ╚══════════════╝╚═════════╝           ║                                                              ║║║
║║║╚═════════════════════════════════════════════════╝                                                              ║║║
║║║╔═S7Payload/S7PayloadReadVarResponse/payload════════════════════════════════════════════════════════════════╗    ║║║
║║║║╔═Items═══════════════════════════════════════════════════════════════════════════════════════════════════╗║    ║║║
║║║║║╔═S7VarPayloadDataItem══════════════════════════════════════════════════════════════════════════════════╗║║    ║║║
║║║║║║╔═DataTransportErrorCode/returnCode╗╔═DataTransportSize/transportSize╗╔═DataLength╗╔═Data═══╗╔═padding╗║║║    ║║║
║║║║║║║             0xff OK              ║║            0x03 BIT            ║║ 0x0001 1  ║║╔══════╗║║ 0x00 0 ║║║║    ║║║
║║║║║║╚══════════════════════════════════╝╚════════════════════════════════╝╚═══════════╝║║0x01 1║║╚════════╝║║║    ║║║
║║║║║║                                                                                   ║╚══════╝║          ║║║    ║║║
║║║║║║                                                                                   ╚════════╝          ║║║    ║║║
║║║║║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════╝║║    ║║║
║║║║╚═════════════════════════════════════════════════════════════════════════════════════════════════════════╝║    ║║║
║║║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════╝    ║║║
║║╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║║
║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║
╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
`,
			wantStringSerialized: `
╔═TPKTPacket═════════════════════════════════════════════════════════════════════════════════════════════════════╗
║╔═protocolId╗╔═reserved╗╔═len═════╗                                                                             ║
║║  0x03 3   ║║ 0x00 0  ║║0x001d 29║                                                                             ║
║╚═══════════╝╚═════════╝╚═════════╝                                                                             ║
║╔═COTPPacket═══════════════════════════════════════════════════════════════════════════════════════════════════╗║
║║╔═headerLength╗╔═tpduCode╗╔═COTPPacketData═════╗╔═parameters═════════════════════════════════════════════════╗║║
║║║   0x05 5    ║║0xf0 240 ║║╔═eot════╗╔═tpduRef╗║║╔═COTPParameter════════════════════════════════════════════╗║║║
║║╚═════════════╝╚═════════╝║║b0 false║║ 0xd 13 ║║║║╔═parameterType╗╔═parameterLength╗╔═COTPParameterTpduSize╗║║║║
║║                          ║╚════════╝╚════════╝║║║║   0xc0 192   ║║     0x01 1     ║║ ╔═COTPTpduSize════╗  ║║║║║
║║                          ╚════════════════════╝║║╚══════════════╝╚════════════════╝║ ║0x0c 12 SIZE_4096║  ║║║║║
║║                                                ║║                                  ║ ╚═════════════════╝  ║║║║║
║║                                                ║║                                  ╚══════════════════════╝║║║║
║║                                                ║╚══════════════════════════════════════════════════════════╝║║║
║║                                                ╚════════════════════════════════════════════════════════════╝║║
║║╔═S7Message══════════════════════════════════════════════════════════════════════════════╗                    ║║
║║║╔═protocolId╗╔═messageType╗╔═reserved╗╔═tpduReference╗╔═parameterLength╗╔═payloadLength╗║                    ║║
║║║║  0x32 50  ║║   0x03 3   ║║0x0000 0 ║║  0x000b 11   ║║    0x0002 2    ║║   0x0005 5   ║║                    ║║
║║║╚═══════════╝╚════════════╝╚═════════╝╚══════════════╝╚════════════════╝╚══════════════╝║                    ║║
║║║╔═S7MessageResponseData═══╗╔═S7Parameter═════════════════════════════════╗              ║                    ║║
║║║║╔═errorClass╗╔═errorCode╗║║╔═parameterType╗╔═S7ParameterReadVarResponse╗║              ║                    ║║
║║║║║  0x00 0   ║║  0x00 0  ║║║║    0x04 4    ║║        ╔═numItems╗        ║║              ║                    ║║
║║║║╚═══════════╝╚══════════╝║║╚══════════════╝║        ║ 0x01 1  ║        ║║              ║                    ║║
║║║╚═════════════════════════╝║                ║        ╚═════════╝        ║║              ║                    ║║
║║║                           ║                ╚═══════════════════════════╝║              ║                    ║║
║║║                           ╚═════════════════════════════════════════════╝              ║                    ║║
║║║╔═S7Payload══════════════════════════════════════════════════════════════════════════╗  ║                    ║║
║║║║╔═S7PayloadReadVarResponse═════════════════════════════════════════════════════════╗║  ║                    ║║
║║║║║╔═items══════════════════════════════════════════════════════════════════════════╗║║  ║                    ║║
║║║║║║╔═S7VarPayloadDataItem═════════════════════════════════════════════════════════╗║║║  ║                    ║║
║║║║║║║╔═DataTransportErrorCode╗╔═DataTransportSize╗╔═dataLength╗╔═data═══╗╔═padding╗║║║║  ║                    ║║
║║║║║║║║      0xff 255 OK      ║║    0x03 3 BIT    ║║ 0x0001 1  ║║╔══════╗║║        ║║║║║  ║                    ║║
║║║║║║║╚═══════════════════════╝╚══════════════════╝╚═══════════╝║║0x01 1║║╚════════╝║║║║  ║                    ║║
║║║║║║║                                                          ║╚══════╝║          ║║║║  ║                    ║║
║║║║║║║                                                          ╚════════╝          ║║║║  ║                    ║║
║║║║║║╚══════════════════════════════════════════════════════════════════════════════╝║║║  ║                    ║║
║║║║║╚════════════════════════════════════════════════════════════════════════════════╝║║  ║                    ║║
║║║║╚══════════════════════════════════════════════════════════════════════════════════╝║  ║                    ║║
║║║╚════════════════════════════════════════════════════════════════════════════════════╝  ║                    ║║
║║╚════════════════════════════════════════════════════════════════════════════════════════╝                    ║║
║╚══════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║
╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
`,
			wantStringXml: `
<TPKTPacket>
  <protocolId dataType="uint8" bitLength="8">3</protocolId>
  <reserved dataType="uint8" bitLength="8">0</reserved>
  <len dataType="uint16" bitLength="16">29</len>
  <COTPPacket>
    <headerLength dataType="uint8" bitLength="8">5</headerLength>
    <tpduCode dataType="uint8" bitLength="8">240</tpduCode>
    <COTPPacketData>
      <eot dataType="bit" bitLength="1">false</eot>
      <tpduRef dataType="uint8" bitLength="7">13</tpduRef>
    </COTPPacketData>
    <parameters>
      <COTPParameter>
        <parameterType dataType="uint8" bitLength="8">192</parameterType>
        <parameterLength dataType="uint8" bitLength="8">1</parameterLength>
        <COTPParameterTpduSize>
          <COTPTpduSize dataType="int8" bitLength="8" stringRepresentation="SIZE_4096">12</COTPTpduSize>
        </COTPParameterTpduSize>
      </COTPParameter>
    </parameters>
    <S7Message>
      <protocolId dataType="uint8" bitLength="8">50</protocolId>
      <messageType dataType="uint8" bitLength="8">3</messageType>
      <reserved dataType="uint16" bitLength="16">0</reserved>
      <tpduReference dataType="uint16" bitLength="16">11</tpduReference>
      <parameterLength dataType="uint16" bitLength="16">2</parameterLength>
      <payloadLength dataType="uint16" bitLength="16">5</payloadLength>
      <S7MessageResponseData>
        <errorClass dataType="uint8" bitLength="8">0</errorClass>
        <errorCode dataType="uint8" bitLength="8">0</errorCode>
      </S7MessageResponseData>
      <S7Parameter>
        <parameterType dataType="uint8" bitLength="8">4</parameterType>
        <S7ParameterReadVarResponse>
          <numItems dataType="uint8" bitLength="8">1</numItems>
        </S7ParameterReadVarResponse>
      </S7Parameter>
      <S7Payload>
        <S7PayloadReadVarResponse>
          <items>
            <S7VarPayloadDataItem>
              <DataTransportErrorCode dataType="uint8" bitLength="8" stringRepresentation="OK">255</DataTransportErrorCode>
              <DataTransportSize dataType="uint8" bitLength="8" stringRepresentation="BIT">3</DataTransportSize>
              <dataLength dataType="uint16" bitLength="16">1</dataLength>
              <data>
                <value dataType="int8" bitLength="8">1</value>
              </data>
              <padding></padding>
            </S7VarPayloadDataItem>
          </items>
        </S7PayloadReadVarResponse>
      </S7Payload>
    </S7Message>
  </COTPPacket>
</TPKTPacket>
`,
			wantDump: `
00|03 00 00 1d 05 f0 0d c0 01 0c '..........'
10|32 03 00 00 00 0b 00 02 00 05 '2.........'
20|00 00 04 01 ff 03 00 01 01    '......... '
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantString = strings.Trim(tt.wantString, "\n")
			if got := tt.args.debuggable.String(); got != tt.wantString {
				t.Errorf("String() = '\n%v\n', want '\n%v\n'", got, tt.wantString)
			}
			boxWriter := utils.NewBoxedWriteBuffer()
			if err := tt.args.debuggable.Serialize(boxWriter); err != nil {
				t.Error(err)
			}
			tt.wantStringSerialized = strings.Trim(tt.wantStringSerialized, "\n")
			if got := string(boxWriter.GetBox()); got != tt.wantStringSerialized {
				t.Errorf("Serialize() = '\n%v\n', want '\n%v\n'", got, tt.wantStringSerialized)
			}
			xmlWriteBuffer := utils.NewXmlWriteBuffer()
			if err := tt.args.debuggable.Serialize(xmlWriteBuffer); err != nil {
				t.Error(err)
			}
			tt.wantStringXml = strings.Trim(tt.wantStringXml, "\n")
			if got := xmlWriteBuffer.GetXmlString(); got != tt.wantStringXml {
				t.Errorf("Serialize() = '\n%v\n', want '\n%v\n'", got, tt.wantStringXml)
			}
			buffer := utils.NewWriteBuffer()
			if err := tt.args.debuggable.Serialize(buffer); err != nil {
				t.Error(err)
			}
			tt.wantDump = strings.Trim(tt.wantDump, "\n")
			if got := utils.Dump(buffer.GetBytes()); !reflect.DeepEqual(got, tt.wantDump) {
				t.Errorf("Serialize() = '\n%v\n', want '\n%v\n'", got, tt.wantDump)
			}
			// and at least a roundtip
			reader := strings.NewReader(tt.wantStringXml)
			readBuffer := utils.NewXmlReadBuffer(reader)
			if got, err := model.TPKTPacketParse(readBuffer); err != nil || !reflect.DeepEqual(got, tt.args.debuggable) {
				if err != nil {
					t.Error(err)
				} else {
					t.Errorf("Roundtrip = '\n%v\n', want '\n%v\n'", got, tt.wantDump)
				}
			}
		})
	}
}
