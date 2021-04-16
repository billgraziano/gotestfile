go build 
gotestfile ./parse/mod_test.go ./parse/tests_test.go
gotestfile -debug -timeout 60m ./parse/mod_test.go
