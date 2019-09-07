# u2
double underscore binding

## install

```bash
$ go get -u github.com/luthfikw/u2
```

```go
import "github.com/luthfikw/u2"
```

## usage

### ScanPrefix
get all u2 binding with a prefix

`ScanPrefix(q string, prefixs []string) (items map[string][]string)`


example: [on playground](https://play.golang.org/p/Kg4v2_NVMVz)

```go
q := `
    hi __@name__, how are you today
    i am __:me__, nice to meet you
`

items := u2.ScanPrefix(q, []string{"@", ":"})
fmt.Printf("%+v\n", items)

// output
// map[::[me] @:[name]]
```

### ScanTags

get all u2tag binding

`ScanTags(q string, tag string) (nq string, tags []TagInfo, err error)`

example: [on playground](https://play.golang.org/p/TDRuF1SYK0h)

```go
q := `hi <foo:bar[index]{ meta1: value; meta2: "value 2"; }>hello world!</foo:bar>`

newq, items, err := u2.ScanTags(q, "foo")
if err != nil {
    panic(err)
}

fmt.Printf("%s\n\n%+v\n", newq, items)

// output
// hi __#foo:0__
//
// [{Tag:foo Key:bar Index:index Value:hello world! Meta:map[meta1:value meta2:value 2]}]
```

## contributing

we are open become contributor!
