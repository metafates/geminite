package afs

import "github.com/spf13/afero"

var AFS = afero.Afero{Fs: afero.NewOsFs()}
