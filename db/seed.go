package db

import (
	"sort"

	"api-recipe-keeper/models"
)

func Seed() {
	hasNewSeed := false
	setting := models.Setting{Key: "db.seed.version"}
	db.Where(models.Setting{Key: setting.Key}).FirstOrCreate(&setting)

	index := make([]string, 0)
	for i, _ := range seed {
		index = append(index, i)
	}
	sort.Strings(index)
	for _, i := range index {
		if setting.Value == "" || setting.Value < i {
			seed[i]()
			setting.Value = i
			hasNewSeed = true
		}
	}
	if hasNewSeed {
		db.Where(models.Setting{Key: setting.Key}).Assign(setting).FirstOrCreate(&setting)
	}
}

var seed = map[string]func(){}
