package parray

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUniqObjectID(t *testing.T) {
	type args struct {
		arr []primitive.ObjectID
	}
	id1, _ := primitive.ObjectIDFromHex("61640d3b32edbb4e39ec7db3")
	id2, _ := primitive.ObjectIDFromHex("615ffdc4c2c4924d5bd8947f")
	tests := []struct {
		name string
		args args
		want []primitive.ObjectID
	}{
		{
			name: "1",
			args: args{
				arr: []primitive.ObjectID{
					id1,
					id2,
					id1,
				},
			},
			want: []primitive.ObjectID{id1, id2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqObjectID(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqObjectID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqString(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				arr: []string{"1", "2", "1"},
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqString(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqString() = %v, want %v", got, tt.want)
			}
		})
	}
}
