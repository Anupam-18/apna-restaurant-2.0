// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: menus.sql

package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const checkExistingMenuitem = `-- name: CheckExistingMenuitem :one
SELECT COUNT(*) AS menuitem_count
FROM menuitems
WHERE id = $1
`

func (q *Queries) CheckExistingMenuitem(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.queryRow(ctx, q.checkExistingMenuitemStmt, checkExistingMenuitem, id)
	var menuitem_count int64
	err := row.Scan(&menuitem_count)
	return menuitem_count, err
}

const createMenu = `-- name: CreateMenu :one

WITH inserted AS (
    INSERT INTO menus (category, menu_item_ids)
    VALUES (
        $1,
        (SELECT ARRAY[$2])
    )
    RETURNING id
)
SELECT id, $1 AS category, $2 AS menu_item_ids
FROM inserted
`

type CreateMenuParams struct {
	Category    string      `json:"category"`
	MenuItemIds []uuid.UUID `json:"menu_item_ids"`
}

type CreateMenuRow struct {
	ID          uuid.UUID   `json:"id"`
	Category    interface{} `json:"category"`
	MenuItemIds interface{} `json:"menu_item_ids"`
}

// -- name: CreateMenu :one
// INSERT INTO menus (category, menu_item_ids)
// VALUES (
//
//	$1,
//	(SELECT ARRAY[$2])
//
// )
// RETURNING *;
// params: category:string, menu_item_ids:array[string]
func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) (CreateMenuRow, error) {
	row := q.queryRow(ctx, q.createMenuStmt, createMenu, arg.Category, pq.Array(arg.MenuItemIds))
	var i CreateMenuRow
	err := row.Scan(&i.ID, &i.Category, &i.MenuItemIds)
	return i, err
}

const createMenuitem = `-- name: CreateMenuitem :one
INSERT INTO menuitems (
  name,
  price,
  image_url,
  menu_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, price, image_url, menu_id
`

type CreateMenuitemParams struct {
	Name     string    `json:"name"`
	Price    int32     `json:"price"`
	ImageUrl string    `json:"image_url"`
	MenuID   uuid.UUID `json:"menu_id"`
}

func (q *Queries) CreateMenuitem(ctx context.Context, arg CreateMenuitemParams) (Menuitem, error) {
	row := q.queryRow(ctx, q.createMenuitemStmt, createMenuitem,
		arg.Name,
		arg.Price,
		arg.ImageUrl,
		arg.MenuID,
	)
	var i Menuitem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.ImageUrl,
		&i.MenuID,
	)
	return i, err
}

const deleteMenuByID = `-- name: DeleteMenuByID :exec
DELETE FROM menus
WHERE id = $1
`

// param: id: uuid
func (q *Queries) DeleteMenuByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteMenuByIDStmt, deleteMenuByID, id)
	return err
}

const deleteMenuitem = `-- name: DeleteMenuitem :exec
DELETE FROM menuitems
WHERE id = $1
`

func (q *Queries) DeleteMenuitem(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteMenuitemStmt, deleteMenuitem, id)
	return err
}

const getAllMenus = `-- name: GetAllMenus :many
SELECT id, category, menu_item_ids FROM menus
ORDER BY id
`

func (q *Queries) GetAllMenus(ctx context.Context) ([]Menu, error) {
	rows, err := q.query(ctx, q.getAllMenusStmt, getAllMenus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Menu
	for rows.Next() {
		var i Menu
		if err := rows.Scan(&i.ID, &i.Category, pq.Array(&i.MenuItemIds)); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMenuByID = `-- name: GetMenuByID :one
SELECT id, category, menu_item_ids
FROM menus
WHERE id = $1
`

// param: id: uuid
func (q *Queries) GetMenuByID(ctx context.Context, id uuid.UUID) (Menu, error) {
	row := q.queryRow(ctx, q.getMenuByIDStmt, getMenuByID, id)
	var i Menu
	err := row.Scan(&i.ID, &i.Category, pq.Array(&i.MenuItemIds))
	return i, err
}

const getMenuitemsById = `-- name: GetMenuitemsById :one
SELECT id, name, price, image_url, menu_id FROM menuitems
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMenuitemsById(ctx context.Context, id uuid.UUID) (Menuitem, error) {
	row := q.queryRow(ctx, q.getMenuitemsByIdStmt, getMenuitemsById, id)
	var i Menuitem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.ImageUrl,
		&i.MenuID,
	)
	return i, err
}

const listMenuitems = `-- name: ListMenuitems :many
SELECT id, name, price, image_url, menu_id FROM menuitems
ORDER BY id
`

func (q *Queries) ListMenuitems(ctx context.Context) ([]Menuitem, error) {
	rows, err := q.query(ctx, q.listMenuitemsStmt, listMenuitems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Menuitem
	for rows.Next() {
		var i Menuitem
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.ImageUrl,
			&i.MenuID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMenu = `-- name: UpdateMenu :one
UPDATE menus
SET
    category = $1,
    menu_item_ids = $2
WHERE
    id = $3
RETURNING id, category, menu_item_ids
`

type UpdateMenuParams struct {
	Category    string      `json:"category"`
	MenuItemIds []uuid.UUID `json:"menu_item_ids"`
	ID          uuid.UUID   `json:"id"`
}

// param: category: string
// param: menu_item_ids: []string
// param: id: uuid
func (q *Queries) UpdateMenu(ctx context.Context, arg UpdateMenuParams) (Menu, error) {
	row := q.queryRow(ctx, q.updateMenuStmt, updateMenu, arg.Category, pq.Array(arg.MenuItemIds), arg.ID)
	var i Menu
	err := row.Scan(&i.ID, &i.Category, pq.Array(&i.MenuItemIds))
	return i, err
}

const updateMenuitem = `-- name: UpdateMenuitem :one
UPDATE menuitems
SET name = $2,
    price = $3,
    image_url = $4
WHERE id = $1
RETURNING id, name, price, image_url, menu_id
`

type UpdateMenuitemParams struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    int32     `json:"price"`
	ImageUrl string    `json:"image_url"`
}

func (q *Queries) UpdateMenuitem(ctx context.Context, arg UpdateMenuitemParams) (Menuitem, error) {
	row := q.queryRow(ctx, q.updateMenuitemStmt, updateMenuitem,
		arg.ID,
		arg.Name,
		arg.Price,
		arg.ImageUrl,
	)
	var i Menuitem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.ImageUrl,
		&i.MenuID,
	)
	return i, err
}
