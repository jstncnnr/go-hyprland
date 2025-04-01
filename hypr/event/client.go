package events

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Listener func(Event)

type Client struct {
	connection net.Conn
	listeners  []Listener
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

	socketPath := fmt.Sprintf("%s/hypr/%s/.socket2.sock", runtimeDir, hyprlandInstance)
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to open hyprland socket: %v", err))
	}

	return &Client{
		connection: conn,
		listeners:  make([]Listener, 0),
	}, nil
}

func (c *Client) RegisterListener(listener Listener) {
	c.listeners = append(c.listeners, listener)
}

func (c *Client) Listen(ctx context.Context) error {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
			// We set a read timeout so we have a chance to check for context cancellation
			// and shutdown gracefully
			_ = c.connection.SetReadDeadline(time.Now().Add(time.Second))

			read, err := c.connection.Read(buffer)
			if err != nil {
				var netErr net.Error
				if errors.As(err, &netErr) && netErr.Timeout() {
					// Just continue the loop if we timeout on a read
					continue
				}

				return err
			}

			events := make([]Event, 0)
			for _, event := range strings.Split(string(buffer[:read]), "\n") {
				// The last event has a trailing \n which will give us an empty string to skip
				if event == "" {
					continue
				}

				events = append(events, Parse(event))
			}

			for _, event := range events {
				for _, listener := range c.listeners {
					listener(event)
				}
			}
		}
	}
}

func (c *Client) Close() error {
	return c.connection.Close()
}
