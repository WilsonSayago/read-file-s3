package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	totalCount int
	mu         sync.Mutex
	wg         sync.WaitGroup
)

func main() {
	// tiempo inicial
	startTime := time.Now()
	bucket := "test-read-file-golang"
	key := "test_error_lines.csv"

	// Cargar configuración AWS para us-east-1
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic(fmt.Sprintf("Error cargando configuración AWS: %v", err))
	}

	// Crear cliente S3
	client := s3.NewFromConfig(cfg)

	// Descargar el archivo desde S3
	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		panic(fmt.Sprintf("Error accediendo al bucket S3: %v", err))
	}
	defer resp.Body.Close()

	/// Leer línea por línea
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()

		// Ignorar cabecera
		if strings.TrimSpace(line) == "linea" {
			continue
		}

		wg.Add(1)
		go processLine(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	wg.Wait()
	fmt.Println("Total de líneas que contienen 'error':", totalCount)
	elapsed := time.Since(startTime)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}

func processLine(line string) {
	defer wg.Done()

	if strings.Contains(strings.ToLower(line), "error") {
		mu.Lock()
		totalCount++
		mu.Unlock()
	}
}
