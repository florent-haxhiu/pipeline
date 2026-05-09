package pipeline

// TODO ConfigParser should read the pipeline.yml file as this will be the bass
// ConfigParser should
// 	read a file
//  produce a data structure
//  know nothing about the excution

type PipelineStatus int

const (
	Pending PipelineStatus = iota
	Running
	Stopped
	Succeeded
	Failed
	Skipped
)

var PipelineStatusName = map[string]PipelineStatus{
	"pending":   Pending,
	"running":   Running,
	"stopped":   Stopped,
	"succeeded": Succeeded,
	"failed":    Failed,
	"skipped":   Skipped,
}

type ConfigParser interface {
	readPipelineFile() (Pipeline, error)
}

type Job struct {
	Name         string   `yaml:"name"`
	Commands     []string `yaml:"commands"`
	Dependencies []string `yaml:"dependencies"`
}

type Pipeline struct {
	Name string `yaml:"name"`
	Jobs []Job `yaml:"jobs"`
}

func ParseStatus(s string) (PipelineStatus, bool) {

}

func ParseFile(fileName string) (Pipeline, error) {

}
