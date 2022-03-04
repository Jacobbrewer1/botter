FROM golang:1.17

WORKDIR $home\source\repos\botter

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get github.com/Jacobbrewer1/botter

RUN go build -a -v -work -o /botterexe

CMD [ "/botterexe" ]
