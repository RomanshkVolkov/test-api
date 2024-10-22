package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	schema "github.com/RomanshkVolkov/test-api/internal/core/domain/schemas"
)

func (server Server) SignIn(username string, password string) (domain.APIResponse[domain.SignInResponse, any], error) {
	repo := repository.GetDBConnection(server.Host)
	user, err := repo.FindByUsername(username)
	if err != nil || user.ID == 0 {
		return domain.APIResponse[domain.SignInResponse, any]{
			Success: false,
			Message: domain.Message{
				En: "User not found",
				Es: "Usuario no encontrado",
			},
			Error: err,
		}, err
	}

	validatedPassword := repository.ComparePasswords(user.Password, password)

	if !validatedPassword {
		return domain.APIResponse[domain.SignInResponse, any]{
			Success: false,
			Message: domain.Message{
				En: "Invalid credentials",
				Es: "Credenciales inválidas",
			},
		}, nil
	}

	token, err := repository.SigninJWT(user)

	if err != nil {
		return domain.APIResponse[domain.SignInResponse, any]{
			Success: false,
			Message: domain.Message{
				En: "Error on sign in JWT",
				Es: "Error al iniciar sesión JWT",
			},
		}, nil
	}

	return domain.APIResponse[domain.SignInResponse, any]{
		Success: true,
		Message: domain.Message{
			En: "Welcome back",
			Es: "Bienvenido de nuevo",
		},
		Data: domain.SignInResponse{
			User: domain.UserData{
				Username: user.Username,
				Name:     user.Name,
				Email:    user.Email,
			},
			Profile: user.Profile,
			Token:   token,
		},
	}, nil
}

func (server Server) SignUp(request *domain.NewUser) (domain.APIResponse[domain.UserData, any], error) {
	fields := schema.GenericForm[domain.NewUser]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			Data: domain.UserData{
				Name:     request.Name,
				Username: request.Username,
				Email:    request.Email,
			},
			SchemaError: failValidatedFields,
		}, nil
	}

	repo := repository.GetDBConnection(server.Host)
	existUser, err := repo.FindByUsernameOrEmail(request.Username, request.Email)
	if err != nil {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: domain.Message{
				En: "User not found",
				Es: "Usuario no encontrado",
			},
			Error: err,
		}, err
	}

	if existUser.ID != 0 {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: domain.Message{
				En: "User already exists",
				Es: "El usuario ya existe",
			},
		}, nil
	}

	user, err := repo.NewUser(request)

	if err != nil {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: domain.Message{
				En: "Error on save user",
				Es: "Error al guardar el usuario",
			},
		}, nil
	}

	return domain.APIResponse[domain.UserData, any]{
		Success: true,
		Message: domain.Message{
			En: "User created successfully",
			Es: "Usuario creado exitosamente",
		},
		Data: user,
	}, nil
}

func (server Server) ResetPasswordRequest(request *domain.PasswordResetRequest) (domain.APIResponse[string, any], error) {
	repo := repository.GetDBConnection(server.Host)
	user, err := repo.SaveOTPCode(request.Username)
	if err != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Error on save OTP code",
				Es: "Error al guardar el código OTP",
			},
			Error: err,
		}, nil
	}

	if user.ID == 0 {
		return repository.UserNotFound(), nil
	}

	t, err := template.ParseFiles("/srv/internal/adapters/templates/forgotten-password.html")
	if err != nil {
		panic(err)
	}

	var Domain = ""

	var body bytes.Buffer
	t.Execute(&body, struct {
		Name        string
		Code        string
		AppName     string
		SupporEmail string
		Domain      string
	}{
		Name:        user.Name,
		Code:        user.OTP,
		AppName:     "Test API",
		SupporEmail: "sistemas@dwitmexico.com",
		Domain:      Domain,
	})

	mailOptions := &MailOptions{
		To:      []string{user.Email},
		Subject: "Test API - Código de recuperacion",
		Body:    body.String(),
	}
	done, err := SendMail(mailOptions)
	fmt.Println(err)
	if !done {
		return domain.APIResponse[string, any]{
			Success: true,
			Message: domain.Message{
				En: "Error on send email",
				Es: "Error al enviar el correo",
			},
			Error: err,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: domain.Message{
			En: "Email sent with the OTP code",
			Es: "Correo enviado con el código OTP",
		},
	}, nil
}

