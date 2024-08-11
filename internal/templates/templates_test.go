package templates

import (
	"bytes"
	"testing"
	"text/template"
)

// Mocking the original template.ParseGlob call by using a temporary template setup for testing
func TestTemplates_Render(t *testing.T) {
	// Create a new Templates instance
	tmpl := NewTemplates("../../views/*.html")

	// Check that the template is correctly initialized
	if tmpl.templates == nil {
		t.Fatal("Expected non-nil Templates.templates")
	}

	// Prepare a test template in-memory for testing
	tmpl.templates = template.Must(template.New("test_template").Parse("Hello, {{.Name}}"))

	// Prepare a buffer to capture the output
	var buf bytes.Buffer

	// Call the Render function with the test template and data
	err := tmpl.Render(&buf, "test_template", map[string]string{
		"Name": "John Doe",
	})

	if err != nil {
		t.Errorf("Render() returned an error: %v", err)
	}

	// Check if the rendered content is what we expect
	expectedOutput := "Hello, John Doe"
	if buf.String() != expectedOutput {
		t.Errorf("Render() output = %v, want %v", buf.String(), expectedOutput)
	}
}
