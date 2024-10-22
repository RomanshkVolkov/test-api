package service

import (
	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
)

func (server Server) GetUsersProfiles() (domain.APIResponse[domain.UserProfiles, any], error) {
	repo := repository.GetDBConnection(server.Host)
	profiles, err := repo.GetUsersProfiles()
	if err != nil {
		return domain.APIResponse[domain.UserProfiles, any]{
			Success: false,
			Message: domain.Message{
				En: "Error on get users profiles",
				Es: "Error al obtener perfiles de usuarios",
			},
			Error: err,
		}, err
	}

	return domain.APIResponse[domain.UserProfiles, any]{
		Success: true,
		Message: domain.Message{
			En: "Users profiles",
			Es: "Perfiles de usuarios",
		},
		Data: profiles,
	}, nil
}
