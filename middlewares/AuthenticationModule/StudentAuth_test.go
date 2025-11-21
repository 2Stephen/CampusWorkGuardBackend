package middlewares

import (
	"CampusWorkGuardBackend/models/request"
	"testing"
)

func TestCHSIStudentInfo_StudentAuth(t *testing.T) {
	type fields struct {
		Name         string
		Gender       string
		Birthday     string
		Nation       string
		School       string
		Level        string
		Major        string
		Duration     string
		DegreeType   string
		StudyMode    string
		College      string
		Department   string
		EntranceDate string
		Status       string
		ExpectedGrad string
		Vcode        string
		StudentID    string
		Email        string
	}
	type args struct {
		params request.StudentAuthParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// 测试
		{
			name: "Valid student authentication",
			fields: fields{
				School: "武汉理工大学",
				Vcode:  "AMTVU01H39D07FX2",
			},
			args: args{
				params: request.StudentAuthParams{
					ID:     "123456789",
					School: "武汉理工大学",
					Vcode:  "AMTVU01H39D07FX2",
					Email:  "2432757689@qq.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cHSIStudentInfo := &CHSIStudentInfo{
				Name:         tt.fields.Name,
				Gender:       tt.fields.Gender,
				Birthday:     tt.fields.Birthday,
				Nation:       tt.fields.Nation,
				School:       tt.fields.School,
				Level:        tt.fields.Level,
				Major:        tt.fields.Major,
				Duration:     tt.fields.Duration,
				DegreeType:   tt.fields.DegreeType,
				StudyMode:    tt.fields.StudyMode,
				College:      tt.fields.College,
				Department:   tt.fields.Department,
				EntranceDate: tt.fields.EntranceDate,
				Status:       tt.fields.Status,
				ExpectedGrad: tt.fields.ExpectedGrad,
				Vcode:        tt.fields.Vcode,
				StudentID:    tt.fields.StudentID,
				Email:        tt.fields.Email,
			}
			if got := cHSIStudentInfo.StudentAuth(tt.args.params); got != tt.want {
				t.Errorf("StudentAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
