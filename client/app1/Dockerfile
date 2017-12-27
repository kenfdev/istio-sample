FROM node:8-alpine as builder
WORKDIR /workspace

COPY . .

RUN npm install
RUN npm run build

FROM nginx:1-alpine
WORKDIR /usr/share/nginx/html
COPY --from=builder /workspace/dist .
CMD ["nginx", "-g", "daemon off;"]