// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package auto

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/siderolabs/kres/internal/project/helm"
)

// DetectHelm checks the helm settings.
// It returns true if helm is enabled and the chart path is set.
func (builder *builder) DetectHelm() (bool, error) {
	var helm Helm

	if err := builder.meta.Config.Load(&helm); err != nil {
		return false, err
	}

	if helm.Enabled {
		if helm.ChartDir == "" {
			return false, fmt.Errorf("chart directory is not set")
		}

		if _, err := os.Stat(filepath.Join(builder.rootPath, helm.ChartDir, "Chart.yaml")); err != nil {
			return false, fmt.Errorf("chart.yaml not found in %s: %w", helm.ChartDir, err)
		}

		builder.meta.HelmChartDir = helm.ChartDir

		return true, nil
	}

	return false, nil
}

func (builder *builder) BuildHelm() error {
	helm := helm.NewBuild(builder.meta)

	builder.targets = append(builder.targets, helm)
	helm.AddInput(builder.commonInputs...)

	return nil
}
