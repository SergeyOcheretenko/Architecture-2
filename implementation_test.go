package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Init go-check

func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixToPostfix(c *C) {
	resultTest1, errTest1 := PrefixToPostfix("+ 5 * - 4 2 3")
	if errTest1 == nil {
		c.Check(resultTest1, Equals, "3 2 4 - * 5 +")
	} else {
		c.Fatal(errTest1)
	}
	
	resultTest2, errTest2 := PrefixToPostfix("+ 2 2")
	if errTest2 == nil {
		c.Check(resultTest2, Equals, "2 2 +")
	} else {
		c.Fatal("Error in PrefixToPostfix work")
	}

	resultTest3, errTest3 := PrefixToPostfix("- 3 / 5 10")
	if errTest3 == nil {
		c.Check(resultTest3, Equals, "10 5 / 3 -")
	} else {
		c.Fatal("Error in PrefixToPostfix work")
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- 3 / 5 10")
	fmt.Println(res)

	// Output:
	// 10 5 / 3 -
}
