package flyweight

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestGetFlyweightDistinguishesFontFamily(t *testing.T) {
	factory := NewCharacterFlyweightFactory()
	arial := factory.GetFlyweight('A', "Arial", 12, "black")
	times := factory.GetFlyweight('A', "Times", 12, "black")
	if arial == times {
		t.Fatal("same glyph with different font families reused one flyweight")
	}
	if got := factory.GetFlyweightCount(); got != 2 {
		t.Fatalf("got %d flyweights, want 2", got)
	}
}

func TestRenderTextAdvancesByRune(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	originalStdout := os.Stdout
	os.Stdout = writer
	NewTextEditorClient().RenderText("A界B", 0, 0, "Arial", 12, "black")
	os.Stdout = originalStdout
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if err := reader.Close(); err != nil {
		t.Fatal(err)
	}
	for _, position := range []string{"at (0, 0)", "at (12, 0)", "at (24, 0)"} {
		if !strings.Contains(string(output), position) {
			t.Errorf("render output %q does not contain %q", output, position)
		}
	}
}
