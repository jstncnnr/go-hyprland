# go-hyprland
A simple IPC wrapper around the two unix sockets used by Hyprland for events
and control.

These sockets are commonly found at `$XDG_RUNTIME_DIR/hypr/$HYPRLAND_INSTANCE_SIGNATURE/`.
The two sockets used are `.socket.sock` and `.socket2.sock` with the first being
used to issue commands to Hyprland while the second is used to stream events from
Hyprland.
