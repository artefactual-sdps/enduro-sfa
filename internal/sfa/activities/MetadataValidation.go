package activities

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

const MetadataValidationActivityName = "metadata-validation"

type MetadataValidationActivity struct{}

func NewMetadataValidationActivity() *MetadataValidationActivity {
	return &MetadataValidationActivity{}
}

type MetadataValidationParams struct {
	SipPath string
}

type MetadataValidationResult struct {
	Out string
}

func (md *MetadataValidationActivity) Execute(ctx context.Context, params *MetadataValidationParams) (*MetadataValidationResult, error) {
	res := &MetadataValidationResult{}
	e, err := exec.Command("python3",
		"xsdval.py",
		// Arguments.
		filepath.Join(params.SipPath, "/header/metadata.xml"),
		"arelda.xsd").CombinedOutput()
	if err != nil {
		fmt.Println(string(e))
		return nil, err
	}

	res.Out = string(e)
	return res, nil
}
