package enum
type Role int

const (
  Assistant Role = iota + 1
  User
)