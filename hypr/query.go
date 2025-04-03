package hypr

import "encoding/json"

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
