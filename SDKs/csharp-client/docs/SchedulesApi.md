# IO.Swagger.Api.SchedulesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountSchedule**](SchedulesApi.md#getaccountschedule) | **GET** /accounts/{account_id}/schedules/{schedule_id} | Show details of an individual schedule
[**ListAccountSchedules**](SchedulesApi.md#listaccountschedules) | **GET** /accounts/{account_id}/schedules | Get a list of schedules for an account


<a name="getaccountschedule"></a>
# **GetAccountSchedule**
> ScheduleFull GetAccountSchedule (int? accountId, int? scheduleId)

Show details of an individual schedule

This service shows the details of an individual schedule.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountScheduleExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SchedulesApi();
            var accountId = 56;  // int? | Account ID
            var scheduleId = 56;  // int? | Schedule ID

            try
            {
                // Show details of an individual schedule
                ScheduleFull result = apiInstance.GetAccountSchedule(accountId, scheduleId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchedulesApi.GetAccountSchedule: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **scheduleId** | **int?**| Schedule ID | 

### Return type

[**ScheduleFull**](ScheduleFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountschedules"></a>
# **ListAccountSchedules**
> ListSchedules ListAccountSchedules (int? accountId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of schedules for an account

See Intro to Schedules for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountSchedulesExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SchedulesApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersName = new List<string>(); // List<string> | Name filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortName = sortName_example;  // string | Name sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of schedules for an account
                ListSchedules result = apiInstance.ListAccountSchedules(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SchedulesApi.ListAccountSchedules: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **filtersId** | [**List&lt;string&gt;**](string.md)| ID filter | [optional] 
 **filtersName** | [**List&lt;string&gt;**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListSchedules**](ListSchedules.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

