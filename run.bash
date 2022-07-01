docker build -t vskgr/controller .
docker stop controller
docker rm controller
docker run -p 5000:5000 -it --name controller vskgr/controller