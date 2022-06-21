package lor

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/kittcodes/golio/api"
	"github.com/kittcodes/golio/internal"
	"github.com/kittcodes/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), logrus.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
