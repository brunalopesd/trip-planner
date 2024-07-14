FROM golang:1.20.4-alpine

WORKDIR /planner-golang

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR /planner-golang/cmd/planner

RUN go build -o /planner-golang/bin/planner .

EXPOSE 8080
ENTRYPOINT [ "/planner-golang/bin/planner" ]


