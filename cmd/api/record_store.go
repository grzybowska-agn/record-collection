package main

import "context"

func (app *application) insertRecord(ctx context.Context, record *Record) error {
	query := `
		INSERT INTO records (
			artist, title, original_year, cover_url,
			pressing_year, pressing_label, is_reissue,
			size_inches, release_type, disc_count, pressing_country
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, created_at
	`
	return app.db.QueryRowContext(
		ctx,
		query,
		record.Artist,
		record.Title,
		record.OriginalYear,
		record.CoverURL,
		record.PressingYear,
		record.PressingLabel,
		record.IsReissue,
		record.SizeInches,
		record.ReleaseType,
		record.DiscCount,
		record.PressingCountry,
	).Scan(&record.ID, &record.CreatedAt)
}

func (app *application) listRecords(ctx context.Context) ([]Record, error) {
	query := `
		SELECT id, artist, title, original_year,
				COALESCE(cover_url, ''),
		       	pressing_year, COALESCE(pressing_label, ''), is_reissue,
		       	size_inches, release_type, disc_count,
		       	COALESCE(pressing_country, ''), created_at
		FROM records
		ORDER BY created_at DESC, id DESC
	`

	rows, err := app.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records := make([]Record, 0)

	for rows.Next() {
		var record Record

		err := rows.Scan(&record.ID, &record.Artist, &record.Title,
			&record.OriginalYear, &record.CoverURL,
			&record.PressingYear, &record.PressingLabel,
			&record.IsReissue, &record.SizeInches,
			&record.ReleaseType, &record.DiscCount,
			&record.PressingCountry, &record.CreatedAt)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