func (server Server) VerifyForgottenPasswordCode(request *domain.ForgottenPasswordCode) (domain.APIResponse[string, any], error) {
	fields := schema.GenericForm[domain.ForgottenPasswordCode]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			SchemaError: failValidatedFields,
		}, nil
	}

	repo := repository.GetDBConnection(server.Host)
	user, scheme, err := repo.FindAndValidateOTP(fields.Data.Username, fields.Data.OTP)
	if err != nil || user.ID == 0 {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Your code is incorrect",
				Es: "Tu código es incorrecto",
			},
			SchemaError: scheme,
			Error:       err,
		}, nil
	}

	if len(scheme) > 0 {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			SchemaError: scheme,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: domain.Message{
			En: "Your code is correct",
			Es: "Tu código es correcto",
		},
		Data: user.OTP,
	}, nil
}

func (server Server) ResetForgottenPassword(request *domain.ResetForgottenPassword) (domain.APIResponse[string, any], error) {
	fields := schema.GenericForm[domain.ResetForgottenPassword]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			SchemaError: failValidatedFields,
		}, nil
	}

	equalPasswords := request.Password == request.ConfirmPassword
	if !equalPasswords {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Passwords do not match",
				Es: "Las contraseñas no coinciden",
			},
			SchemaError: map[string][]string{
				"password":             {"Las contraseñas no coinciden"},
				"passwordConfirmation": {"Las contraseñas no coinciden"},
			},
		}, nil
	}

	repo := repository.GetDBConnection(server.Host)
	user, scheme, err := repo.FindAndValidateOTP(request.Username, request.OTP)
	if err != nil {
		scheme["otp"] = []string{"Tu código es incorrecto"}
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Your code is incorrect",
				Es: "Tu código es incorrecto",
			},
			SchemaError: scheme,
			Error:       err,
		}, nil
	}

	if user.ID == 0 {
		return repository.UserNotFound(), nil
	}

	if len(scheme) > 0 {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			SchemaError: scheme,
		}, nil
	}

	err = repo.UpdatePassword(user.ID, request.Password)
	if err != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Error on update password",
				Es: "Error al actualizar la contraseña",
			},
			Error: err,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: domain.Message{
			En: "Password updated",
			Es: "Contraseña actualizada",
		},
	}, nil
}

func (server Server) ChangePassword(user repository.CustomClaims, request *domain.ChangePassword) (domain.APIResponse[string, any], error) {
	fields := schema.GenericForm[domain.ChangePassword]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Check the fields",
				Es: "Verifica los campos",
			},
			SchemaError: failValidatedFields,
		}, nil
	}

	repo := repository.GetDBConnection(server.Host)
	userData, err := repo.FindByID(user.ID)
	if err != nil {
		return repository.UserNotFound(), nil
	}

	validatedPassword := repository.ComparePasswords(userData.Password, request.CurrentPassword)
	if !validatedPassword {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Your current password is incorrect",
				Es: "Tu contraseña actual es incorrecta",
			},
			SchemaError: map[string][]string{
				"oldPassword": {"Tu contraseña actual es incorrecta"},
			},
		}, nil
	}

	equalPasswords := request.Password == request.ConfirmPassword
	if !equalPasswords {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Passwords do not match",
				Es: "Las contraseñas no coinciden",
			},
			SchemaError: map[string][]string{
				"password":        {"Las contraseñas no coinciden"},
				"confirmPassword": {"Las contraseñas no coinciden"},
			},
		}, nil
	}

	onUpdatedError := repo.UpdatePassword(user.ID, request.Password)
	if onUpdatedError != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: domain.Message{
				En: "Error on update password",
				Es: "Error al actualizar la contraseña",
			},
			Error: onUpdatedError,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: domain.Message{
			En: "Password updated",
			Es: "Contraseña actualizada",
		},
	}, nil
}
