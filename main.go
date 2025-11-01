package main

import (
	"fmt"
	"os"
	"go-reloaded/functions"
)

// Pipeline represents the text processing pipeline
type Pipeline struct {
	stages []Stage
}

// Stage interface for pipeline stages
type Stage interface {
	Process(text string) string
}

// NewPipeline creates a new pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{stages: []Stage{}}
}

// AddStage adds a stage to the pipeline
func (p *Pipeline) AddStage(stage Stage) {
	p.stages = append(p.stages, stage)
}

// Process runs text through all pipeline stages
func (p *Pipeline) Process(text string) string {
	for _, stage := range p.stages {
		text = stage.Process(text)
	}
	return text
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input file
	content, err := functions.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Create pipeline with stages
	pipeline := NewPipeline()
	pipeline.AddStage(&functions.NumericConversionStage{})
	pipeline.AddStage(&functions.ArticleStage{})
	pipeline.AddStage(&functions.CaseTransformStage{})
	pipeline.AddStage(&functions.QuotationStage{})
	pipeline.AddStage(&functions.PunctuationStage{})

	// Process content
	result := pipeline.Process(content)

	// Write output file
	err = functions.WriteFile(outputFile, result)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing complete: %s -> %s\n", inputFile, outputFile)
}