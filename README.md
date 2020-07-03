# gotestfile.exe

GOTESTFILE.exe reads the tests from a series of test files and runs just those tests.  It simulates the behavior of VSCode's "Run File Tests" command.

## Usage
`gotestfile -debug -env ABC path/file_test.go [...files]`

* `-debug` prints debug information and runs the tests in verbose mode
* `-env` will print all environment variables with that prefix before the test.

## Notes

* It has only been tested on Windws but it should work on other operating systems.
* It runs each test with `-count=1` and `-p=1` since that's what I needed.
* It runs `go test` for each file





