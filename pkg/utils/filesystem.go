package utils

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/djherbis/times.v1"
	"io/ioutil"
	"os"
	//"syscall"
	"time"
)

func (u *Utils) FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (u *Utils) DeleteFileByAge(path string, minAgeForDeletion time.Duration) (bool,error) {
	u.Log.Infof("Deleting file: %s", path)
	/*fileStat,*/_, fileStatErr := os.Stat(path)
	if os.IsNotExist(fileStatErr) {
		u.Log.Errorf("Could not find file %s. %s", path, fileStatErr.Error())
		return false, fileStatErr
	} else {
		t, err := times.Stat(path)
		if err != nil {
			log.Fatal(err.Error())
		}

		//tFileCreationTime := fileStat.Sys().(*syscall.Stat_t).Ctimespec
		//fileCreationTime := time.Unix(tFileCreationTime.Sec, tFileCreationTime.Nsec)
		var fileCreationTime time.Time
		if t.HasBirthTime() {
			fileCreationTime = t.BirthTime()
		}
		if t.HasChangeTime() {
			fileCreationTime = t.ChangeTime()
		}

		tCurrentTime := time.Now()
		tFileAgeForDeletion := int64(minAgeForDeletion) // 10 secs X 60 = 600 secs  OR 10 mins

		if (tCurrentTime.Unix() - fileCreationTime.Unix()) >= tFileAgeForDeletion {
			delFileErr := os.Remove(path)
			if delFileErr != nil {
				u.Log.Errorf("FAILED to remove file %s. %s", path, delFileErr.Error())
				return false, delFileErr
			} else {
				u.Log.Errorf("Successfully Removed file %s", path)
				return true, nil
			}
		} else {
			u.Log.Errorf("Specified File: %s is newer than specified deletion age of - %d second(s). WONT DELETE!", path, minAgeForDeletion)
			return false, errors.New(fmt.Sprintf("Specified File: %s is newer than specified deletion age of - %d second(s). WONT DELETE!", path, minAgeForDeletion))
		}
	}
}


func (u *Utils) GetFileList( directoryPath string ) map[int]string {
	filesList := make(map[int]string)
	filesListIterator := 0

	u.Log.Infof("directoryPath = %s", directoryPath)
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		u.Log.Errorf("Error reading directories: %s", err.Error())
		return filesList
	}

	for fileKey, file := range files {
		fileStat, fileStatErr := os.Stat(file.Name())
		if fileStatErr != nil {
			u.Log.Errorf("Error stating file #%d: %s. %s", fileKey, file.Name(), fileStatErr.Error())
			continue
		}

		if fileStat.IsDir() {
			tFilesList := u.GetFileList(file.Name())
			tFilesListLen := len(tFilesList)
			if tFilesListLen > 0 {
				for _, value := range tFilesList {
					filesList[filesListIterator] = value
					filesListIterator++
				}
			}
			continue
		}

		filesList[filesListIterator] = file.Name()
		filesListIterator++
	}
	return filesList
}

func (u *Utils) DeleteFileList( fileExtToClean, directoryPath string ) int {
	deletionCount := 0
	if fileExtToClean == "" {
		fileExtToClean = "pdf"
	}

	fileExtToClean = fmt.Sprintf("*%s", fileExtToClean)

	if directoryPath == "" {
		u.Log.Errorf("Must specify Directory Path where files need to deleted")
	} else {
		if u.FileExists( directoryPath ) {
			filesList := u.GetFileList(fmt.Sprintf("%s/%s", directoryPath, fileExtToClean) )
			filesListLen := len(filesList)
			if filesListLen > 0 {
				for _, filePath := range filesList {
					fileDel, fileDelErr := u.DeleteFileByAge(filePath, 600*time.Second)
					if fileDelErr == nil {
						if fileDel {
							deletionCount++
						}
					} else {
						u.Log.Errorf("Error deleting file %s. %s", filePath, fileDelErr.Error())
					}
				}
			} else {
				u.Log.Infof("No files found in directory %s", directoryPath)
			}
		} else {
			u.Log.Infof("Path %s does not exist", directoryPath)
		}
	}

	u.Log.Infof("Deleted %d %s files at path %s", deletionCount, fileExtToClean, directoryPath)
	return deletionCount
}

// DeleteFileListE( fileExtToClean, directoryPath string) int
// Empathetic version of the DeleteFileList() function. Checks system load average
// for overages and depending upon the result pauses or keeps deletion going
func (u *Utils) DeleteFileListE( fileExtToClean, directoryPath string ) int {
	deletionCount := 0
	if fileExtToClean == "" {
		fileExtToClean = "pdf"
	}

	fileExtToClean = fmt.Sprintf("*%s", fileExtToClean)

	if directoryPath == "" {
		u.Log.Errorf("Must specify Directory Path where files need to deleted")
	} else {
		if u.FileExists( directoryPath ) {
			filesList := u.GetFileList(fmt.Sprintf("%s/%s", directoryPath, fileExtToClean))
			filesListLen := len(filesList)
			if filesListLen > 0 {
				for _, filePath := range filesList {
					if u.LoadAvgCheck() == LAVG_TREND_NORMAL {
						fileDel, fileDelErr := u.DeleteFileByAge(filePath, 600*time.Second)
						if fileDelErr == nil {
							if fileDel {
								deletionCount++
							}
						} else {
							u.Log.Errorf("Error deleting file %s. %s", filePath, fileDelErr.Error())
						}
					} else {
						loadAvg, loadAvgErr := u.LoadAvg()
						if loadAvgErr != nil {
							u.Log.Errorf("Unable to read System Load Average. %s", loadAvgErr.Error())
						} else {
							u.Log.Infof("Load Average (1), (5), (15) = (%f), (%f), (%f)", loadAvg.Load1, loadAvg.Load5, loadAvg.Load15)
						}
						u.Log.Infof("Sleeping for %d seconds...", DEL_FILE_LIST_E_SLEEP)
						time.Sleep(DEL_FILE_LIST_E_SLEEP * time.Second)
						u.Log.Infof("Woke-up!!!")
					}
				}
			} else {
				u.Log.Infof("No files found in directory %s", directoryPath)
			}
		} else {
			u.Log.Infof("Path %s does not exist", directoryPath)
		}
	}

	u.Log.Infof("Deleted %d %s files at path %s", deletionCount, fileExtToClean, directoryPath)
	return deletionCount
}

