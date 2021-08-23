package media

import (
	"github.com/pion/webrtc/v2"
	"stream_server/pkg/util"
)


var (
	webRTCengine *WebRTCEngine
)

//实例化
func init()  {
	webRTCengine = NewWebRTCEngine()
}

var defaultPeerCfg =webrtc.Configuration{
	ICEServers : []webrtc.ICEServer{
		{
			URLs:[] string {"stun:stun.stunprotocol.org:3478"},
		},
	},
}

const (
	//一个媒体传送单元是1400 分成7个包 每侦所需要的RTP包的个数
	averageRtpPacketsPerFrame = 7
)

type WebRTCEngine struct {
	cfg webrtc.Configuration

	mediaEngine webrtc.MediaEngine

	api *webrtc.API
}

func NewWebRTCEngine() *WebRTCEngine  {
	urls:= []string{}
	w :=&WebRTCEngine{
		mediaEngine: webrtc.MediaEngine{},
		cfg: webrtc.Configuration{
			SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
			ICEServers: []webrtc.ICEServer{
				{
					URLs: urls,
				},
			},
		},
	}
	w.mediaEngine.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8,90000))
	w.mediaEngine.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus,48000))
	w.api = webrtc.NewAPI(webrtc.WithMediaEngine(w.mediaEngine))
	return w

}
//创建发送数据对象 朝接收者发送数据
func (s WebRTCEngine) CreateSender(offer webrtc.SessionDescription, pc **webrtc.PeerConnection, addVideoTrack, addAudioTrack **webrtc.Track, stop chan int) (answer webrtc.SessionDescription, err error) {

	*pc, err = s.api.NewPeerConnection(s.cfg)
	util.Infof("WebRTCEngine.CreateSender pc=%p", *pc)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}

	if *addVideoTrack != nil && *addAudioTrack != nil {
		(*pc).AddTrack(*addVideoTrack)
		(*pc).AddTrack(*addAudioTrack)
		err = (*pc).SetRemoteDescription(offer)
		if err != nil {
			return webrtc.SessionDescription{}, err
		}
	}

	//创建应答Answer
	answer, err = (*pc).CreateAnswer(nil)
	//设置本地SDP
	err = (*pc).SetLocalDescription(answer)
	util.Infof("WebRTCEngine.CreateReceiver ok")
	return answer, err

}

//创建接收者对象
func (s WebRTCEngine) CreateReceiver(offer webrtc.SessionDescription, pc **webrtc.PeerConnection, videoTrack, audioTrack **webrtc.Track, stop chan int, pli chan int) (answer webrtc.SessionDescription, err error) {

	*pc, err = s.api.NewPeerConnection(s.cfg)
	util.Infof("WebRTCEngine.CreateReceiver pc=%p", *pc)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}

	_, err = (*pc).AddTransceiver(webrtc.RTPCodecTypeVideo)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}

	_, err = (*pc).AddTransceiver(webrtc.RTPCodecTypeAudio)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}

	//监听OnTrack事件
	(*pc).OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {



	})

	//设置远端SDP
	err = (*pc).SetRemoteDescription(offer)
	if err != nil {
		return webrtc.SessionDescription{}, err
	}

	//创建应答Answer
	answer, err = (*pc).CreateAnswer(nil)
	//设置本地SDP
	err = (*pc).SetLocalDescription(answer)
	util.Infof("WebRTCEngine.CreateReceiver ok")
	return answer, err

}
