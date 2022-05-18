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

func (s *MySuite) TestPrefixToPostfixAnswers(c *C) {
	resultTest1, errTest1 := PrefixToPostfix("+ 5 * - 4 2 3")
	c.Check(resultTest1, Equals, "3 2 4 - * 5 +")
	c.Check(errTest1, IsNil)

	resultTest2, errTest2 := PrefixToPostfix("+ 2 2")
	c.Check(resultTest2, Equals, "2 2 +")
	c.Check(errTest2, IsNil)

	resultTest3, errTest3 := PrefixToPostfix("- 3 / 5 10")
	c.Check(resultTest3, Equals, "10 5 / 3 -")
	c.Check(errTest3, IsNil)

	resultTest4, errTest4 := PrefixToPostfix("- 3 / 5 + 4 * 3 - 3 / + 2 3 10")
	c.Check(resultTest4, Equals, "10 3 2 + / 3 - 3 * 4 + 5 / 3 -")
	c.Check(errTest4, IsNil)

	resultTest5, errTest5 := PrefixToPostfix("+ / + ^ % * 5 7 3 8 0 2 10")
	c.Check(resultTest5, Equals, "10 2 0 8 3 7 5 * % ^ + / +")
	c.Check(errTest5, IsNil)
}

func (s *MySuite) TestPrefixToPostfixErrors(c *C) {
	_, errTest1 := PrefixToPostfix("$ 2 2")
	c.Check(errTest1, NotNil)
	c.Check(fmt.Sprintf("%s", errTest1), Equals, "Should only be numbers or mathematical signs")

	_, errTest2 := PrefixToPostfix("2 2 +")
	c.Check(errTest2, NotNil)
	c.Check(fmt.Sprintf("%s", errTest2), Equals, "Error in input text")

	_, errTest3 := PrefixToPostfix("+ 2 3 4")
	c.Check(errTest3, NotNil)
	c.Check(fmt.Sprintf("%s", errTest3), Equals, "Error in input text")

	_, errTest4 := PrefixToPostfix("+ 2 / 6 2 4")
	c.Check(errTest4, NotNil)
	c.Check(fmt.Sprintf("%s", errTest4), Equals, "Error in input text")
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- 3 / 5 10")
	fmt.Println(res)

	// Output:
	// 10 5 / 3 -
}
