# 🔗 CONFIGURACIÓN DE AWS

## 📋 PASO 1: Crear bucket S3

1. Ve a AWS S3 > "Create bucket"
2. Asigna un nombre: `test-read-file-golang`
3. En "Block public access", Desmarcar:
  - Block public access to buckets and objects granted through new public bucket or access point policies
  - Block public and cross-account access to buckets and objects through any public bucket or access point policies
4. Avanza y crea el bucket

## 📋 PASO 2: Subir el archivo

1. Ve a tu bucket > "Upload"
2. Sube el archivo `test_error_lines.csv`

## 📋 PASO 3: Hacer un archivo específico público

1. Ve a tu bucket > selecciona el archivo
2. Ve a la pestaña "Permissions"
3. Baja a "Bucket Policy" y agrega:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicRead",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::test-read-file-golang/test_error_lines.csv"
    }
  ]
}
```

#### 👨‍💼 3. Crear usuario IAM

1. Ve a IAM > Users > Create user
2. Asigna nombre: `go-s3-reader`
3. Marca "Access key - Programmatic access"
4. Asigna permisos: `AmazonS3ReadOnlyAccess`
5. Finaliza y guarda el Access Key ID y Secret Access Key

#### ⚖️ 4. Configurar en tu equipo local

```bash
aws configure
```

Responde:

```
AWS Access Key ID: [tu clave]
AWS Secret Access Key: [tu secreto]
Default region name: us-east-1
Default output format: json
```

---
