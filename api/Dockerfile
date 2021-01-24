FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/user/api-b2w-planetas/api
WORKDIR /go/src/github.com/user/api-b2w-planetas/api

# RUN go get ./vendor/database

# RUN go get ./
# RUN go build

RUN go run api/main.go

CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	go get github.com/gorilla/mux && \
	fresh; \
	fi
	
EXPOSE 10000
