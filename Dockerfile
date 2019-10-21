FROM ubuntu
WORKDIR /myapp

COPY db /myapp/db
COPY todoApp /myapp
COPY .env /myapp
COPY app/build /myapp/app/build

CMD ["./todoApp"]
