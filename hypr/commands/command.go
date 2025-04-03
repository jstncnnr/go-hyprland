package commands

import (
	"fmt"
	"strings"
)

type Command interface {
	String() string
}

// DispatchCommand issues a dispatch to call a keybind dispatcher with an argument.
// See https://wiki.hyprland.org/Configuring/Dispatchers for a list of dispatchers.
type DispatchCommand struct {
	Dispatcher string
	Args       []string
}

func (cmd DispatchCommand) String() string {
	return fmt.Sprintf("dispatch %s %s", cmd.Dispatcher, strings.Join(cmd.Args, " "))
}

// KeywordCommand issues a keyword to call a config keyword dynamically.
type KeywordCommand struct {
	Command string
}

func (cmd KeywordCommand) String() string {
	return fmt.Sprintf("keyword %s", cmd.Command)
}

// ReloadCommand issues a reload to force reload the config.
type ReloadCommand struct{}

func (cmd ReloadCommand) String() string {
	return "reload"
}

// KillCommand issues a kill to get into a kill mode, where you can kill an app by clicking on it.
// You can exit it with ESCAPE.
//
// Kind of like xkill.
type KillCommand struct{}

func (cmd KillCommand) String() string {
	return "kill"
}

// SetCursorCommand sets the cursor theme and reloads the cursor manager. Will set the theme for
// everything except GTK, because GTK.
//
// Please note that since 0.37.0, this only accepts hyprcursor themes. For legacy xcursor themes,
// use the XCURSOR_THEME and XCURSOR_SIZE env vars.
type SetCursorCommand struct {
	Name string
	Size int
}

func (cmd SetCursorCommand) String() string {
	return fmt.Sprintf("setcursor %s %d", cmd.Name, cmd.Size)
}

type OutputBackend string

const (
	BackendWayland  OutputBackend = "wayland"
	BackendHeadless OutputBackend = "headless"
	BackendAuto     OutputBackend = "auto"
)

// CreateOutputCommand allows you to add and remove fake outputs to your preferred backend.
//
// Where Name is an optional name for the output. If Name is not specified, the default
// naming scheme will be used (HEADLESS-2, WL-1, etc.)
type CreateOutputCommand struct {
	Backend OutputBackend
	Name    string
}

func (cmd CreateOutputCommand) String() string {
	if cmd.Name == "" {
		return fmt.Sprintf("output create %s", cmd.Backend)
	}

	return fmt.Sprintf("output create %s %s", cmd.Backend, cmd.Name)
}

// RemoveOutputCommand allows you to remove a fake output by name.
type RemoveOutputCommand struct {
	Name string
}

func (cmd RemoveOutputCommand) String() string {
	return fmt.Sprintf("output remove %s", cmd.Name)
}

// SwitchXkbLayoutCommand Sets the xkb layout index for a keyboard.
//
// For example, if you set:
//
//	device {
//	   name = my-epic-keyboard-v1
//	   kb_layout = us,pl,de
//	}
//
// You can use this command to switch between them. where Command is either 'next' for next,
// 'prev' for previous, or ID for a specific one (in the above case, us: 0, pl: 1, de: 2).
// You can find the Device using 'hyprctl devices' command.
//
// Device can also be 'current' or 'all', self-explanatory. Current is the main keyboard from devices.
type SwitchXkbLayoutCommand struct {
	Device  string
	Command string
}

func (cmd SwitchXkbLayoutCommand) String() string {
	return fmt.Sprintf("switchxkblayout %s %s", cmd.Device, cmd.Command)
}

// SetErrorCommand sets the hyprctl error string.
// Will reset when Hyprland’s config is reloaded.
type SetErrorCommand struct {
	Color   string
	Message string
}

func (cmd SetErrorCommand) String() string {
	return fmt.Sprintf("seterror '%s' %s", cmd.Color, cmd.Message)
}

// DisableErrorCommand disables the hyprctl error string.
// Will reset when Hyprland's config is reloaded.
type DisableErrorCommand struct{}

func (cmd DisableErrorCommand) String() string {
	return "seterror disable"
}

type NotifyIcon int

const (
	IconNone     NotifyIcon = -1
	IconWarning  NotifyIcon = 0
	IconInfo     NotifyIcon = 1
	IconHint     NotifyIcon = 2
	IconError    NotifyIcon = 3
	IconConfused NotifyIcon = 4
	IconOk       NotifyIcon = 5

	NotifyColorDefault string = "0"
)

// NotifyCommand sends a notification using the built-in Hyprland notification system.
// Color of 0 means “Default color for icon”
//
// Optionally, you can specify a font size of the notification like so:
//
// Message: "fontsize:35 This text is big"
type NotifyCommand struct {
	Icon    NotifyIcon
	Timeout int
	Color   string
	Message string
}

func (cmd NotifyCommand) String() string {
	if cmd.Color == "" || cmd.Color == "0" {
		return fmt.Sprintf("notify %d %d 0 \"%s\"", cmd.Icon, cmd.Timeout, cmd.Message)
	}

	return fmt.Sprintf("notify %d %d \"%s\" \"%s\"", cmd.Icon, cmd.Timeout, cmd.Color, cmd.Message)
}

// DismissNotifyCommand dismisses all or up to Count notifications.
// Count of 0 or -1 will dismiss all notifications
type DismissNotifyCommand struct {
	Count int
}

func (cmd DismissNotifyCommand) String() string {
	count := cmd.Count
	if count == 0 {
		count = -1
	}

	return fmt.Sprintf("dismissnotify %d", count)
}
