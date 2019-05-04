## Install

```shell
go get github.com/noahhai/jbuild
```

## Usage

We can build a map with the two main methods, Add (which accepts value interface{}) and AddMap (value is a JMap).

For each, the first argument is the value (end node) and the variadic args are the sequence of keys (path)  
```go
func (n Jmap) AddMap(node Jmap, path ...string) {}
func (n Jmap) Add(val interface{}, path ...string) {}
```

For example
```go
j := jbuild.Jmap{}
j.Add("B", "A")
j.Add("C", "D")
j.Add("G", "E", "F")
j.AddMap(jbuild.Jmap{"H": "I"}, "E", "F")
```

Creates

```json
{"A" : "B",
 "C" : "D",
 "E" : {
    "F" : {
      "H" : "I"
    }
 }
}
```

And we can merge two maps together. 

If ErrorOnKeyConflict is set to false and there is a conflict, rather than return an error, the key of the second map wins.

```go
j2 := jbuild.Jmap{}
j2.Add("B-new", "A")
j2.Add("L", "E", "G")
j2.AddMap(jbuild.Jmap{"H_2": "I_2"}, "E", "F")

j.Merge(j2, &jbuild.MergeOptions{ErrorOnKeyConflict: errOnConflict})
```

Creates

```json
{"A" : "B",
 "C" : "D",
 "E" : {
    "F" : {
      "H" : "I",
      "H2" : "I2"
    },
    "G" : "L"
 }
}
```