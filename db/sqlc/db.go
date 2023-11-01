// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package repo

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.checkExistingMenuStmt, err = db.PrepareContext(ctx, checkExistingMenu); err != nil {
		return nil, fmt.Errorf("error preparing query CheckExistingMenu: %w", err)
	}
	if q.checkExistingMenuitemStmt, err = db.PrepareContext(ctx, checkExistingMenuitem); err != nil {
		return nil, fmt.Errorf("error preparing query CheckExistingMenuitem: %w", err)
	}
	if q.checkExistingOrderStmt, err = db.PrepareContext(ctx, checkExistingOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CheckExistingOrder: %w", err)
	}
	if q.checkExistingTableStmt, err = db.PrepareContext(ctx, checkExistingTable); err != nil {
		return nil, fmt.Errorf("error preparing query CheckExistingTable: %w", err)
	}
	if q.checkExistingUserStmt, err = db.PrepareContext(ctx, checkExistingUser); err != nil {
		return nil, fmt.Errorf("error preparing query CheckExistingUser: %w", err)
	}
	if q.createMenuStmt, err = db.PrepareContext(ctx, createMenu); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMenu: %w", err)
	}
	if q.createMenuitemStmt, err = db.PrepareContext(ctx, createMenuitem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMenuitem: %w", err)
	}
	if q.createOrderStmt, err = db.PrepareContext(ctx, createOrder); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrder: %w", err)
	}
	if q.createTableStmt, err = db.PrepareContext(ctx, createTable); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTable: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteMenuByIDStmt, err = db.PrepareContext(ctx, deleteMenuByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMenuByID: %w", err)
	}
	if q.deleteMenuitemStmt, err = db.PrepareContext(ctx, deleteMenuitem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMenuitem: %w", err)
	}
	if q.deleteOrderByIDStmt, err = db.PrepareContext(ctx, deleteOrderByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOrderByID: %w", err)
	}
	if q.deleteTableByIDStmt, err = db.PrepareContext(ctx, deleteTableByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTableByID: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAllMenusStmt, err = db.PrepareContext(ctx, getAllMenus); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllMenus: %w", err)
	}
	if q.getAllOrdersStmt, err = db.PrepareContext(ctx, getAllOrders); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllOrders: %w", err)
	}
	if q.getAllTablesStmt, err = db.PrepareContext(ctx, getAllTables); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllTables: %w", err)
	}
	if q.getMenuByIDStmt, err = db.PrepareContext(ctx, getMenuByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetMenuByID: %w", err)
	}
	if q.getMenuitemsByIdStmt, err = db.PrepareContext(ctx, getMenuitemsById); err != nil {
		return nil, fmt.Errorf("error preparing query GetMenuitemsById: %w", err)
	}
	if q.getOrderByIDStmt, err = db.PrepareContext(ctx, getOrderByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetOrderByID: %w", err)
	}
	if q.getTableByIDStmt, err = db.PrepareContext(ctx, getTableByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetTableByID: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.listMenuitemsStmt, err = db.PrepareContext(ctx, listMenuitems); err != nil {
		return nil, fmt.Errorf("error preparing query ListMenuitems: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateMenuStmt, err = db.PrepareContext(ctx, updateMenu); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMenu: %w", err)
	}
	if q.updateMenuitemStmt, err = db.PrepareContext(ctx, updateMenuitem); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMenuitem: %w", err)
	}
	if q.updateOrderStmt, err = db.PrepareContext(ctx, updateOrder); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOrder: %w", err)
	}
	if q.updateTableStmt, err = db.PrepareContext(ctx, updateTable); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTable: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkExistingMenuStmt != nil {
		if cerr := q.checkExistingMenuStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkExistingMenuStmt: %w", cerr)
		}
	}
	if q.checkExistingMenuitemStmt != nil {
		if cerr := q.checkExistingMenuitemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkExistingMenuitemStmt: %w", cerr)
		}
	}
	if q.checkExistingOrderStmt != nil {
		if cerr := q.checkExistingOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkExistingOrderStmt: %w", cerr)
		}
	}
	if q.checkExistingTableStmt != nil {
		if cerr := q.checkExistingTableStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkExistingTableStmt: %w", cerr)
		}
	}
	if q.checkExistingUserStmt != nil {
		if cerr := q.checkExistingUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkExistingUserStmt: %w", cerr)
		}
	}
	if q.createMenuStmt != nil {
		if cerr := q.createMenuStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMenuStmt: %w", cerr)
		}
	}
	if q.createMenuitemStmt != nil {
		if cerr := q.createMenuitemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMenuitemStmt: %w", cerr)
		}
	}
	if q.createOrderStmt != nil {
		if cerr := q.createOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrderStmt: %w", cerr)
		}
	}
	if q.createTableStmt != nil {
		if cerr := q.createTableStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTableStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteMenuByIDStmt != nil {
		if cerr := q.deleteMenuByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMenuByIDStmt: %w", cerr)
		}
	}
	if q.deleteMenuitemStmt != nil {
		if cerr := q.deleteMenuitemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMenuitemStmt: %w", cerr)
		}
	}
	if q.deleteOrderByIDStmt != nil {
		if cerr := q.deleteOrderByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOrderByIDStmt: %w", cerr)
		}
	}
	if q.deleteTableByIDStmt != nil {
		if cerr := q.deleteTableByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTableByIDStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAllMenusStmt != nil {
		if cerr := q.getAllMenusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllMenusStmt: %w", cerr)
		}
	}
	if q.getAllOrdersStmt != nil {
		if cerr := q.getAllOrdersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllOrdersStmt: %w", cerr)
		}
	}
	if q.getAllTablesStmt != nil {
		if cerr := q.getAllTablesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllTablesStmt: %w", cerr)
		}
	}
	if q.getMenuByIDStmt != nil {
		if cerr := q.getMenuByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMenuByIDStmt: %w", cerr)
		}
	}
	if q.getMenuitemsByIdStmt != nil {
		if cerr := q.getMenuitemsByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMenuitemsByIdStmt: %w", cerr)
		}
	}
	if q.getOrderByIDStmt != nil {
		if cerr := q.getOrderByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOrderByIDStmt: %w", cerr)
		}
	}
	if q.getTableByIDStmt != nil {
		if cerr := q.getTableByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTableByIDStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.listMenuitemsStmt != nil {
		if cerr := q.listMenuitemsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMenuitemsStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateMenuStmt != nil {
		if cerr := q.updateMenuStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMenuStmt: %w", cerr)
		}
	}
	if q.updateMenuitemStmt != nil {
		if cerr := q.updateMenuitemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMenuitemStmt: %w", cerr)
		}
	}
	if q.updateOrderStmt != nil {
		if cerr := q.updateOrderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOrderStmt: %w", cerr)
		}
	}
	if q.updateTableStmt != nil {
		if cerr := q.updateTableStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTableStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                        DBTX
	tx                        *sql.Tx
	checkExistingMenuStmt     *sql.Stmt
	checkExistingMenuitemStmt *sql.Stmt
	checkExistingOrderStmt    *sql.Stmt
	checkExistingTableStmt    *sql.Stmt
	checkExistingUserStmt     *sql.Stmt
	createMenuStmt            *sql.Stmt
	createMenuitemStmt        *sql.Stmt
	createOrderStmt           *sql.Stmt
	createTableStmt           *sql.Stmt
	createUserStmt            *sql.Stmt
	deleteMenuByIDStmt        *sql.Stmt
	deleteMenuitemStmt        *sql.Stmt
	deleteOrderByIDStmt       *sql.Stmt
	deleteTableByIDStmt       *sql.Stmt
	deleteUserStmt            *sql.Stmt
	getAllMenusStmt           *sql.Stmt
	getAllOrdersStmt          *sql.Stmt
	getAllTablesStmt          *sql.Stmt
	getMenuByIDStmt           *sql.Stmt
	getMenuitemsByIdStmt      *sql.Stmt
	getOrderByIDStmt          *sql.Stmt
	getTableByIDStmt          *sql.Stmt
	getUserByEmailStmt        *sql.Stmt
	getUserByIdStmt           *sql.Stmt
	listMenuitemsStmt         *sql.Stmt
	listUsersStmt             *sql.Stmt
	updateMenuStmt            *sql.Stmt
	updateMenuitemStmt        *sql.Stmt
	updateOrderStmt           *sql.Stmt
	updateTableStmt           *sql.Stmt
	updateUserStmt            *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                        tx,
		tx:                        tx,
		checkExistingMenuStmt:     q.checkExistingMenuStmt,
		checkExistingMenuitemStmt: q.checkExistingMenuitemStmt,
		checkExistingOrderStmt:    q.checkExistingOrderStmt,
		checkExistingTableStmt:    q.checkExistingTableStmt,
		checkExistingUserStmt:     q.checkExistingUserStmt,
		createMenuStmt:            q.createMenuStmt,
		createMenuitemStmt:        q.createMenuitemStmt,
		createOrderStmt:           q.createOrderStmt,
		createTableStmt:           q.createTableStmt,
		createUserStmt:            q.createUserStmt,
		deleteMenuByIDStmt:        q.deleteMenuByIDStmt,
		deleteMenuitemStmt:        q.deleteMenuitemStmt,
		deleteOrderByIDStmt:       q.deleteOrderByIDStmt,
		deleteTableByIDStmt:       q.deleteTableByIDStmt,
		deleteUserStmt:            q.deleteUserStmt,
		getAllMenusStmt:           q.getAllMenusStmt,
		getAllOrdersStmt:          q.getAllOrdersStmt,
		getAllTablesStmt:          q.getAllTablesStmt,
		getMenuByIDStmt:           q.getMenuByIDStmt,
		getMenuitemsByIdStmt:      q.getMenuitemsByIdStmt,
		getOrderByIDStmt:          q.getOrderByIDStmt,
		getTableByIDStmt:          q.getTableByIDStmt,
		getUserByEmailStmt:        q.getUserByEmailStmt,
		getUserByIdStmt:           q.getUserByIdStmt,
		listMenuitemsStmt:         q.listMenuitemsStmt,
		listUsersStmt:             q.listUsersStmt,
		updateMenuStmt:            q.updateMenuStmt,
		updateMenuitemStmt:        q.updateMenuitemStmt,
		updateOrderStmt:           q.updateOrderStmt,
		updateTableStmt:           q.updateTableStmt,
		updateUserStmt:            q.updateUserStmt,
	}
}
