package models

import (
	"api-recipe-keeper/helpers"
	"reflect"
	"testing"
)

func TestRecipeIngridient_GetById(t *testing.T) {
	type args struct {
		ctx    helpers.Context
		id     string
		params map[string][]string
	}
	tests := []struct {
		name string
		o    *RecipeIngridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetById(tt.args.ctx, tt.args.id, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecipeIngridient.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipeIngridient_Create(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *RecipeIngridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Create(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecipeIngridient.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecipeIngridient_UpdateById(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *RecipeIngridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.UpdateById(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecipeIngridient.UpdateById() = %v, want %v", got, tt.want)
			}
		})
	}
}
