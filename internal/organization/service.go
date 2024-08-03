package organization

import (
	"errors"
	user "github.com/sittaendah/aegis/internal/user"
)

type OrganizationService struct {
	Repo     *OrganizationRepository
	UserRepo *user.UserRepository
}

func (s *OrganizationService) CreateOrganization(org Organization) (int, error) {
	return s.Repo.CreateOrganization(org)
}

func (s *OrganizationService) GetOrganization(id int) (Organization, error) {
	return s.Repo.GetOrganization(id)
}

func (s *OrganizationService) UpdateOrganization(org Organization, username string) error {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user.Role != "ADMIN" {
		return errors.New("insufficient permissions")
	}

	existingOrg, err := s.Repo.GetOrganization(org.ID)
	if err != nil {
		return err
	}
	if existingOrg.CreatedBy != user.Username && user.Role != "ADMIN" {
		return errors.New("user not authorized")
	}
	return s.Repo.UpdateOrganization(org)
}

func (s *OrganizationService) DeleteOrganization(id int, username string) error {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user.Role != "ADMIN" {
		return errors.New("insufficient permissions")
	}

	org, err := s.Repo.GetOrganization(id)
	if err != nil {
		return err
	}
	if org.CreatedBy != user.Username && user.Role != "ADMIN" {
		return errors.New("user not authorized")
	}
	return s.Repo.DeleteOrganization(id)
}

func (s *OrganizationService) GetAllOrganizations() ([]Organization, error) {
	return s.Repo.GetAllOrganizations()
}
