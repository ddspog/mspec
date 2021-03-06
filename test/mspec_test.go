package test

import (
	"math"
	"testing"

	"github.com/ddsgok/bdd"
	"github.com/ddsgok/bdd/spec"
)

func Test_MSpec_Instances(t *testing.T) {
	given := bdd.Sentences().Given()

	spec.SetSilent()

	given(t, "an spec instance", func(when bdd.When) {

		f := "feature value"

		c := &spec.Configuration{
			LastFeature: f,
			LastGiven:   "context value",
			LastWhen:    "when value",
			LastIt:    "title value",
		}

		when("calling reset()", func(it bdd.It) {

			c.ResetLasts()

			it("should not reset lastFeature as that is used globally", func(assert bdd.Assert) {
				assert.NotEmpty(c.LastFeature)
			})

			it("should keep the lastFeature value", func(assert bdd.Assert) {
				assert.Equal(c.LastFeature, f)
			})

			it("should set lastGiven to zero value", func(assert bdd.Assert) {
				assert.Empty(c.LastGiven)
			})

			it("should set lastWhen to zero value", func(assert bdd.Assert) {
				assert.Empty(c.LastWhen)
			})

			it("should set lastIt to zero value", func(assert bdd.Assert) {
				assert.Empty(c.LastIt)
			})
		})
	})
}

func BenchmarkGivenStub(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()

	given := bdd.Sentences().Given()

	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		given(t, "a single given")
	}
}

func BenchmarkWhenStub(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()

	given := bdd.Sentences().Given()

	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		given(t, "a single given", func(when bdd.When) {
			when("a single when")
		})
	}
}

func BenchmarkThenStub(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()

	given := bdd.Sentences().Given()

	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		given(t, "a single given", func(when bdd.When) {
			when("a single when", func(it bdd.It) {
				it("should have a single then")
			})
		})
	}
}

func BenchmarkError(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()

	given := bdd.Sentences().Given()

	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		given(t, "a context to fail", func(when bdd.When) {
			when("prepping to call the thing to fail", func(it bdd.It) {
				it("should fail", func(assert bdd.Assert) {
					assert.True(false)
				})
			})
		})
	}
}

func BenchmarkSimpleMspec(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()

	given := bdd.Sentences().Given()

	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		given(t, "xyz", func(when bdd.When) {
			ii := int8(10)
			when("123", func(it bdd.It) {
				it("should be this", func(assert bdd.Assert) {
					if !assert.Equal(10, int(ii)) {
						b.Fail()
					}
				})
			})
		})
	}
}

func BenchmarkSimpleTest(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{} // get accurate GC numbers
		if t == nil {
			b.Fail()
		}
		ii := int8(10)
		if 10 != int(ii) {
			b.Fail()
		}
	}
}

func BenchmarkComplexMspec(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()

	given := bdd.Sentences().Given()

	for i := 0; i < b.N; i++ {
		t := &testing.T{}
		given(t, "a struct we create often", func(when bdd.When) {

			x := struct {
				Value string
				Log   float64
			}{}

			when("we do something more complex", func(it bdd.It) {

				x.Value = "a string to be set"
				x.Log = math.Log(20)
				ii := int8(10)

				it("should have x.Value be what we expect", func(assert bdd.Assert) {
					if !assert.Equal("a string to be set", x.Value) {
						b.Fail()
					}
				})

				it("should have x.Log be what we expect", func(assert bdd.Assert) {
					if !assert.Equal(2.995732273553991, x.Log) {
						b.Fail()
					}
				})

				it("should be true", func(assert bdd.Assert) {
					if !assert.Equal(10, int(ii)) {
						b.Fail()
					}
				})
			})

			when("we do something else", func(it bdd.It) {

				x.Value = "another string to be set"
				x.Log = math.Log(15)

				it("should have x.Value be what we expect", func(assert bdd.Assert) {
					if !assert.Equal("another string to be set", x.Value) {
						b.Fail()
					}
				})

				it("should have x.Log be what we expect", func(assert bdd.Assert) {
					if !assert.Equal(2.70805020110221, x.Log) {
						b.Fail()
					}
				})

				it("should be true", func(assert bdd.Assert) {
					if !assert.Equal(math.Log2E, 1/math.Ln2) {
						b.Fail()
					}
				})
			})
		})
	}
}

func BenchmarkComplexTest(b *testing.B) {
	spec.SetSilent()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := &testing.T{} // get accurate GC numbers
		if t == nil {
			b.Fail()
		}
		x := struct {
			Value string
			Log   float64
		}{}

		x.Value = "a string to be set"
		x.Log = math.Log(20)
		ii := int8(10)

		if x.Value != "a string to be set" {
			b.Fail()
		}

		if x.Log != 2.995732273553991 {
			b.Fail()
		}

		if 10 != int(ii) {
			b.Fail()
		}

		x.Value = "another string to be set"
		x.Log = math.Log(15)

		if x.Value != "another string to be set" {
			b.Fail()
		}

		if x.Log != 2.70805020110221 {
			b.Fail()
		}

		if math.Log2E != 1/math.Ln2 {
			b.Fail()
		}
	}
}
