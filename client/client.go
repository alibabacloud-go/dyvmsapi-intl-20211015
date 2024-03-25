// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BackendCallGroupRequest struct {
	// The called numbers. You can specify up to 50,000 phone numbers.
	CalledNumber []*string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty" type:"Repeated"`
	// The calling number.
	//
	// If you do not specify this parameter, the system uses a local random number as the display number. If you use a dedicated number for outbound calls, you must specify the purchased number. You can specify only one number. You can log on to the VMS console and choose Number Management to view the purchased phone numbers.
	CallerIdNumber *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	// The ISO2 country code.
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller in a receipt message.
	//
	// The value must be of the STRING type and 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times the audio file is played. Valid values: 1 to 3.
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The delivery type. Valid values: 1 and 2.
	//
	// 1: The audio file is delivered immediately.
	//
	// 2: The audio file is delivered at a scheduled time. If you specify SendType as 2, you must specify TimingStart.
	SendType *int64 `json:"SendType,omitempty" xml:"SendType,omitempty"`
	// The playback speed. Valid values: -500 to 500.
	//
	// You must specify this parameter when the audio type is text-to-speech (TTS). You do not need to specify this parameter when you use recordings.
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The time when the audio file is scheduled to be delivered.
	TimingStart *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	// The voice template ID of the audio file.
	//
	// You can log on to the VMS console and choose Voice Call Template > Audio File to view the template ID.
	//
	// You must specify either TtsCode or VoiceCode. You can specify TtsCode to use the audio file as the call content. Alternatively, you can specify VoiceCode to use preset text as the call content.
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// The TTS template ID.
	//
	// You can log on to the VMS console and choose Voice Call Template > TTS Template to view the template ID.
	//
	// You must specify either TtsCode or VoiceCode. You can specify TtsCode to use the audio file as the call content. Alternatively, you can specify VoiceCode to use preset text as the call content.
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The playback volume of the audio file. Valid values: 0 to 100. Default value: 100.
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s BackendCallGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupRequest) GoString() string {
	return s.String()
}

func (s *BackendCallGroupRequest) SetCalledNumber(v []*string) *BackendCallGroupRequest {
	s.CalledNumber = v
	return s
}

func (s *BackendCallGroupRequest) SetCallerIdNumber(v string) *BackendCallGroupRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *BackendCallGroupRequest) SetCountryId(v string) *BackendCallGroupRequest {
	s.CountryId = &v
	return s
}

func (s *BackendCallGroupRequest) SetOutId(v string) *BackendCallGroupRequest {
	s.OutId = &v
	return s
}

func (s *BackendCallGroupRequest) SetOwnerId(v int64) *BackendCallGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *BackendCallGroupRequest) SetPlayTimes(v int64) *BackendCallGroupRequest {
	s.PlayTimes = &v
	return s
}

func (s *BackendCallGroupRequest) SetResourceOwnerAccount(v string) *BackendCallGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BackendCallGroupRequest) SetResourceOwnerId(v int64) *BackendCallGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BackendCallGroupRequest) SetSendType(v int64) *BackendCallGroupRequest {
	s.SendType = &v
	return s
}

func (s *BackendCallGroupRequest) SetSpeed(v int64) *BackendCallGroupRequest {
	s.Speed = &v
	return s
}

func (s *BackendCallGroupRequest) SetTaskName(v string) *BackendCallGroupRequest {
	s.TaskName = &v
	return s
}

func (s *BackendCallGroupRequest) SetTimingStart(v string) *BackendCallGroupRequest {
	s.TimingStart = &v
	return s
}

func (s *BackendCallGroupRequest) SetTtsCode(v string) *BackendCallGroupRequest {
	s.TtsCode = &v
	return s
}

func (s *BackendCallGroupRequest) SetVoiceCode(v string) *BackendCallGroupRequest {
	s.VoiceCode = &v
	return s
}

