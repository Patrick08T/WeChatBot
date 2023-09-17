FROM golang as build

# go mod Installation source, container environment variable addition will override the default variable value
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# Set up the working directory
WORKDIR /WeChatBot
# add all files to the container
COPY . .

WORKDIR /WeChatBot/script
RUN chmod +x *.sh

RUN /bin/sh -c ./build_all_service.sh

#Blank image Multi-Stage Build
FROM ubuntu

RUN rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install apt-transport-https && apt-get install procps\
&&apt-get install net-tools
#Non-interactive operation
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get install -y vim curl tzdata gawk
#Time zone adjusted to East eighth District
RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && dpkg-reconfigure -f noninteractive tzdata


#set directory to map logs,config file,script file.
VOLUME ["/WeChatBot/logs","/WeChatBot/script"]

#Copy scripts files and binary files to the blank image
COPY --from=build /WeChatBot/script /WeChatBot/script
COPY --from=build /WeChatBot/bin /WeChatBot/bin

WORKDIR /WeChatBot/script

CMD ["./docker_start_all.sh"]
