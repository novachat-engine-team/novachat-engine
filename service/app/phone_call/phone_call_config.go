/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/11 14:33
 * @File : phone_call_config.go
 */

package phone_call

import jsoniter "github.com/json-iterator/go"

/*
#if TARGET_OS_IPHONE
	flag=ServerConfig::GetSharedInstance()->GetBoolean("use_ios_vpio_agc", true) ? 1 : 0;
#else
	flag=ServerConfig::GetSharedInstance()->GetBoolean("use_osx_vpio_agc", true) ? 1 : 0;
#endif

cwnd=(size_t) ServerConfig::GetSharedInstance()->GetInt("audio_congestion_window", 1024);

#ifdef __APPLE__
	switch(ServerConfig::GetSharedInstance()->GetInt("webrtc_ns_level_vpio", 0)){
#else
    switch (ServerConfig::GetSharedInstance()->GetInt("webrtc_ns_level", 2)) {
#endif

if (enableAGC) {
    config.gain_controller1.mode = webrtc::AudioProcessing::Config::GainController1::kAdaptiveDigital;
    config.gain_controller1.target_level_dbfs = ServerConfig::GetSharedInstance()->GetInt("webrtc_agc_target_level", 9);
    config.gain_controller1.enable_limiter = ServerConfig::GetSharedInstance()->GetBoolean("webrtc_agc_enable_limiter", true);
    config.gain_controller1.compression_gain_db = ServerConfig::GetSharedInstance()->GetInt("webrtc_agc_compression_gain", 20);
}

if(step<30){
    minMinDelay=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_min_delay_20", 6);
    maxMinDelay=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_max_delay_20", 25);
    maxUsedSlots=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_max_slots_20", 50);
}else if(step<50){
    minMinDelay=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_min_delay_40", 4);
    maxMinDelay=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_max_delay_40", 15);
    maxUsedSlots=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_max_slots_40", 30);
}else{
    minMinDelay=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_min_delay_60", 2);
    maxMinDelay=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_max_delay_60", 10);
    maxUsedSlots=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_max_slots_60", 20);
}
lossesToReset=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_losses_to_reset", 20);
resyncThreshold=ServerConfig::GetSharedInstance()->GetDouble("jitter_resync_threshold", 1.0);

ipv6Timeout=ServerConfig::GetSharedInstance()->GetDouble("nat64_fallback_timeout", 3);

vadNoVoiceBitrate=static_cast<uint32_t>(ServerConfig::GetSharedInstance()->GetInt("audio_vad_no_voice_bitrate", 6000));
vadModeVoiceBandwidth=serverConfigValueToBandwidth(ServerConfig::GetSharedInstance()->GetInt("audio_vad_bandwidth", 3));
vadModeNoVoiceBandwidth=serverConfigValueToBandwidth(ServerConfig::GetSharedInstance()->GetInt("audio_vad_no_voice_bandwidth", 0));
secondaryEnabledBandwidth=serverConfigValueToBandwidth(ServerConfig::GetSharedInstance()->GetInt("audio_extra_ec_bandwidth", 2));

maxAudioBitrate=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_max_bitrate", 20000);
maxAudioBitrateGPRS=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_max_bitrate_gprs", 8000);
maxAudioBitrateEDGE=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_max_bitrate_edge", 16000);
maxAudioBitrateSaving=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_max_bitrate_saving", 8000);
initAudioBitrate=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_init_bitrate", 16000);
initAudioBitrateGPRS=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_init_bitrate_gprs", 8000);
initAudioBitrateEDGE=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_init_bitrate_edge", 8000);
initAudioBitrateSaving=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_init_bitrate_saving", 8000);
audioBitrateStepIncr=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_bitrate_step_incr", 1000);
audioBitrateStepDecr=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_bitrate_step_decr", 1000);
minAudioBitrate=(uint32_t) ServerConfig::GetSharedInstance()->GetInt("audio_min_bitrate", 8000);
relaySwitchThreshold=ServerConfig::GetSharedInstance()->GetDouble("relay_switch_threshold", 0.8);
p2pToRelaySwitchThreshold=ServerConfig::GetSharedInstance()->GetDouble("p2p_to_relay_switch_threshold", 0.6);
relayToP2pSwitchThreshold=ServerConfig::GetSharedInstance()->GetDouble("relay_to_p2p_switch_threshold", 0.8);
reconnectingTimeout=ServerConfig::GetSharedInstance()->GetDouble("reconnecting_state_timeout", 2.0);
needRateFlags=static_cast<uint32_t>(ServerConfig::GetSharedInstance()->GetInt("rate_flags", 0xFFFFFFFF));
rateMaxAcceptableRTT=ServerConfig::GetSharedInstance()->GetDouble("rate_min_rtt", 0.6);
rateMaxAcceptableSendLoss=ServerConfig::GetSharedInstance()->GetDouble("rate_min_send_loss", 0.2);
packetLossToEnableExtraEC=ServerConfig::GetSharedInstance()->GetDouble("packet_loss_for_extra_ec", 0.02);
maxUnsentStreamPackets=static_cast<uint32_t>(ServerConfig::GetSharedInstance()->GetInt("max_unsent_stream_packets", 2));

ServerConfig::GetSharedInstance()->GetBoolean("bad_call_rating", false);

if(stm->frameDuration>50)
    stm->jitterBuffer->SetMinPacketCount((uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_initial_delay_60", 2));
else if(stm->frameDuration>30)
    stm->jitterBuffer->SetMinPacketCount((uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_initial_delay_40", 4));
else
    stm->jitterBuffer->SetMinPacketCount((uint32_t) ServerConfig::GetSharedInstance()->GetInt("jitter_initial_delay_20", 6));

ServerConfig::GetSharedInstance()->GetDouble("established_delay_if_no_stream_data", 1.5)

ServerConfig::GetSharedInstance()->GetBoolean("use_tcp", true);

if(type==Type::UDP_RELAY && ServerConfig::GetSharedInstance()->GetBoolean("force_tcp", false))
    this->type=Type::TCP_RELAY;
*/

type Config struct {
	WebRtc bool `json:"webrtc"`
	ForceTcp bool `json:"force_tcp"`
	UseTcp bool `json:"use_tcp"`

}

var _config Config

//TODO:(Coder)
func init() {
	_config.WebRtc = false
	_config.UseTcp = true
}

type PhoneCallConfig struct {}

func NewPhoneCallConfig(webRtc bool) *PhoneCallConfig{
	_config.WebRtc = webRtc
	return &PhoneCallConfig{}
}

func (m *PhoneCallConfig) PhoneCallsAvailableConf() bool {
	//TODO:(Coder)
	return true
}

func (m *PhoneCallConfig) PhoneCallsPrivateConf() bool {
	//TODO:(Coder)
	return false
}

func (m *PhoneCallConfig) GetPhoneCallConfig() Config {

	return _config
}

func (m *PhoneCallConfig) GetPhoneCallConfigString() string {
	c, _ := jsoniter.Marshal(_config)
	return string(c)
}

