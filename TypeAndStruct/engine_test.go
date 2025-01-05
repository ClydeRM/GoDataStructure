package TypeAndStruct

import "testing"

func TestCalculateRange(t *testing.T) {
    engine := GasEngine{
        mpg:       10,
        gallons:   20,
        OwnerInfo: Owner{Name: "TestOwner"},
    }

    expected := uint(200)
    result := engine.CalculateRange()

    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}