func (s *BackendCallGroupRequest) SetVolume(v int64) *BackendCallGroupRequest {
	s.Volume = &v
	return s
}

type BackendCallGroupShrinkRequest struct {
	// The called numbers. You can specify up to 50,000 phone numbers.
	CalledNumberShrink *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// If you do not specify this parameter, the system uses a local random number as the display number. If you use a dedicated number for outbound calls, you must specify the purchased number. You can specify only one number. You can log on to the VMS console and choose Number Management to view the purchased phone numbers.
	CallerIdNumber *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	// The ISO2 country code.
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller in a receipt message.
	//
	// The value must be of the STRING type and 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times the audio file is played. Valid values: 1 to 3.
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The delivery type. Valid values: 1 and 2.
	//
	// 1: The audio file is delivered immediately.
	//
	// 2: The audio file is delivered at a scheduled time. If you specify SendType as 2, you must specify TimingStart.
	SendType *int64 `json:"SendType,omitempty" xml:"SendType,omitempty"`
	// The playback speed. Valid values: -500 to 500.
	//
	// You must specify this parameter when the audio type is text-to-speech (TTS). You do not need to specify this parameter when you use recordings.
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The time when the audio file is scheduled to be delivered.
	TimingStart *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	// The voice template ID of the audio file.
	//
	// You can log on to the VMS console and choose Voice Call Template > Audio File to view the template ID.
	//
	// You must specify either TtsCode or VoiceCode. You can specify TtsCode to use the audio file as the call content. Alternatively, you can specify VoiceCode to use preset text as the call content.
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// The TTS template ID.
	//
	// You can log on to the VMS console and choose Voice Call Template > TTS Template to view the template ID.
	//
	// You must specify either TtsCode or VoiceCode. You can specify TtsCode to use the audio file as the call content. Alternatively, you can specify VoiceCode to use preset text as the call content.
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The playback volume of the audio file. Valid values: 0 to 100. Default value: 100.
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s BackendCallGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *BackendCallGroupShrinkRequest) SetCalledNumberShrink(v string) *BackendCallGroupShrinkRequest {
	s.CalledNumberShrink = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetCallerIdNumber(v string) *BackendCallGroupShrinkRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetCountryId(v string) *BackendCallGroupShrinkRequest {
	s.CountryId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetOutId(v string) *BackendCallGroupShrinkRequest {
	s.OutId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetOwnerId(v int64) *BackendCallGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetPlayTimes(v int64) *BackendCallGroupShrinkRequest {
	s.PlayTimes = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetResourceOwnerAccount(v string) *BackendCallGroupShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetResourceOwnerId(v int64) *BackendCallGroupShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetSendType(v int64) *BackendCallGroupShrinkRequest {
	s.SendType = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetSpeed(v int64) *BackendCallGroupShrinkRequest {
	s.Speed = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetTaskName(v string) *BackendCallGroupShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetTimingStart(v string) *BackendCallGroupShrinkRequest {
	s.TimingStart = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetTtsCode(v string) *BackendCallGroupShrinkRequest {
	s.TtsCode = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetVoiceCode(v string) *BackendCallGroupShrinkRequest {
	s.VoiceCode = &v
	return s
}

func (s *BackendCallGroupShrinkRequest) SetVolume(v int64) *BackendCallGroupShrinkRequest {
	s.Volume = &v
	return s
}

type BackendCallGroupResponseBody struct {
	// The response code.
	//
	// The value OK indicates that the request was successful. Other values indicate that the request failed. For more information, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID. You can use this ID to query the details of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BackendCallGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BackendCallGroupResponseBody) SetCode(v string) *BackendCallGroupResponseBody {
	s.Code = &v
	return s
}

func (s *BackendCallGroupResponseBody) SetMessage(v string) *BackendCallGroupResponseBody {
	s.Message = &v
	return s
}

func (s *BackendCallGroupResponseBody) SetRequestId(v string) *BackendCallGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackendCallGroupResponseBody) SetTaskId(v string) *BackendCallGroupResponseBody {
	s.TaskId = &v
	return s
}

type BackendCallGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackendCallGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackendCallGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BackendCallGroupResponse) GoString() string {
	return s.String()
}

func (s *BackendCallGroupResponse) SetHeaders(v map[string]*string) *BackendCallGroupResponse {
	s.Headers = v
	return s
}

func (s *BackendCallGroupResponse) SetStatusCode(v int32) *BackendCallGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *BackendCallGroupResponse) SetBody(v *BackendCallGroupResponseBody) *BackendCallGroupResponse {
	s.Body = v
	return s
}

type BackendCallSignalRequest struct {
	// The phone number that receives the voice notification.
	//
	// You must add the country code to the beginning of the phone number. Example: 85200\*\*\*\*00.
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// If you do not specify this parameter, the system uses a local random number as the display number. If you use a dedicated number for outbound calls, you must specify the purchased number. You can specify only one number. You can log on to the VMS console and choose Number Management to view the purchased phone numbers.
	CallerIdNumber *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	// The ISO2 country code.
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller in a receipt message.
	//
	// The value must be of the STRING type and 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times the voice notification is played back in a call. Valid values: 1 to 3. Default value: 3.
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The playback speed. Valid values: -500 to 500.
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The ID of the approved voice verification code template.
	//
	// You can log on to the VMS console and choose Voice Call Template to view the template ID.
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// The variables in the template, in the JSON format.
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The playback volume of the voice notification. Valid values: 0 to 100. Default value: 100.
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s BackendCallSignalRequest) String() string {
	return tea.Prettify(s)
}

func (s BackendCallSignalRequest) GoString() string {
	return s.String()
}

func (s *BackendCallSignalRequest) SetCalledNumber(v string) *BackendCallSignalRequest {
	s.CalledNumber = &v
	return s
}

func (s *BackendCallSignalRequest) SetCallerIdNumber(v string) *BackendCallSignalRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *BackendCallSignalRequest) SetCountryId(v string) *BackendCallSignalRequest {
	s.CountryId = &v
	return s
}

func (s *BackendCallSignalRequest) SetOutId(v string) *BackendCallSignalRequest {
	s.OutId = &v
	return s
}

func (s *BackendCallSignalRequest) SetOwnerId(v int64) *BackendCallSignalRequest {
	s.OwnerId = &v
	return s
}

func (s *BackendCallSignalRequest) SetPlayTimes(v int64) *BackendCallSignalRequest {
	s.PlayTimes = &v
	return s
}

func (s *BackendCallSignalRequest) SetResourceOwnerAccount(v string) *BackendCallSignalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BackendCallSignalRequest) SetResourceOwnerId(v int64) *BackendCallSignalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BackendCallSignalRequest) SetSpeed(v int64) *BackendCallSignalRequest {
	s.Speed = &v
	return s
}

func (s *BackendCallSignalRequest) SetTtsCode(v string) *BackendCallSignalRequest {
	s.TtsCode = &v
	return s
}

func (s *BackendCallSignalRequest) SetTtsParam(v string) *BackendCallSignalRequest {
	s.TtsParam = &v
	return s
}

func (s *BackendCallSignalRequest) SetVolume(v int64) *BackendCallSignalRequest {
	s.Volume = &v
	return s
}

type BackendCallSignalResponseBody struct {
	// The unique receipt ID for the call. You can use this ID to query the details of a single call.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// The value OK indicates that the request was successful. Other values indicate that the request failed. For more information, see Error codes.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BackendCallSignalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BackendCallSignalResponseBody) GoString() string {
	return s.String()
}

func (s *BackendCallSignalResponseBody) SetCallId(v string) *BackendCallSignalResponseBody {
	s.CallId = &v
	return s
}

func (s *BackendCallSignalResponseBody) SetCode(v string) *BackendCallSignalResponseBody {
	s.Code = &v
	return s
}

func (s *BackendCallSignalResponseBody) SetMessage(v string) *BackendCallSignalResponseBody {
	s.Message = &v
	return s
}

func (s *BackendCallSignalResponseBody) SetRequestId(v string) *BackendCallSignalResponseBody {
	s.RequestId = &v
	return s
}

type BackendCallSignalResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BackendCallSignalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BackendCallSignalResponse) String() string {
	return tea.Prettify(s)
}

func (s BackendCallSignalResponse) GoString() string {
	return s.String()
}

func (s *BackendCallSignalResponse) SetHeaders(v map[string]*string) *BackendCallSignalResponse {
	s.Headers = v
	return s
}

func (s *BackendCallSignalResponse) SetStatusCode(v int32) *BackendCallSignalResponse {
	s.StatusCode = &v
	return s
}

func (s *BackendCallSignalResponse) SetBody(v *BackendCallSignalResponseBody) *BackendCallSignalResponse {
	s.Body = v
	return s
}

type GroupCallRequest struct {
	// 被叫号码。上限为5万个。
	CalledNumber []*string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty" type:"Repeated"`
	// 主叫号码。  若您不填该参数，系统将会使用当地随机号码作为外显号码。 若您专属号码外呼，则必须传入已购买的号码，仅支持一个号码。您可以登录国际语音服务控制台，选择"号码管理"查看已购买的号码。
	CallerIdNumber *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	// 国家/地区二字码，ISO2。
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// 预留给调用方使用的ID，最终会通过在回执消息中将此ID带回给调用方。  字符串类型，长度为1~15个字节。
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 语音文件的播放次数。取值范围：1~3。
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 发送类型：取值[1,2]。  1： 立即开始发送任务，不等待。  2： 定时开始发送任务。如果传该类型，TimingStart为必选字段。
	SendType       *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	SignatureNonce *string `json:"SignatureNonce,omitempty" xml:"SignatureNonce,omitempty"`
	// 语速控制，取值范围：-500~500。  音频类型为TTS时必传。录音文件可不传。
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// 任务名称。
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// 2022-05-01T08:00:00+08:00
	TimingStart *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	// 文本转语音的模板ID。  您可以登录国际语音服务控制台，选择"语音模板管理"-"文本转语音模板"，查看模板ID。  此参数与VoiceCode二选一必填。分别代表使用语音文件作为呼叫内容呼叫或者使用固定模板文字作为呼叫内容。
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// 模板中的变量参数，JSON格式。
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// 语音文件的模板ID。  您可以登录国际语音服务控制台，选择"语音模板管理"-"语音文件"，查看模板ID。  此参数与TtsCode二选一必填。分别代表使用语音文件作为呼叫内容呼叫或者使用固定模板文字作为呼叫内容。
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// 语音文件播放的音量。取值范围：0~100，默认取值100。
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GroupCallRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCallRequest) GoString() string {
	return s.String()
}

func (s *GroupCallRequest) SetCalledNumber(v []*string) *GroupCallRequest {
	s.CalledNumber = v
	return s
}

func (s *GroupCallRequest) SetCallerIdNumber(v string) *GroupCallRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *GroupCallRequest) SetCountryId(v string) *GroupCallRequest {
	s.CountryId = &v
	return s
}

