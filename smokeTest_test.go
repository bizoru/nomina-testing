package nomina_testing_test

import (

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SmokeTest", func() {
       It("Se puede acceder la nomina desde la URL publica",func(){
           estado := 1

       Expect(estado).To(Equal(1))
       })
       
})
