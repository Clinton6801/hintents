# Erst By Hintents

**Erst** is a premium developer toolset for the Stellar network, designed to provide high-fidelity "glass-box" debugging and simulation for Soroban smart contracts.

> **Status**: Active Development (Phase 4: Advanced Diagnostics)
> **Documentation**: [https://dotandev-hintents-75.mintlify.app/](https://dotandev-hintents-75.mintlify.app/)
> **Focus**: High-Fidelity Simulation, Auth Tracing, and Security Auditing

## Scope & Objective

The primary goal of `erst` is to eliminate the opaque "black box" experience of failed Stellar smart contract transactions. By providing local-first, high-fidelity replay and tracing, `erst` maps generic network errors back to human-readable diagnostic events and source code.

**Core Features (Planned):**

1.  **Transaction Replay**: Fetch a failed transaction's envelope and ledger state from an RPC provider.
2.  **Local Simulation**: Re-execute the transaction logically in a local environment.
3.  **Trace decoding**: Map execution steps and failures back to readable instructions or Rust source lines.
4.  **Source Mapping**: Map WASM instruction failures to specific Rust source code lines using debug symbols.
5.  **GitHub Source Links**: Automatically generate clickable GitHub links to source code locations in traces (when in a Git repository).
6.  **Error Suggestions**: Heuristic-based engine that suggests potential fixes for common Soroban errors.

## Usage (MVP)

### Debugging a Transaction

Fetches a transaction envelope from the Stellar Public network and prints its XDR size (Simulation pending).

```bash
./erst debug <transaction-hash> --network testnet
```

Debug an offline envelope from stdin (no RPC):

```bash
./erst debug < tx.xdr
```

### Interactive Trace Viewer

Launch an interactive terminal UI to explore transaction execution traces with search functionality.

```bash
./erst debug <transaction-hash> --interactive
# or
./erst debug <transaction-hash> -i
```

**Features:**

- **Search**: Press `/` to search through traces (contract IDs, functions, errors)
- **Help overlay**: Press `?` or `h` to see all keyboard shortcuts
- **Tree Navigation**: Expand/collapse nodes, navigate with arrow keys
- **Syntax Highlighting**: Color-coded contract IDs, functions, and errors
- **Fast Navigation**: Jump between search matches with `n`/`N`
- **Match Counter**: See "Match 2 of 5" status while searching

See [internal/trace/README.md](internal/trace/README.md) for detailed documentation.

### Performance Profiling

Generate interactive flamegraphs to visualize CPU and memory consumption during contract execution:

```bash
./erst debug --profile <transaction-hash>
```

This generates an interactive HTML file (`<tx-hash>.flamegraph.html`) with:
- **Hover tooltips** showing frame details (function name, duration, percentage)
- **Click-to-zoom** to focus on specific call stacks
- **Search/highlight** to find frames by name
- **Dark mode support** that adapts to your system theme

**Export Formats:**

```bash
# Interactive HTML (default)
./erst debug --profile --profile-format html <transaction-hash>

# Raw SVG with dark mode support
./erst debug --profile --profile-format svg <transaction-hash>
```

See [docs/INTERACTIVE_FLAMEGRAPH.md](docs/INTERACTIVE_FLAMEGRAPH.md) for detailed documentation and [docs/examples/sample_flamegraph.html](docs/examples/sample_flamegraph.html) for a live demo.

### Audit log signing (software / HSM)

`erst` includes a small utility command to generate a deterministic, signed audit log from a JSON payload.

#### Software signing (Ed25519 private key)

Provide a PKCS#8 PEM Ed25519 private key via env or CLI:

- Env: `ERST_AUDIT_PRIVATE_KEY_PEM`
- CLI: `--software-private-key <pem>`

Example:

```bash
node dist/index.js audit:sign \
  --payload '{"input":{},"state":{},"events":[],"timestamp":"2026-01-01T00:00:00.000Z"}' \
  --software-private-key "$(cat ./ed25519-private-key.pem)"
```

#### PKCS#11 HSM signing

Select the PKCS#11 provider with `--hsm-provider pkcs11` and configure the module/token/key via env vars.

Required env vars:

- `ERST_PKCS11_MODULE` (path to the PKCS#11 module `.so`)
- `ERST_PKCS11_PIN`
- `ERST_PKCS11_KEY_LABEL` **or** `ERST_PKCS11_KEY_ID` (hex)
- `ERST_PKCS11_PUBLIC_KEY_PEM` (SPKI PEM public key for verification/audit metadata)

Optional:

- `ERST_PKCS11_SLOT` (numeric index into the slot list)
- `ERST_PKCS11_TOKEN_LABEL`
- `ERST_PKCS11_PIV_SLOT` (YubiKey PIV slot such as 9a, 9c, 9d, 9e, 82-95, f9)

Example:

```bash
export ERST_PKCS11_MODULE=/usr/lib/softhsm/libsofthsm2.so
export ERST_PKCS11_PIN=1234
export ERST_PKCS11_KEY_LABEL=erst-audit-ed25519
export ERST_PKCS11_PUBLIC_KEY_PEM="$(cat ./ed25519-public-key-spki.pem)"

node dist/index.js audit:sign \
  --hsm-provider pkcs11 \
  --payload '{"input":{},"state":{},"events":[],"timestamp":"2026-01-01T00:00:00.000Z"}'
```

The command prints the signed audit log JSON to stdout so it can be redirected to a file.

For YubiKey PIV (YKCS11) usage details, see [docs/pkcs11-yubikey.md](docs/pkcs11-yubikey.md).

### Protocol Handler

Erst registers a custom `erst://` URI scheme, allowing external tools (browsers, dashboards) to deep-link directly into a debug session.

Register the protocol handler:

```bash
./erst protocol:register
```

Open a debug session via URI:

```bash
./erst protocol:handle "erst://debug/<transaction-hash>?network=testnet"
```

With an optional operation index:

```bash
./erst protocol:handle "erst://debug/<transaction-hash>?network=mainnet&op=0"
```

Unregister the handler when no longer needed:

```bash
./erst protocol:unregister
```

## Documentation

- **[Architecture Overview](docs/architecture.md)**: Deep dive into how the Go CLI communicates with the Rust simulator, including data flow, IPC mechanisms, and design decisions.
- **[Project Proposal](docs/proposal.md)**: Detailed project proposal and roadmap.
- **[Source Mapping](docs/source-mapping.md)**: Implementation details for mapping WASM failures to Rust source code.
- **[Debug Symbols Guide](docs/debug-symbols-guide.md)**: How to compile Soroban contracts with debug symbols.
- **[Error Suggestions](docs/ERROR_SUGGESTIONS.md)**: Heuristic-based error suggestion engine for common Soroban errors.
- **[Canonical JSON Serialization](docs/CANONICAL_JSON.md)**: Deterministic JSON serialization for audit log hashing.
- **[Interactive Trace Showcase](docs/showcase/README.md)**: Try out the interactive trace explorer online.
- **[Time Travel Guide](docs/TIME_TRAVEL_GUIDE.md)**: How to use Magic Rewind to replay transactions across time, save sessions to disk, and share reproducible debug files.

## Technical Analysis

### The Challenge

Stellar's `soroban-env-host` executes WASM. When it traps (crashes), the specific reason is often sanitized or lost in the XDR result to keep the ledger size small.

### The Solution Architecture

`erst` operates by:

1.  **Fetching Data**: Using the Stellar RPC to get the `TransactionEnvelope` and `LedgerFootprint` (read/write set) for the block where the tx failed.
2.  **Simulation Environment**: A Rust binary (`erst-sim`) that integrates with `soroban-env-host` to replay transactions.
3.  **Execution**: Feeding the inputs into the VM and capturing `diagnostic_events`.

For a detailed explanation of the architecture, see [docs/architecture.md](docs/architecture.md).

## Contributors

Thanks goes to these wonderful people:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/dotandev"><img src="https://avatars.githubusercontent.com/u/105521093?v=4" width="100px;" alt="dotdev."/><br /><sub><b>dotdev.</b></sub></a><br /><a href="#code-dotandev" title="Code">Code</a> <a href="#doc-dotandev" title="Documentation">Documentation</a> <a href="#ideas-dotandev" title="Ideas & Planning">Ideas & Planning</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

---

_Erst is an open-source initiative. Contributions, PRs, and Issues are welcome._
