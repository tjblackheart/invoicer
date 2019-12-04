# dev
FROM node:13-alpine as dev
WORKDIR /app
COPY src .
RUN chown -R 1000.1000 /app && \
    npm install
CMD ["npm", "run", "serve"]
