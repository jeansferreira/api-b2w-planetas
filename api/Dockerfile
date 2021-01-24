FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/user/api-b2w-planetas/api
WORKDIR /go/src/github.com/user/api-b2w-planetas/api

# RUN go get ./vendor/database

# RUN go get ./
# RUN go build

RUN go get github.com/gorilla/mux

CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	fresh; \
	fi

RUN go run main.go

	
EXPOSE 10000
