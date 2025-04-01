package hypr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
)

type Client struct {
	connection net.Conn
}

func NewClient() (*Client, error) {
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

	return &Client{
		connection: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.connection.Close()
}

// SendRequest sends a low level request directly to the socket. This should only
// be used when no other option is available.
func (c *Client) SendRequest(command string) ([]byte, error) {
	written, err := c.connection.Write([]byte(command))
	if err != nil {
		return nil, err
	}

	if written != len(command) {
		return nil, fmt.Errorf("expected to write %d bytes, wrote %d", len(command), written)
	}

	buffer := make([]byte, 4096)
	read, err := c.connection.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:read], nil
}

// SendJSONRequest sends a low level request directly to the socket and requests a
// JSON response. This should only be used when no other option is available.
func (c *Client) SendJSONRequest(command string) ([]byte, error) {
	return c.SendRequest("j/" + command)
}

func (c *Client) GetMonitors() ([]Monitor, error) {
	resp, err := c.SendJSONRequest("monitors")
	if err != nil {
		return nil, err
	}

	monitors := make([]Monitor, 0)
	err = json.Unmarshal(resp, &monitors)
	if err != nil {
		return nil, err
	}

	return monitors, nil
}

func (c *Client) GetWorkspaces() ([]Workspace, error) {
	resp, err := c.SendJSONRequest("workspaces")
	if err != nil {
		return nil, err
	}

	workspaces := make([]Workspace, 0)
	err = json.Unmarshal(resp, &workspaces)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (c *Client) GetWindows() ([]Window, error) {
	resp, err := c.SendJSONRequest("clients")
	if err != nil {
		return nil, err
	}

	windows := make([]Window, 0)
	err = json.Unmarshal(resp, &windows)
	if err != nil {
		return nil, err
	}

	return windows, nil
}

func (c *Client) GetActiveWorkspace() (*Workspace, error) {
	resp, err := c.SendJSONRequest("activeworkspace")
	if err != nil {
		return nil, err
	}

	workspace := new(Workspace)
	err = json.Unmarshal(resp, workspace)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (c *Client) GetActiveWindow() (*Window, error) {
	resp, err := c.SendJSONRequest("activewindow")
	if err != nil {
		return nil, err
	}

	window := new(Window)
	err = json.Unmarshal(resp, window)
	if err != nil {
		return nil, err
	}

	return window, nil
}

func (c *Client) GetDeviceTable() (*DeviceTable, error) {
	resp, err := c.SendJSONRequest("devices")
	if err != nil {
		return nil, err
	}

	devices := new(DeviceTable)
	err = json.Unmarshal(resp, devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}
