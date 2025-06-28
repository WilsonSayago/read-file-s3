# CONCURRENCIA - PREVENCIÓN DE RACE CONDITIONS

### ⚠️ El Problema de Concurrencia

Cuando múltiples goroutines modifican la misma variable:

```go
// 🚨 CÓDIGO PELIGROSO (Race Condition)
var totalCount int

func processLine(line string) {
    defer wg.Done()
    if strings.Contains(strings.ToLower(line), "error") {
        totalCount++  // ⚠️ PELIGRO: Múltiples goroutines escribiendo
    }
}
```

### 💥 **¿Qué puede salir mal?**

| Escenario | Goroutine A | Goroutine B | Resultado |
|-----------|-------------|-------------|-----------|
| 1 | Lee: `totalCount = 5` | Lee: `totalCount = 5` | |
| 2 | Incrementa: `5 + 1 = 6` | Incrementa: `5 + 1 = 6` | |
| 3 | Escribe: `totalCount = 6` | Escribe: `totalCount = 6` | ❌ **Perdimos un conteo** |

### ✅ La Solución: sync.Mutex

```go
// ✅ CÓDIGO SEGURO
var (
    totalCount int
    mu         sync.Mutex  // 🔒 Candado para proteger la variable
)

func processLine(line string) {
    defer wg.Done()
    if strings.Contains(strings.ToLower(line), "error") {
        mu.Lock()           // 🔒 Bloquear acceso
        totalCount++        // ✅ Operación segura
        mu.Unlock()         // 🔓 Liberar acceso
    }
}
```

### 🛡️ **Cómo Funciona el Mutex:**

1. **Lock()**: Solo una goroutine puede acceder a la vez
2. **Operación Crítica**: Incremento seguro del contador
3. **Unlock()**: Libera el acceso para otras goroutines

### 📊 **Resultado con Mutex:**

| Escenario | Goroutine A | Goroutine B | Resultado |
|-----------|-------------|-------------|-----------|
| 1 | `mu.Lock()` ✅ | Espera... ⏳ | |
| 2 | Lee, incrementa, escribe | Espera... ⏳ | |
| 3 | `mu.Unlock()` ✅ | `mu.Lock()` ✅ | |
| 4 | Termina | Lee, incrementa, escribe | ✅ **Conteo correcto** |

---

## 🎯 COMPARACIÓN FINAL

### 📊 Rendimiento y Precisión

| Método | Velocidad | Precisión | Memoria | Complejidad |
|--------|-----------|-----------|---------|-------------|
| **Chunks Fijos** | ⚡⚡⚡ | ❌ 0-70% | ⭐⭐⭐ | 🟡 Media |
| **Línea por Línea** | ⚡⚡ | ✅ 100% | ⭐⭐⭐ | 🟢 Baja |
| **Cargar Todo** | ⚡ | ✅ 100% | ❌ | 🟢 Baja |

### 🏆 **Ganador: Línea por Línea + Concurrencia**

**Razones:**
- ✅ **Precisión perfecta** (100%)
- ✅ **Velocidad excelente** (goroutines)
- ✅ **Uso eficiente de memoria** (streaming)
- ✅ **Código simple y mantenible**
- ✅ **Escalable** para archivos de cualquier tamaño

---

