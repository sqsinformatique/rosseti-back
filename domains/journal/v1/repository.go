package journalv1

import (
	"time"

	actsdetailv1 "github.com/sqsinformatique/rosseti-back/domains/acts_detail/v1"
	"github.com/sqsinformatique/rosseti-back/models"
	"github.com/sqsinformatique/rosseti-back/types"
)

func (j *JournalV1) GetJournalData() (data *ArrayOfJournalData, err error) {
	data = &ArrayOfJournalData{}

	var startTime, endTime types.NullTime
	startTime.Time = time.Now()
	startTime.Valid = true
	endTime = startTime
	endTime.Valid = true

	startTime.Time = startTime.Time.Add(-time.Hour * 24 * 30)

	acts, err := j.actV1.GetActsByDate(startTime, endTime)
	if err != nil {
		return nil, err
	}

	for _, v := range *acts {
		for _, w := range *v.ActDetailDesc.(*actsdetailv1.ArrayOfActsDetailData) {
			var item models.JournalItem

			item.DefectsDesc = w.DefectsDesc
			item.Category = w.Category
			item.ObjectDesc = v.ObjectDesc
			item.FindAd = v.StaffSignAt
			item.ElementDesc = w.ElementDesc
			item.ElementTypeDesk, err = j.elementequipmentV1.GetElementEquipmentByID(int64(w.ElementDesc.(*models.ObjectsDetail).ElementEqupment))
			if err != nil {
				return nil, err
			}
			item.StaffDesc = v.StaffDesc

			var maxCategory int
			for _, v := range item.DefectsDesc.([]*models.Defect) {
				if maxCategory < v.Сategory {
					maxCategory = v.Сategory
				}
			}
			item.Category = maxCategory
			category, err := j.categoryV1.GetCategoryByID(int64(maxCategory))
			if err != nil {
				return nil, err
			}

			item.RapairAt.Time = item.FindAd.Time.Add(time.Duration(category.RapairPeriod) * time.Hour * 24)

			*data = append(*data, item)
		}
	}

	return data, nil
}
