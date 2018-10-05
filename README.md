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