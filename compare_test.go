package compare

import (
	"reflect"
	"testing"
)

func Test_compare(t *testing.T) {
	type args struct {
		src         *bar
		comparision *bar
	}
	tests := []struct {
		name    string
		args    args
		want    *bar
		wantErr bool
	}{
		{
			"test 1",
			args{
				&bar{
					[]foo{
						foo{0x001, 1},
						foo{0x002, 2},
						foo{0x003, 3},
						foo{0x004, 4},
						foo{0x005, 5},
					},
				},
				&bar{
					[]foo{
						foo{0x001, 0xFFFFFFFF},
						foo{0x003, 0xFFFFFFFF},
						foo{0x005, 0xFFFFFFFF},
					},
				},
			},
			&bar{
				[]foo{
					foo{0x001, 0xFFFFFFFF},
					foo{0x002, 2},
					foo{0x003, 0xFFFFFFFF},
					foo{0x004, 4},
					foo{0x005, 0xFFFFFFFF},
				},
			},
			false,
		},
		{
			"test 2",
			args{
				&bar{
					[]foo{
						foo{0x001, 1},
						foo{0x002, 2},
						foo{0x003, 3},
						foo{0x004, 4},
						foo{0x005, 5},
					},
				},
				&bar{
					[]foo{
						foo{0x002, 0xFFFFFFFF},
						foo{0x004, 0xFFFFFFFF},
					},
				},
			},
			&bar{
				[]foo{
					foo{0x001, 1},
					foo{0x002, 0xFFFFFFFF},
					foo{0x003, 3},
					foo{0x004, 0xFFFFFFFF},
					foo{0x005, 5},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareAndUpdate(tt.args.src, tt.args.comparision)
			if (err != nil) != tt.wantErr {
				t.Errorf("compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