func (s *GroupCallRequest) SetOutId(v string) *GroupCallRequest {
	s.OutId = &v
	return s
}

func (s *GroupCallRequest) SetOwnerId(v int64) *GroupCallRequest {
	s.OwnerId = &v
	return s
}

func (s *GroupCallRequest) SetPlayTimes(v int64) *GroupCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *GroupCallRequest) SetResourceOwnerAccount(v string) *GroupCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GroupCallRequest) SetResourceOwnerId(v int64) *GroupCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GroupCallRequest) SetSendType(v int64) *GroupCallRequest {
	s.SendType = &v
	return s
}

func (s *GroupCallRequest) SetSignature(v string) *GroupCallRequest {
	s.Signature = &v
	return s
}

func (s *GroupCallRequest) SetSignatureNonce(v string) *GroupCallRequest {
	s.SignatureNonce = &v
	return s
}

func (s *GroupCallRequest) SetSpeed(v int64) *GroupCallRequest {
	s.Speed = &v
	return s
}

func (s *GroupCallRequest) SetTaskName(v string) *GroupCallRequest {
	s.TaskName = &v
	return s
}

func (s *GroupCallRequest) SetTimestamp(v string) *GroupCallRequest {
	s.Timestamp = &v
	return s
}

func (s *GroupCallRequest) SetTimingStart(v string) *GroupCallRequest {
	s.TimingStart = &v
	return s
}

