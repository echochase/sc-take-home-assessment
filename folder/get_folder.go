package folder

import "github.com/gofrs/uuid"
import "regexp"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

/*
 * Gets all the child folders of a given folder.
 * Parameters:
 * - orgID, the ID of the organisation to which the folders belong.
 * - name, the name of the folder for which we're finding child folders.
 * Returns:
 * - a list of all child folders to the given folder.
 */
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.GetFoldersByOrgID(orgID)
	var parent Folder
	res := []Folder{}
	for _, folder := range folders {
		if name == folder.Name {
			parent = folder
		}
	}
	for _, folder := range folders {
		matched, _ := regexp.MatchString(name, folder.Paths)
		if matched && parent.Paths != folder.Paths {
			res = append(res, folder)
		}
	}
	
	return res
}
