package count

import (
	"testing"

	"n3dr/internal/app/n3dr/goswagger/models"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	asset := models.AssetXO{ContentType: "text/plain"}
	assets := []*models.AssetXO{&asset, &asset, &asset}

	n := Nexus3{}

	repositoriesTotalArtifacts := 0
	repositoriesTotalArtifactsPointer := &repositoriesTotalArtifacts
	n.assets(assets, repositoriesTotalArtifactsPointer)

	assert.Equal(t, 3, *repositoriesTotalArtifactsPointer)
	assert.Equal(t, 3, total)
}
