# Loverage

*coverage made with love, literally*

Loverage is a tool for Golang projects that verbalizes a test coverage basing
on a names of test functions.

## Installation

Install **loverage** using **go get**:

```
go get github.com/kovetskiy/loverage
```

## Usage

Create testcases and declare test functions with following naming convention:

```go
func TestF_B()   { ... }
func TestT_B()   { ... }
func TestT_M_B() { ... }
```

Where, 
- **F** &ndash; tested function
- **T** &ndash; tested type
- **M** &ndash; tested type's method
- **B** &ndash; expected behavior

Examples:

```go
func TestNewLogger_ReturnsLoggerInstance()             { ... }
func TestLogger_ImplementsWriterInterface()            { ... }
func TestLogger_Infof_WritesMessageWithInfoSeverity()  { ... }
```

Spawn **loverage** in a directory with test files and see output

```
$ loverage
NewLogger      Returns Logger Instance
Logger         Implements Writer Interface
Logger.Infof   Writes Message With Info Severity
```

Add loverage's output to your readme file as markdown table:

```
$ loverage --table >> README.md
```

And see resulting table:

| Subject        | Behavior |
| -------------- | -------- |
| NewLogger      | Returns Logger Instance |
| Logger         | Implements Writer Interface |
| Logger.Infof   | Writes Message With Info Severity |

### License
MIT.
