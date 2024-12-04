# Go Type Assertion Panic

This example demonstrates a common error in Go: panics caused by type assertions.  The `i.(int32)` line will panic if `i` does not hold an `int32` value.

The `bug.go` file shows the problematic code. The `bugSolution.go` file shows how to safely handle type assertions using type switches.

**How to Reproduce:**
1. Run `bug.go`.
2. Observe the panic.

**Solution:** Use type switches to handle potential type mismatches gracefully.