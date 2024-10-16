package folder

import "regexp"
import "errors"

/*
 * Helper function for MoveFolder
 * Given a folder name, retrieves the folder struct that matches the name.
 * Parameters:
 * - name, the name of the folder
 * Returns:
 * - the folder struct that matches the given name (or an empty folder struct if there is an error).
 * - an error (or nil if no error).
 */

func (f *driver) GetFolderByName(name string) (Folder, error) {
	for _, folder := range f.folders {
		if name == folder.Name {
			return folder, nil
		}
	}
	return Folder{}, errors.New("folder not found")
}

/*
 * Moves one folder into another, given they are from the same organisation.
 * Parameters:
 * - name, the name of the source folder that we're moving.
 * - dst, the destination of the source folder.
 * Returns:
 * - the updated list of folders after the move (or an empty list if there is an error).
 * - an error (or nil if no error).
 * Will throw an error if:
 * - the two folders belong to different organisations.
 * - the two folders are the same.
 * - one or both of the source and destination folders do not exist.
 * - the destination is the child of the source.
 */
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folder, err := f.GetFolderByName(name)
	if err != nil {
		return []Folder{}, errors.New("Error: Source folder does not exist")
	}
	dstFolder, dstErr := f.GetFolderByName(dst)
	if dstErr != nil {
		return []Folder{}, errors.New("Error: Destination folder does not exist")
	}
	if folder.OrgId != dstFolder.OrgId {
		return []Folder{}, errors.New("Error: Cannot move a folder to a different organization")
	}
	if folder.Paths == dstFolder.Paths {
		return []Folder{}, errors.New("Error: Cannot move a folder to itself")
	}
	// If folder is a parent of dstFolder (already checked the case where their paths are equal)
	matched, _ := regexp.MatchString(folder.Paths, dstFolder.Paths)
	if matched {
		return []Folder{}, errors.New("Error: Cannot move a folder to a child of itself")
	}
	folder.Paths = dstFolder.Paths + "." + folder.Name
	childFolders := f.GetAllChildFolders(folder.OrgId, folder.Name)
	for _, child := range childFolders {
		child.Paths = folder.Paths + "." + child.Name
		for i, fldr := range f.folders {
			if fldr.Name == child.Name {
				f.folders[i] = child
			}
		}
	}
	// Update the folder we just moved
	for i, fldr := range f.folders {
		if fldr.Name == folder.Name {
			f.folders[i] = folder
		}
	}

	return f.folders, nil
}
