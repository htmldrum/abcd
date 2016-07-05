package fs

import (
	"github.com/spf13/afero"
)

var AppFs afero.Fs = afero.NewOsFs()

func NewMockFs() afero.Fs {
	return afero.NewMemMapFs()
}

func Assert(e error){
	if IsNil(e) {
		panic(e)
	}
}
func IsNil(e error) bool {
	return e != nil
}
