# Running a lambda on localstack

1. Create a zip
```
zip lambda.zip lambda.py
```

1. Start localstack
```
SERVICES=lambda localstack start
```

1. Upload zip and create a function definition
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws \
--endpoint-url=http://localhost:4574 lambda create-function \
--function-name=f1 --runtime=python2.7 --role=r1 \
--handler=lambda.handler --zip-file fileb://lambda.zip
```

1. Call the function
```
AWS_ACCESS_KEY_ID=999999 AWS_SECRET_ACCESS_KEY=1231 aws lambda \
--endpoint-url=http://localhost:4574 invoke \
--function-name f1 result.log
```

1. View output
```
cat result.log
```

You should see:
```
{
  "foo": "bar"
}
```
