package exception

var (
	GENERAL_BAD_REQUEST                   = Code("GEN4001")
	GENERAL_NOT_FOUND                     = Code("GEN4002")
	GENERAL_TOO_MANY_REQUEST              = Code("GEN4003")
	GENERAL_ENDPOINT_IS_UNDER_MAINTENANCE = Code("GEN4004")
	GENERAL_ACCESS_DENIED                 = Code("GEN4005")
	GENERAL_SERVICE_IS_NOT_AVAILABLE      = Code("GEN4006")
	GENERAL_FORMAT_REQUEST_NOT_VALID      = Code("GEN4009")

	GENERAL_INTERNAL_SERVER_ERROR = Code("GEN5001")
)

func init() {

	except[GENERAL_BAD_REQUEST] = Exception{
		ErrorCode:  string(GENERAL_BAD_REQUEST),
		Message:    "bad request",
		StatusCode: 400,
	}
	except[GENERAL_NOT_FOUND] = Exception{
		ErrorCode:  string(GENERAL_NOT_FOUND),
		Message:    "not found",
		StatusCode: 404,
	}
	except[GENERAL_TOO_MANY_REQUEST] = Exception{
		ErrorCode:  string(GENERAL_TOO_MANY_REQUEST),
		Message:    "too many request",
		StatusCode: 429,
	}
	except[GENERAL_ENDPOINT_IS_UNDER_MAINTENANCE] = Exception{
		ErrorCode:  string(GENERAL_ENDPOINT_IS_UNDER_MAINTENANCE),
		Message:    "saat ini endpoint sedang dalam pemeliharaan",
		StatusCode: 400,
	}
	except[GENERAL_ACCESS_DENIED] = Exception{
		ErrorCode:  string(GENERAL_ACCESS_DENIED),
		Message:    "akses ditolak",
		StatusCode: 403,
	}
	except[GENERAL_SERVICE_IS_NOT_AVAILABLE] = Exception{
		ErrorCode:  string(GENERAL_SERVICE_IS_NOT_AVAILABLE),
		Message:    "saat ini layanan sedang tidak tersedia",
		StatusCode: 400,
	}
	except[GENERAL_FORMAT_REQUEST_NOT_VALID] = Exception{
		ErrorCode:  string(GENERAL_FORMAT_REQUEST_NOT_VALID),
		Message:    "data payload tidak benar",
		StatusCode: 400,
	}

	except[GENERAL_INTERNAL_SERVER_ERROR] = Exception{
		ErrorCode:  string(GENERAL_INTERNAL_SERVER_ERROR),
		Message:    "internal server error",
		StatusCode: 500,
	}
}
