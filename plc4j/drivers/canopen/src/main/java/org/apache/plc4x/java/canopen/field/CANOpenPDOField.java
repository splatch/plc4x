/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.canopen.field;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.canopen.readwrite.types.CANOpenDataType;
import org.apache.plc4x.java.canopen.readwrite.types.CANOpenService;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class CANOpenPDOField extends CANOpenField implements CANOpenSubscriptionField {

    public static final Pattern ADDRESS_PATTERN = Pattern.compile("(?<pdo>(?:RECEIVE|TRANSMIT)_PDO_[1-4]):" + NODE_PATTERN + ":(?<canDataType>\\w+)(\\[(?<numberOfElements>\\d)])?");
    private final CANOpenService service;
    private final CANOpenDataType canOpenDataType;

    public CANOpenPDOField(int node, CANOpenService service, CANOpenDataType canOpenDataType) {
        super(node);
        this.service = service;
        this.canOpenDataType = canOpenDataType;
    }

    public CANOpenDataType getCanOpenDataType() {
        return canOpenDataType;
    }

    public static boolean matches(String addressString) {
        return ADDRESS_PATTERN.matcher(addressString).matches();
    }

    public static Matcher getMatcher(String addressString) throws PlcInvalidFieldException {
        Matcher matcher = ADDRESS_PATTERN.matcher(addressString);
        if (matcher.matches()) {
            return matcher;
        }

        throw new PlcInvalidFieldException(addressString, ADDRESS_PATTERN);
    }

    public static CANOpenPDOField of(String addressString) {
        Matcher matcher = getMatcher(addressString);
        int nodeId = Integer.parseInt(matcher.group("nodeId"));

        String pdo = matcher.group("pdo");
        CANOpenService service = CANOpenService.valueOf(pdo);
        if (service == null) {
            throw new IllegalArgumentException("Invalid PDO detected " + pdo);
        }

        String canDataTypeString = matcher.group("canDataType");
        CANOpenDataType canOpenDataType = CANOpenDataType.valueOf(canDataTypeString);

        return new CANOpenPDOField(nodeId, service, canOpenDataType);
    }

    public CANOpenService getService() {
        return service;
    }

    @Override
    public boolean isWildcard() {
        return false;
    }

    @Override
    public void xmlSerialize(Element parent) {
        Document doc = parent.getOwnerDocument();
        Element messageElement = doc.createElement(getClass().getSimpleName());
        parent.appendChild(messageElement);

        Element serviceElement = doc.createElement("service");
        serviceElement.appendChild(doc.createTextNode(service.name()));
        messageElement.appendChild(serviceElement);

        Element nodeElement = doc.createElement("node");
        nodeElement.appendChild(doc.createTextNode(Integer.toString(getNodeId())));
        messageElement.appendChild(nodeElement);

        Element dataType = doc.createElement("dataType");
        dataType.appendChild(doc.createTextNode(getCanOpenDataType().name()));
        messageElement.appendChild(dataType);
    }
}
