package models

import (
	"api-recipe-keeper/helpers"
	"reflect"
	"testing"
)

func TestIngridient_GetById(t *testing.T) {
	type args struct {
		ctx    helpers.Context
		id     string
		params map[string][]string
	}
	tests := []struct {
		name string
		o    *Ingridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetById(tt.args.ctx, tt.args.id, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ingridient.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngridient_GetPaginated(t *testing.T) {
	type args struct {
		ctx    helpers.Context
		params map[string][]string
	}
	tests := []struct {
		name string
		o    *Ingridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetPaginated(tt.args.ctx, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ingridient.GetPaginated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngridient_Create(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *Ingridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Create(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ingridient.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngridient_UpdateById(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *Ingridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.UpdateById(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ingridient.UpdateById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngridient_DeleteById(t *testing.T) {
	type args struct {
		ctx helpers.Context
	}
	tests := []struct {
		name string
		o    *Ingridient
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.DeleteById(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ingridient.DeleteById() = %v, want %v", got, tt.want)
			}
		})
	}
}
