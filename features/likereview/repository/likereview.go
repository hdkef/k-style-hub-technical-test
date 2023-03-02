package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"kstylehub/features/likereview"
	"log"
)

type LikeReviewRepository struct {
	db *sql.DB
}

// Delete implements likereview.LikeReviewRepository_
func (u *LikeReviewRepository) Delete(idR, idM int) error {
	res, err := u.db.Exec("DELETE FROM like_review WHERE id_member = ? AND id_review = ?", idM, idR)
	if err != nil {
		return errors.New("internal database error")
	}

	val, err := res.RowsAffected()
	if err != nil {
		return errors.New("internal database error")
	}
	if val == 0 {
		if err != nil {
			return errors.New("no rows affected")
		}
	}
	return nil
}

// Upsert implements likereview.LikeReviewRepository_
func (u *LikeReviewRepository) Upsert(idR, idM int) error {
	//start transaction
	tx, err := u.db.Begin()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Print(err.Error())
		}
		return errors.New("internal database error")
	}

	//if not exist then insert
	var tmp int
	err = tx.QueryRow("SELECT id_review FROM like_review WHERE id_member = ? AND id_review = ?", idM, idR).
		Scan(&tmp)

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		err = tx.Rollback()
		if err != nil {
			log.Print(err.Error())
		}
		return errors.New("already like review")
	}

	res, err := tx.Exec("INSERT INTO like_review(id_member,id_review) VALUES (?,?)", idM, idR)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Print(err.Error())
		}
		fmt.Println("DEBUG 1")
		return errors.New("internal database error")
	}

	val, err := res.RowsAffected()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Print(err.Error())
		}
		fmt.Println("DEBUG 2")
		return errors.New("internal database error")
	}
	if val == 0 {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				log.Print(err.Error())
			}
			return errors.New("no rows affected")
		}
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Print(err.Error())
		}
		fmt.Println("DEBUG 3")
		return errors.New("internal database error")
	}

	return nil
}

func NewLikeReviewRepository(db *sql.DB) likereview.LikeReviewRepository_ {
	return &LikeReviewRepository{
		db: db,
	}
}
