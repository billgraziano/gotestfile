package parse

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	src := `package foo

			import (
				"fmt"
				"time"
			)

			func bar() {
				fmt.Println(time.Now())
			}
			
			func foo() {
				fmt.Println(time.Now())
			}
			`

	fn, err := parse("", src)
	require.NoError(err)
	assert.Equal(2, len(fn))
	assert.Equal("bar", fn[0])
	assert.Equal("foo", fn[1])
}

func TestParseFile(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	wd, err := os.Getwd()
	require.NoError(err)
	names, err := Tests(filepath.Join(wd, "tests_test.go"), nil)
	require.NoError(err)
	assert.GreaterOrEqual(2, len(names))
}
