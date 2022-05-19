package lab2

import (
	"bytes"
	"fmt"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestComputeAnswers(c *C) {
	handler := &ComputeHandler{}
	var (
		got, want    string
		err, ptp_err error
	)

	// TEST 1
	var bufer1 bytes.Buffer

	handler.Output = &bufer1
	handler.Input = strings.NewReader("+ 2 2")

	err = handler.Compute()
	got = bufer1.String()
	want, ptp_err = PrefixToPostfix("+ 2 2")
	c.Check(ptp_err, IsNil)
	c.Check(err, IsNil)
	c.Check(got, Equals, want)

	// TEST 2
	var bufer2 bytes.Buffer

	handler.Output = &bufer2
	handler.Input = strings.NewReader("- 3 / 5 + 4 * 3 - 3 / + 2 3 10")

	err = handler.Compute()
	got = bufer2.String()
	want, ptp_err = PrefixToPostfix("- 3 / 5 + 4 * 3 - 3 / + 2 3 10")
	c.Check(ptp_err, IsNil)
	c.Check(err, IsNil)
	c.Check(got, Equals, want)

	// TEST 3
	var bufer3 bytes.Buffer

	handler.Output = &bufer3
	handler.Input = strings.NewReader("+ / 2 2 + 5 5")

	err = handler.Compute()
	got = bufer3.String()
	want, ptp_err = PrefixToPostfix("+ / 2 2 + 5 5")
	c.Check(ptp_err, IsNil)
	c.Check(err, IsNil)
	c.Check(got, Equals, want)

}

func (s *MySuite) TestComputeErrors(c *C) {
	handler := &ComputeHandler{}
	var (
		got string
		err error
	)

	// TEST 1
	var bufer1 bytes.Buffer

	handler.Output = &bufer1
	handler.Input = strings.NewReader("+ 2 / - * 5 - 3 + 3 2")

	err = handler.Compute()
	got = bufer1.String()
	c.Check(err, NotNil)
	c.Check(fmt.Sprintf("%s", err), Equals, "Error in input text")
	c.Check(got, Equals, "")

	// TEST 2
	var bufer2 bytes.Buffer

	handler.Output = &bufer2
	handler.Input = strings.NewReader("2 2 +")

	err = handler.Compute()
	got = bufer2.String()
	c.Check(err, NotNil)
	c.Check(fmt.Sprintf("%s", err), Equals, "Error in input text")
	c.Check(got, Equals, "")

	// TEST 3
	var bufer3 bytes.Buffer

	handler.Output = &bufer3
	handler.Input = strings.NewReader("+ 2 $ 5 6")

	err = handler.Compute()
	got = bufer3.String()
	c.Check(err, NotNil)
	c.Check(fmt.Sprintf("%s", err), Equals, "Should only be numbers or mathematical signs")
	c.Check(got, Equals, "")
}
