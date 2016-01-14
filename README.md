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

[more examples](https://github.com/gebv/goco/tree/master/examples)

``` bash
$ ./bin/goco -config=./examples/model_user.json -out=./examples/_model_user.go -type=file_models
```


## how to use in your project?

Via go generate tools, follow code
``` go
...
//go:generate goco -type=consts -out=consts.go -config=./config/config.json
...
```

## Helpful functions

[Extends](EXTENDS.md)

### FromJson(obj interface{}, data interface{}) error

``` golang
func FromJson(obj interface{}, data interface{}) error {
    switch data.(type) {
    case io.Reader:
        decoder := json.NewDecoder(data.(io.Reader))
        return decoder.Decode(obj)
    case []byte:
        return json.Unmarshal(data.([]byte), obj)
    }

    return ErrNotSupported
}
```

Test
``` golang
func TestFromJson(t *testing.T) {
    obj := struct {
        A string `json:"a"`
    }{}

    if err := FromJson(&obj, []byte(`{"a": "b"}`)); err != nil {
        t.Error(err)
    }

    if obj.A != "b" {
        t.Error("not expected value")
    }

    data := bytes.NewReader([]byte(`{"a": "c"}`))

    if err := FromJson(&obj, data); err != nil {
        t.Error(err)
    }

    if obj.A != "c" {
        t.Error("not expected value")
    }
}
```

### ExtractFieldsFromMap(m map[string]interface{}, without ...string) (keys []string, fields []interface{})

``` golang
func ExtractFieldsFromMap(m map[string]interface{}, without ...string) (keys []string, fields []interface{}) {
    _without := make(map[string]bool)

    for _, v := range without {
        _without[v] = true
    }

    for fieldName, field := range m {
        if !_without[fieldName] {
            continue
        }

        keys = append(keys, fieldName)
        fields = append(fields, field)
    }

    return
}
```

# Supported types

* consts - ...
* file_models - ...

# TODO

* testing all types
* generate code test for generated code
