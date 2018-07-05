FROM java:8-jre-alpine

RUN apk --no-cache add openssl

RUN mkdir /usr/local/dynamodb \
    && cd /usr/local/dynamodb \
    && wget https://s3-ap-northeast-1.amazonaws.com/dynamodb-local-tokyo/dynamodb_local_latest.tar.gz \
    && tar zxvf dynamodb_local_latest.tar.gz

CMD ["java", "-Djava.library.path=/usr/local/dynamodb/DynamoDBLocal_lib", "-jar", "/usr/local/dynamodb/DynamoDBLocal.jar", "-dbPath", "/usr/local/dynamodb/db"]

EXPOSE 8000
