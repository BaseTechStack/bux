package generators

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateComposable generates the composable for an entity
func GenerateComposable(baseDir, composablesDir, entityName, pluralName string, fields []Field) error {
	// Define the output path for the composable
	// Preserve proper capitalization for multi-word entities by ensuring PascalCase
	pascalPluralName := ToPascalCase(pluralName)
	outputPath := filepath.Join(composablesDir, "use"+pascalPluralName+".ts")
	fmt.Printf("Creating composable: use%s.ts\n", pascalPluralName)

	// Load the template from the embedded filesystem
	templateContent, err := loadTemplate("use_entities.ts.tmpl")
	if err != nil {
		return err
	}

	// Create the composables directory if it doesn't exist
	if err := os.MkdirAll(composablesDir, 0755); err != nil {
		return fmt.Errorf("error creating composables directory: %v", err)
	}

	// Create the template with common functions
	tmpl, err := createTemplate("composable", templateContent)
	if err != nil {
		return err
	}

	// Create a file to write the processed template
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating composable file: %v", err)
	}
	defer file.Close()

	// Execute the template with the data
	data := struct {
		StructName string
		PluralName string
		Fields     []Field
	}{
		StructName: entityName,
		PluralName: pluralName,
		Fields:     fields,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing composable template: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	return nil
}
