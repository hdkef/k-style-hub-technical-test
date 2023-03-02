package repository

import (
	"database/sql"
	"errors"
	"kstylehub/features/member"
	"kstylehub/features/member/models"
)

type MemberRepository struct {
	db *sql.DB
}

// Delete implements member.MemberRepository_
func (u *MemberRepository) Delete(id int) error {
	res, err := u.db.Exec("DELETE FROM member WHERE id_member = ?", id)
	if err != nil {
		return errors.New("internal database error")
	}
	val, err := res.RowsAffected()
	if err != nil {
		return errors.New("internal database error")
	}
	if val == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

// GetAll implements member.MemberRepository_
func (u *MemberRepository) GetAll(page int, limit int) ([]member.Member, error) {
	var data []models.Member
	offset := (page - 1) * limit
	rows, err := u.db.Query("SELECT id_member, username, gender, skintype, skincolor FROM member LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, errors.New("internal database error")
	}
	for rows.Next() {
		var member models.Member
		err = rows.Scan(&member.ID, &member.Username, &member.Gender, &member.SkinType, &member.SkinColor)
		if err != nil {
			return nil, errors.New("internal database error")
		}
		data = append(data, member)
	}
	return convertToCoreList(data), nil
}

// GetOneByID implements member.MemberRepository_
func (u *MemberRepository) GetOneByID(id int) (member.Member, error) {
	var data models.Member
	err := u.db.QueryRow("SELECT id_member, username, gender, skintype, skincolor FROM member WHERE id_member = ?", id).
		Scan(&data.ID, &data.Username, &data.Gender, &data.SkinType, &data.SkinColor)
	if err != nil {
		return member.Member{}, errors.New("internal database error")
	}
	return convertToCore(&data), nil
}

// Insert implements member.MemberRepository_
func (u *MemberRepository) Insert(member member.Member) error {
	data := convertToModels(&member)
	res, err := u.db.Exec("INSERT INTO member(id_member,username,gender,skintype,skincolor) VALUES (?,?,?,?,?)",
		data.ID, data.Username, data.Gender, data.SkinType, data.SkinColor)
	if err != nil {
		return errors.New("internal database error")
	}
	val, err := res.RowsAffected()
	if err != nil {
		return errors.New("internal database error")
	}
	if val == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

// Update implements member.MemberRepository_
func (u *MemberRepository) Update(id int, member member.Member) error {
	data := convertToModels(&member)
	res, err := u.db.Exec("UPDATE member SET username = ?,gender = ?,skintype = ?,skincolor = ? WHERE id_member = ?",
		data.Username, data.Gender, data.SkinType, data.SkinColor, id)
	if err != nil {
		return errors.New("internal database error")
	}
	val, err := res.RowsAffected()
	if err != nil {
		return errors.New("internal database error")
	}
	if val == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func NewMemberRespository(db *sql.DB) member.MemberRepository_ {
	return &MemberRepository{
		db: db,
	}
}
