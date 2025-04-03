package hypr

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
)

type client struct {
	connection net.Conn
}

func newClient() (*client, error) {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		runtimeDir = "/tmp"
	}

	hyprlandInstance := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if hyprlandInstance == "" {
		return nil, errors.New("HYPRLAND_INSTANCE_SIGNATURE environment variable not set. Please ensure Hyprland is running")
	}

	socketPath := fmt.Sprintf("%s/hypr/%s/.socket.sock", runtimeDir, hyprlandInstance)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to open hyprland socket: %v", err))
	}

	return &client{
		connection: conn,
	}, nil
}

func (c *client) Close() error {
	return c.connection.Close()
}

// SendRequest sends a low level request directly to the socket. This should only
// be used when no other option is available.
func (c *client) SendRequest(command string) ([]byte, error) {
	written, err := c.connection.Write([]byte(command))
	if err != nil {
		return nil, err
	}

	if written != len(command) {
		return nil, fmt.Errorf("expected to write %d bytes, wrote %d", len(command), written)
	}

	response := bytes.Buffer{}

	buffer := make([]byte, 4096)
	for {
		read, err := c.connection.Read(buffer)
		if err != nil {
			return nil, err
		}

		response.Write(buffer[:read])

		if read != 4096 {
			break
		}
	}

	return response.Bytes(), nil
}

// SendJSONRequest sends a low level request directly to the socket and requests a
// JSON response. This should only be used when no other option is available.
func (c *client) SendJSONRequest(command string) ([]byte, error) {
	return c.SendRequest("j/" + command)
}
