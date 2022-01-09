package temp

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go-template/utils"
)

func TestTemp_DoSomething(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := NewMockRepository(ctrl)

	type args struct {
		arg string
	}

	tests := []struct {
		name     string
		args     args
		mockArgs utils.MockData
		wantErr  error
	}{
		{
			name: "Test one",
			args: args{arg: "test"},
			mockArgs: utils.MockData{
				Calls: 1,
				ReturnObjs: []interface{}{
					nil,
				},
			},
			wantErr: nil,
		},
	}

	s := Temp{repo: repoMock}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock.EXPECT().
				DoSomethingTemp(gomock.Any()).
				Times(tt.mockArgs.Calls).
				Return(tt.mockArgs.ReturnObjs...)

			err := s.DoSomething(tt.args.arg)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}
