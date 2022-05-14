rsync --delete --exclude=docker-data -r ./ node1.nihao.int:/opt/crossfit/backend-golang
ssh kshatrov@node1.nihao.int sudo /opt/crossfit/backend-golang/prod-build.sh