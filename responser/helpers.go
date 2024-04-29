package responser

import (
	"fmt"
	"runtime"
)

func RuntimeToString(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("Archivo %v, en línea %v.", file, line)
	}
	return "Ubicación desconocida"
}
