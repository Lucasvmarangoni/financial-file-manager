FROM nginx@sha256:67682bda769fae1ccf5183192b8daf37b64cae99c6c3302650f6f8bf5f0f95df AS nginx

EXPOSE 80
EXPOSE 443

COPY nginx/nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /etc/nginx/ \
    && chown -R nginx:nginx /etc/nginx/ \
    && chmod -R 750 /etc/nginx/

RUN mkdir -p /var/cache/nginx/client_temp \
    && chown -R nginx:nginx /var/cache/nginx \
    && chmod -R 750 /var/cache/nginx \
    && mkdir -p /var/run/nginx \
    && touch /var/run/nginx.pid \
    && chown nginx:nginx /var/run/nginx.pid \
    && chmod 640 /var/run/nginx.pid \
    && chown -R nginx:nginx /var/run/nginx \
    && chmod 750 /var/run/nginx

USER nginx:nginx

CMD ["nginx", "-g", "daemon off;"]



