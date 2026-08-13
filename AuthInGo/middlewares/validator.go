package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"net/http"
)

func LoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
			var payload dto.LoginUserRequestDto

		if jsonErr := utils.ReadJsonBody(r,&payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Something went wrong while logging in",jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload);validationErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid Input Data",validationErr)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"LoginUserPayload",
			payload,
		)

		next.ServeHTTP(w,r.WithContext(ctx))
		})
}

func CreateUserRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		var payload dto.CreateUserRequestDto

		if jsonErr := utils.ReadJsonBody(r,&payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Something went wrong while creating user",jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload);validationErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid Input Data",validationErr)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"CreateUserPayload",
			payload,
		)

		next.ServeHTTP(w,r.WithContext(ctx))
	})

}