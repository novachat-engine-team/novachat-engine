/*
void VoIPController::SendPublicEndpointsRequest(const Endpoint& relay){
	if(!useUDP)
		return;
	LOGD("Sending public endpoints request to %s:%d", relay.address.ToString().c_str(), relay.port);
	publicEndpointsReqTime=GetCurrentTime();
	waitingForRelayPeerInfo=true;
	unsigned char buf[32];
	memcpy(buf, relay.peerTag, 16);
	memset(buf+16, 0xFF, 16);
	NetworkPacket pkt={0};
	pkt.data=buf;
	pkt.length=32;
	pkt.address=(NetworkAddress*)&relay.address;
	pkt.port=relay.port;
	pkt.protocol=PROTO_UDP;
	udpSocket->Send(&pkt);
}
*/