build_server:
	@read -p "Enter tag version:" tag; \
	cd server; \
	docker build -t hello-server .; \
	docker tag hello-server discofighter47/hello-server:$$tag; \
	docker push discofighter47/hello-server:$$tag