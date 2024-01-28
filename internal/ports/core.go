package ports

type ArithmeticPort interface {
	Addition(x, y int32) (int32, error)
	Subtraction(x, y int32) (int32, error)
	Multiplication(x, y int32) (int32, error)
	Division(x, y int32) (int32, error)
}
