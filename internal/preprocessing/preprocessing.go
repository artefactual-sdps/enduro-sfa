package preprocessing

type Config struct {
	// Enable preprocessing child workflow.
	Enabled bool
	// Extract package in preprocessing.
	Extract bool
	// Local path shared with Enduro.
	SharedPath string
	// Temporal configuration.
	Temporal Temporal
}

type Temporal struct {
	Namespace    string
	TaskQueue    string
	WorkflowName string
}

type WorkflowParams struct {
	// Relative path to the shared path.
	RelativePath string
}

type WorkflowResult struct {
	// Relative path to the shared path.
	RelativePath string
}
