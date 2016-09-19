package pathutil

import (
	"io/ioutil"
	"testing"

	"github.com/frozzare/go-assert"
	"github.com/kardianos/osext"
)

func TestRealPath(t *testing.T) {
	assert.Equal(t, "public", RealPath("public"))
}

func TestRealPathTmp(t *testing.T) {
	path, _ := osext.ExecutableFolder()
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(path+"/dat", d1, 0644)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, path+"/dat", RealPath("dat"))
}
