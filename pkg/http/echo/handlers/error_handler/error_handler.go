package echoErrorHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	httpError "github.com/samannsr/vending-machine-control-system/pkg/error/http"
	"github.com/samannsr/vending-machine-control-system/pkg/logger"
	loggerConstant "github.com/samannsr/vending-machine-control-system/pkg/logger/constant"
)

func ErrorHandler(err error, c echo.Context) {
	// default echo errors
	echoHttpError, ok := err.(*echo.HTTPError)
	var httpResponseError httpError.HttpErr

	if ok {
		httpResponseError = httpError.NewHttpError(echoHttpError.Code, echoHttpError.Code, http.StatusText(echoHttpError.Code), http.StatusText(echoHttpError.Code), nil)
	} else {
		// parse as a custom error
		httpResponseError = httpError.ParseError(err)
	}

	if !c.Response().Committed {
		if _, err := httpResponseError.WriteTo(c.Response()); err != nil {
			logger.Zap.Sugar().Error(`error while writing http error response: %v`, err)
		}

		logger.Zap.Error(
			"request failed",
			zap.Error(err),
			zap.String(loggerConstant.TYPE, loggerConstant.HTTP),
			zap.String(loggerConstant.TITILE, httpResponseError.GetTitle()),
			zap.Int(loggerConstant.CODE, httpResponseError.GetCode()),
			zap.String(loggerConstant.STATUS, http.StatusText(httpResponseError.GetStatus())),
			zap.Time(loggerConstant.TIME, httpResponseError.GetTimestamp()),
			zap.Any(loggerConstant.DETAILS, httpResponseError.GetDetails()),
		)
	}
}
