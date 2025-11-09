[![CI](https://github.com/bitstep-ie/mango-go/actions/workflows/ci.yml/badge.svg)](https://github.com/bitstep-ie/mango-go/actions/workflows/ci.yml)
[![CodeQL](https://github.com/bitstep-ie/mango-go/actions/workflows/codeql.yml/badge.svg)](https://github.com/bitstep-ie/mango-go/actions/workflows/codeql.yml)
[![Dependabot](https://github.com/bitstep-ie/mango-go/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/bitstep-ie/mango-go/actions/workflows/dependabot/dependabot-updates)
[![codecov](https://codecov.io/github/bitstep-ie/mango-go/graph/badge.svg?token=L6EJH29N5L)](https://codecov.io/github/bitstep-ie/mango-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitstep-ie/mango-go)](https://goreportcard.com/report/github.com/bitstep-ie/mango-go)

<br />
<div align="center">
    <a href="https://github.com/bitstep-ie/mango-go">
    <picture>
        <source srcset="images/mango-with-text-black.png" media="(prefers-color-scheme: light)">
        <source srcset="images/mango-with-text-white.png" media="(prefers-color-scheme: dark)">
        <img src="images/mango-with-text-black.png" alt="mango Logo">
    </picture>
    </a>
    <h3 align="center">mango-go</h3>
    <p align="center">
        A collection of utility packages for go
        <br />
        <a href="#"><strong>📚 Explore the docs »</strong></a>
        <br />
        <br />
        <a href="#">🔎 View Examples</a>
        &middot;
        <a href="https://github.com/bitstep-ie/mango-go/issues/new?labels=bug&template=bug-report---.md">
            🐛 Report Bug
        </a>
        &middot;
        <a href="https://github.com/bitstep-ie/mango-go/issues/new?labels=enhancement&template=feature-request---.md">
            💡 Request Feature
        </a>
    </p>
</div>
<br />
<!-- TABLE OF CONTENTS -->
    <details>
      <summary>📂 Table of Contents</summary>
      <ol>
        <li>
            <a href="#about-the-project">📋 About The Project</a>
        </li>
        <li>
            <a href="#getting-started">🚀 Getting Started</a>
            <ul>
                <li><a href="#installation">🛠️ Installation</a></li>
                <li><a href="#packages">📦 Packages</a></li>
                <li><a href="#developer-guide">🧑‍💻 Developer Guide</a></li>
            </ul>
        </li>
        <li>
            <a href="#usage">👨‍💻 Usage</a>
        </li>
        <li>
            <a href="#contributing">📝 Contributing</a>
            <ul>
                <li><a href="#make">🔧 Make</a></li>
                <li><a href="#structure">📐 Structure</a></li>
                <li><a href="#filename-convention">🔖 Filename convention</a></li>
                <li><a href="#gremlins-coverage">👹 Gremlins coverage</a></li>
            </ul>
        </li>
        <li>
            <a href="#license">📜 License</a>
        </li>
        <li>
            <a href="#acknowledgments">👏 Acknowledgments</a>
            <ul>
                <a href="#contributors">🤝 Contributors</a>
            </ul>
        </li>
      </ol>
    </details>



## <a id="about-the-project"></a>📋 About the project

`mango-go` is a grab-bag of small, dependency-light utilities we found ourselves rewriting across services. Every package is:

- **Focused** – each folder solves a single problem (logging, env parsing, random data, etc.).
- **Drop-in** – import paths live under `github.com/bitstep-ie/mango-go/pkg/...`.
- **Well-documented** – every package ships with dedicated docs plus a [developer guide](docs/developer-guide.md) full of copy-paste examples.
- **CI-backed** – linted, tested, and mutation-tested in CI so helpers stay reliable.

## <a id="getting-started"></a>🚀 Getting started

### <a id="installation"></a>🛠️ Installation

```bash
go get github.com/bitstep-ie/mango-go@latest
```

Modules are versioned, so you can pin a specific tag in `go.mod` if required.

### <a id="packages"></a>📦 Packages 

| Package | What it does | Docs |
| --- | --- | --- |
| `env` | read env vars with defaults or panic-on-missing helpers | [docs](docs/mango-go/packages/env.md) |
| `io` | delete/backup/restore files by extension for safe inline edits | [docs](docs/mango-go/packages/io.md) |
| `mango_logger` | opinionated slog handler with CLI/file/syslog outputs | [docs](docs/mango-go/packages/mclogger.md) |
| `random` | math/crypto random helpers for fixtures, passwords, timestamps | [docs](docs/mango-go/packages/random.md) |
| `slices` | generic slice utilities (contains, chunk, unique, etc.) | [docs](docs/mango-go/packages/slices.md) |
| `testutils` | test helpers for temp files and UUID/token assertions | [docs](docs/mango-go/packages/testutils.md) |
| `time` | start/end-of-day helpers, duration parsing, “time ago” strings | [docs](docs/mango-go/packages/time.md) |

Looking for a tour that stitches these together?  
👉 [Developer Guide](docs/developer-guide.md)

### 🔁 Quick start

```go
package main

import (
    "context"
    "log/slog"
    "time"
    mangoenv "github.com/bitstep-ie/mango-go/pkg/env"
    mangolog "github.com/bitstep-ie/mango-go/pkg/mango_logger"
    mangotime "github.com/bitstep-ie/mango-go/pkg/time"
)

func main() {
    cfg := &mangolog.LogConfig{
        MangoConfig: &mangolog.MangoConfig{
            Strict: true,
            CorrelationId: &mangolog.CorrelationIdConfig{AutoGenerate: true},
        },
        Out: &mangolog.OutConfig{
            Enabled: true,
            Cli:   &mangolog.CliConfig{Enabled: true, Friendly: true, Verbose: true},
            File:  &mangolog.FileOutputConfig{Enabled: false},
        },
    }

    logger := slog.New(mangolog.NewMangoLogger(cfg))
    ctx := context.Background()
    ctx = context.WithValue(ctx, mangolog.APPLICATION, "billing-api")
    ctx = context.WithValue(ctx, mangolog.OPERATION, "invoice-create")
    ctx = context.WithValue(ctx, mangolog.TYPE, mangolog.BusinessType)

    timeout := mangoenv.EnvAsInt("HTTP_TIMEOUT", 15)
    deadline := mangotime.TimeAgo(mangotime.EndOfDay(time.Now()))

    logger.InfoContext(ctx, "ready to serve",
        slog.Int("timeoutSeconds", timeout),
        slog.String("deadline", deadline),
    )
}
```

Run the snippet to see CLI-friendly output plus structured JSON (when file logging is enabled).


### <a id="developer-guide"></a>🧑‍💻 Developer Guide

Looking for end-to-end examples that combine logging, environment loading, random data generation, time helpers, and more?  
👉 Jump into [docs/developer-guide.md](docs/developer-guide.md).


## <a id="usage"></a>👨‍💻 Usage

The packages are intentionally orthogonal, so feel free to mix and match:

```go
import (
    "log/slog"
    "time"
    mangorand "github.com/bitstep-ie/mango-go/pkg/random"
    mangoslices "github.com/bitstep-ie/mango-go/pkg/slices"
    mangotime "github.com/bitstep-ie/mango-go/pkg/time"
)

func demo() {
    orders := []int{1, 2, 2, 3}
    if mangoslices.EqualsIgnoreOrder(orders, []int{3, 2, 2, 1}) {
        token := mangorand.Password(20, mangorand.PasswordOptions{Letters: true, Digits: true})
        start := mangotime.StartOfDay(time.Now())
        end := mangotime.EndOfDay(time.Now())

        slog.Info("processing window",
            slog.String("token", token),
            slog.String("start", start.Format(time.RFC3339)),
            slog.String("end", end.Format(time.RFC3339)),
        )
    }
}
```

Check each package doc (table above) for deeper walkthroughs and additional helpers.


## <a id="contributing"></a>📝 Contributing

Do you have a useful function? Do you have something that could be useful to others?

Yes please! [See our starting guidelines](contributing.md). Your help is very welcome!

### <a id="make"></a>🔧 Make

You can use `make all` to ensure all the checks are performed before you push the code on a remote branch and open PR which will execute the github actions.

The makefile is to help with local development of the library by giving you the exact steps that the ci will execute.

This makefile will NOT be used as part of builds.

It is up to you if you deviate from the github actions, do so at your own risk and should not be committed back into the project.

### <a id="structure"></a>📐 Structure

Try to keep common functions in existing packages, and follow the same pattern if required to create a new package. Update documentation as required, and make sure to note any breaking changes clearly in the PR and ofc debate on the mango teams channel if it requires and how to handle version increase.

### <a id="filename-convention"></a>🔖 Filename convention
For basic packages try to match to this convention:
`smallCaseUtils.go`
`smallCaseUtils_test.go`

Larger packages requiring multiple files (e.g: `mclogger`), has no current structure convention.

### <a id="gremlins-coverage"></a>👹 Gremlins coverage

Both `efficacy` & `mutant-coverage` sit at `95%`. Aim for this or higher, the build will fail if these thresholds are not met. Under committee review if necessary these thresholds will be reviewed.


## <a id="license"></a>📜 License


## <a id="acknowledgments"></a>👏 Acknowledgments


### <a id="contributors"></a>🤝 Contributors

Thanks goes to these wonderful people:

<table align="center">
  <tr>
    <td align="center"><a href="https://github.com/Ronan-L-OByrne"><img src="https://github.com/Ronan-L-OByrne.png?size=100" width="100px;" alt="Ronan"/><br /><sub><b>Ronan</b></sub></a></td>
    <td align="center"><a href="https://github.com/bencarroll1"><img src="https://github.com/bencarroll1.png?size=100" width="100px;" alt="Ben"/><br /><sub><b>Ben</b></sub></a></td>
  </tr>
</table>
