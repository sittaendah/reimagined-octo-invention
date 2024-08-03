package organization

import (
	"database/sql"
)

type OrganizationRepository struct {
	DB *sql.DB
}

func (r *OrganizationRepository) CreateOrganization(org Organization) (int, error) {
	var id int
	err := r.DB.QueryRow(`INSERT INTO organizations(name, description) VALUES($1, $2) RETURNING id`, org.Name, org.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrganizationRepository) GetOrganization(id int) (Organization, error) {
	var org Organization
	row := r.DB.QueryRow(`SELECT id, name, description FROM organizations WHERE id = $1`, id)
	err := row.Scan(&org.ID, &org.Name, &org.Description)
	if err != nil {
		return Organization{}, err
	}
	return org, nil
}

func (r *OrganizationRepository) UpdateOrganization(org Organization) error {
	_, err := r.DB.Exec(`UPDATE organizations SET name = $1, description = $2 WHERE id = $3`, org.Name, org.Description, org.ID)
	return err
}

func (r *OrganizationRepository) DeleteOrganization(id int) error {
	_, err := r.DB.Exec(`DELETE FROM organizations WHERE id = $1`, id)
	return err
}

func (r *OrganizationRepository) GetAllOrganizations() ([]Organization, error) {
	rows, err := r.DB.Query(`SELECT id, name, description FROM organizations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orgs []Organization
	for rows.Next() {
		var org Organization
		if err := rows.Scan(&org.ID, &org.Name, &org.Description); err != nil {
			return nil, err
		}
		orgs = append(orgs, org)
	}
	return orgs, nil
}
