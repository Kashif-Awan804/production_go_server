package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Listing struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func GetListings(db *pgxpool.Pool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(
			r.Context(),
			`SELECT id, title, description, price FROM listings ORDER BY id`,
		)

		if err != nil {
			http.Error(w, "Failed to fetch listings", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var listings []Listing

		for rows.Next() {

			var listing Listing

			err := rows.Scan(
				&listing.ID,
				&listing.Title,
				&listing.Description,
				&listing.Price,
			)

			if err != nil {
				http.Error(w, "Failed to read listing", http.StatusInternalServerError)
				return
			}

			listings = append(listings, listing)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Failed to read database rows", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(listings)
	}
}