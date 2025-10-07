package events

import (
	"fmt"
	"strconv"
	"strings"
)

var parsers = map[string]func([]string) (Event, error){
	"workspace":          parseWorkspaceEvent,
	"workspacev2":        parseWorkspaceV2Event,
	"focusedmon":         parseFocusedMonEvent,
	"focusedmonv2":       parseFocusedMonV2Event,
	"activewindow":       parseActiveWindowEvent,
	"activewindowv2":     parseActiveWindowV2Event,
	"fullscreen":         parseFullscreenEvent,
	"monitorremoved":     parseMonitorRemovedEvent,
	"monitorremovedv2":   parseMOnitorRemovedV2Event,
	"monitoradded":       parseMonitorAddedEvent,
	"monitoraddedv2":     parseMonitorAddedV2Event,
	"createworkspace":    parseCreateWorkspaceEvent,
	"createworkspacev2":  parseCreateWorkspaceV2Event,
	"destroyworkspace":   parseDestroyWorkspaceEvent,
	"destroyworkspacev2": parseDestroyWorkspaceV2Event,
	"moveworkspace":      parseMoveWorkspaceEvent,
	"moveworkspacev2":    parseMoveWorkspaceV2Event,
	"renameworkspace":    parseRenameWorkspaceEvent,
	"activespecial":      parseActiveSpecialEvent,
	"activespecialv2":    parseActiveSpecialV2Event,
	"activelayout":       parseActiveLayoutEvent,
	"openwindow":         parseOpenWindowEvent,
	"closewindow":        parseCloseWindowEvent,
	"movewindow":         parseMoveWindowEvent,
	"movewindowv2":       parseMoveWindowV2Event,
	"openlayer":          parseOpenLayerEvent,
	"closelayer":         parseCloseLayerEvent,
	"submap":             parseSubmapEvent,
	"changefloatingmode": parseChangeFloatingModeEvent,
	"urgent":             parseUrgentEvent,
	"screencast":         parseScreencastEvent,
	"windowtitle":        parseWindowTitleEvent,
	"windowtitlev2":      parseWindowTitleV2Event,
	"togglegroup":        parseToggleGroupEvent,
	"moveintogroup":      parseMoveIntoGroupEvent,
	"moveoutofgroup":     parseMoveOutOfGroupEvent,
	"ignoregrouplock":    parseIgnoreGroupLockEvent,
	"lockgroups":         parseLockGroupsEvent,
	"configreloaded":     parseConfigReloadedEvent,
	"pin":                parsePinEvent,
	"minimized":          parseMinimizedEvent,
	"bell":               parseBellEvent,
}

// Parse will take the raw event string in the format
// {type}>>{arg0},{arg1},...{argN} and return an Event.
//
// Returns UnhandledEvent when the event type is unknown.
// This includes the raw event string for the user to handle.
//
// Returns MalformedEvent when there is an error parsing the event.
// This includes the raw event string and error message for the user.
func Parse(raw string) Event {
	parts := strings.Split(raw, ">>")
	eventName := parts[0]

	parserFunc, ok := parsers[eventName]
	if ok {
		event, err := parserFunc(strings.Split(parts[1], ","))
		if err != nil {
			return MalformedEvent{Raw: raw, Error: err}
		}

		return event
	}

	return UnhandledEvent{
		Raw: raw,
	}
}

func parseWorkspaceEvent(args []string) (Event, error) {
	return WorkspaceEvent{
		WorkspaceName: args[0],
	}, nil
}

func parseWorkspaceV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return WorkspaceV2Event{
		WorkspaceID:   id,
		WorkspaceName: args[1],
	}, nil
}

func parseFocusedMonEvent(args []string) (Event, error) {
	return FocusedMonitorEvent{
		MonitorName:   args[0],
		WorkspaceName: args[1],
	}, nil
}

func parseFocusedMonV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return FocusedMonitorV2Event{
		MonitorName: args[0],
		WorkspaceID: id,
	}, nil
}

func parseActiveWindowEvent(args []string) (Event, error) {
	return ActiveWindowEvent{
		WindowClass: args[0],
		WindowTitle: args[1],
	}, nil
}

func parseActiveWindowV2Event(args []string) (Event, error) {
	return ActiveWindowV2Event{
		WindowAddress: args[0],
	}, nil
}

func parseFullscreenEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing fullscreen state: %v", err)
	}

	return FullscreenEvent{
		FullscreenMode: FullscreenMode(state),
	}, nil
}

func parseMonitorRemovedEvent(args []string) (Event, error) {
	return MonitorRemovedEvent{
		MonitorName: args[0],
	}, nil
}

func parseMOnitorRemovedV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing monitor id: %v", err)
	}

	return MonitorRemovedV2Event{
		MonitorID:          id,
		MonitorName:        args[1],
		MonitorDescription: args[2],
	}, nil
}

func parseMonitorAddedEvent(args []string) (Event, error) {
	return MonitorAddedEvent{
		MonitorName: args[0],
	}, nil
}

func parseMonitorAddedV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing monitor id: %v", err)
	}

	return MonitorAddedV2Event{
		MonitorID:          id,
		MonitorName:        args[1],
		MonitorDescription: args[2],
	}, nil
}

func parseCreateWorkspaceEvent(args []string) (Event, error) {
	return CreateWorkspaceEvent{
		WorkspaceName: args[0],
	}, nil
}

