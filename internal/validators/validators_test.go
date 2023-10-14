package validators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

func TestValidate(t *testing.T) {
	require.EqualError(
		t,
		Validate(&config.Config{
			Backend:  nil,
			Frontend: nil,
		}),
		fmt.Sprintf(constants.ErrorValidateConfigBlockNotFound, "backend"),
	)
	require.EqualError(
		t,
		Validate(&config.Config{
			Backend:  &config.Backend{ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: nil},
			Frontend: nil,
		}),
		fmt.Sprintf(constants.ErrorValidateConfigOptionInBlockNotFound, "timeout", "backend"),
	)
	require.EqualError(
		t,
		Validate(&config.Config{
			Backend: &config.Backend{
				ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0},
			},
			Frontend: nil,
		}),
		fmt.Sprintf(constants.ErrorValidateConfigBlockNotFound, "frontend"),
	)
	require.NoError(
		t,
		Validate(&config.Config{
			Backend: &config.Backend{
				ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0},
			},
			Frontend: &config.Frontend{
				PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
				HTMX: "latest", Hyperscript: "latest", Manifest: nil,
			},
		}),
	)
	require.Error(
		t,
		Validate(&config.Config{
			Backend: &config.Backend{
				ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0},
			},
			Frontend: &config.Frontend{
				PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
				HTMX: "latest", Hyperscript: "latest", Manifest: &config.Manifest{},
			},
		}),
	)
	require.Error(
		t,
		Validate(&config.Config{
			Backend: &config.Backend{
				ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0},
			},
			Frontend: &config.Frontend{
				PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
				HTMX: "latest", Hyperscript: "latest", Manifest: &config.Manifest{
					Name: "project", ShortName: "project", Description: "project", BackgroundColor: "project",
					ThemeColor: "project", Display: "project", Orientation: "project", StartURL: "project",
					Icons: nil,
				},
			},
		}),
	)
	require.Error(
		t,
		Validate(&config.Config{
			Backend: &config.Backend{
				ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0},
			},
			Frontend: &config.Frontend{
				PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
				HTMX: "latest", Hyperscript: "latest", Manifest: &config.Manifest{
					Name: "project", ShortName: "project", Description: "project", BackgroundColor: "#ffffff",
					ThemeColor: "#ffffff", Display: "standalone", Orientation: "portrait", StartURL: ".",
					Icons: []*config.Icon{
						{Src: "src", Type: "type", Sizes: "sizes"},
					},
				},
			},
		}),
	)
	require.NoError(
		t,
		Validate(&config.Config{
			Backend: &config.Backend{
				ModuleName: "project", GoFramework: "default", Port: 5000, Timeout: &config.Timeout{Read: 0, Write: 0},
			},
			Frontend: &config.Frontend{
				PackageName: "project", CSSFramework: "default", RuntimeEnvironment: "default",
				HTMX: "latest", Hyperscript: "latest", Manifest: &config.Manifest{
					Name: "project", ShortName: "project", Description: "project", BackgroundColor: "#ffffff",
					ThemeColor: "#ffffff", Display: "standalone", Orientation: "portrait", StartURL: ".",
					Icons: []*config.Icon{
						{Src: "src", Type: "image/svg+xml", Sizes: "sizes"},
					},
				},
			},
		}),
	)
}
