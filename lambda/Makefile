LOCAL_ENDPOINT_HOST    ?=localhost
LOCAL_ENDPOINT         =--endpoint-url=http://${LOCAL_ENDPOINT_HOST}
SNS_ENDPOINT           =${LOCAL_ENDPOINT}:4575
ES_ENDPOINT            =${LOCAL_ENDPOINT}:4578
LAMBDA_ENDPOINT        =${LOCAL_ENDPOINT}:4574
PROJ_ES_DOMAIN         =records
PROJ_SNS_TOPIC         =data-update
PROJ_AWS_REGION        =us-east-1
PROJ_AWS_ACCESS_KEY_ID =999999
PROJ_AWS_SECRET_KEY    =999999
PROJ_ENV_VARS          =AWS_REGION=${PROJ_AWS_REGION} AWS_ACCESS_KEY_ID=${PROJ_AWS_ACCESS_KEY_ID} AWS_SECRET_ACCESS_KEY=${PROJ_AWS_SECRET_KEY}

default: package upload subscribe

package:
	GOOS=linux go build -o main
	zip lambda.zip main

upload:
	$(eval FUNCTION=$(shell aws $(LAMBDA_ENDPOINT) lambda list-functions | jq '.Functions[0].FunctionName // ""'))
	@if [[ $(FUNCTION) != "" ]]; \
	  then echo "$(FUNCTION) function found"; \
	  else aws $(LAMBDA_ENDPOINT) lambda create-function \
	    --function-name f1 \
		--runtime go1.x \
		--role r1 \
		--handler main \
		--environment "Variables={ES_HOST=$(ES_ENDPOINT)}" \
		--zip-file fileb://lambda.zip && \
		echo "f1 function created"; \
fi;

subscribe:
	$(eval FUNCTION_ARN=$(shell aws $(LAMBDA_ENDPOINT) lambda list-functions | jq '.Functions[0].FunctionArn // ""'))
	@if [[ $(FUNCTION_ARN) == "" ]]; then upload && subscribe; fi;
	$(eval TOPIC=$(shell aws $(SNS_ENDPOINT) sns list-topics --region=$(PROJ_AWS_REGION) | jq '.Topics[0].TopicArn // ""'))
	@if [[ $(TOPIC) != "" ]]; \
	  then aws $(SNS_ENDPOINT) sns subscribe \
	    --topic-arn $(TOPIC) \
		--protocol lambda \
		--notification-endpoint $(FUNCTION_ARN) \
		--region=$(PROJ_AWS_REGION); \
fi;