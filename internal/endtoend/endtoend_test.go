package endtoend

import (
	"os"
	"os/exec"
	"path/filepath"
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

func TestGenerate(t *testing.T) {
	sqlc := LookPath(t, "sqlc-dev", "sqlc")

	for _, dir := range FindTests(t, "testdata") {
		dir := dir
		t.Run(dir, func(t *testing.T) {
			want := ExpectedOutput(t, dir)
			cmd := exec.Command(sqlc, "diff")
			cmd.Dir = dir
			got, err := cmd.CombinedOutput()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("sqlc diff mismatch (-want +got):\n%s", diff)
			}
			if len(want) == 0 && err != nil {
				t.Error(err)
			}
		})
	}
}
