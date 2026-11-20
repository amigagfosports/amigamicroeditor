# GFOpenMicro

GFOpenMicro packages the MIT-licensed Micro editor for the GFOS appliance host.
It is deliberately a host terminal application, available for Linux x86-64 and
ARM64 (including the Pi 5 appliance). It is not presented as an m68k Workbench
binary and does not depend on the not-yet-landed ACX86 application personality.

Run `GFOpenMicro [file]` from a desktop launcher or shell. From a shell with an
attached terminal it opens in place; from the desktop it tries the system's
configured terminal, then `xterm`, then GNOME Terminal. The `micro` binary can
also be invoked directly.

The editor remains fully offline except when a user explicitly invokes one of
Micro's own network features, such as installing a plug-in.
