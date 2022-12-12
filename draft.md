# Connect AppSheet to Apigee

- Create an Apigee evaluation account.
- Create an API proxy.
  - In the Apigee Edge management UI, navigate to the API proxies page and click on the "+API Proxy" button.
  - Select "REST" as the type of API that you want to create.
  - Enter a name for the API proxy and select a base path for its URLs.
  - Click on the "Next" button to continue.
  - Select a blank proxy as the starting point for the API proxy, and then click on the "Create" button.
  - In the API proxy editor, click on the "Proxy Endpoints" tab and then click on the "+Target Endpoint" button.
  - Enter a name for the target endpoint and specify its URL.
  - Click on the "Save" button to save the target endpoint.
  - Note: These steps are based on the information provided in the article and may vary depending on your specific setup and the version of Apigee that you are using. Additionally, these steps only cover the basic process of creating an API proxy. In a real-world scenario, you would likely need to configure additional settings such as authentication, request and response handling, and more.
- Configure the API proxy as a data source in AppSheet.
  - In the Apigee Edge management UI, navigate to the API proxies page and click on the API proxy that you want to use as a data source in AppSheet.
  - Click on the Develop tab, and then click on the API proxy's OpenAPI specification file to view its contents.
  - Copy the entire contents of the OpenAPI specification file.
  - In AppSheet, open the data source configuration page for the app that you want to connect to the API proxy.
  - Select "Open API Spec" as the data source type, and then paste the contents of the OpenAPI specification file into the provided text box.
  - Configure any additional settings as needed, such as authentication or request headers, and then save the data source configuration.
  - Test the connection to the API proxy by making a request to one of its endpoints using AppSheet's built-in test functionality.
  - Note: These steps are based on the information provided in the article and may vary depending on your specific setup and the version of AppSheet that you are using.
- Test the connection between AppSheet and Apigee.
  - In AppSheet, open the app that you want to test the connection with.
  - Navigate to the data source configuration page for the app and make sure that the API proxy is properly configured as a data source.
  - In the data source configuration page, click on the "Test" button next to the endpoint that you want to test.
  - If the endpoint requires authentication, enter the necessary credentials in the authentication dialog that appears.
  - Click on the "Send" button to send the test request to the API proxy.
  - If the connection is successful, you should see the response from the API proxy in the test results window.
  - If the connection is not successful, check the error messages and troubleshoot as needed.
  - Note: These steps are based on the information provided in the article and may vary depending on your specific setup and the version of AppSheet that you are using. Additionally, these steps only cover the basic process of testing the connection between AppSheet and Apigee. In a real-world scenario, you would likely need to test the connection with multiple endpoints and perform more advanced testing to ensure that the integration is working properly.

---

> Note: The article recommends using the Open API Spec method for connecting AppSheet to Apigee. It also suggests using a mock API (one that returns fixed values) for testing purposes.