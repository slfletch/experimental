package action

import (
	"errors"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/tektoncd/experimental/oci/pkg/oci"
)

// Push will perform the `push` action by recursively reading all of the
// Tekton specs passed in, bundling it into an image, and pushing the result
// to an OCI-compliant repository.
func Push(r string, filePaths []string) error {
	// Validate the parameters.
	if r == "" || len(filePaths) == 0 {
		return errors.New("must specify a valid image name and file paths")
	}

	resources, err := oci.ReadPaths(filePaths)
	if err != nil {
		return err
	}

	ref, err := name.ParseReference(r)
	if err != nil {
		return err
	}

	return oci.PushImage(ref, resources)
}
