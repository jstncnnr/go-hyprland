package hypr

import (
	"errors"
	"fmt"
	"github.com/jstncnnr/go-hyprland/hypr/commands"
	"strings"
)

type Request struct {
	commands []commands.Command
}

func NewRequest() *Request {
	return &Request{
		commands: make([]commands.Command, 0),
	}
}

func (req *Request) AddCommand(command commands.Command) *Request {
	req.commands = append(req.commands, command)
	return req
}

// Dispatch issues a dispatch to call a keybind dispatcher with an argument.
// See https://wiki.hyprland.org/Configuring/Dispatchers for a list of dispatchers.
func (req *Request) Dispatch(dispatcher string, args ...string) *Request {
	if len(args) == 0 {
		args = []string{"unused"}
	}

	return req.AddCommand(&commands.DispatchCommand{
		Dispatcher: dispatcher,
		Args:       args,
	})
}

// Keyword issues a keyword to call a config keyword dynamically.
// See https://wiki.hyprland.org/Configuring/Keywords/ for more information.
func (req *Request) Keyword(command string) *Request {
	return req.AddCommand(&commands.KeywordCommand{
		Command: command,
	})
}

// Reload issues a reload to force reload the config.
func (req *Request) Reload() *Request {
	return req.AddCommand(&commands.ReloadCommand{})
}

// Kill issues a kill to get into a kill mode, where you can kill an app by clicking on it.
// You can exit it with ESCAPE.
//
// Kind of like xkill.
func (req *Request) Kill() *Request {
	return req.AddCommand(&commands.KillCommand{})
}

// SetCursor sets the cursor theme and reloads the cursor manager. Will set the theme for
// everything except GTK, because GTK.
//
// Please note that since 0.37.0, this only accepts hyprcursor themes. For legacy xcursor themes,
// use the XCURSOR_THEME and XCURSOR_SIZE env vars.
func (req *Request) SetCursor(name string, size int) *Request {
	return req.AddCommand(&commands.SetCursorCommand{
		Name: name,
		Size: size,
	})
}

// CreateOutput allows you to add and remove fake outputs to your preferred backend.
//
// Where Name is an optional name for the output. If Name is not specified, the default
// naming scheme will be used (HEADLESS-2, WL-1, etc.)
func (req *Request) CreateOutput(backend commands.OutputBackend, name string) *Request {
	return req.AddCommand(&commands.CreateOutputCommand{
		Backend: backend,
		Name:    name,
	})
}

// RemoveOutput allows you to remove a fake output by name.
func (req *Request) RemoveOutput(name string) *Request {
	return req.AddCommand(&commands.RemoveOutputCommand{
		Name: name,
	})
}

// SwitchXkbLayout Sets the xkb layout index for a keyboard.
//
// For example, if you set:
//
//	device {
//	   name = my-epic-keyboard-v1
//	   kb_layout = us,pl,de
//	}
//
// You can use this command to switch between them. where command is either 'next' for next,
// 'prev' for previous, or ID for a specific one (in the above case, us: 0, pl: 1, de: 2).
// You can find the device using 'hyprctl devices' command.
//
// Device can also be 'current' or 'all', self-explanatory. Current is the main keyboard from devices.
func (req *Request) SwitchXkbLayout(device string, command string) *Request {
	return req.AddCommand(&commands.SwitchXkbLayoutCommand{
		Device:  device,
		Command: command,
	})
}

// SetError sets the hyprctl error string.
// Will reset when Hyprland’s config is reloaded.
func (req *Request) SetError(color string, message string) *Request {
	return req.AddCommand(commands.SetErrorCommand{
		Color:   color,
		Message: message,
	})
}

// DisableError disables the hyprctl error string.
// Will reset when Hyprland's config is reloaded.
func (req *Request) DisableError() *Request {
	return req.AddCommand(commands.DisableErrorCommand{})
}

// Notify sends a notification using the built-in Hyprland notification system.
// Color of "0" or "" means “Default color for icon”
//
// Optionally, you can specify a font size of the notification like so:
//
// Message: "fontsize:35 This text is big"
func (req *Request) Notify(icon commands.NotifyIcon, timeout int, color string, message string) *Request {
	return req.AddCommand(commands.NotifyCommand{
		Icon:    icon,
		Timeout: timeout,
		Color:   color,
		Message: message,
	})
}

// DismissNotify dismisses all or up to Count notifications.
// Count of 0 or -1 will dismiss all notifications
func (req *Request) DismissNotify(count int) *Request {
	return req.AddCommand(commands.DismissNotifyCommand{
		Count: count,
	})
}

func (req *Request) Send() error {
	c, err := newClient()
	if err != nil {
		return err
	}

	defer func(c *client) {
		_ = c.Close()
	}(c)

	if len(req.commands) == 0 {
		return errors.New("request has no commands")
	}

	var request = ""
	if len(req.commands) > 1 {
		request += "[[BATCH]]"
	}

	for index, command := range req.commands {
		if index > 0 {
			request += " ; "
		}

		request += command.String()
	}

	resp, err := c.SendRequest(request)
	if err != nil {
		return err
	}

	err = nil
	responses := strings.Split(string(resp), "\n\n\n")
	for index, response := range responses {
		if response != "ok" {
			if err == nil {
				err = fmt.Errorf("error running command %T: %v\n", req.commands[index], response)
			} else {
				err = fmt.Errorf("%verror running command %T: %v\n", err, req.commands[index], response)
			}
		}
	}

	return err
}
