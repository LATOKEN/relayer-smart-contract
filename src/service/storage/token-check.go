package storage

func (d *DataBase) SaveTokenChecks(checks []interface{}) {
	for _, rID := range checks {
		checkMap := rID.(map[string]interface{})
		check := TokenCheck{
			OriginChainID:      checkMap["origin_chain_id"].(string),
			DestinationChainID: checkMap["destination_chain_id"].(string),
			ResourceID:         checkMap["resource_id"].(string),
			Amount:             int64(checkMap["amount"].(float64)),
		}
		if err := d.saveTokenCheck(&check); err != nil {
			continue
		}
	}
}

func (d *DataBase) FetchSpecificTokenCheck(origin, dest, resourceID string) (tc TokenCheck) {
	d.db.Model(TokenCheck{}).Where("origin_chain_id = ? and destination_chain_id = ? and resource_id = ?", origin, dest, resourceID).First(&tc)
	return tc
}

func (d *DataBase) saveTokenCheck(tc *TokenCheck) error {
	if err := d.db.Model(TokenCheck{}).Create(tc).Error; err != nil {
		return err
	}
	return nil
}
