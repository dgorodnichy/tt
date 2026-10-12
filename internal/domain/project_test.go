package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestNewProject(t *testing.T) {
	tests := []struct {
		testName        string
		id              string
		name            string
		wantID          string
		wantName        string
		wantDescription string
		wantErr         error
	}{
		{testName: "валидное имя", id: "1", wantID: "1", name: "firstProject", wantName: "firstProject", wantDescription: "firstDescription", wantErr: nil},
		{testName: "пустой id", id: "", wantID: "1", name: "firstProject", wantName: "firstProject", wantDescription: "firstDescription", wantErr: ErrBlankID},
		{testName: "нормализация", id: "1", wantID: "1", name: " \n\t firstProject ", wantName: "firstProject", wantDescription: "firstDescription", wantErr: nil},
		{testName: "длинное имя(латиница)", id: "1", wantID: "1", name: strings.Repeat("a", 101), wantName: "", wantDescription: "", wantErr: ErrTooLongName},
		{testName: "длинное имя(кириллица)", id: "1", wantID: "1", name: strings.Repeat("ф", 101), wantName: "", wantDescription: "", wantErr: ErrTooLongName},
		{testName: "пустое имя", id: "1", wantID: "1", name: "", wantName: "", wantDescription: "", wantErr: ErrBlankName},
		{testName: "имя из пробелов", id: "1", wantID: "1", name: "   ", wantName: "", wantDescription: "", wantErr: ErrBlankName},
		{testName: "имя 100 символов(латиница)", id: "1", wantID: "1", name: strings.Repeat("a", 100), wantName: strings.Repeat("a", 100), wantDescription: "firstDescription", wantErr: nil},
		{testName: "имя 100 символов(кириллица)", id: "1", wantID: "1", name: strings.Repeat("ф", 100), wantName: strings.Repeat("ф", 100), wantDescription: "firstDescription", wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := NewProject(tt.id, tt.name, "firstDescription")
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got error: %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr != nil {
				return
			}

			if got.ID() != tt.wantID {
				t.Errorf("got id: %q, want %q", got.ID(), tt.wantID)
			}

			if got.Name() != tt.wantName {
				t.Errorf("got name: %q, want %q", got.Name(), tt.wantName)
			}

			if got.Description() != tt.wantDescription {
				t.Errorf("got description: %q, want %q", got.Description(), tt.wantDescription)
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
			project, err := NewProject("1", "Microservice", "")
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
