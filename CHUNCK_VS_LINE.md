# 🎥 ERRORS CODE

### ⚡️ CUANDO IR POR CHUNKS NO ES TAN BUENA IDEA

```go
buffer := make([]byte, 4096)
for {
    n, err := reader.Read(buffer)
    if err == io.EOF {
        break
    }
    chunk := buffer[:n]
    go processChunk(chunk)
}
```

### 💥 **Resultado del Procesamiento:**

| Chunk | Contenido | ¿Contiene "error"? | Resultado |
|-------|-----------|-------------------|-----------|
| 1 | `"hubo un err"` | ❌ NO | No cuenta |
| 2 | `"or en el sistema"` | ❌ NO | No cuenta |

### 🚨 **PROBLEMA IDENTIFICADO:**
- La palabra `"error"` se dividió en `"err"` + `"or"`
- **Resultado:** 0 errores detectados
- **Esperado:** 1 error detectado
- **Precisión:** 0% ❌

---
**ACCION GENERADA:** "Posibilidad de que las líneas se corten y no se detecte el error."

**CONCLUSION:** "❌ El conteo falla."

---

### 🔹 LA SOLUCIÓN REAL: LECTURA LÍNEA POR LÍNEA

```go
scanner := bufio.NewScanner(resp.Body)
for scanner.Scan() {
    line := scanner.Text()
    wg.Add(1)
    go processLine(line)
}
```

**CONCLUSION:** "Cada línea se procesa sin riesgo de estar cortada."

---