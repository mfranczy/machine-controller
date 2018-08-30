package ubuntu

import (
	"testing"

	"github.com/go-test/deep"
)

func TestGetDockerInstallCandidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		version string
		resErr  error
		resPkg  string
		resVer  string
	}{
		{
			name:    "no version found",
			version: "foo-does-not-exist",
			resErr:  errNoInstallCandidateAvailable,
			resPkg:  "",
			resVer:  "",
		},
		{
			name:    "get patch version",
			version: "17.12.1",
			resErr:  nil,
			resPkg:  "docker.io",
			resVer:  "17.12.1-0ubuntu1",
		},
		{
			name:    "get minor version",
			version: "17.12",
			resErr:  nil,
			resPkg:  "docker.io",
			resVer:  "17.12.1-0ubuntu1",
		},
		{
			name:    "get different package for newer version",
			version: "18.06.1",
			resErr:  nil,
			resPkg:  "docker-ce",
			resVer:  "18.06.1~ce~3-0~ubuntu",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pkg, version, err := getDockerInstallCandidate(test.version)
			if diff := deep.Equal(err, test.resErr); diff != nil {
				t.Errorf("expected to get %v instead got: %v", test.resErr, err)
			}
			if err != nil {
				return
			}
			if pkg != test.resPkg {
				t.Errorf("expected to get %v instead got: %v", test.resPkg, pkg)
			}
			if version != test.resVer {
				t.Errorf("expected to get %v instead got: %v", test.resVer, version)
			}
		})
	}

}
