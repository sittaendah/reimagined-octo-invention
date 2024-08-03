package organization

import (
	"database/sql"
)

type OrganizationRepository struct {
	DB *sql.DB
}

func (r *OrganizationRepository) CreateOrganization(org Organization) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO organizations(name, description, created_by) VALUES($1, $2, $3) RETURNING id`,
		org.Name, org.Description, org.CreatedBy,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrganizationRepository) GetOrganization(id int) (Organization, error) {
	var org Organization
	row := r.DB.QueryRow(`SELECT id, name, description, created_by FROM organizations WHERE id = $1`, id)
	err := row.Scan(&org.ID, &org.Name, &org.Description, &org.CreatedBy)
	if err != nil {
		return Organization{}, err
	}
	return org, nil
}

func (r *OrganizationRepository) UpdateOrganization(org Organization) error {
	_, err := r.DB.Exec(
		`UPDATE organizations SET name = $1, description = $2, updated_by = $3 WHERE id = $4`,
		org.Name, org.Description, org.UpdatedBy, org.ID,
	)
	return err
}

func (r *OrganizationRepository) DeleteOrganization(id int) error {
	_, err := r.DB.Exec(`DELETE FROM organizations WHERE id = $1`, id)
	return err
}

func (r *OrganizationRepository) GetAllOrganizations() ([]Organization, error) {
	rows, err := r.DB.Query(`SELECT id, name, description, created_by FROM organizations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orgs []Organization
	for rows.Next() {
		var org Organization
		if err := rows.Scan(&org.ID, &org.Name, &org.Description, &org.CreatedBy); err != nil {
			return nil, err
		}
		orgs = append(orgs, org)
	}
	return orgs, nil
}
