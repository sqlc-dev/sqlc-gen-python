package endtoend

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FindTests(t *testing.T, root string) []string {
	t.Helper()
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "sqlc.yaml" {
			dirs = append(dirs, filepath.Dir(path))
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return dirs
}

func LookPath(t *testing.T, cmds ...string) string {
	t.Helper()
	for _, cmd := range cmds {
		path, err := exec.LookPath(cmd)
		if err == nil {
			return path
		}
	}
	t.Fatalf("could not find command(s) in $PATH: %s", cmds)
	return ""
}

func ExpectedOutput(t *testing.T, dir string) []byte {
	t.Helper()
	path := filepath.Join(dir, "stderr.txt")
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return []byte{}
		} else {
			t.Fatal(err)
		}
	}
	output, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return output
}

var pattern = regexp.MustCompile(`sha256: ".*"`)

func TestGenerate(t *testing.T) {
	// The SHA256 is required, so we calculate it and then update all of the
	// sqlc.yaml files.
	// TODO: Remove this once sqlc v1.24.0 has been released
	wasmpath := filepath.Join("..", "..", "bin", "sqlc-gen-python.wasm")
	if _, err := os.Stat(wasmpath); err != nil {
		t.Fatalf("sqlc-gen-python.wasm not found: %s", err)
	}
	wmod, err := os.ReadFile(wasmpath)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(wmod)
	sha256 := fmt.Sprintf("%x", sum)

	sqlc := LookPath(t, "sqlc-dev", "sqlc")

	for _, dir := range FindTests(t, "testdata") {
		dir := dir
		t.Run(dir, func(t *testing.T) {
			// Check if sqlc.yaml has the correct SHA256 for the plugin. If not, update the file
			// TODO: Remove this once sqlc v1.24.0 has been released
			yaml, err := os.ReadFile(filepath.Join(dir, "sqlc.yaml"))
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Contains(yaml, []byte(sha256)) {
				yaml = pattern.ReplaceAllLiteral(yaml, []byte(`sha256: "`+sha256+`"`))
				if err := os.WriteFile(filepath.Join(dir, "sqlc.yaml"), yaml, 0644); err != nil {
					t.Fatal(err)
				}
			}

			want := ExpectedOutput(t, dir)
			cmd := exec.Command(sqlc, "diff")
			cmd.Dir = dir
			got, err := cmd.CombinedOutput()
			if diff := cmp.Diff(string(want), string(got)); diff != "" {
				t.Errorf("sqlc diff mismatch (-want +got):\n%s", diff)
			}
			if len(want) == 0 && err != nil {
				t.Error(err)
			}
		})
	}
}
