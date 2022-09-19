package clientdb

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	"time"
)

// DBModel holds the database
type DBModel struct {
	DB *sql.DB
}

// GetPTMember gets a member
func (m *DBModel) GetPTMember(id int) (clientmodels.PTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s clientmodels.PTMember

	query := `select e.id, e.first_name, e.email, e.voted
			from pt_members e
			where e.id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&s.ID,
		&s.FirstName,
		&s.Email,
		&s.Voted,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}

// GetFTMember gets a member
func (m *DBModel) GetFTMember(id int) (clientmodels.FTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s clientmodels.FTMember

	query := `select e.id, e.first_name, e.email, e.voted
			from ft_members e
			where e.id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&s.ID,
		&s.FirstName,
		&s.Email,
		&s.Voted,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}

// GetFTMember gets a member
func (m *DBModel) GetPTMemberByEmail(email string) (clientmodels.PTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s clientmodels.PTMember

	query := `select e.id, e.first_name, e.email, e.voted
			from pt_members e
			where e.email = $1`
	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&s.ID,
		&s.FirstName,
		&s.Email,
		&s.Voted,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}

// GetFTMember gets a member
func (m *DBModel) GetFTMemberByEmail(email string) (clientmodels.FTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s clientmodels.FTMember

	query := `select e.id, e.first_name, e.email, e.voted
			from ft_members e
			where e.email = $1`
	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&s.ID,
		&s.FirstName,
		&s.Email,
		&s.Voted,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}

// VoteYesPT votes yes
func (m *DBModel) VoteYesFT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set yes = yes + 1 where id = 1`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update ft_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// VoteNoFT votes no
func (m *DBModel) VoteNoFT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set no = no + 1 where id = 1`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update ft_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// VoteYesPT votes yes
func (m *DBModel) VoteYesPT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set yes = yes + 1 where id = 2`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update pt_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// VoteNoPT votes no
func (m *DBModel) VoteNoPT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set no = no + 1 where id = 2`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update pt_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllFTMembers get all ft members
func (m *DBModel) GetAllFTMembers() ([]clientmodels.FTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var members []clientmodels.FTMember

	query := "select id, first_name, email, voted from ft_members order by id"

	rows, err := m.DB.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		s := &clientmodels.FTMember{}
		err = rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.Email,
			&s.Voted,
		)
		members = append(members, *s)
	}
	return members, nil
}

// GetAllPTMembers gets all pt members
func (m *DBModel) GetAllPTMembers() ([]clientmodels.PTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var members []clientmodels.PTMember

	query := "select id, first_name, email, voted from pt_members order by id"

	rows, err := m.DB.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		s := &clientmodels.PTMember{}
		err = rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.Email,
			&s.Voted,
		)
		members = append(members, *s)
	}
	return members, nil
}

func (m *DBModel) GetPTResults() (int, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var y, n int

	query := "select yes, no from vote_totals where id = 2"

	row := m.DB.QueryRowContext(ctx, query)
	err := row.Scan(
		&y,
		&n,
	)
	if err != nil {
		return y, n, err
	}
	return y, n, nil
}

func (m *DBModel) GetFTResults() (int, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var y, n int

	query := "select yes, no from vote_totals where id = 1"

	row := m.DB.QueryRowContext(ctx, query)
	err := row.Scan(
		&y,
		&n,
	)
	if err != nil {
		return y, n, err
	}
	return y, n, nil
}
