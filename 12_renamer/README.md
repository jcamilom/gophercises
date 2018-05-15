# 12_renamer

A cli tool to change filenames using regular expresions.

## How to use it

```
$ go get -u github.com/jcamilom/gophercises/12_renamer

# run it
$ 12_renamer "^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$" "\$2 - \$1 - \$3 of \$4.\$5"
```

usage: 12_renamer [<path>] \"<match_regexp>\" \"<replace_string>\
