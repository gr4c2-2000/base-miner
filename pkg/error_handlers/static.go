package error_handlers

import "github.com/rotisserie/eris"

var (
	NO_TASK_ERROR                        = eris.New("No tasks")
	TO_MANY_RECORDS_IN_ES_ERROR          = eris.New("To Many Records")
	DATA_COULD_NOT_BE_FIX_IN_ES_ERROR    = eris.New("Data COULDN'T Be Fix")
	UNKNOW_ERROR_TYPE_FROM_COMPERE_ERROR = eris.New("Unknow Error Type")
	NO_IMPLEMENTATION_ERROR              = eris.New("Function Not Implemented")
	RETRT_FUNCTION_RETRY_ERROR           = eris.New("Func Retry  internal error")
	NO_MYSQL_DOC_FOR_ID_ERROR            = eris.New("NO mysql Document for id")
	NO_DATA_IN_ES_ERROR                  = eris.New("No Data In ES")
	CONNECTION_NOT_EXISTS_ERROR          = eris.New("Conection dosen't Exists")
	AGG_NOT_EXISTS_ERROR                 = eris.New("Agregat dosen't Exists")
	EP_NOT_EXISTS_ERROR                  = eris.New("Endpoint dosen't Exists")
	TASK_NOT_EXISTS_ERROR                = eris.New("Task dosen't Exists")
	US_NOT_EXISTS_ERROR                  = eris.New("UpdateShema dosen't Exists")
	ES_VERSION_NOT_SUPPORTED             = eris.New("ES version not supported")
)
