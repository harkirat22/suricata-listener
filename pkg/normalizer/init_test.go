package normalizer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNormalizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Normalizer Suite")
}
