package Interface

import (
	"reflect"
	"testing"
)


func TestImplement_Car(t *testing.T) {
	type args struct {
		speedList []uint
		passengerList []uint
	}
	tests := []struct {
		name string
		args args
		want []Car
	}{
		// test cases
		{
			"Car_Test",
			args{
				speedList: []uint{100, 120, 150},
				passengerList: []uint{4,4,2},
			},
			[]Car{{100,4}, {120, 4}, {150, 2}},
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			speedList := tt.args.speedList
			passengerList := tt.args.passengerList
			carList := make([]Car, 3)
			for index := range speedList {
				carList[index] = Car{speedList[index], passengerList[index]}
			}

			for index, _ := range carList {
				if !reflect.DeepEqual(carList[index], tt.want[index]) {
					t.Errorf("Car: got = %v, want = %v", carList[index], tt.want[index])
				}

				if !reflect.DeepEqual(carList[index].Speed() , tt.want[index].Speed()) {
					t.Errorf("Car Speed: got = %v, want = %v", carList[index].Speed(), tt.want[index].Speed())
				}

				if !reflect.DeepEqual(carList[index].Capacity() , tt.want[index].Capacity()) {
					t.Errorf("Car Capacity: got = %v, want = %v", carList[index].Capacity(), tt.want[index].Capacity())
				}
			}
		})
	}
}