func (s *GroupCallRequest) SetTtsCode(v string) *GroupCallRequest {
	s.TtsCode = &v
	return s
}

func (s *GroupCallRequest) SetTtsParam(v string) *GroupCallRequest {
	s.TtsParam = &v
	return s
}

func (s *GroupCallRequest) SetVoiceCode(v string) *GroupCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *GroupCallRequest) SetVolume(v int64) *GroupCallRequest {
	s.Volume = &v
	return s
}

type GroupCallShrinkRequest struct {
	// 被叫号码。上限为5万个。
	CalledNumberShrink *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// 主叫号码。  若您不填该参数，系统将会使用当地随机号码作为外显号码。 若您专属号码外呼，则必须传入已购买的号码，仅支持一个号码。您可以登录国际语音服务控制台，选择"号码管理"查看已购买的号码。
	CallerIdNumber *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	// 国家/地区二字码，ISO2。
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// 预留给调用方使用的ID，最终会通过在回执消息中将此ID带回给调用方。  字符串类型，长度为1~15个字节。
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 语音文件的播放次数。取值范围：1~3。
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 发送类型：取值[1,2]。  1： 立即开始发送任务，不等待。  2： 定时开始发送任务。如果传该类型，TimingStart为必选字段。
	SendType       *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	SignatureNonce *string `json:"SignatureNonce,omitempty" xml:"SignatureNonce,omitempty"`
	// 语速控制，取值范围：-500~500。  音频类型为TTS时必传。录音文件可不传。
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// 任务名称。
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// 2022-05-01T08:00:00+08:00
	TimingStart *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	// 文本转语音的模板ID。  您可以登录国际语音服务控制台，选择"语音模板管理"-"文本转语音模板"，查看模板ID。  此参数与VoiceCode二选一必填。分别代表使用语音文件作为呼叫内容呼叫或者使用固定模板文字作为呼叫内容。
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// 模板中的变量参数，JSON格式。
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// 语音文件的模板ID。  您可以登录国际语音服务控制台，选择"语音模板管理"-"语音文件"，查看模板ID。  此参数与TtsCode二选一必填。分别代表使用语音文件作为呼叫内容呼叫或者使用固定模板文字作为呼叫内容。
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// 语音文件播放的音量。取值范围：0~100，默认取值100。
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s GroupCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *GroupCallShrinkRequest) SetCalledNumberShrink(v string) *GroupCallShrinkRequest {
	s.CalledNumberShrink = &v
	return s
}

