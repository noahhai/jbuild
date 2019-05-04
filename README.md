## Install

```shell
go get github.com/noahhai/jbuild
```

## Usage

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