package utils

import (
	"context"
	"strings"

	repo "apna-restaurant-2.0/db/sqlc"
	"github.com/google/uuid"
)

func ValidateMenuItem(item uuid.UUID, db *repo.Queries) (string, bool) {
	menuitemCount, err := db.CheckExistingMenuitem(context.Background(), item)
	if err != nil {
		return "Internal server error", false
	}
	if menuitemCount != 1 {
		return "One of the Menuitems does not exist", false
	}
	return "", true
}

func ValidateAddMenuRequest(menu *repo.Menu, db *repo.Queries) (string, bool) {
	if len(strings.TrimSpace(menu.Category)) == 0 {
		return "Category required", false
	} else if len(strings.TrimSpace(menu.Category)) < 5 {
		return "Name should be at least 5 chars", false
	}
	if (len(menu.MenuItemIds)) > 0 {
		for _, menuitem := range menu.MenuItemIds {
			if resp, ok := ValidateMenuItem(menuitem, db); !ok {
				return resp, false
			}
		}
	}
	return "requirement passed", true
}
