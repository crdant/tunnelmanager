package tunnel

import (
	"testing"
)

// TestNewTunnelManager tests the initialization of TunnelManager.
func TestNewTunnelManager(t *testing.T) {
	host := "localhost"
	port := 22

	tm := NewTunnelManager(host, port)

	if tm.host != host {
		t.Errorf("Expected host %s, got %s", host, tm.host)
	}

	if tm.port != port {
		t.Errorf("Expected port %d, got %d", port, tm.port)
	}
}
