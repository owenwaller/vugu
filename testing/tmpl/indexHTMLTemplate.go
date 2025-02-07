package tmpl

import (
	"html/template"
	"os"
	"testing"
)

// Reads a index.html.tmpl file in the currect directory - which should be the same directory as the
// Go test that is executing i.e. te "testdriver" directory of the test in question
// It creates an index.html in the directory of the wasm test i.e. the parent directory of the "testdriver".
// which is then served by the containerized nginx when the test is run.
func CreateIndexHtml(t *testing.T, pkgName string) {
	type TestPath struct {
		TestDir string
	}
	// record where we are on the FS
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// ensure we always return the the CWD on exit
	defer func() {
		_ = os.Chdir(cwd)

	}()

	tp := TestPath{TestDir: pkgName}

	tmpl, err := template.ParseFiles(cwd + "/index.html.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	// cd to the parent dir where the wasm test is built and served from
	parentDir := ".."
	err = os.Chdir(parentDir)
	// get the full path
	parentDir, err = os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// remove any existing "index.html" - we don't care if the file does not exist
	err = os.Remove(parentDir + "/index.html")
	if err != nil {
		t.Logf("rm error (not fatal) %s", err)
	}
	// create the new index.html in the parent (i.e the tests not ht test driver) directory
	indexHTML, err := os.Create(parentDir + "/index.html")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err = indexHTML.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	// build the index.html
	err = tmpl.Execute(indexHTML, tp)
	if err != nil {
		t.Fatal(err)
	}
	// and then write it to disk
	err = indexHTML.Sync() // ensure we flush to disk
	if err != nil {
		t.Fatal(err)
	}
}
