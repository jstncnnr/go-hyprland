package hypr

import (
	"bytes"
	"encoding/json"
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

func GetMonitors() ([]Monitor, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

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

func GetWorkspaces() ([]Workspace, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

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

func GetWindows() ([]Window, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

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

func GetActiveWorkspace() (*Workspace, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

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

func GetActiveWindow() (*Window, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

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

func GetDeviceTable() (*DeviceTable, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

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

// Dispatch issues a dispatch to call a keybind dispatcher.
// See https://wiki.hyprland.org/Configuring/Dispatchers for
// a list of dispatchers.
//
// This is functionally equivalent to running
// hyprctl dispatch {command}
func Dispatch(command string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("dispatch " + command)
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}

// Keyword issues a keyword to call a config keyword dynamically.
//
// This is functionally equivalent to running
// hyprctl keyword {command}
func Keyword(command string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("keyword " + command)
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}

// Reload issues a reloada to force reload the
// config.
//
// This is functionally equivalent to running
// hyprctl reload
func Reload() error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("reload")
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}

// SetCursor sets the cursor theme and reloads the
// cursor manager. Will set the theme for everything
// except GTK, because GTK.
//
// This is functionally equivalent to running
// hyprctl setcursor {command}
func SetCursor(command string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("setcursor " + command)
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}

// SwitchXkbLayout sets the xkb layout index for a keyboard.
//
// cmd is either "next" or "prev" or an index to a layout
// defined in your Hyprland config.
//
// device can either be the name of a device from hyprctl devices
// or "current" or "all". Current is the main keyboard from devices.
//
// This is equivalent to calling hyprctl switchxkblayout {device} {cmd}
func SwitchXkbLayout(device string, cmd string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("switchxkblayout " + device + " " + cmd)
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}

// SetError sets the hyprctl error string. This will reset when
// Hyprland's config is reloaded.
//
// This is equivalent to calling hyprctl seterror {color} {message}
func SetError(color string, message string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("seterror " + color + " " + message)
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}

// DisableError disables the hyprctl error string. This will reset when
// Hyprland's config is reloaded.
//
// This is equivalent to calling hyprctl seterror disable
func DisableError() error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	resp, err := c.SendRequest("seterror disable")
	if err != nil {
		return err
	}

	if string(resp) != "ok" {
		return errors.New(string(resp))
	}

	return nil
}
