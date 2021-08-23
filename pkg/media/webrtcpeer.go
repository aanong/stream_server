package media

import (
	"github.com/pion/webrtc/v2"
	"stream_server/pkg/util"
)

type WebRTCPeer struct {
	 ID         string
	 PC         *webrtc.PeerConnection
	 //视频
	 VideoTrack *webrtc.Track
	 //音频
	 AudioTrack *webrtc.Track
	 //通道停止使用
	 stop chan  int
	 //关键帧丢包重传
	 pli  chan  int
}

func NewWebRTCPeer(id string) *WebRTCPeer {
	return &WebRTCPeer{
		ID: id,
		stop: make(chan int),
		pli: make(chan int),
	}
	
}
// 停止
func (p *WebRTCPeer) Stop()  {
	close(p.stop)
	close(p.pli)
}
//响应发送方
func (p *WebRTCPeer) AnswerSender(offer webrtc.SessionDescription) (answer webrtc.SessionDescription,err error) {

	util.Infof("WebRTC.AnswerSender")
	//TODO

}

//响应接收方
func (p *WebRTCPeer) AnswerReceiver(offer webrtc.SessionDescription,addVideoTrack **webrtc.Track,addAudioTrack **webrtc.Track) (answer webrtc.SessionDescription,err error) {

	util.Infof("WebRTC.AnswerReceiver")
	//TODO

}

func (p *WebRTCPeer) SendPli()  {
	// TODO
}