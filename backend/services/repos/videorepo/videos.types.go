package videorepo

func (v *VideoRepository) CreateTypes() (bool, error) {
	_, err := v.DB.Exec(`
		DO $$
		BEGIN
			CREATE TYPE VIDEOSTATUS AS ENUM (
     			'Uploading',
     			'Encoding',
     			'Ready',
     			'Hosting',
     			'Deleted'
			);
			EXCEPTION
				WHEN	duplicate_object THEN NULL;
			END
			$$;
	`)

	if err != nil {
		return false, err
	}
	return true, nil
}