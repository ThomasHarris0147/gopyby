Golang import and use go code in python.
Before everything install:
1. python3 -m pip install pybindgen
2. go install golang.org/x/tools/cmd/goimports@latest
3. go install github.com/go-python/gopy@latest

Then

1. write .go and .mod file
2. `gopy build -output=./output_dir package_name`
3. import from python as `from output_dir import package_name`
