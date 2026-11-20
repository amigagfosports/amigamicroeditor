# GFOpenMicro 0.1.0-rc1 receipt

GFOpenMicro packages upstream Micro `3281b6cc74d3` (`2.0.16-dev.106`) for
the GFOS appliance host. Upstream history is preserved through the GitHub fork,
and both the application and port additions are MIT licensed.

## Candidate contents

- static Linux x86-64 binary;
- static Linux ARM64 binary for the Pi appliance;
- `GFOpenMicro` terminal/desktop launcher;
- upstream MIT and third-party notices; and
- reproducible archive timestamps derived from the upstream commit.

This is a host terminal application. It is not an m68k Workbench binary and it
does not claim that the specified ACX86 application personality has landed.

## Current evidence

- upstream Go test suite: passed locally in Go 1.24.13;
- x86-64 archive: built, extracted and launcher/version checked;
- ARM64 archive: built and identified as a static AArch64 ELF;
- both archive checksums: passed; and
- LumaVault CI: required before publication.
