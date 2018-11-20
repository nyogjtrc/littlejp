package littlejp

// ThrowRecord save player throw money
type ThrowRecord struct {
	ID       int
	UserID   int
	Amount   float32
	FromGame string
}
