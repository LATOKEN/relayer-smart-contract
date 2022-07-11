package storage

func (d *DataBase) SaveTokenChecks(checks []interface{}) {
	for _, rID := range checks {
		checkMap := rID.(map[string]interface{})
		check := TokenCheck{
			OriginChain:      checkMap["origin_chain"].(string),
			DestinationChain: checkMap["destination_chain"].(string),
			ResourceID:       checkMap["resource_id"].(string),
			Amount:           int64(checkMap["amount"].(float64)),
		}
		_, err := d.FetchSpecificTokenCheck(check.OriginChain, check.DestinationChain, check.ResourceID)
		if err != nil {
			if err := d.saveTokenCheck(&check); err != nil {
				continue
			}
		} else {
			if err := d.updateSpecificTokenCheck(check); err != nil {
				println("error updating token check")
			}
		}

	}
}

func (d *DataBase) FetchSpecificTokenCheck(origin, dest, resourceID string) (tc TokenCheck, err error) {
	if err = d.db.Model(TokenCheck{}).Where("origin_chain = ? and destination_chain = ? and resource_id = ?", origin, dest, resourceID).First(&tc).Error; err != nil {
		return tc, err
	}
	return tc, nil
}

func (d *DataBase) updateSpecificTokenCheck(tc TokenCheck) error {
	if err := d.db.Model(TokenCheck{}).Where("origin_chain = ? and destination_chain = ? and resource_id = ?", tc.OriginChain, tc.DestinationChain, tc.ResourceID).Updates(tc).Error; err != nil {
		return err
	}
	return nil
}

func (d *DataBase) saveTokenCheck(tc *TokenCheck) error {
	if err := d.db.Model(TokenCheck{}).Create(tc).Error; err != nil {
		return err
	}
	return nil
}
