package ports

type APIPort interface {
	GetAddition(x, y int32) (int32, error)
	GetSubtraction(x, y int32) (int32, error)
	GetMultiplication(x, y int32) (int32, error)
	GetDivision(x, y int32) (int32, error)
}
