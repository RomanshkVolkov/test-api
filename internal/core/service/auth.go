package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	schema "github.com/RomanshkVolkov/test-api/internal/core/domain/schemas"
)

func SignIn(username string, password string) (domain.APIResponse[string, any], error) {

	user, err := repository.FindByUsername(username)
	if err != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Error al buscar usuario",
			Error:   err,
		}, err
	}

	if user.ID == 0 {
		return repository.UserNotFound(), nil
	}

	validatedPassword := repository.ComparePasswords(user.Password, password)

	fmt.Println(password)
	fmt.Println(validatedPassword)

	if !validatedPassword {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Contraseña incorrecta",
		}, nil
	}

	token, err := repository.SigninJWT(user)

	if err != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Se produjo un error al iniciar sesión (jwt error)",
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: fmt.Sprintf("Bienvenido de nuevo %s", user.Name),
		Data:    token,
	}, nil
}

func SignUp(request *domain.NewUser) (domain.APIResponse[domain.UserData, any], error) {
	fields := schema.GenericForm[domain.NewUser]{Data: *request}
	failValidatedFields := schema.FormValidator[domain.NewUser](fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: "Verifica los campos en rojo",
			Data: domain.UserData{
				Name:     request.Name,
				Username: request.Username,
				Email:    request.Email,
				Role:     request.Role,
			},
			SchemaError: failValidatedFields,
		}, nil
	}
	existUser, err := repository.FindByUsernameOrEmail(request.Username, request.Email)
	if err != nil {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: "Error al buscar usuario",
			Error:   err,
		}, err
	}

	if existUser.ID != 0 {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: "El usuario o correo ya existe",
		}, nil
	}

	user, err := repository.NewUser(request)

	if err != nil {
		return domain.APIResponse[domain.UserData, any]{
			Success: false,
			Message: "Error al registrar usuario",
		}, nil
	}

	return domain.APIResponse[domain.UserData, any]{
		Success: true,
		Message: "Successfully registered",
		Data:    user,
	}, nil
}

func ResetPasswordRequest(request *domain.PasswordResetRequest) (domain.APIResponse[string, any], error) {
	user, err := repository.SaveOTPCode(request.Username)
	if err != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Error al guardar el código OTP",
			Error:   err,
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
		SupporEmail: "joseguzmandev@gmail.com",
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
			Message: "Error on send mail",
			Error:   err,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: "Successfully sent email",
	}, nil
}

func VerifyForgottenPasswordCode(request *domain.ForgottenPasswordCode) (domain.APIResponse[string, any], error) {
	fields := schema.GenericForm[domain.ForgottenPasswordCode]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Verifica los campos ",
			SchemaError: failValidatedFields,
		}, nil
	}

	user, scheme, err := repository.FindAndValidateOTP(request.Username, request.OTP)
	if err != nil {
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Tu código es incorrecto",
			SchemaError: scheme,
			Error:       err,
		}, nil
	}

	if user.ID == 0 {
		return repository.UserNotFound(), nil
	}

	if len(scheme) > 0 {
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Verifica los campos",
			SchemaError: scheme,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: "Code verified",
		Data:    user.OTP,
	}, nil
}

func ResetForgottenPassword(request *domain.ResetForgottenPassword) (domain.APIResponse[string, any], error) {
	fields := schema.GenericForm[domain.ResetForgottenPassword]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Verifica los campos ",
			SchemaError: failValidatedFields,
		}, nil
	}

	equalPasswords := request.Password == request.ConfirmPassword
	if !equalPasswords {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Las contraseñas no coinciden",
			SchemaError: map[string][]string{
				"password":             {"Las contraseñas no coinciden"},
				"passwordConfirmation": {"Las contraseñas no coinciden"},
			},
		}, nil
	}

	user, scheme, err := repository.FindAndValidateOTP(request.Username, request.OTP)
	if err != nil {
		scheme["otp"] = []string{"Tu código es incorrecto"}
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Tu código es incorrecto",
			SchemaError: scheme,
			Error:       err,
		}, nil
	}

	if user.ID == 0 {
		return repository.UserNotFound(), nil
	}

	if len(scheme) > 0 {
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Verifica los campos",
			SchemaError: scheme,
		}, nil
	}

	err = repository.UpdatePassword(user.ID, request.Password)
	if err != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Error al actualizar la contraseña",
			Error:   err,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: "Tu contraseña ha sido actualizada",
	}, nil
}

func ChangePassword(user repository.CustomClaims, request *domain.ChangePassword) (domain.APIResponse[string, any], error) {
	fields := schema.GenericForm[domain.ChangePassword]{Data: *request}
	failValidatedFields := schema.FormValidator(fields)
	if len(failValidatedFields) > 0 {
		return domain.APIResponse[string, any]{
			Success:     false,
			Message:     "Verifica los campos ",
			SchemaError: failValidatedFields,
		}, nil
	}

	userData, err := repository.FindByID(user.ID)
	if err != nil {
		return repository.UserNotFound(), nil
	}

	validatedPassword := repository.ComparePasswords(userData.Password, request.CurrentPassword)
	if !validatedPassword {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Tu contraseña actual es incorrecta",
			SchemaError: map[string][]string{
				"oldPassword": {"Tu contraseña actual es incorrecta"},
			},
		}, nil
	}

	equalPasswords := request.Password == request.ConfirmPassword
	if !equalPasswords {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Las contraseñas no coinciden",
			SchemaError: map[string][]string{
				"password":        {"Las contraseñas no coinciden"},
				"confirmPassword": {"Las contraseñas no coinciden"},
			},
		}, nil
	}

	onUpdatedError := repository.UpdatePassword(user.ID, request.Password)
	if onUpdatedError != nil {
		return domain.APIResponse[string, any]{
			Success: false,
			Message: "Error al actualizar la contraseña",
			Error:   onUpdatedError,
		}, nil
	}

	return domain.APIResponse[string, any]{
		Success: true,
		Message: "Tu contraseña ha sido actualizada",
	}, nil
}
