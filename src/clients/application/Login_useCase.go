package application

import (
	"ejercicio1/src/clients/domain"
	"ejercicio1/src/clients/domain/entities"
	"ejercicio1/src/core/security"
	"errors"
)

type AuthService struct {
	clientRepo domain.IClient
}

func NewAuthService(clientRepo domain.IClient) *AuthService {
	return &AuthService{clientRepo: clientRepo}
}

// Login autenticando al usuario
func (as *AuthService) Login(email, password string) (map[string]interface{}, error) {
	client, err := as.clientRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	// Verificar la contraseña
	if !security.CheckPassword(client.Password, password) {
		return nil, errors.New("contraseña incorrecta")
	}

	// Generar el token JWT
	token, err := security.GenerateJWT(int(client.ID), client.Email)
	if err != nil {
		return nil, err
	}

	// Regresar tanto el token como el ID del usuario
	return map[string]interface{}{
		"token":  token,
		"userId": client.ID,
	}, nil
}

// Register registra un nuevo cliente, asegurando que la contraseña esté hasheada
func (as *AuthService) Register(client entities.Client) error {
	// Hashear la contraseña antes de guardar el cliente
	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword

	// Guardar cliente en la base de datos
	return as.clientRepo.Save(client)
}
