# Ubuntu File Explorer (Go + GTK4)

A production-grade Ubuntu file explorer with column view similar to macOS Finder.

## Architecture
- **Core**: Go
- **UI**: GTK4 (Native)
- **Engine**: Async I/O, UDisks2 device monitoring, DBus event bus.

## Features
- Column-based navigation.
- Async file operations (Rename, Copy, Move, Delete, Trash).
- Preview panel (Images, Text).
- Hardware auto-detection (USB, SD, etc.).

## Build
```bash
# Install dependencies (Ubuntu/Debian)
sudo apt install libgtk-4-dev libdbus-1-dev libgirepository1.0-dev pkg-config

make build
```
