package events

import (
	"reflect"
	"testing"
)

var eventTests = map[string]Event{
	"workspace>>test":                           WorkspaceEvent{WorkspaceName: "test"},
	"workspacev2>>1,test":                       WorkspaceV2Event{WorkspaceID: 1, WorkspaceName: "test"},
	"focusedmon>>DP-1,test":                     FocusedMonitorEvent{MonitorName: "DP-1", WorkspaceName: "test"},
	"focusedmonv2>>DP-1,1":                      FocusedMonitorV2Event{MonitorName: "DP-1", WorkspaceID: 1},
	"activewindow>>class,Title":                 ActiveWindowEvent{WindowClass: "class", WindowTitle: "Title"},
	"activewindowv2>>62c8246947c0":              ActiveWindowV2Event{WindowAddress: "62c8246947c0"},
	"fullscreen>>0":                             FullscreenEvent{FullscreenMode: FullscreenExit},
	"monitorremoved>>DP-1":                      MonitorRemovedEvent{MonitorName: "DP-1"},
	"monitorremovedv2>>1,DP-1,Description":      MonitorRemovedV2Event{MonitorID: 1, MonitorName: "DP-1", MonitorDescription: "Description"},
	"monitoradded>>DP-1":                        MonitorAddedEvent{MonitorName: "DP-1"},
	"monitoraddedv2>>1,DP-1,Description":        MonitorAddedV2Event{MonitorID: 1, MonitorName: "DP-1", MonitorDescription: "Description"},
	"createworkspace>>test":                     CreateWorkspaceEvent{WorkspaceName: "test"},
	"createworkspacev2>>1,test":                 CreateWorkspaceV2Event{WorkspaceID: 1, WorkspaceName: "test"},
	"moveworkspace>>test,DP-1":                  MoveWorkspaceEvent{WorkspaceName: "test", MonitorName: "DP-1"},
	"moveworkspacev2>>1,test,DP-1":              MoveWorkspaceV2Event{WorkspaceID: 1, WorkspaceName: "test", MonitorName: "DP-1"},
	"renameworkspace>>1,test":                   RenameWorkspaceEvent{WorkspaceID: 1, NewWorkspaceName: "test"},
	"activespecial>>test,DP-1":                  ActiveSpecialEvent{WorkspaceName: "test", MonitorName: "DP-1"},
	"activespecialv2>>1,test,DP-1":              ActiveSpecialV2Event{WorkspaceID: 1, WorkspaceName: "test", MonitorName: "DP-1"},
	"activelayout>>keyboard,us":                 ActiveLayoutEvent{KeyboardName: "keyboard", LayoutName: "us"},
	"openwindow>>62c8246947c0,test,class,Title": OpenWindowEvent{WindowAddress: "62c8246947c0", WorkspaceName: "test", WindowClass: "class", WindowTitle: "Title"},
	"closewindow>>62c8246947c0":                 CloseWindowEvent{WindowAddress: "62c8246947c0"},
	"movewindow>>62c8246947c0,test":             MoveWindowEvent{WindowAddress: "62c8246947c0", WorkspaceName: "test"},
	"movewindowv2>>62c8246947c0,1,test":         MoveWindowV2Event{WindowAddress: "62c8246947c0", WorkspaceID: 1, WorkspaceName: "test"},
	"openlayer>>namespace":                      OpenLayerEvent{Namespace: "namespace"},
	"closelayer>>namespace":                     CloseLayerEvent{Namespace: "namespace"},
	"submap>>name":                              SubmapEvent{SubmapName: "name"},
	"changefloatingmode>>62c8246947c0,1":        ChangeFloatingModeEvent{WindowAddress: "62c8246947c0", Floating: true},
	"urgent>>62c8246947c0":                      UrgentEvent{WindowAddress: "62c8246947c0"},
	"screencast>>0,0":                           ScreencastEvent{ScreencastState: 0, Owner: ScreencastMonitor},
	"windowtitle>>62c8246947c0":                 WindowTitleEvent{WindowAddress: "62c8246947c0"},
	"windowtitlev2>>62c8246947c0,Title":         WindowTitleV2Event{WindowAddress: "62c8246947c0", WindowTitle: "Title"},
	"moveintogroup>>62c8246947c0":               MoveIntoGroupEvent{WindowAddress: "62c8246947c0"},
	"moveoutofgroup>>62c8246947c0":              MoveOutOfGroupEvent{WindowAddress: "62c8246947c0"},
	"ignoregrouplock>>1":                        IgnoreGroupLockEvent{Ignored: true},
	"lockgroups>>1":                             LockGroupsEvent{Locked: true},
	"configreloaded>>":                          ConfigReloadEvent{},
	"pin>>62c8246947c0,1":                       PinEvent{WindowAddress: "62c8246947c0", Pinned: true},
	"minimized>>62c8246947c0,0":                 MinimizedEvent{WindowAddress: "62c8246947c0", Minimized: false},
	"bell>>62c8246947c0":                        BellEvent{WindowAddress: "62c8246947c0"},
	"bell>>":                                    BellEvent{WindowAddress: ""},
}

func TestParseWithValidInput(t *testing.T) {
	for input, expected := range eventTests {
		result := Parse(input)

		if result != expected {
			t.Errorf("Parse(%q): expected %v, got %v", input, expected, result)
		}
	}
}

func TestParseToggleGroupEvent(t *testing.T) {
	result := Parse("togglegroup>>0,62c8246947c0,62c8246947c1")

	if reflect.TypeOf(result) != reflect.TypeOf(ToggleGroupEvent{}) {
		t.Errorf("Did not receive ToggleGroupEvent")
	}

	event := result.(ToggleGroupEvent)
	if event.GroupState != 0 {
		t.Errorf("Did not parse GroupState correctly")
	}

	addresses := []string{"62c8246947c0", "62c8246947c1"}
	for index, _ := range addresses {
		if event.WindowAddresses[index] != addresses[index] {
			t.Errorf("Did not parse WindowAddresses correctly")
		}
	}
}

func TestUnhandledEvent(t *testing.T) {
	result := Parse("unhandled>>")
	if reflect.TypeOf(result) != reflect.TypeOf(UnhandledEvent{}) {
		t.Errorf("Did not receive UnhandledEvent")
	}
}

func TestMalformedEvent(t *testing.T) {
	result := Parse("workspacev2>>test,1")
	if reflect.TypeOf(result) != reflect.TypeOf(MalformedEvent{}) {
		t.Errorf("Did not receive MalformedEvent")
	}
}
