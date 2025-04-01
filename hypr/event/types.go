package events

type Event interface{}

// UnhandledEvent Pass the raw event along in case the user wants to implement it.
type UnhandledEvent struct {
	Raw string
}

// MalformedEvent Pass the raw event along and the error message in case the user
// wants to handle it.
type MalformedEvent struct {
	Raw   string
	Error error
}

// FullscreenMode represents if we are entering or exiting fullscreen mode.
type FullscreenMode int

// ScreencastOwner represents if the screencast event originated from a monitor
// share or a window share.
type ScreencastOwner int

const (
	FullscreenExit    FullscreenMode  = 0
	FullscreenEnter   FullscreenMode  = 1
	ScreencastMonitor ScreencastOwner = 0
	ScreencastWindow  ScreencastOwner = 1
)

// WorkspaceEvent emitted on workspace change. Is emitted ONLY when a user requests
// a workspace change, and is not emitted on mouse movements (see focusedmon).
type WorkspaceEvent struct {
	WorkspaceName string
}

// WorkspaceV2Event emitted on workspace change. Is emitted ONLY when a user requests
// a workspace change, and is not emitted on mouse movements (see focusedmon).
type WorkspaceV2Event struct {
	WorkspaceID   int
	WorkspaceName string
}

// FocusedMonitorEvent emitted on the active monitor being changed.
type FocusedMonitorEvent struct {
	MonitorName   string
	WorkspaceName string
}

// FocusedMonitorV2Event emitted on the active monitor being changed.
type FocusedMonitorV2Event struct {
	MonitorName string
	WorkspaceID int
}

// ActiveWindowEvent emitted on the active window being changed.
type ActiveWindowEvent struct {
	WindowClass string
	WindowTitle string
}

// ActiveWindowV2Event emitted on the active window being changed.
type ActiveWindowV2Event struct {
	WindowAddress string
}

// FullscreenEvent emitted when a fullscreen status of a window changes.
//
// A fullscreen event is not guaranteed to fire on/off once in succession.
// Some windows may fire multiple requests to be fullscreened, resulting in
// multiple fullscreen events.
type FullscreenEvent struct {
	FullscreenMode FullscreenMode
}

// MonitorRemovedEvent emitted when a monitor is removed (disconnected).
type MonitorRemovedEvent struct {
	MonitorName string
}

// MonitorAddedEvent emitted when a monitor is added (connected).
type MonitorAddedEvent struct {
	MonitorName string
}

// MonitorAddedV2Event emitted when a monitor is added (connected).
type MonitorAddedV2Event struct {
	MonitorID          int
	MonitorName        string
	MonitorDescription string
}

// CreateWorkspaceEvent emitted when a workspace is created.
type CreateWorkspaceEvent struct {
	WorkspaceName string
}

// CreateWorkspaceV2Event emitted when a workspace is created.
type CreateWorkspaceV2Event struct {
	WorkspaceID   int
	WorkspaceName string
}

// DestroyWorkspaceEvent emitted when a workspace is destroyed.
type DestroyWorkspaceEvent struct {
	WorkspaceName string
}

// DestroyWorkspaceV2Event emitted when a workspace is destroyed.
type DestroyWorkspaceV2Event struct {
	WorkspaceID   int
	WorkspaceName string
}

// MoveWorkspaceEvent emitted when a workspace is moved to a different monitor.
type MoveWorkspaceEvent struct {
	WorkspaceName string
	MonitorName   string
}

// MoveWorkspaceV2Event emitted when a workspace is moved to a different monitor.
type MoveWorkspaceV2Event struct {
	WorkspaceID   int
	WorkspaceName string
	MonitorName   string
}

// RenameWorkspaceEvent emitted when a workspace is renamed.
type RenameWorkspaceEvent struct {
	WorkspaceID      int
	NewWorkspaceName string
}

// ActiveSpecialEvent emitted when the special workspace opened in a monitor changes
// (closing results in an empty WorkspaceName).
type ActiveSpecialEvent struct {
	WorkspaceName string
	MonitorName   string
}

// ActiveSpecialV2Event emitted when the special workspace opened in a monitor changes
// (closing results in empty WorkspaceID and WorkspaceName values).
type ActiveSpecialV2Event struct {
	WorkspaceID   int
	WorkspaceName string
	MonitorName   string
}

// ActiveLayoutEvent emitted on a layout change of the active keyboard.
type ActiveLayoutEvent struct {
	KeyboardName string
	LayoutName   string
}

// OpenWindowEvent emitted when a window is opened.
type OpenWindowEvent struct {
	WindowAddress string
	WorkspaceName string
	WindowClass   string
	WindowTitle   string
}

// CloseWindowEvent emitted when a window is closed.
type CloseWindowEvent struct {
	WindowAddress string
}

// MoveWindowEvent emitted when a window is moved to a workspace.
type MoveWindowEvent struct {
	WindowAddress string
	WorkspaceName string
}

// MoveWindowV2Event emitted when a window is moved to a workspace.
type MoveWindowV2Event struct {
	WindowAddress string
	WorkspaceID   int
	WorkspaceName string
}

// OpenLayerEvent emitted when a layerSurface is mapped.
type OpenLayerEvent struct {
	Namespace string
}

// CloseLayerEvent emitted when a layerSurface is unmapped.
type CloseLayerEvent struct {
	Namespace string
}

// SubmapEvent emitted when a keybind submap changes.
// Empty means default.
type SubmapEvent struct {
	SubmapName string
}

// ChangeFloatingModeEvent emitted when a window changes its floating mode.
type ChangeFloatingModeEvent struct {
	WindowAddress string
	Floating      bool
}

// UrgentEvent emitted when a window requests an urgent state.
type UrgentEvent struct {
	WindowAddress string
}

// ScreencastEvent emitted when a screencopy state of a client changes.
// Keep in mind there might be multiple separate clients. ScreencastState is 0/1,
// Owner is 0 - ScreencastMonitor, 1 - ScreencastWindow
type ScreencastEvent struct {
	ScreencastState int
	Owner           ScreencastOwner
}

// WindowTitleEvent emitted when a window title changes.
type WindowTitleEvent struct {
	WindowAddress string
}

// WindowTitleV2Event emitted when a window title changes.
type WindowTitleV2Event struct {
	WindowAddress string
	WindowTitle   string
}

// ToggleGroupEvent emitted when togglegroup command is used.
// The GroupState is a toggle status where 0 means the group has been destroyed.
type ToggleGroupEvent struct {
	GroupState      int
	WindowAddresses []string
}

// MoveIntoGroupEvent emitted when the window is merged into a group.
type MoveIntoGroupEvent struct {
	WindowAddress string
}

// MoveOutOfGroupEvent emitted when the window is removed from a group.
type MoveOutOfGroupEvent struct {
	WindowAddress string
}

// IgnoreGroupLockEvent emitted when ignoregrouplock is toggled.
type IgnoreGroupLockEvent struct {
	Ignored bool
}

// LockGroupsEvent emitted when lockgroups is toggled.
type LockGroupsEvent struct {
	Locked bool
}

// ConfigReloadEvent emitted when the config is done reloading.
type ConfigReloadEvent struct{}

// PinEvent emitted when a window is pinned or unpinned.
type PinEvent struct {
	WindowAddress string
	Pinned        bool
}

// MinimizedEvent emitted when an external taskbar-like app requests
// a window to be minimized.
type MinimizedEvent struct {
	WindowAddress string
	Minimized     bool
}
