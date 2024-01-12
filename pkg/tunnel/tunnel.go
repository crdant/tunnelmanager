package tunnel

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

type TunnelManager struct {
	host     string
	port     int
	tempdir  string
	socket   string
	proxyURL string
}

func NewTunnelManager(host string, port int) *TunnelManager {
	if port == 0 {
		port = 8080
	}
	tempdir := "/tmp"
	socket := fmt.Sprintf("%s/%s.sock", tempdir, host)

	return &TunnelManager{
		host:    host,
		port:    port,
		tempdir: tempdir,
		socket:  socket,
	}
}

func (t *TunnelManager) Establish(user, credential string) error {
	sshCmd := []string{"ssh", "-q", "-o", "StrictHostKeyChecking=no", "-o", "ServerAliveInterval=60", "-M", "-S", t.socket, "-D", fmt.Sprint(t.port), "-NCf", "-l", user, t.host}

	log.Println(sshCmd)
	cmd := exec.Command(sshCmd[0], sshCmd[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Unable to establish SSH tunnel: %v", err)
	}
	t.proxyURL = fmt.Sprintf("socks5://localhost:%d", t.port)

  return nil
}

func (t *TunnelManager) Teardown() error {
	sshCmd := []string{"ssh", "-S", t.socket, "-O", "exit", t.host}
	log.Println(sshCmd)
	cmd := exec.Command(sshCmd[0], sshCmd[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Unable to teardown SSH tunnel: %v", err)
	}
	t.proxyURL = ""

  return nil
}

func (t *TunnelManager) isKey(credential string) bool {
	checker := regexp.MustCompile("-----BEGIN.*PRIVATE KEY-----")
	return checker.MatchString(credential)
}

func main() {
	// Example usage
	tunnelManager := NewTunnelManager("example.com", 8080)
	tunnelManager.Establish("user", "credential")
	defer tunnelManager.Teardown()
}
