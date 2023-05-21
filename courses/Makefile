build_image:
	docker build -t melvinkimathi/courses-app:v1.0.3 . && \
		docker push melvinkimathi/courses-app:v1.0.3
build_postgresql_image:
	docker build -t melvinkimathi/postgresql:v1 -f Dockerfile.postgresql .
docker_hub_login:
	docker login || true
push_courses_image: docker_hub_login
	docker push melvinkimathi/courses-app:v1.0.3
push_postgresql_image: docker_hub_login
	docker push melvinkimathi/postgresql:v1
run_image:
	docker run --name test-multistage-courses test-multistage-courses
run_test:
	go test -v ./... 