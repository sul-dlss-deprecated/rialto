# Running a lambda on localstack

1. Create a zip
```
GOOS=linux go build -o main
zip lambda.zip main
```

1. Start localstack. If you're on a Mac, ensure you are running the docker daemon.
```
SERVICES=lambda,es,sns LAMBDA_EXECUTOR=docker localstack start
```

1. Upload zip and create a function definition
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws \
--endpoint-url http://localhost:4574 lambda create-function \
--function-name f1 \
--runtime go1.x \
--role r1 \
--handler main \
--environment "Variables={ES_HOST=192.168.2.7}" \
--zip-file fileb://lambda.zip
```

1. Create ES domain
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws es \
--endpoint-url=http://localhost:4578 create-elasticsearch-domain \
--domain-name records \
--elasticsearch-version 6.2 \
--elasticsearch-cluster-config InstanceType=t2.small.elasticsearch,InstanceCount=1 \
--ebs-options EBSEnabled=true,VolumeType=standard,VolumeSize=10
```

1. Create SNS topic
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws sns \
--endpoint-url=http://localhost:4575 create-topic \
--name data-update
```

1. Subscribe to SNS events
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws sns \
--endpoint-url=http://localhost:4575 subscribe \
--topic-arn arn:aws:sns:us-east-1:123456789012:data-update \
--protocol lambda \
--notification-endpoint arn:aws:lambda:us-east-1:000000000000:function:f1
```

1. Publish a Message
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws sns \
--endpoint-url=http://localhost:4575 publish \
--topic-arn arn:aws:sns:us-east-1:123456789012:data-update \
--message '{"Records": [{"EventSource": "foo", "Sns": { "Timestamp": "2014-05-16T08:28:06.801Z",
"Message": "Hello world!" }}]}'
```

1. View output
When you go to http://localhost:4571/records/_search

You should see an item record with:
```
"_source":{"foo": "barfoo"}
```

1. Cleanup (necessary before you upload a newer version of the function)

```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws lambda \
--endpoint-url=http://localhost:4574 delete-function \
--function-name f1
```
