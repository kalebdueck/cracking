package chapter2

import (
	"reflect"
	"testing"
)

func Test_kthLastElement(t *testing.T) {
	type args struct {
		ll *nodeSingly
		k  int
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
								val: "d",
								next: &nodeSingly{
									val: "d",
									next: &nodeSingly{
										val:  "a",
										next: nil,
									},
								},
							},
						},
					},
				},
				k: 2,
			},
			want: &nodeSingly{
				val: "d",
				next: &nodeSingly{
					val:  "a",
					next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLastElement(tt.args.ll, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthLastElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
