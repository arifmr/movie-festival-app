// package rest

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"

// 	"github.com/labstack/echo/v4"
// 	"go.uber.org/dig"

// 	"github.com/koinworks/asgard-koinbpr/internal/service/business_saving_account_creation"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/check_status_user_creation"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/doku_generate_virtual_account"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/get_account_balance"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/get_tnc"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/get_virtual_accounts"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/personal_saving_account_creation"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/update_latest_user_creation_status"
// 	"github.com/koinworks/asgard-koinbpr/internal/service/update_status_user_creation"
// 	"github.com/koinworks/asgard-koinbpr/pkg/entity"
// 	"github.com/koinworks/asgard-koinbpr/pkg/middleware"
// )

// var ErrInvalidKWUserID = errors.New("invalid koinworks user id")

// type (
// 	ServicesImpl struct {
// 		Services
// 	}

// 	Services struct {
// 		dig.In
// 		business_saving_account_creation.BPRSavingAccountRegistrationSvc
// 		personal_saving_account_creation.PsaCreationSvc
// 		check_status_user_creation.CheckStatusUserCreationSvc
// 		update_status_user_creation.UpdateStatusUserCreationSvc
// 		get_tnc.GetTncSvc
// 		get_account_balance.GetAccountBalance
// 		update_latest_user_creation_status.UpdateLatestUserCreationStatus
// 		doku_generate_virtual_account.DokuGenerateVirtualAccountSvc
// 		get_virtual_accounts.GetVirtualAccountsSvc
// 	}
// )

// func NewRestHandler(
// 	e *echo.Echo,
// 	services Services,
// ) {
// 	handler := ServicesImpl{
// 		Services: services,
// 	}

// 	e.GET("/tnc", handler.getTnc, middleware.ProtectMiddleware)
// 	e.POST("/latest-user-creation-status", handler.updateLatestUserCreationStatus)
// 	e.POST("/saving-account-business", handler.businessRegistration, middleware.ProtectMiddleware)
// 	e.POST("/saving-account-personal", handler.personalRegistration, middleware.ProtectMiddleware)
// 	e.GET("/status-user-creation", handler.checkStatusUserCreation, middleware.ProtectMiddleware)
// 	e.POST("/status-user-creation", handler.updateStatusUserCreation, middleware.ProtectMiddleware)
// 	e.GET("/balances", handler.getAccountBalance, middleware.ProtectMiddleware)
// 	e.GET("/virtual-accounts", handler.getVirtualAccounts, middleware.ProtectMiddleware)
// }

// func (x *ServicesImpl) checkStatusUserCreation(c echo.Context) (err error) {
// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	var kwUserID int64

// 	if kwUserID, err = strconv.ParseInt(c.Request().Header.Get(middleware.RestHeaderKeyUserID), 10, 64); err != nil {
// 		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidKWUserID)))
// 	}

// 	resp, httpCode, err := x.CheckStatusUserCreationSvc.CheckStatusUserCreation(ctx, check_status_user_creation.CheckStatusUserCreationRequest{KWUserID: kwUserID})
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
// }

// func (x *ServicesImpl) businessRegistration(c echo.Context) (err error) {
// 	var request business_saving_account_creation.BPRSavingAccountRegistrationSvcRequest

// 	for key, values := range c.Request().Header {
// 		if strings.ToLower(key) == middleware.RestHeaderKeyUserID || strings.ToLower(key) == middleware.RestHeaderKeyUserCode {
// 			for _, value := range values {
// 				if strings.ToLower(key) == middleware.RestHeaderKeyUserID {
// 					v, _ := strconv.Atoi(value)
// 					request.KWUserID = int64(v)
// 				} else {
// 					request.KWUserCode = value
// 				}
// 			}
// 		}
// 	}

// 	err = c.Bind(&request)
// 	if err != nil {
// 		splittedErr := strings.Split(err.Error(), ",")
// 		statusCode, _ := strconv.Atoi(strings.Split(splittedErr[0], "=")[1])
// 		field := strings.Split(splittedErr[3], "=")[1]
// 		expected := strings.Split(splittedErr[1], "expected=")[1]
// 		newErr := fmt.Sprintf("%s must be %s", field, expected)

// 		return c.JSON(http.StatusUnprocessableEntity, NewResponseError(statusCode, msgFailed, text(statusCode), text(statusCode), unwrapFirstError(errors.New(newErr))))
// 	}

// 	if err := c.Validate(request); err != nil {
// 		splittedErr := strings.Split(err.Error(), ",")
// 		statusCode, _ := strconv.Atoi(strings.Split(splittedErr[0], "=")[1])
// 		field := strings.Split(splittedErr[1], "'")[3]
// 		newErr := fmt.Sprintf("%s is required", field)

