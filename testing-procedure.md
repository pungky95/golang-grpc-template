# Testing Procedure for gRPC Endpoint in Postman

## 1. Install Postman
Make sure you have Postman installed on your machine. If not, you can download it from [Postman's official website](https://www.postman.com/).

## 2. Create a New Request
Open Postman and create a new request.

## 3. Set Up the Request
- Select the request type as `POST`.
- Enter your gRPC endpoint. Replace `https` with `grpc` in the URL.

## 4. Add Metadata
- Click on the "Headers" tab.
- Add a new header with the key `x-api-key` and the value `QbWwtX15K5EDp0gNvcxMRkzrC7f9Qxqn`.

## 5. Configure gRPC Settings
- Fetch by `Use Server Reflection`
- Click RPC you want and click use example message
- Modify the payload as you want and click invoke

If you encounter difficulties, consider using dedicated gRPC tools or libraries for testing, such as [BloomRPC](https://github.com/uw-labs/bloomrpc) or [grpcurl](https://github.com/fullstorydev/grpcurl).
