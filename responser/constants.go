package responser

const (
	StatusCodeOkResponse          = 200
	StatusCodeInternalServerError = 500
	StatusCodeBadRequestError     = 400
	StatusCodeNotFoundError       = 404
	//Código de estado interno que indica una solicitud exitosa
	InternalStatusCodeOkResponse = "T000"
	//Código de estado interno que indica una conexión fallida a la base de datos
	InternalStatusCodeConnectionError = "T101"
	//Código de estado interno que indica el ingreso de datos erróneos al endpoint/API
	InternalStatusCodeBadInputData = "T109"
	//Código de estado interno que indica que el solicitante no tiene permisos para ejecutar un proceso de la API
	InternalStatusCodeCustomerNotAuthorized = "T111"
	//Código de estado interno que indica un error de servidor
	InternalStatusCodeServerError = "T112"
	//Código de estado interno que indica que el recurso buscado no fue encontrado
	InternalStatusCodeNotFound = "T113"
	//Código de estado interno que indica que el área de Cumplimiento debe autorizar al cliente para operar con Tanner
	InternalStatusCodeComplianceMustCheck = "T114"
)
