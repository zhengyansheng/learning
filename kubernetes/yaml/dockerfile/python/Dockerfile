FROM registry.cn-beijing.aliyuncs.com/zhengyansheng/python:3.10.12

WORKDIR /app

COPY . .

RUN pip3 install -r requirement.txt -i https://mirrors.aliyun.com/pypi/simple

EXPOSE 8000

WORKDIR /app/neo4jApp

CMD ["gunicorn", "-c", "neo4jApp/config/gunicorn_config.py", "neo4jApp.wsgi:application"]
