package nomina_testing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNominaTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NominaTesting Suite")
}
