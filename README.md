# TaskTracker

A simple and powerful command-line task management tool built with Go for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## Features

- ✅ Add new tasks
- ✏️ Update task descriptions
- 🗑️ Delete tasks
- 📋 List all tasks or filter by status
- 🔄 Mark task status (in-progress, done)
- 💾 Persistent data storage (JSON format)
- 🆔 Auto-incrementing task IDs

## Project Structure

```plaintext
TaskTracker/
├── cmd/
│   └── main.go          # Main program entry point
├── internal/
│   ├── domain/
│   │   └── task.go      # Domain models
│   ├── model/
│   │   ├── task.go      # Data models
│   │   └── status.go    # Status enums
│   └── service/
│       └── task.go      # Business logic services
├── json/                # JSON data storage directory
├── go.mod
└── README.md
```

## Installation and Setup

### Prerequisites

- Go 1.23.4 or higher

### Clone the Repository

```bash
git clone https://github.com/to404hanga/TaskTracker.git
cd TaskTracker
```

### Build the Program

```bash
cd cmd
go build -o task-cli.exe
```

## Usage

### Basic Commands

#### Add a Task

```bash
./task-cli.exe add "Task 1"
```

#### List Tasks

```bash
./task-cli.exe list
```

#### List Tasks by Status

```bash
# List todo tasks
./task-cli.exe list todo

# List in-progress tasks
./task-cli.exe list in-progress

# List done tasks
./task-cli.exe list done
```

#### Update Task Description

```bash
./task-cli.exe update 1 "Updated Task 1"
```

#### Delete a Task

```bash
./task-cli.exe delete 1
```

#### Mark Task as Done

```bash
./task-cli.exe mark-done 1
```

#### Mark Task as In-Progress

```bash
./task-cli.exe mark-in-progress 1
```

#### Show Help

```bash
./task-cli.exe help
```

### Command Reference

| Command | Syntax | Description |
| --- | --- | --- |
| add | add \<description\> | Add a new task |
| update | update \<id\> \<description\> | Update the description of a task |
| delete | delete \<id\> | Delete a task |
| list | list [status] | List tasks with optional status filter (todo, in-progress, done) |
| mark-in-progress | mark-in-progress \<id\> | Mark a task as in-progress |
| mark-done | mark-done \<id\> | Mark a task as done |
| help | help | Show help for available commands |

## Task Status

Tasks support three statuses:

- **todo**: Pending (default status)
- **in-progress**: Currently being worked on
- **done**: Completed

## Data Storage

- Task data is stored in JSON format in the `json/` directory
- Tasks of different statuses are stored in separate JSON files
- Auto-incrementing ID is stored in `auto-increment.txt` file
- Tasks include creation and update timestamps

## Example Workflow

```bash
# 1. Add some tasks
./task-cli.exe add "Learn Go programming"
./task-cli.exe add "Complete project development"
./task-cli.exe add "Write documentation"

# 2. View all tasks
./task-cli.exe list

# 3. Start working on the first task
./task-cli.exe mark-in-progress 1

# 4. Complete the first task
./task-cli.exe mark-done 1

# 5. Update the second task description
./task-cli.exe update 2 "Complete TaskTracker project development"

# 6. View in-progress tasks
./task-cli.exe list in-progress
```

## Technology Stack

- **Language**: Go 1.23.4
- **Architecture**: Layered achitecture (Domain, Model, Service)
- **Data Storage**: JSON files
- **CLI Parsing**: Native `os.Args`