func (s *GroupCallShrinkRequest) SetCallerIdNumber(v string) *GroupCallShrinkRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *GroupCallShrinkRequest) SetCountryId(v string) *GroupCallShrinkRequest {
	s.CountryId = &v
	return s
}

func (s *GroupCallShrinkRequest) SetOutId(v string) *GroupCallShrinkRequest {
	s.OutId = &v
	return s
}

func (s *GroupCallShrinkRequest) SetOwnerId(v int64) *GroupCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *GroupCallShrinkRequest) SetPlayTimes(v int64) *GroupCallShrinkRequest {
	s.PlayTimes = &v
	return s
}

func (s *GroupCallShrinkRequest) SetResourceOwnerAccount(v string) *GroupCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GroupCallShrinkRequest) SetResourceOwnerId(v int64) *GroupCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GroupCallShrinkRequest) SetSendType(v int64) *GroupCallShrinkRequest {
	s.SendType = &v
	return s
}

func (s *GroupCallShrinkRequest) SetSignature(v string) *GroupCallShrinkRequest {
	s.Signature = &v
	return s
}

func (s *GroupCallShrinkRequest) SetSignatureNonce(v string) *GroupCallShrinkRequest {
	s.SignatureNonce = &v
	return s
}

func (s *GroupCallShrinkRequest) SetSpeed(v int64) *GroupCallShrinkRequest {
	s.Speed = &v
	return s
}

