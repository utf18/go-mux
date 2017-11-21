#### Sample Project for a gorilla/mux based go rest API with prometheus metrics enabled and simple request logging.

#### the main app is a sample golang application with a gorilla/mux based router.
port 8080 is used and you can build it yourself with:
`go get -d`
***linux:***
`CGO_ENABLED=0 GOOS=linux go build -o mux`
***OSX:***
`CGO_ENABLED=0 GOOS=darwin go build -o mux`

after this you should be able to start it with:
`./mux`

and test it with
`curl localhost:8080`

or for the metrics endpoint:
`curl localhost:8080/metrics`

If you do not want to build it yourself just use docker or even more convenient docker-compose

`docker-compose up`

this will launch a stack with grafana,prometheus and the mux application

You can reach your applications under:

***grafana:***
  - localhost:3000/login
  - admin:password

***prometheus:***
  - localhost:9090

***mux:***
  - localhost:8080

If you want to check if prometheus scrapes the mux application correctly,
you can use the util/configure-grafana-datasource.sh script to configure the prometheus datasource.
After this you only have to import the dashboards in the *grafana* folder.
The mux-dashboard.json is a very basic dashboard which displays the go_threads of both prometheus and mux.

i added a simple prometheus gauge in the gauge.go file to show how to get updated metrics while the http server is running and blocking.

I guess i cloud have used a middleware to take care of this, but since this repo is solely based for my learnings, i will stay with this.


### Contributing

All patches are welcome!
If you find something is missing or broken please file an issue
