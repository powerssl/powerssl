#!/bin/bash

set -euo pipefail

docker run -e URLS="[{url:\"$GRPC_PROXY_URI/swagger/v1/acme_account.swagger.json\",name:\"ACMEAccount\"},{url:\"$GRPC_PROXY_URI/swagger/v1/acme_server.swagger.json\",name:\"ACMEServer\"},{url:\"$GRPC_PROXY_URI/swagger/v1/certificate.swagger.json\",name:\"Certificate\"},{url:\"$GRPC_PROXY_URI/swagger/v1/certificate_issue.swagger.json\",name:\"CertificateIssue\"},{url:\"$GRPC_PROXY_URI/swagger/v1/user.swagger.json\",name:\"User\"}]" -p 80:8080 swaggerapi/swagger-ui
