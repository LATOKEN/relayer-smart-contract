package storage

func (d *DataBase) SaveTokenChecks(checks []*TokenCheck) {
	for _, rID := range checks {
		if err := d.saveTokenCheck(rID); err != nil {
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
