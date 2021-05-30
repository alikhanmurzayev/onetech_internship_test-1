package full_outer_join

import (
	"os"
	"testing"
)

func TestFullOuterJoin(t *testing.T) {
	f1Path := "f1.txt"
	f2Path := "f2.txt"
	resultPath := "result.txt"

	f1Data := `first
second
forth
fifth`
	f2Data := `forth
fifth
seventh
eighth`
	expectedData := `eighth
first
second
seventh`

	if err := os.WriteFile(f1Path, []byte(f1Data), os.ModePerm); err != nil {
		t.Fatalf("could not write f1: %s", err)
	}
	if err := os.WriteFile(f2Path, []byte(f2Data), os.ModePerm); err != nil {
		t.Fatalf("could not write f2: %s", err)
	}
	defer func() {
		if err := os.Remove(f1Path); err != nil {
			t.Fatalf("could not remove f1: %s", err)
		}
		if err := os.Remove(f2Path); err != nil {
			t.Fatalf("could not remove f2: %s", err)
		}
	}()

	FullOuterJoin(f1Path, f2Path, resultPath)

	got, err := os.ReadFile(resultPath)
	if err != nil {
		t.Fatalf("could not read resultant file: %s", err)
	}

	if string(got) != expectedData {
		t.Fatalf("got:\n%s\nexpected:\n%s", string(got), expectedData)
	}

	if err := os.Remove(resultPath); err != nil {
		t.Fatalf("could not remove resultant file: %s", err)
	}
}
