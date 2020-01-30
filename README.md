# Go Latest
A simple tool that updates go.mod dependencies to their latest version.

## How to get it?
Simply install it using `go get`:
```
$> go get github.com/matang28/go-latest
```

## How to use it?
`go-latest` takes 2 command-line arguments:  
* Regex expression that matches the module name  
* Root path to scan for `go.mod` files (recursively)
 
So if we have the following go.mod file:
```
module github.com/my/module

require (
    github.com/some/dep1 v0.0.0-20190101-askj298jkhasd
    github.com/some/dep2 v1.2.3
    github.com/some/dep3 v0.1.0-20190101-askj298jkhasd
)

replace github.com/some/dep1 => github.com/some/dep v.1.2.3
```

and you want to upgrade all dependencies that matches `github.com/some/dep1` you'll run the following command:
```
$> go-latest "dep1" .
```

will change the go.mod file to:
```
module github.com/my/module

require (
    github.com/some/dep1 latest
    github.com/some/dep2 v1.2.3
    github.com/some/dep3 v0.1.0-20190101-askj298jkhasd
)

replace github.com/some/dep1 => github.com/some/dep latest
```