# Fork Notes

This fork adds UCI chess engines built on top of the `chess` library.

## Structure

* **`internal/uciserver/`**: Shared UCI protocol loop (`uciserver.Run`) handling stdin/stdout communication with chess GUIs.
* **`cmd/<engine>/`**: Engine source code (e.g. `cmd/random-engine/`).
* **`engines/`**: Destination directory for compiled engine binaries (ignored by Git).

## Building

```bash
./build.sh                     # Build all engines (incremental / cached)
./build.sh clean               # Rebuild all engines from source (no cache)
./build.sh <engine-name>       # Build a specific engine (e.g. ./build.sh random-engine)
./build.sh clean <engine-name> # Rebuild a specific engine from source
```

## Cute Chess Setup

In **Cute Chess**, add a new engine pointing to:
* **Protocol**: `UCI`
* **Command**: `/Users/stefan/src/Chess/chess/engines/random-engine`
