package organization

type OrganizationService struct {
	Repo *OrganizationRepository
}

func (s *OrganizationService) CreateOrganization(org Organization) (int, error) {
	return s.Repo.CreateOrganization(org)
}

func (s *OrganizationService) GetOrganization(id int) (Organization, error) {
	return s.Repo.GetOrganization(id)
}

func (s *OrganizationService) UpdateOrganization(org Organization) error {
	return s.Repo.UpdateOrganization(org)
}

func (s *OrganizationService) DeleteOrganization(id int) error {
	return s.Repo.DeleteOrganization(id)
}

func (s *OrganizationService) GetAllOrganizations() ([]Organization, error) {
	return s.Repo.GetAllOrganizations()
}
