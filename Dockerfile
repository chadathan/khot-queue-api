# ใช้ golang 1.23.0 image เป็น base image
FROM docker.io/library/golang:1.23.0-alpine AS builder

# กำหนด working directory ใน container
WORKDIR /app

# คัดลอก go.mod และ go.sum เพื่อดาวน์โหลด dependencies ก่อน
COPY go.mod go.sum ./

# ดาวน์โหลด dependencies ที่จำเป็น
RUN go mod tidy

# คัดลอก source code ทั้งหมด
COPY . .

# คอมไพล์แอปพลิเคชัน Go
RUN go build -o intakes-api .

# ใช้ image ที่เล็กกว่าในการรันแอปพลิเคชัน (สำหรับ production)
FROM alpine:latest

# ติดตั้ง dependencies ที่จำเป็น
RUN apk --no-cache add ca-certificates

# กำหนด working directory ใน container
WORKDIR /root/

# คัดลอกแอปพลิเคชันที่คอมไพล์แล้วจาก builder stage
COPY --from=builder /app/intakes-api .
COPY environment.prod.json ./environment.json

# เปิดพอร์ตที่แอปพลิเคชันจะรัน
EXPOSE 80

# รันแอปพลิเคชัน
CMD ["./intakes-api"]