func (s *GroupCallShrinkRequest) SetTaskName(v string) *GroupCallShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *GroupCallShrinkRequest) SetTimestamp(v string) *GroupCallShrinkRequest {
	s.Timestamp = &v
	return s
}

func (s *GroupCallShrinkRequest) SetTimingStart(v string) *GroupCallShrinkRequest {
	s.TimingStart = &v
	return s
}

func (s *GroupCallShrinkRequest) SetTtsCode(v string) *GroupCallShrinkRequest {
	s.TtsCode = &v
	return s
}

func (s *GroupCallShrinkRequest) SetTtsParam(v string) *GroupCallShrinkRequest {
	s.TtsParam = &v
	return s
}

func (s *GroupCallShrinkRequest) SetVoiceCode(v string) *GroupCallShrinkRequest {
	s.VoiceCode = &v
	return s
}

func (s *GroupCallShrinkRequest) SetVolume(v int64) *GroupCallShrinkRequest {
	s.Volume = &v
	return s
}

type GroupCallResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码。  返回OK代表请求成功。 其他错误码，请参见API错误码。
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息描述
	Message *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GroupCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GroupCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupCallResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCallResponseBody) SetAccessDeniedDetail(v string) *GroupCallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GroupCallResponseBody) SetCode(v string) *GroupCallResponseBody {
	s.Code = &v
	return s
}

func (s *GroupCallResponseBody) SetMessage(v string) *GroupCallResponseBody {
	s.Message = &v
	return s
}

func (s *GroupCallResponseBody) SetModel(v *GroupCallResponseBodyModel) *GroupCallResponseBody {
	s.Model = v
	return s
}

func (s *GroupCallResponseBody) SetRequestId(v string) *GroupCallResponseBody {
	s.RequestId = &v
	return s
}

type GroupCallResponseBodyModel struct {
	// 任务ID，可以通过此ID查询任务的详情。
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GroupCallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s GroupCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GroupCallResponseBodyModel) SetTaskId(v string) *GroupCallResponseBodyModel {
	s.TaskId = &v
	return s
}

type GroupCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupCallResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCallResponse) GoString() string {
	return s.String()
}

func (s *GroupCallResponse) SetHeaders(v map[string]*string) *GroupCallResponse {
	s.Headers = v
	return s
}

func (s *GroupCallResponse) SetStatusCode(v int32) *GroupCallResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupCallResponse) SetBody(v *GroupCallResponseBody) *GroupCallResponse {
	s.Body = v
	return s
}

