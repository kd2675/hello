package model

type CarList struct {
	cars []*Car
}

func NewCarList(carNames []string) (*CarList, error) {
	cars := make([]*Car, 0, len(carNames))

	for _, name := range carNames {
		car, err := NewCar(name)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return &CarList{
		cars: cars,
	}, nil
}

func (cl *CarList) Cars() []*Car {
	return cl.cars
}

func (cl *CarList) MoveAll(randomValues []int) {
	for i, car := range cl.cars {
		if i < len(randomValues) {
			car.MoveForward(randomValues[i])
		}
	}
}
