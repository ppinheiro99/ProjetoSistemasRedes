#FROM nginx:latest

FROM openresty/openresty:alpine

## remove default.conf and copy custom conf to nginx
RUN rm /etc/nginx/conf.d/default.conf
##COPY conf /etc/nginx
COPY nginx.conf /etc/nginx/conf.d