// 		return c.JSON(http.StatusUnprocessableEntity, NewResponseError(statusCode, msgFailed, text(statusCode), text(statusCode), unwrapFirstError(errors.New(newErr))))
// 	}

// 	w := io.Writer(os.Stdout)
// 	ctx := entity.OTel.
// 		NewLogger(c.Request().Context(), w).
// 		Z().WithContext(c.Request().Context())

// 	res, httpCode, err := x.BPRSavingAccountRegistrationSvc.BPRSavingAccountRegistrationSvc(ctx, request)
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: res})
// }

// func (x *ServicesImpl) personalRegistration(c echo.Context) (err error) {
// 	var request personal_saving_account_creation.PsaCreationSvcRequest

// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(err)))
// 	}

// 	if err := c.Validate(request); err != nil {
// 		return c.JSON(http.StatusUnprocessableEntity, NewResponseError(http.StatusUnprocessableEntity, msgFailed, text(http.StatusUnprocessableEntity), text(http.StatusUnprocessableEntity), unwrapFirstError(err)))
// 	}

// 	v := fmt.Sprintf("%v", c.Get(middleware.RestHeaderKeyUserID))

// 	request.KWUserID, err = strconv.ParseInt(v, 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusUnprocessableEntity, NewResponseError(http.StatusUnprocessableEntity, msgFailed, text(http.StatusUnprocessableEntity), text(http.StatusUnprocessableEntity), unwrapFirstError(err)))
// 	}

// 	request.KWUserCode = fmt.Sprintf("%v", c.Get(middleware.RestHeaderKeyUserCode))

// 	ctx := entity.OTel.
// 		NewLogger(c.Request().Context(), os.Stdout).
// 		Z().WithContext(c.Request().Context())

// 	res, httpCode, err := x.PsaCreationSvc.PsaCreationSvc(ctx, request)
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(http.StatusCreated, Response{
// 		Status:  http.StatusCreated,
// 		Message: msgSuccess,
// 		Data:    res,
// 	})
// }

// func (x *ServicesImpl) updateStatusUserCreation(c echo.Context) (err error) {
// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	var (
// 		kwUserID int64
// 		payload  update_status_user_creation.UpdateStatusUserCreationRequest
// 	)

// 	if err := c.Bind(&payload); err != nil {
// 		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(err)))
// 	}

// 	if kwUserID, err = strconv.ParseInt(c.Request().Header.Get(middleware.RestHeaderKeyUserID), 10, 64); err != nil {
// 		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidKWUserID)))
// 	}

// 	payload.KWUserID = kwUserID

// 	resp, httpCode, err := x.UpdateStatusUserCreationSvc.UpdateStatusUserCreation(ctx, payload)
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
// }

// func (x *ServicesImpl) getTnc(c echo.Context) (err error) {
// 	language := c.QueryParam("language")

// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	resp, httpCode, err := x.GetTncSvc.GetTnc(ctx, get_tnc.GetTncRequest{Language: language})
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
// }

// func (x *ServicesImpl) getAccountBalance(c echo.Context) (err error) {
// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	var kwUserID int64

// 	if kwUserID, err = strconv.ParseInt(c.Request().Header.Get(middleware.RestHeaderKeyUserID), 10, 64); err != nil {
// 		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(ErrInvalidKWUserID)))
// 	}

// 	resp, httpCode, err := x.GetAccountBalance.GetAccountBalance(ctx, get_account_balance.GetAccountBalanceRequest{KWUserID: kwUserID})
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
// }

// func (x *ServicesImpl) updateLatestUserCreationStatus(c echo.Context) (err error) {
// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	httpCode, err := x.UpdateLatestUserCreationStatus.UpdateLatestUserCreationStatus(ctx)
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess})
// }

// func (x *ServicesImpl) generateDokuVirtualAccounts(c echo.Context) (err error) {
// 	request := doku_generate_virtual_account.DokuGenerateVirtualAccountSvcRequest{}

// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(err)))
// 	}

// 	resp, httpCode, err := x.DokuGenerateVirtualAccountSvc.DokuGenerateVirtualAccountSvc(ctx, request)
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
// }

// func (x *ServicesImpl) getVirtualAccounts(c echo.Context) (err error) {
// 	userID := c.Request().Header.Get(middleware.RestHeaderKeyUserID)

// 	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

// 	id, _ := strconv.ParseInt(userID, 10, 64)

// 	resp, httpCode, err := x.GetVirtualAccountsSvc.GetVirtualAccounts(ctx, get_virtual_accounts.GetVirtualAccountsRequest{UserID: id})
// 	if err != nil {
// 		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
// 	}

// 	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
// }
