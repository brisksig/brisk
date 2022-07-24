package brisk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	//  build instance
	instance := New("./test/config_test.json")
	// assert router
	assert.Equal(t, NewRouter(), instance.Router)
}
