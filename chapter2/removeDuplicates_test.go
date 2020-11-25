package chapter2

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		ll *nodeSingly
	}
	tests := []struct {
		name string
		args args
		want *nodeSingly
	}{
		{
			name: "ll",
			args: args{
				ll: &nodeSingly{
					val: "a",
					next: &nodeSingly{
						val: "b",
						next: &nodeSingly{
							val: "a",
							next: &nodeSingly{
								val:  "d",
								next: nil,
							},
						},
					},
				},
			},
			want: &nodeSingly{
				val: "a",
				next: &nodeSingly{
					val: "b",
					next: &nodeSingly{
						val:  "d",
						next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if got := removeDuplicatesNoSpace(tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
