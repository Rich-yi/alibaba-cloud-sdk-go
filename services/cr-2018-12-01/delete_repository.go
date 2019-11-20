package cr_20181201

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteRepository invokes the cr.DeleteRepository API synchronously
// api document: https://help.aliyun.com/api/cr/deleterepository.html
func (client *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
	response = CreateDeleteRepositoryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepositoryWithChan invokes the cr.DeleteRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepositoryWithChan(request *DeleteRepositoryRequest) (<-chan *DeleteRepositoryResponse, <-chan error) {
	responseChan := make(chan *DeleteRepositoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepository(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteRepositoryWithCallback invokes the cr.DeleteRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepositoryWithCallback(request *DeleteRepositoryRequest, callback func(response *DeleteRepositoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepositoryResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepository(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteRepositoryRequest is the request struct for api DeleteRepository
type DeleteRepositoryRequest struct {
	*requests.RpcRequest
	RepoId     string `position:"Query" name:"RepoId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DeleteRepositoryResponse is the response struct for api DeleteRepository
type DeleteRepositoryResponse struct {
	*responses.BaseResponse
	DeleteRepositoryIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                      string `json:"Code" xml:"Code"`
	RequestId                 string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRepositoryRequest creates a request to invoke DeleteRepository API
func CreateDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
	request = &DeleteRepositoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "DeleteRepository", "cr", "openAPI")
	return
}

// CreateDeleteRepositoryResponse creates a response to parse from DeleteRepository response
func CreateDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
	response = &DeleteRepositoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
