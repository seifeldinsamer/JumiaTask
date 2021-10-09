package pg

import (
	"backend/entity"
	"backend/pkgs/filters"
	"fmt"
	"github.com/pkg/errors"
)

func (i *Instance) Migrate() {
	i.orm.Migrator().DropTable(&entity.User{})
	i.orm.Migrator().CreateTable(&entity.User{})
}

func (i *Instance) LoadUsers(users []entity.User, filters []filters.Filter) (unHandledReCords []error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Loading csv panic", err)
			panic(err)
		}
	}()

	unHandledReCords = make([]error, 100)

OUTER:
	for recordIndex, record := range users[] {

		for _, filter := range filters {
			if err := filter(&record); err != nil {
				unHandledReCords = append(unHandledReCords, errors.Errorf("Can't filter record of index %d with error %s", recordIndex, err))
				continue OUTER
			}
		}
		i.orm.Create(&record)

	}
	return
}

func (i *Instance) GetTotalCountryUsers(country string) int64 {
	var count int64
	i.orm.Model(&entity.User{}).Where("country = ?", country).Count(&count)

	return count
}

func (i *Instance) GetCountryUsers(country string) ([]entity.User, error) {
	users := make([]entity.User, 100)
	rows, err := i.orm.Model(&entity.User{}).Where("country = ?", country).Select("id, phone_number, email, parcel_weight, country").Rows()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var user entity.User
	for rows.Next() {
		i.orm.ScanRows(rows, &user)
		users = append(users, user)
	}
	return users, nil
}
