# Multiple Parameters

When multiple arguments are of the same type, the type only needs to be declared after the last one, assuming they are in order.

For example:

```go
func add(x, y int) int {
  return x + y
}

func createUser(firstName, lastName string, age int) {
  return firstName + lastName + age
}

```

If they are not in order they need to be defined separately.
