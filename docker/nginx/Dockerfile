FROM nginx:latest

ARG NGINX_CONTAINER_STATIC_DIRECTORY=/var/www/static
ARG NGINX_CONF_LOCALPATH
ARG ROOT_STATIC_LOCALPATH

RUN mkdir -p ${NGINX_CONTAINER_STATIC_DIRECTORY}
COPY ${ROOT_STATIC_LOCALPATH} /var/www/static/
COPY ${NGINX_CONF_LOCALPATH} /etc/nginx/nginx.conf
