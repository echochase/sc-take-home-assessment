package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Single folder matches orgID",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("98765432-4321-4321-8765-987654321000"), Name: "TestFolder2"},
			},
			want: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1"},
			},
		},
		{
			name:  "Multiple folders match orgID",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("98765432-4321-4321-8765-987654321000"), Name: "TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder4"},
			},
			want: []folder.Folder{
				
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder4"},
			},
		},
		{
			name:  "No folder matches orgID",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("98765432-4321-4321-8765-987654321000"), Name: "TestFolder2"},
			},
			want: []folder.Folder{},
		},
		{
			name:    "Empty folder list",
			orgID:   uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		fname	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "No children",
			fname: "TestFolder1",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("98765432-4321-4321-8765-987654321000"), Name: "TestFolder2", Paths: "TestFolder2"},
			},
			want: []folder.Folder{},
		},
		{
			name:  "Returns correct list",
			fname: "TestFolder1",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder4", Paths: "TestFolder4"},
			},
			want: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
			},
		},
		{
			name:  "Deeper level children",
			fname: "TestFolder1",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder4", Paths: "TestFolder1.TestFolder3.TestFolder4"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder5", Paths: "TestFolder2.TestFolder5"},
			},
			want: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder4", Paths: "TestFolder1.TestFolder3.TestFolder4"},
			},
		},
		{
			name:  "Correctly checks OrgID",
			fname: "TestFolder1",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"),
			folders: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder1", Paths: "TestFolder1"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
				{OrgId: uuid.FromStringOrNil("98765432-4321-4321-8765-987654321000"), Name: "TestFolder4", Paths: "TestFolder1.TestFolder3.TestFolder4"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder5", Paths: "TestFolder2.TestFolder5"},
			},
			want: []folder.Folder{
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder2", Paths: "TestFolder1.TestFolder2"},
				{OrgId: uuid.FromStringOrNil("12345678-1234-1234-5678-123456789000"), Name: "TestFolder3", Paths: "TestFolder1.TestFolder3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.fname)
			assert.Equal(t, tt.want, get)
		})
	}
}