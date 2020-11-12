package utils

import (
	"database/sql"
	"fmt"
	"github.com/gchaincl/dotsql"
	log "github.com/sirupsen/logrus"
)

type Migrator struct {
	ProceedOnError bool
}

func (m *Migrator) Values(proceedOnError bool) {
	m.ProceedOnError = proceedOnError
}

func (m *Migrator) Migrate(db *sql.DB, filePath string) error {
	var retErr error
	dot, dotErr := dotsql.LoadFromFile(filePath)

	if dotErr != nil {
		log.Errorf("Error loading migration file %s. %s", filePath, dotErr.Error())
		return dotErr
	}

	queryMap := dot.QueryMap()

	totalQueries := len(queryMap)

	logErr := true

	for i := 0; i < totalQueries; i++ {
		queryName := "Q" + fmt.Sprintf("%03d", i)
		log.Infof("SQLEXEC:: %s - %s", queryName, queryMap[queryName])
		res, resErr := dot.Exec(db, queryName)
		if resErr != nil {
			log.Errorf("FAILED to execute query %s", queryName)
			log.Errorf("%s", resErr.Error())
			if m.ProceedOnError == false {
				return resErr
			}
		} else {
			lastInertID, lastInsertIDErr := res.LastInsertId()
			if lastInsertIDErr != nil {
				if logErr {
					log.Errorf("Error getting LAST_INSERT_ID(). %s", lastInsertIDErr.Error())
				}
				retErr = lastInsertIDErr
			} else {
				log.Infof("%s - LAST_INSERT_ID() = %d", queryName, lastInertID)
				logErr = false
			}

			rowsAffected, rowsAffectedErr := res.RowsAffected()
			if rowsAffectedErr != nil {
				if logErr {
					log.Errorf("Error getting affected rows. %s", rowsAffectedErr.Error())
				}
				retErr = rowsAffectedErr
			} else {
				log.Infof("%s - Rows Affected = %d", queryName, rowsAffected)
				logErr = false
			}
		}
	}
	return retErr
}