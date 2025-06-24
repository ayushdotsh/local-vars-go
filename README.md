# Projector - Local Project Variables Manager

A Go-based command-line tool for managing local project variables in a hierarchical directory structure. Projector allows you to store, retrieve, and manage key-value pairs that are associated with specific directories, with automatic inheritance from parent directories.

## Use Case

Projector is designed to solve the problem of managing project-specific configuration and variables across different directories in your workspace. It's particularly useful for:

- **Project Configuration Management**: Store project-specific settings like database URLs, API endpoints, or environment-specific configurations
- **Development Environment Variables**: Manage different settings for development, staging, and production environments within the same project structure
- **Hierarchical Variable Inheritance**: Variables defined in parent directories are automatically available in subdirectories, allowing for cascading configuration
- **Multi-Project Workspaces**: Manage different configurations for multiple projects within the same workspace

### Example Scenarios

1. **Microservices Architecture**: Each service directory can have its own configuration while inheriting common settings from the root
2. **Environment-Specific Settings**: Different directories for dev/staging/prod with appropriate variables
3. **Team Development**: Share common project variables while allowing individual developer overrides

## Features

- 🗂️ **Hierarchical Variable Storage**: Variables are inherited from parent directories
- 🔍 **Smart Lookup**: Automatically searches up the directory tree for variables
- 💾 **JSON Storage**: Variables are stored in a centralized JSON configuration file
- 🎯 **Directory-Specific**: Each directory can have its own set of variables
- 🔧 **Simple CLI**: Easy-to-use command-line interface

## Installation

### Prerequisites

- Go 1.24.3 or later

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/ayushdotsh/local-vars-go.git
cd local-vars-go
```

2. Install dependencies:
```bash
go mod download
```

3. Build the binary:
```bash
go build -o projector cmd/projector/main.go
```

4. (Optional) Install globally:
```bash
# Move to a directory in your PATH
sudo mv projector /usr/local/bin/
```

## Getting Started

### Basic Usage

The projector tool supports three main operations:

#### 1. Print Variables
Display all variables available in the current directory (including inherited ones):
```bash
./projector
```

Display a specific variable:
```bash
./projector [variable-name]
```

#### 2. Add Variables
Add a new key-value pair to the current directory:
```bash
./projector add [key] [value]
```

Example:
```bash
./projector add DATABASE_URL "postgresql://localhost:5432/mydb"
./projector add API_ENDPOINT "https://api.example.com"
```

#### 3. Remove Variables
Remove a variable from the current directory:
```bash
./projector remove [key]
```

Example:
```bash
./projector remove DATABASE_URL
```

### Command Line Options

- `-c, --config`: Specify a custom config file path (default: `~/.config/projector/projector.json`)
- `-p, --pwd`: Specify a custom working directory (default: current directory)

Example with options:
```bash
./projector -c /path/to/custom/config.json -p /path/to/project add PORT 8080
```

### Example Workflow

1. **Initialize a project with common variables**:
```bash
cd /path/to/project
./projector add PROJECT_NAME "MyAwesomeProject"
./projector add LOG_LEVEL "info"
```

2. **Create environment-specific configurations**:
```bash
cd /path/to/project/environments/development
./projector add DATABASE_URL "postgresql://localhost:5432/mydb_dev"
./projector add LOG_LEVEL "debug"  # Override parent value

cd /path/to/project/environments/production
./projector add DATABASE_URL "postgresql://prod-server:5432/mydb_prod"
```

3. **View variables in different contexts**:
```bash
# In development directory - shows inherited + local variables
cd /path/to/project/environments/development
./projector
# Output: PROJECT_NAME=MyAwesomeProject, DATABASE_URL=postgresql://localhost:5432/mydb_dev, LOG_LEVEL=debug

# In production directory
cd /path/to/project/environments/production
./projector
# Output: PROJECT_NAME=MyAwesomeProject, DATABASE_URL=postgresql://prod-server:5432/mydb_prod, LOG_LEVEL=info
```

## Configuration

Projector stores all variables in a JSON configuration file. By default, this file is located at:
- **macOS/Linux**: `~/.config/projector/projector.json`
- **Windows**: `%APPDATA%\projector\projector.json`

### Configuration File Structure

```json
{
  "projector": {
    "/path/to/project": {
      "PROJECT_NAME": "MyAwesomeProject",
      "LOG_LEVEL": "info"
    },
    "/path/to/project/environments/development": {
      "DATABASE_URL": "postgresql://localhost:5432/mydb_dev",
      "LOG_LEVEL": "debug"
    }
  }
}
```

## Development

### Project Structure

```
.
├── cmd/
│   └── projector/
│       └── main.go          # Main entry point
├── pkg/
│   └── projector/
│       ├── config.go        # Configuration management
│       ├── config_test.go   # Configuration tests
│       ├── opts.go          # Command-line options parsing
│       ├── projector.go     # Core projector logic
│       └── projector_test.go # Core logic tests
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
└── README.md               # This file
```

### Running Tests

```bash
go test ./pkg/projector/...
```

### Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o projector-linux cmd/projector/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o projector-windows.exe cmd/projector/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o projector-macos cmd/projector/main.go
```

## Dependencies

- [hellflame/argparse](https://github.com/hellflame/argparse) - Command-line argument parsing

## License

This project is part of a learning exercise from Frontend Masters course on TypeScript, Go, and Rust.

## Contributing

This is a learning project, but feel free to submit issues or pull requests if you find bugs or have suggestions for improvements.
