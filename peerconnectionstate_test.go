package webrtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPeerConnectionState(t *testing.T) {
	testCases := []struct {
		stateString   string
		expectedState PeerConnectionState
	}{
		{unknownStr, PeerConnectionStateUnknown},
		{"new", PeerConnectionStateNew},
		{"connecting", PeerConnectionStateConnecting},
		{"connected", PeerConnectionStateConnected},
		{"disconnected", PeerConnectionStateDisconnected},
		{"failed", PeerConnectionStateFailed},
		{"closed", PeerConnectionStateClosed},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.expectedState,
			newPeerConnectionState(testCase.stateString),
			"testCase: %d %v", i, testCase,
		)
	}
}

func TestPeerConnectionState_String(t *testing.T) {
	testCases := []struct {
		state          PeerConnectionState
		expectedString string
	}{
		{PeerConnectionStateUnknown, unknownStr},
		{PeerConnectionStateNew, "new"},
		{PeerConnectionStateConnecting, "connecting"},
		{PeerConnectionStateConnected, "connected"},
		{PeerConnectionStateDisconnected, "disconnected"},
		{PeerConnectionStateFailed, "failed"},
		{PeerConnectionStateClosed, "closed"},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.expectedString,
			testCase.state.String(),
			"testCase: %d %v", i, testCase,
		)
	}
}
