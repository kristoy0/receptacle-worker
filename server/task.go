package server

// Resources - Resource spec
type Resources struct {
	Memory    string   `json:"memory,omitempty"`
	CPU       string   `json:"cpu,omitempty"`
	Instances int      `json:"instances"`
	Volumes   []string `json:"volumes"`
}

// Task - task spec
type Task struct {
	Name      string    `json:"name,omitempty"`
	Image     string    `json:"image,omitempty"`
	Command   []string  `json:"command,omitempty"`
	Resources Resources `json:"resources"`
	Env       []string  `json:"env,omitempty"`
}
