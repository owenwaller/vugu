package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func PkgNameOfTest(t *testing.T) string {
	t.Helper()
	// The magefile mounts the test's parent directory as the directory into the nginx container
	// But this code is called by the test which resides in the "testdriver" subdir of the test in question
	// So we need to know the package name, of the test we apply to which is the last component of the directory path of our parent directory.
	// This testdriver is always a dir (strictly a module) rooted at the test in question
	// We then need to append the package name onto the end of the URL passed to chromedp

	// this gives us the "testdriver" directory
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// ensure we are always left in the test driver directory
	defer func() {
		_ = os.Chdir(cwd)
	}()
	// change dir to the parent dir - which will be the test in question
	parentDir := ".."
	err = os.Chdir(parentDir)
	if err != nil {
		t.Fatal(err)
	}
	// find out the pathname on disk...
	parent, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// return the last part which will be the package name of the test we are part of.
	return filepath.Base(parent)
}