func parseCreateWorkspaceV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return CreateWorkspaceV2Event{
		WorkspaceID:   id,
		WorkspaceName: args[1],
	}, nil
}

func parseDestroyWorkspaceEvent(args []string) (Event, error) {
	return DestroyWorkspaceEvent{
		WorkspaceName: args[0],
	}, nil
}

func parseDestroyWorkspaceV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return DestroyWorkspaceV2Event{
		WorkspaceID:   id,
		WorkspaceName: args[1],
	}, nil
}

func parseMoveWorkspaceEvent(args []string) (Event, error) {
	return MoveWorkspaceEvent{
		WorkspaceName: args[0],
		MonitorName:   args[1],
	}, nil
}

func parseMoveWorkspaceV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return MoveWorkspaceV2Event{
		WorkspaceID:   id,
		WorkspaceName: args[1],
		MonitorName:   args[2],
	}, nil
}

func parseRenameWorkspaceEvent(args []string) (Event, error) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return RenameWorkspaceEvent{
		WorkspaceID:      id,
		NewWorkspaceName: args[1],
	}, nil
}

func parseActiveSpecialEvent(args []string) (Event, error) {
	return ActiveSpecialEvent{
		WorkspaceName: args[0],
		MonitorName:   args[1],
	}, nil
}

func parseActiveSpecialV2Event(args []string) (Event, error) {
	var id int
	id, err := strconv.Atoi(args[0])
	if err != nil && args[0] != "" {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return ActiveSpecialV2Event{
		WorkspaceID:   id,
		WorkspaceName: args[1],
		MonitorName:   args[2],
	}, nil
}

func parseActiveLayoutEvent(args []string) (Event, error) {
	return ActiveLayoutEvent{
		KeyboardName: args[0],
		LayoutName:   args[1],
	}, nil
}

func parseOpenWindowEvent(args []string) (Event, error) {
	return OpenWindowEvent{
		WindowAddress: args[0],
		WorkspaceName: args[1],
		WindowClass:   args[2],
		WindowTitle:   args[3],
	}, nil
}

func parseCloseWindowEvent(args []string) (Event, error) {
	return CloseWindowEvent{
		WindowAddress: args[0],
	}, nil
}

func parseMoveWindowEvent(args []string) (Event, error) {
	return MoveWindowEvent{
		WindowAddress: args[0],
		WorkspaceName: args[1],
	}, nil
}

func parseMoveWindowV2Event(args []string) (Event, error) {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing workspace id: %v", err)
	}

	return MoveWindowV2Event{
		WindowAddress: args[0],
		WorkspaceID:   id,
		WorkspaceName: args[2],
	}, nil
}

func parseOpenLayerEvent(args []string) (Event, error) {
	return OpenLayerEvent{
		Namespace: args[0],
	}, nil
}

func parseCloseLayerEvent(args []string) (Event, error) {
	return CloseLayerEvent{
		Namespace: args[0],
	}, nil
}

func parseSubmapEvent(args []string) (Event, error) {
	return SubmapEvent{
		SubmapName: args[0],
	}, nil
}

func parseChangeFloatingModeEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing floating state: %v", err)
	}

	return ChangeFloatingModeEvent{
		WindowAddress: args[0],
		Floating:      state == 1,
	}, nil
}

func parseUrgentEvent(args []string) (Event, error) {
	return UrgentEvent{
		WindowAddress: args[0],
	}, nil
}

func parseScreencastEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing screencast state: %v", err)
	}

	owner, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing screencast owner: %v", err)
	}

	return ScreencastEvent{
		ScreencastState: state,
		Owner:           ScreencastOwner(owner),
	}, nil
}

func parseWindowTitleEvent(args []string) (Event, error) {
	return WindowTitleEvent{
		WindowAddress: args[0],
	}, nil
}

func parseWindowTitleV2Event(args []string) (Event, error) {
	return WindowTitleV2Event{
		WindowAddress: args[0],
		WindowTitle:   args[1],
	}, nil
}

func parseToggleGroupEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing group state: %v", err)
	}

	return ToggleGroupEvent{
		GroupState:      state,
		WindowAddresses: args[1:],
	}, nil
}

func parseMoveIntoGroupEvent(args []string) (Event, error) {
	return MoveIntoGroupEvent{
		WindowAddress: args[0],
	}, nil
}

func parseMoveOutOfGroupEvent(args []string) (Event, error) {
	return MoveOutOfGroupEvent{
		WindowAddress: args[0],
	}, nil
}

func parseIgnoreGroupLockEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing ignore group lock state: %v", err)
	}

	return IgnoreGroupLockEvent{
		Ignored: state == 1,
	}, nil
}

func parseLockGroupsEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing locked group state: %v", err)
	}

	return LockGroupsEvent{
		Locked: state == 1,
	}, nil
}

func parseConfigReloadedEvent(_ []string) (Event, error) {
	return ConfigReloadEvent{}, nil
}

func parsePinEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing pinned state: %v", err)
	}

	return PinEvent{
		WindowAddress: args[0],
		Pinned:        state == 1,
	}, nil
}

func parseMinimizedEvent(args []string) (Event, error) {
	state, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing minimized state: %v", err)
	}

	return MinimizedEvent{
		WindowAddress: args[0],
		Minimized:     state == 1,
	}, nil
}

func parseBellEvent(args []string) (Event, error) {
	return BellEvent{
		WindowAddress: args[0],
	}, nil
}
