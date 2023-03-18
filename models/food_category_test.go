package models

import (
	"api-recipe-keeper/helpers"
	"reflect"
	"testing"
)

func TestFoodCategory_GetById(t *testing.T) {
	type args struct {
		ctx    helpers.Context
		id     string
		params map[string][]string
	}
	tests := []struct {
		name string
		o    *FoodCategory
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetById(tt.args.ctx, tt.args.id, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoodCategory.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodCategory_GetPaginated(t *testing.T) {
	type args struct {
		ctx    helpers.Context
		params map[string][]string
	}
	tests := []struct {
		name string
		o    *FoodCategory
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetPaginated(tt.args.ctx, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoodCategory.GetPaginated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodCategory_Create(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *FoodCategory
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Create(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoodCategory.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodCategory_UpdateById(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *FoodCategory
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.UpdateById(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoodCategory.UpdateById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoodCategory_DeleteById(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *FoodCategory
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.DeleteById(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoodCategory.DeleteById() = %v, want %v", got, tt.want)
			}
		})
	}
}
