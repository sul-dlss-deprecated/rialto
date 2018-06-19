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
AWS_REGION=localstack AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws \
--endpoint-url http://localhost:4574 lambda create-function \
--function-name f1 \
--runtime go1.x \
--role r1 \
--handler main \
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

1. Create the SNS Topic
```
awslocal sns create-topic --name rialto-msg
```

1. Subscribe the lamba function to the SNS Topic
```
awslocal sns subscribe \
  --topic-arn arn:aws:sns:us-east-1:123456789012:rialto-msg \
  --protocol lambda \
  --notification-endpoint arn:aws:lambda:localstack:000000000000:function:f1
```

1. Publish to the SNS Topic
```
awslocal sns publish \
   --topic-arn arn:aws:sns:us-east-1:123456789012:test-topic \
   --message 'Test Message!'
```

1. View output
When you go to http://localhost:4571/records/_search

You should see an item record with:
```
"_source":{"foo": "bar -- Test Message!"}
```

1. Cleanup

```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws lambda \
--endpoint-url=http://localhost:4574 delete-function \
--function-name f1
```
