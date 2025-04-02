# go-hyprland
A simple IPC wrapper around the two unix sockets used by Hyprland for events
and control.

These sockets are commonly found at `$XDG_RUNTIME_DIR/hypr/$HYPRLAND_INSTANCE_SIGNATURE/`.
The two sockets used are `.socket.sock` and `.socket2.sock` with the first being
used to issue commands to Hyprland while the second is used to stream events from
Hyprland.

## Event Client
The event client is used to listen to events from Hyprland. The many events can be found:
https://wiki.hyprland.org/IPC/#events-list

It is highly recommended you use a cancelable context so the socket can be cleaned up.

Please see [examples/dynamic-windows/main.go](examples/dynamic-windows/main.go) for a more in-depth example.

```go
import "github.com/jstncnnr/go-hyprland/hypr/events"

client, err := events.NewClient()
if err != nil {
    fmt.Printf("Error creating event client: %v\n", err)
    os.Exit(1)
}

client.RegisterListener(func (event events.Event) {
	//Handle events here
})

if err := client.Listen(context.Background()); err != nil && !errors.Is(err, context.Canceled) {
    fmt.Printf("Error running event client: %v\n", err)
    os.Exit(1)
}
```

## Hyprctl Client
This client is used to issue commands to Hyprland. It functions similarly to `hyprctl` itself.
Most of the useful commands are implemented, however not everything is implemented yet.

Please see [examples/dynamic-windows/main.go](examples/dynamic-windows/main.go) for a more in-depth example.

```go
import "github.com/jstncnnr/go-hyprland/hypr"

workspaces, err := hypr.GetWorkspaces()
if err != nil {
	fmt.Printf("Error requesting workspaces: %v", err)
	os.Exit(1)
}

for _, workspace := range workspaces {
	fmt.Printf("Workspace: \n%v", workspace)
}
```