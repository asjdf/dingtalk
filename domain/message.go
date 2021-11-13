/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package domain

type message struct {
	MsgType string `json:"msgtype" validate:"required,oneof=text image voice file link oa markdown action_card feedCard"`
}

//MessageResponse:发送消息返回
type MessageResponse struct {
	Response
	MessageId string `json:"messageId"` //指定员工的部门信息。
}

//SendToConversationResponse:发送普通消息返回
type SendToConversationResponse struct {
	Response
	Receiver string `json:"receiver"`
}