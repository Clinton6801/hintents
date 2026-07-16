# Interactive Trace Viewer

The interactive trace viewer provides a powerful search and navigation interface for exploring Stellar Soroban transaction execution traces.

> **Full Documentation**: [https://dotandev-hintents-75.mintlify.app/](https://dotandev-hintents-75.mintlify.app/)

## Features

### Search Functionality

- **Case-insensitive search** by default
- **Highlights all matches** in yellow
- **Current match highlighted** in green with arrow indicator
- **Search across all fields**: contract IDs, function names, errors, events, and types
- **Match counter**: Shows "Match X of Y" status
- **Quick navigation**: Jump between matches with `n` and `N`

### Tree Navigation

- **Hierarchical view** of execution trace
- **Expand/collapse** nodes with Enter/Space
- **Expand all** with `e`
- **Collapse all** with `c`
- **Smooth scrolling** with arrow keys, PgUp/PgDn, Home/End
- **Visual indicators** for expanded (v) and collapsed (>) nodes

### Visual Styling

- **Color-coded elements**:
  - Contract IDs in cyan
  - Function names in blue
  - Errors in red
  - Search matches in yellow
  - Current match in green
- **Depth indentation** for clear hierarchy
- **Cursor indicator** (>) for current position

## Usage

### Launching the Viewer

```bash
# Debug a transaction with interactive viewer
./erst debug <transaction-hash> --interactive

# Or use the short flag
./erst debug <transaction-hash> -i
```

### Keyboard Shortcuts

#### Navigation

| Key               | Action                 |
| ----------------- | ---------------------- |
| `↑` / `k`         | Move up                |
| `↓` / `j`         | Move down              |
| `PgUp`            | Scroll up one page     |
| `PgDn`            | Scroll down one page   |
| `Home` / `g`      | Jump to start          |
| `End` / `G`       | Jump to end            |
| `Enter` / `Space` | Toggle expand/collapse |

#### Search

| Key     | Action                      |
| ------- | --------------------------- |
| `/`     | Start search                |
| `Enter` | Execute search              |
| `n`     | Next match                  |
| `N`     | Previous match              |
| `ESC`   | Clear search / Cancel input |

#### Tree Operations

| Key | Action             |
| --- | ------------------ |
| `e` | Expand all nodes   |
| `c` | Collapse all nodes |

#### Other

| Key            | Action      |
| -------------- | ----------- |
| `q` / `Ctrl+C` | Quit viewer |
| `?` / `h`      | Show shortcuts help |

## Search Examples

### Search for Contract ID

```
Press: /
Type: CDLZFC3
Press: Enter
```

Finds all nodes containing that contract ID prefix.

### Search for Error Messages

```
Press: /
Type: insufficient
Press: Enter
```

Finds all error messages containing "insufficient" (case-insensitive).

### Search for Function Names

```
Press: /
Type: transfer
Press: Enter
```

Finds all function calls named "transfer".

### Navigate Between Matches

After searching:

- Press `n` to jump to next match
- Press `N` to jump to previous match
- Navigation wraps around (last → first → second...)

## License

Part of the Erst project - see main LICENSE file.