type SignalCallRequest struct {
	// 接收语音通知的手机号码。  号码格式：国际码+号码： 示例：85200****00。
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// 主叫号码。  若您不填该参数，系统将会使用当地随机号码作为外显号码。 若您专属号码外呼，则必须传入已购买的号码，仅支持一个号码。您可以登录国际语音服务控制台，选择"号码管理"查看已购买的号码。
	CallerIdNumber *string `json:"CallerIdNumber,omitempty" xml:"CallerIdNumber,omitempty"`
	// 国家/地区二字码，ISO2。
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// 预留给调用方使用的ID，最终会通过在回执消息中将此ID带回给调用方。  字符串类型，长度为1~15个字节。
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 一通电话内语音通知内容的播放次数。取值范围：1~3，默认取值3。
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 发送类型：取值[1,2]。  1： 立即开始发送任务，不等待。  2： 定时开始发送任务。如果传该类型，TimingStart为必选字段。
	SendType       *int64  `json:"SendType,omitempty" xml:"SendType,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	SignatureNonce *string `json:"SignatureNonce,omitempty" xml:"SignatureNonce,omitempty"`
	// 语速控制。取值范围为：-500~500。
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// 任务名称。
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// 定时发送的时间。
	TimingStart *string `json:"TimingStart,omitempty" xml:"TimingStart,omitempty"`
	// 文本转语音模板的语音ID。  您可以登录国际语音服务控制台，选择"语音模板管理"-"文本转语音模板"，查看模板ID。  此参数与VoiceCode二选一必填。分别代表使用语音文件作为呼叫内容呼叫或者使用固定模板文字作为呼叫内容。
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// 模板中的变量参数，JSON格式。
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// 语音文件的模板ID。  您可以登录国际语音服务控制台，选择"语音模板管理"-"语音文件"，查看模板ID。  此参数与TtsCode二选一必填。分别代表使用语音文件作为呼叫内容呼叫或者使用固定模板文字作为呼叫内容。
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// 语音通知的播放音量。取值范围：0~100，默认取值100。
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SignalCallRequest) String() string {
	return tea.Prettify(s)
}

func (s SignalCallRequest) GoString() string {
	return s.String()
}

func (s *SignalCallRequest) SetCalledNumber(v string) *SignalCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *SignalCallRequest) SetCallerIdNumber(v string) *SignalCallRequest {
	s.CallerIdNumber = &v
	return s
}

func (s *SignalCallRequest) SetCountryId(v string) *SignalCallRequest {
	s.CountryId = &v
	return s
}

func (s *SignalCallRequest) SetOutId(v string) *SignalCallRequest {
	s.OutId = &v
	return s
}

func (s *SignalCallRequest) SetOwnerId(v int64) *SignalCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SignalCallRequest) SetPlayTimes(v int64) *SignalCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *SignalCallRequest) SetResourceOwnerAccount(v string) *SignalCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SignalCallRequest) SetResourceOwnerId(v int64) *SignalCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SignalCallRequest) SetSendType(v int64) *SignalCallRequest {
	s.SendType = &v
	return s
}

func (s *SignalCallRequest) SetSignature(v string) *SignalCallRequest {
	s.Signature = &v
	return s
}

func (s *SignalCallRequest) SetSignatureNonce(v string) *SignalCallRequest {
	s.SignatureNonce = &v
	return s
}

func (s *SignalCallRequest) SetSpeed(v int64) *SignalCallRequest {
	s.Speed = &v
	return s
}

func (s *SignalCallRequest) SetTaskName(v string) *SignalCallRequest {
	s.TaskName = &v
	return s
}

func (s *SignalCallRequest) SetTimestamp(v string) *SignalCallRequest {
	s.Timestamp = &v
	return s
}

func (s *SignalCallRequest) SetTimingStart(v string) *SignalCallRequest {
	s.TimingStart = &v
	return s
}

func (s *SignalCallRequest) SetTtsCode(v string) *SignalCallRequest {
	s.TtsCode = &v
	return s
}

func (s *SignalCallRequest) SetTtsParam(v string) *SignalCallRequest {
	s.TtsParam = &v
	return s
}

func (s *SignalCallRequest) SetVoiceCode(v string) *SignalCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *SignalCallRequest) SetVolume(v int64) *SignalCallRequest {
	s.Volume = &v
	return s
}

type SignalCallResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码。  返回OK代表请求成功。 其他错误码，请参见API错误码。
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息描述。
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *SignalCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功，成功：true，失败：false。
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SignalCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignalCallResponseBody) GoString() string {
	return s.String()
}

func (s *SignalCallResponseBody) SetAccessDeniedDetail(v string) *SignalCallResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SignalCallResponseBody) SetCode(v string) *SignalCallResponseBody {
	s.Code = &v
	return s
}

func (s *SignalCallResponseBody) SetMessage(v string) *SignalCallResponseBody {
	s.Message = &v
	return s
}

func (s *SignalCallResponseBody) SetModel(v *SignalCallResponseBodyModel) *SignalCallResponseBody {
	s.Model = v
	return s
}

func (s *SignalCallResponseBody) SetRequestId(v string) *SignalCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignalCallResponseBody) SetSuccess(v bool) *SignalCallResponseBody {
	s.Success = &v
	return s
}

type SignalCallResponseBodyModel struct {
	// 任务ID，可以通过此ID查询任务的详情。
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s SignalCallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s SignalCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SignalCallResponseBodyModel) SetCallId(v string) *SignalCallResponseBodyModel {
	s.CallId = &v
	return s
}

type SignalCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignalCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignalCallResponse) String() string {
	return tea.Prettify(s)
}

func (s SignalCallResponse) GoString() string {
	return s.String()
}

func (s *SignalCallResponse) SetHeaders(v map[string]*string) *SignalCallResponse {
	s.Headers = v
	return s
}

func (s *SignalCallResponse) SetStatusCode(v int32) *SignalCallResponse {
	s.StatusCode = &v
	return s
}

func (s *SignalCallResponse) SetBody(v *SignalCallResponseBody) *SignalCallResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dyvmsapi-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BackendCallGroupWithOptions(tmpReq *BackendCallGroupRequest, runtime *util.RuntimeOptions) (_result *BackendCallGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BackendCallGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CalledNumber)) {
		request.CalledNumberShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CalledNumber, tea.String("CalledNumber"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumberShrink)) {
		query["CalledNumber"] = request.CalledNumberShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		query["SendType"] = request.SendType
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TimingStart)) {
		query["TimingStart"] = request.TimingStart
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BackendCallGroup"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BackendCallGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BackendCallGroup(request *BackendCallGroupRequest) (_result *BackendCallGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BackendCallGroupResponse{}
	_body, _err := client.BackendCallGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BackendCallSignalWithOptions(request *BackendCallSignalRequest, runtime *util.RuntimeOptions) (_result *BackendCallSignalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BackendCallSignal"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BackendCallSignalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BackendCallSignal(request *BackendCallSignalRequest) (_result *BackendCallSignalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BackendCallSignalResponse{}
	_body, _err := client.BackendCallSignalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupCallWithOptions(tmpReq *GroupCallRequest, runtime *util.RuntimeOptions) (_result *GroupCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GroupCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CalledNumber)) {
		request.CalledNumberShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CalledNumber, tea.String("CalledNumber"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumberShrink)) {
		query["CalledNumber"] = request.CalledNumberShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		query["SendType"] = request.SendType
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureNonce)) {
		query["SignatureNonce"] = request.SignatureNonce
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["Timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.TimingStart)) {
		query["TimingStart"] = request.TimingStart
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GroupCall"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupCall(request *GroupCallRequest) (_result *GroupCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GroupCallResponse{}
	_body, _err := client.GroupCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignalCallWithOptions(request *SignalCallRequest, runtime *util.RuntimeOptions) (_result *SignalCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallerIdNumber)) {
		query["CallerIdNumber"] = request.CallerIdNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CountryId)) {
		query["CountryId"] = request.CountryId
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendType)) {
		query["SendType"] = request.SendType
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["Signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SignatureNonce)) {
		query["SignatureNonce"] = request.SignatureNonce
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["Timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.TimingStart)) {
		query["TimingStart"] = request.TimingStart
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SignalCall"),
		Version:     tea.String("2021-10-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SignalCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignalCall(request *SignalCallRequest) (_result *SignalCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SignalCallResponse{}
	_body, _err := client.SignalCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
