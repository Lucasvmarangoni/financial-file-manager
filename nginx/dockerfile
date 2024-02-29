FROM nginx:latest as nginx

EXPOSE 80

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/cert.pem /etc/nginx/cert.pem
COPY nginx/key.pem /etc/nginx/key.pem
COPY nginx/certs/ca.crt /etc/nginx/ca.crt