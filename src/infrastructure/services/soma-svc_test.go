package services

import "testing"

func Test_somaServices_Add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		s     *somaServices
		args  args
		wantZ int
	}{
		{
			name: "soma 1",
			args: args{
				x: 1,
				y: 1,
			},
			wantZ: 2,
		},
		{
			name: "soma 2",
			args: args{
				x: 2,
				y: 1,
			},
			wantZ: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &somaServices{}
			if gotZ := s.Add(tt.args.x, tt.args.y); gotZ != tt.wantZ {
				t.Errorf("somaServices.Add() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}
