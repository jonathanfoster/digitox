package pathutil

import (
	"fmt"
	"path"
	"path/filepath"

	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

// FileName sanitizes ID and joins with blocklist directory path to create a
// file path. The file path is then checked to ensure its directory is the
// blocklist directory to prevent directory traversal using relative paths.
func FileName(id string, dirname string) (string, error) {
	filename := path.Join(dirname, validator.SafeFileName(id))
	filename, err := filepath.Abs(filename)
	if err != nil {
		return "", errors.Wrapf(err, "error returning absolute filename file path %s", filename)
	}

	dirname, err = filepath.Abs(dirname)
	if err != nil {
		return "", errors.Wrapf(err, "error returning absolute dirname path %s", dirname)
	}

	if filepath.Dir(filename) != dirname {
		return "", fmt.Errorf("filename path %s not in dirname directory %s", filename, dirname)
	}

	return filename, nil
}
