
# golang:onbuild automatically copies the package source, 
# fetches the application dependencies, builds the program, 
# and configures it to run on startup 
FROM golang:onbuild
LABEL Name=greeter_server Version=0.0.1 
EXPOSE 5001

# For more control, you can copy and build manually
# FROM golang:latest 
# LABEL Name=greeter_server Version=0.0.1 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main .
# EXPOSE 5001 
# CMD ["/app/main"]
