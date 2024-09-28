FROM registry.cn-beijing.aliyuncs.com/zhengyansheng/node:lts-alpine as builder
WORKDIR /app
COPY package*.json ./
RUN npm config set registry https://registry.npmmirror.com
RUN npm install
COPY . .
RUN npm run build

FROM registry.cn-beijing.aliyuncs.com/zhengyansheng/nginx:stable-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]