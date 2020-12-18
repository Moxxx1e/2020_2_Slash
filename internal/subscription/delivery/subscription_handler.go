package delivery

import (
	"github.com/go-park-mail-ru/2020_2_Slash/internal/consts"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/helpers/errors"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/models"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/mwares"
	"github.com/go-park-mail-ru/2020_2_Slash/internal/subscription"
	"github.com/go-park-mail-ru/2020_2_Slash/tools/logger"
	"github.com/go-park-mail-ru/2020_2_Slash/tools/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type SubscriptionHandler struct {
	subUseCase subscription.SubscriptionUseCase
}

func NewSubscriptionHandler(uc subscription.SubscriptionUseCase) *SubscriptionHandler {
	return &SubscriptionHandler{subUseCase: uc}
}

func (sh *SubscriptionHandler) Configure(e *echo.Echo, mw *mwares.MiddlewareManager) {
	e.POST("/api/v1/subscription", sh.CreateSubscriptionHandler(), mw.CheckAuth, mw.CheckCSRF)
	e.GET("/api/v1/subscription", sh.GetSubscriptionHandler(), mw.CheckAuth, mw.CheckCSRF)
	e.DELETE("/api/v1/subscription", sh.DeleteSubscriptionHandler(), mw.CheckAuth, mw.CheckCSRF)
}

func (sh *SubscriptionHandler) CreateSubscriptionHandler() echo.HandlerFunc {
	return func(cntx echo.Context) error {
		userID, ok := cntx.Get("userID").(uint64)
		if !ok {
			customErr := errors.Get(consts.CodeGetFromContextError)
			logger.Error(customErr.Message)
			return cntx.JSON(customErr.HTTPCode, response.Response{Error: customErr})
		}

		subscription := &models.Subscription{
			UserID:  userID,
			Expires: time.Now().AddDate(0, 1, 0),
			Active:  true,
		}

		customErr := sh.subUseCase.Create(subscription)
		if customErr != nil {
			logger.Error(customErr.Message)
			return cntx.JSON(customErr.HTTPCode, response.Response{Error: customErr})
		}

		return cntx.JSON(http.StatusOK, response.Response{
			Body: &response.Body{
				"subscription": subscription,
			},
		})
	}
}

func (sh *SubscriptionHandler) GetSubscriptionHandler() echo.HandlerFunc {
	return func(cntx echo.Context) error {
		userID, ok := cntx.Get("userID").(uint64)
		if !ok {
			customErr := errors.Get(consts.CodeGetFromContextError)
			logger.Error(customErr.Message)
			return cntx.JSON(customErr.HTTPCode, response.Response{Error: customErr})
		}

		subscription, customErr := sh.subUseCase.GetByUserID(userID)
		if customErr != nil {
			logger.Error(customErr.Message)
			return cntx.JSON(customErr.HTTPCode, response.Response{Error: customErr})
		}

		return cntx.JSON(http.StatusOK, response.Response{
			Body: &response.Body{
				"subscription": subscription,
			},
		})
	}
}

func (sh *SubscriptionHandler) DeleteSubscriptionHandler() echo.HandlerFunc {
	return func(cntx echo.Context) error {
		userID, ok := cntx.Get("userID").(uint64)
		if !ok {
			customErr := errors.Get(consts.CodeGetFromContextError)
			logger.Error(customErr.Message)
			return cntx.JSON(customErr.HTTPCode, response.Response{Error: customErr})
		}

		customErr := sh.subUseCase.DeleteByUserID(userID)
		if customErr != nil {
			logger.Error(customErr.Message)
			return cntx.JSON(customErr.HTTPCode, response.Response{Error: customErr})
		}

		return cntx.JSON(http.StatusOK,
			response.Response{Message: "success"})
	}
}
