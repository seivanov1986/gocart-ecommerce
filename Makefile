frontend:
	cd ../gocart-admin-frontend/ && \
	PUBLIC_URL=/admin REACT_APP_HOST=http://localhost:8000 npm run build
	cp -r ../gocart-admin-frontend/build/* ./schemes/admin/

start:
	SERVICE_BASE_PATH=$(PWD) go run cmd/service/main.go
