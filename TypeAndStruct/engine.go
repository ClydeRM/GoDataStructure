package TypeAndStruct

import "fmt"

type GasEngine struct {
     mpg       uint
	 gallons   uint
	 OwnerInfo Owner
}

func NewGasEngine(mpg, gallons uint, ownerName string) (GasEngine, error) {
    if mpg <= 0 || gallons <= 0 {
        return GasEngine{}, fmt.Errorf("mpg and gallons must be greater than 0")
    }

    return GasEngine{
        mpg:       mpg,
        gallons:   gallons,
        OwnerInfo: Owner{Name: ownerName},
    }, nil
}

func (g GasEngine) CalculateRange() uint {
    return g.mpg * g.gallons
}

