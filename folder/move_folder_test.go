package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name 		string
		srcName		string
		dstName		string
		folders		[]folder.Folder
		want   		[]folder.Folder
		wantErr  	bool
		expectedErr	string
	}{
		{
			name: "Successfully moves folder",
			srcName: "TestFolder1",
			dstName: "TestFolder2",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder3"},
			},
			want: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder2.TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder3"},
			},
			wantErr: false,
			expectedErr: "",
		},
		{
			name: "Different OrgID",
			srcName: "TestFolder1",
			dstName: "TestFolder2",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("98765432-4321-4321-8765-987654321000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder3"},
			},
			want: []folder.Folder{},
			wantErr: true,
			expectedErr: "Error: Cannot move a folder to a different organization",
		},
		{
			name: "Nonexistent source",
			srcName: "TestFolder3",
			dstName: "TestFolder2",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
			},
			want: []folder.Folder{},
			wantErr: true,
			expectedErr: "Error: Source folder does not exist",
		},
		{
			name: "Nonexistent destination",
			srcName: "TestFolder1",
			dstName: "TestFolder3",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
			},
			want: []folder.Folder{},
			wantErr: true,
			expectedErr: "Error: Destination folder does not exist",
		},
		{
			name: "Move a folder to itself",
			srcName: "TestFolder1",
			dstName: "TestFolder1",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
			},
			want: []folder.Folder{},
			wantErr: true,
			expectedErr: "Error: Cannot move a folder to itself",
		},
		{
			name: "Move a folder to child of itself",
			srcName: "TestFolder1",
			dstName: "TestFolder2",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
			},
			want: []folder.Folder{},
			wantErr: true,
			expectedErr: "Error: Cannot move a folder to a child of itself",
		},
		{
			name: "Successfully moves folder and children",
			srcName: "TestFolder1",
			dstName: "TestFolder2",
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
			},
			want: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder2.TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder2.TestFolder1.TestFolder3"},
			},
			wantErr: false,
			expectedErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.srcName, tt.dstName)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, get)
			}
		})
	}
}