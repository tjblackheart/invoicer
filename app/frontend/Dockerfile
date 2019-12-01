# dev
FROM node:13-alpine as dev
WORKDIR /app
COPY src .
RUN npm install && \
    chown -R 1000.1000 /app
CMD ["npm", "run", "serve"]

# prod
