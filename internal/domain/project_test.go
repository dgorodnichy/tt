package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNewProject(t *testing.T) {
	tests := []struct {
		testName        string
		name            string
		wantName        string
		wantDescription string
		wantErr         error
	}{
		{testName: "валидное имя", name: "firstProject", wantName: "firstProject", wantDescription: "firstDescription", wantErr: nil},
		{testName: "нормализация", name: " \n\t firstProject ", wantName: "firstProject", wantDescription: "firstDescription", wantErr: nil},
		{testName: "длинное имя(латиница)", name: strings.Repeat("a", 101), wantName: "", wantDescription: "", wantErr: ErrTooLongName},
		{testName: "длинное имя(кириллица)", name: strings.Repeat("ф", 101), wantName: "", wantDescription: "", wantErr: ErrTooLongName},
		{testName: "пустое имя", name: "", wantName: "", wantDescription: "", wantErr: ErrBlankName},
		{testName: "имя из пробелов", name: "   ", wantName: "", wantDescription: "", wantErr: ErrBlankName},
		{testName: "имя 100 символов(латиница)", name: strings.Repeat("a", 100), wantName: strings.Repeat("a", 100), wantDescription: "firstDescription", wantErr: nil},
		{testName: "имя 100 символов(кириллица)", name: strings.Repeat("ф", 100), wantName: strings.Repeat("ф", 100), wantDescription: "firstDescription", wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := NewProject(tt.name, "firstDescription")
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error: %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr != nil {
				return
			}

			if got.Name() != tt.wantName {
				t.Fatalf("got name: %v, want %v", got.Name(), tt.wantName)
			}

			if got.Description() != tt.wantDescription {
				t.Errorf("got description: %v, want %v", got.Description(), tt.wantDescription)
			}
		})
	}
}

func TestSetName(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		wantName string
		wantErr  error
	}{
		{testName: "валидное имя", name: "firstProject", wantName: "firstProject", wantErr: nil},
		{testName: "нормализация", name: " \n\t firstProject ", wantName: "firstProject", wantErr: nil},
		{testName: "длинное имя(латиница)", name: strings.Repeat("a", 101), wantName: "Microservice", wantErr: ErrTooLongName},
		{testName: "длинное имя(кириллица)", name: strings.Repeat("ф", 101), wantName: "Microservice", wantErr: ErrTooLongName},
		{testName: "пустое имя", name: "", wantName: "Microservice", wantErr: ErrBlankName},
		{testName: "имя из пробелов", name: "   ", wantName: "Microservice", wantErr: ErrBlankName},
		{testName: "имя 100 символов(латиница)", name: strings.Repeat("a", 100), wantName: strings.Repeat("a", 100), wantErr: nil},
		{testName: "имя 100 символов(кириллица)", name: strings.Repeat("ф", 100), wantName: strings.Repeat("ф", 100), wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			project, err := NewProject("Microservice", "")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			err = project.SetName(tt.name)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got err: %v, want %v", err, tt.wantErr)
			}

			if project.Name() != tt.wantName {
				t.Errorf("got %v, want %v", project.Name(), tt.wantName)
			}
		})
	}
}
