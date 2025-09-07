[![Actions Status](https://github.com/under-tension/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/under-tension/go-project-244/actions) [![linter](https://github.com/under-tension/go-project-244/actions/workflows/linter.yml/badge.svg)](https://github.com/under-tension/go-project-244/actions/workflows/linter.yml) [![Test and build](https://github.com/under-tension/go-project-244/actions/workflows/test-and-build.yml/badge.svg)](https://github.com/under-tension/go-project-244/actions/workflows/test-and-build.yml) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=under-tension_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=under-tension_go-project-244)

[![GO](https://img.shields.io/badge/go-1.24+-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)

> **gendiff** - Compares two files and shows a difference. Supports **yaml** and **json** formats.

## Requirements

- Go 1.24 or higher.

## 🚀 Get started

1. Cloning
```
git clone https://github.com/under-tension/go-project-244.git gendiff
cd ./gendiff
```

2. Building
```
make build
```

3. Running
```
./bin/gendiff <file_1> <file_2>
```

<img src="./docs/assets/gif/demo-install.gif" alt="Video instruction" width="600" />

### Flags

<table>
    <tr>
        <th>Option name</th>
        <th>Alias</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>--help</td>
        <td>-h</td>
        <td>show help info</td>
    </tr>
    <tr>
        <td>--format</td>
        <td>-f</td>
        <td>output format</td>
    </tr>
</table>

## 🏗️ Architecture

### Technology stack
- **Backend**: Go 1.24

### Project structure
```
RetroGame/
├── bin/                    # Binaries
├── cmd/                    # Console applications
├── pkg/                    # Externally accessible packages
│   ├── fabrics/            # Factories for different packages
│   ├── formatters/         # Formatters for different output styles
│   └── parsers/            # Implementation of parsers for various data formats
└── testdata/               # Testing data
```

## 🧪 Testing

```bash
# Run linter check
make lint

# Run unit test
make test
```