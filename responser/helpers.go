package responser

import (
	"fmt"
	"runtime"
)

func runtimeToString() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("Archivo %v, en línea %v.", file, line)
	}
	return "Ubicación desconocida"
}
