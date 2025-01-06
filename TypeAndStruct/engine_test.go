package TypeAndStruct

import "testing"

func TestNewGasEngine(t *testing.T) {
	type args struct {
		mpg       uint
		gallons   uint
		ownerName string
	}
	tests := []struct {
		name string
		args args
		want GasEngine
	}{
		// test cases
		{
			"validParms",
			args{
				mpg:       10,
				gallons:   20,
				ownerName: "TestCase1",
			},
			GasEngine{10, 20, Owner{"TestCase1"}},
		},
		{
			"invalidParms",
			args{
				mpg:       0,
				gallons:   20,
				ownerName: "TestCase2",
			},
			GasEngine{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewGasEngine(tt.args.mpg, tt.args.gallons, tt.args.ownerName); got != tt.want {
				t.Errorf("NewGasEngine() = %v, want %v", got, tt.want)
			}
		})
	}

}

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
