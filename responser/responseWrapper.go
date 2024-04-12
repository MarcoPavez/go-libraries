package responser

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type response struct {
	Data data `json:"data"`
}

type data struct {
	IdSolicitud string      `json:"idSolicitud"`
	Valido      bool        `json:"valido"`
	Codigo      string      `json:"codigo"`
	Mensaje     interface{} `json:"mensaje"`
}

func Wrapper(endpointHandler func(c *gin.Context) CustomResponseCreator) gin.HandlerFunc {
	return func(c *gin.Context) {

		idSolicitud := uuid.NewString()
		customResponse := endpointHandler(c)

		apiResponse := response{
			data{
				IdSolicitud: idSolicitud,
				Valido:      true,
				Codigo:      InternalStatusCodeOkResponse,
				Mensaje:     "Solicitud completada exitosamente",
			},
		}

		httpRes, isHTTPResponse := customResponse.(*HTTPResponse)

		if isHTTPResponse {
			if httpRes.cause != nil {
				apiResponse.Data.Valido = false
			}
		}

		apiResponse.Data.Codigo = customResponse.GetInternalCode()
		apiResponse.Data.Mensaje = customResponse.GetResponseMessage()

		c.Set("CustomResponse", customResponse)
		log.Println(customResponse.GetResponseMessage())
		c.JSON(customResponse.GetStatusCode(), apiResponse)
	}
}
