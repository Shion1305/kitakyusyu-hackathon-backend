FROM nginx:1.25.3-alpine

COPY default.conf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/nginx.conf
COPY www /var/www
CMD ["nginx", "-g", "daemon off;"]
