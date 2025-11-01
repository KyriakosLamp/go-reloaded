package main

// Stage interface for pipeline stages
type Stage interface {
	Process(text string) string
}

// Pipeline represents the text processing pipeline
type Pipeline struct {
	stages []Stage
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