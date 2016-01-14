# goco
Generate golang code

Generate structs with specific interfaces
If you have ideas as to make the universal it, please write on email egor.workroom@gmail.com

Input params:
* out - strout or file name
* type - type code generate
* config - path to file or config via stdin (see examples)

## Examples

Example generate type **consts**. Config via stdin.
``` bash
$ echo {\"Comment\":\"consts\",\"Consts\":[\"user admin\", \"f\"]} | ./bin/goco -type=consts -out=stdout
// consts
const UserAdmin = "user_admin"
const F = "f"
```

## how to use in your project?

Via go generate tools, follow code
``` go
...
//go:generate goco -type=consts -out=consts.go -config=./config/config.json
...
```

# Supported types

* consts - ...
* file_models - ...

# TODO

* testing all types
* generate code test for generated code
