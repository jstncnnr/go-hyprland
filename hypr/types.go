package hypr

type Monitor struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Make             string    `json:"make"`
	Model            string    `json:"model"`
	Serial           string    `json:"serial"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	RefreshRate      float64   `json:"refresh_rate"`
	X                int       `json:"x"`
	Y                int       `json:"y"`
	ActiveWorkspace  Workspace `json:"activeWorkspace"`
	SpecialWorkspace Workspace `json:"specialWorkspace"`
	ReservedSpace    []int     `json:"reserved"`
	Scale            float64   `json:"scale"`
	Transform        int       `json:"transform"`
	Focused          bool      `json:"focused"`
	DpmsStatus       bool      `json:"dpmsStatus"`
	Vrr              bool      `json:"vrr"`
	Solitary         string    `json:"solitary"`
	ActivelyTearing  bool      `json:"activelyTearing"`
	DirectScanoutTo  string    `json:"directScanoutTo"`
	Disabled         bool      `json:"disabled"`
	CurrentFormat    string    `json:"currentFormat"`
	MirrorOf         string    `json:"mirrorOf"`
	AvailableModes   []string  `json:"availableModes"`
}

type Workspace struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Monitor         string `json:"monitor,omitempty"`
	MonitorID       int    `json:"monitorID,omitempty"`
	Windows         int    `json:"windows,omitempty"`
	HasFullscreen   bool   `json:"hasfullscreen,omitempty"`
	LastWindow      string `json:"lastwindow,omitempty"`
	LastWindowTitle string `json:"lastwindowtitle,omitempty"`
	IsPersistent    bool   `json:"ispersistent,omitempty"`
}

type Window struct {
	Address          string    `json:"address"`
	Mapped           bool      `json:"mapped"`
	Hidden           bool      `json:"hidden"`
	At               []int     `json:"at"`
	Size             []int     `json:"size"`
	Workspace        Workspace `json:"workspace"`
	Floating         bool      `json:"floating"`
	PseudoTiled      bool      `json:"pseudo"`
	MonitorID        int       `json:"monitor"`
	Class            string    `json:"class"`
	Title            string    `json:"title"`
	InitialClass     string    `json:"initialClass"`
	InitialTitle     string    `json:"initialTitle"`
	Pid              string    `json:"pid"`
	XWayland         bool      `json:"xwayland"`
	Pinned           bool      `json:"pinned"`
	Fullscreen       int       `json:"fullscreen"`
	FullscreenClient int       `json:"fullscreenClient"`
	Grouped          []string  `json:"grouped"`
	Tags             []string  `json:"tags"`
	Swallowing       string    `json:"swallowing"`
	FocusHistoryID   int       `json:"focusHistoryID"`
	InhibitingIdle   bool      `json:"inhibitingIdle"`
}

type DeviceTable struct {
	Mice      []Mouse    `json:"mice"`
	Keyboards []Keyboard `json:"keyboards"`
	Tablets   []Tablet   `json:"tablets"`
	Touch     []Touch    `json:"touch"`
	Switches  []Switch   `json:"switches"`
}

type HID struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mouse struct {
	HID
	DefaultSpeed float64 `json:"defaultSpeed"`
}

type Keyboard struct {
	HID
	Rules        string `json:"rules"`
	Model        string `json:"model"`
	Layout       string `json:"layout"`
	Variant      string `json:"variant"`
	Options      string `json:"options"`
	ActiveKeymap string `json:"active_keymap"`
	CapsLock     bool   `json:"caps_lock"`
	NumLock      bool   `json:"num_lock"`
	Main         bool   `json:"main"`
}

type Tablet struct {
	HID
	Type      string `json:"type"`
	BelongsTo HID    `json:"belongsTo,omitempty"`
}

type Touch struct {
	HID
}

type Switch struct {
	HID
}
