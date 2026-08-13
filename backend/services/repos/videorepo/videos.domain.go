package videorepo

func (v *VideoRepository) CreateThumbnailDomain() (bool, error) {
	_, err := v.DB.Exec(`
		DO $$
		BEGIN
			CREATE DOMAIN thumbnail_domain AS TEXT 
				CHECK (VALUE ~* '^thumbnails/[A-Za-z0-9_-]+\.(png|jpg|jpeg|webp)$');
			EXCEPTION
				WHEN duplicate_object THEN NULL;
		END
		$$;
	`)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (v *VideoRepository) CreateVideoDomain() (bool, error) {
	_, err := v.DB.Exec(`
		DO $$
		BEGIN
			CREATE DOMAIN video_domain AS TEXT
				CHECK (VALUE ~* '^videos/[A-Za-z0-9_-]/+\.mp4$');
			EXCEPTION
				WHEN duplicate_object THEN NULL;
			END
			$$;
	`)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (v *VideoRepository) CreateDomains() (error) {
	videodom, err := v.CreateVideoDomain()

	if !videodom {
		return err
	}

	thumbdom, err := v.CreateThumbnailDomain()

	if !thumbdom {
		return err
	}

	return nil
}