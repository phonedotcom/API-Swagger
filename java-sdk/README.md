# swagger-java-client

## Requirements

Building the API client library requires [Maven](https://maven.apache.org/) to be installed.

## Installation

To install the API client library to your local Maven repository, simply execute:

```shell
mvn install
```

To deploy it to a remote Maven repository instead, configure the settings of the repository and execute:

```shell
mvn deploy
```

Refer to the [official documentation](https://maven.apache.org/plugins/maven-deploy-plugin/usage.html) for more information.

### Maven users

Add this dependency to your project's POM:

```xml
<dependency>
    <groupId>io.swagger</groupId>
    <artifactId>swagger-java-client</artifactId>
    <version>1.0.0</version>
    <scope>compile</scope>
</dependency>
```

### Gradle users

Add this dependency to your project's build file:

```groovy
compile "io.swagger:swagger-java-client:1.0.0"
```

### Others

At first generate the JAR by executing:

    mvn package

Then manually install the following JARs:

* target/swagger-java-client-1.0.0.jar
* target/lib/*.jar

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```java

import io.swagger.client.*;
import io.swagger.client.auth.*;
import io.swagger.client.model.*;
import io.swagger.client.api.AccountsApi;

import java.io.File;
import java.util.*;

public class AccountsApiExample {

    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        
        // Configure API key authorization: apiKey
        ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
        apiKey.setApiKey("YOUR API KEY");
        // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
        //apiKey.setApiKeyPrefix("Token");

        AccountsApi apiInstance = new AccountsApi();
        Integer accountId = 56; // Integer | Account ID
        try {
            AccountFull result = apiInstance.getAccount(accountId);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountsApi#getAccount");
            e.printStackTrace();
        }
    }
}

```

## Documentation for API Endpoints

All URIs are relative to *https://api.phone.com/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**getAccount**](docs/AccountsApi.md#getAccount) | **GET** /accounts/{account_id} | Retrieve details of an individual account
*AccountsApi* | [**listAccounts**](docs/AccountsApi.md#listAccounts) | **GET** /accounts | Get a list of accounts visible to the authenticated user or client
*ApplicationsApi* | [**getAccountApplication**](docs/ApplicationsApi.md#getAccountApplication) | **GET** /accounts/{account_id}/applications/{application_id} | Show details of an individual application
*ApplicationsApi* | [**listAccountApplications**](docs/ApplicationsApi.md#listAccountApplications) | **GET** /accounts/{account_id}/applications | Get a list of applications you have defined
*AvailablenumbersApi* | [**listAvailablePhoneNumbers**](docs/AvailablenumbersApi.md#listAvailablePhoneNumbers) | **GET** /phone-numbers/available | 
*CalleridsApi* | [**getCallerIds**](docs/CalleridsApi.md#getCallerIds) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use
*CalllogsApi* | [**listAccountCallLogs**](docs/CalllogsApi.md#listAccountCallLogs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account
*ContactsApi* | [**createAccountExtensionContact**](docs/ContactsApi.md#createAccountExtensionContact) | **POST** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension
*ContactsApi* | [**deleteAccountExtensionContact**](docs/ContactsApi.md#deleteAccountExtensionContact) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 
*ContactsApi* | [**getAccountExtensionContact**](docs/ContactsApi.md#getAccountExtensionContact) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact
*ContactsApi* | [**listAccountExtensionContacts**](docs/ContactsApi.md#listAccountExtensionContacts) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts | Show a list of address book contacts
*ContactsApi* | [**replaceAccountExtensionContact**](docs/ContactsApi.md#replaceAccountExtensionContact) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contacts | 
*DevicesApi* | [**createAccountDevice**](docs/DevicesApi.md#createAccountDevice) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
*DevicesApi* | [**getAccountDevice**](docs/DevicesApi.md#getAccountDevice) | **GET** /accounts/{account_id}/device/{device_id} | Show details of an individual VoIP device
*DevicesApi* | [**listAccountDevices**](docs/DevicesApi.md#listAccountDevices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
*DevicesApi* | [**replaceAccountDevice**](docs/DevicesApi.md#replaceAccountDevice) | **PUT** /accounts/{account_id}/device/{device_id} | Update the settings for an individual VoIP device
*ExpressservicecodesApi* | [**getAccountExpressSrvCode**](docs/ExpressservicecodesApi.md#getAccountExpressSrvCode) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
*ExpressservicecodesApi* | [**listAccountExpressSrvCodes**](docs/ExpressservicecodesApi.md#listAccountExpressSrvCodes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format
*ExtensionsApi* | [**createAccountExtension**](docs/ExtensionsApi.md#createAccountExtension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
*ExtensionsApi* | [**getAccountExtension**](docs/ExtensionsApi.md#getAccountExtension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
*ExtensionsApi* | [**listAccountExtensions**](docs/ExtensionsApi.md#listAccountExtensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
*ExtensionsApi* | [**replaceAccountExtension**](docs/ExtensionsApi.md#replaceAccountExtension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension
*GroupsApi* | [**createAccountExtensionContactGroup**](docs/GroupsApi.md#createAccountExtensionContactGroup) | **POST** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
*GroupsApi* | [**deleteAccountExtensionContactGroup**](docs/GroupsApi.md#deleteAccountExtensionContactGroup) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
*GroupsApi* | [**getAccountExtensionContactGroup**](docs/GroupsApi.md#getAccountExtensionContactGroup) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
*GroupsApi* | [**listAccountExtensionContactGroups**](docs/GroupsApi.md#listAccountExtensionContactGroups) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
*GroupsApi* | [**replaceAccountExtensionContactGroup**](docs/GroupsApi.md#replaceAccountExtensionContactGroup) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
*MediaApi* | [**getAccountMedia**](docs/MediaApi.md#getAccountMedia) | **GET** /accounts/{account_id}/media/{recording_id} | Show details of an individual media recording (Greeting or Hold Music)
*MediaApi* | [**listAccountMedia**](docs/MediaApi.md#listAccountMedia) | **GET** /accounts/{account_id}/media | Get a list of media recordings for an account
*MenusApi* | [**createAccountMenu**](docs/MenusApi.md#createAccountMenu) | **POST** /accounts/{account_id}/menus | Create an individual menu
*MenusApi* | [**deleteAccountMenu**](docs/MenusApi.md#deleteAccountMenu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
*MenusApi* | [**getAccountMenu**](docs/MenusApi.md#getAccountMenu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
*MenusApi* | [**listAccountMenus**](docs/MenusApi.md#listAccountMenus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
*MenusApi* | [**replaceAccountMenu**](docs/MenusApi.md#replaceAccountMenu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu
*NumberregionsApi* | [**listAvailablePhoneNumberRegions**](docs/NumberregionsApi.md#listAvailablePhoneNumberRegions) | **GET** /phone-numbers/available/regions | 
*PhonenumbersApi* | [**createAccountPhoneNumber**](docs/PhonenumbersApi.md#createAccountPhoneNumber) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
*PhonenumbersApi* | [**getAccountPhoneNumber**](docs/PhonenumbersApi.md#getAccountPhoneNumber) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
*PhonenumbersApi* | [**listAccountPhoneNumbers**](docs/PhonenumbersApi.md#listAccountPhoneNumbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
*PhonenumbersApi* | [**replaceAccountPhoneNumber**](docs/PhonenumbersApi.md#replaceAccountPhoneNumber) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account
*QueuesApi* | [**createAccountQueue**](docs/QueuesApi.md#createAccountQueue) | **POST** /accounts/{account_id}/queues | Create a queue
*QueuesApi* | [**deleteAccountQueue**](docs/QueuesApi.md#deleteAccountQueue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
*QueuesApi* | [**getAccountQueue**](docs/QueuesApi.md#getAccountQueue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
*QueuesApi* | [**listAccountQueues**](docs/QueuesApi.md#listAccountQueues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
*QueuesApi* | [**replaceAccountQueue**](docs/QueuesApi.md#replaceAccountQueue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue
*RoutesApi* | [**createRoute**](docs/RoutesApi.md#createRoute) | **POST** /accounts/{account_id}/routes | Add a new address book contact for an extension
*RoutesApi* | [**deleteAccountRoute**](docs/RoutesApi.md#deleteAccountRoute) | **DELETE** /accounts/{account_id}/routes/{route_id} | 
*RoutesApi* | [**getAccountRoute**](docs/RoutesApi.md#getAccountRoute) | **GET** /accounts/{account_id}/routes/{route_id} | Show details of an individual route
*RoutesApi* | [**listAccountRoutes**](docs/RoutesApi.md#listAccountRoutes) | **GET** /accounts/{account_id}/routes | Get a list of routes for an account
*RoutesApi* | [**replaceAccountRoute**](docs/RoutesApi.md#replaceAccountRoute) | **PUT** /accounts/{account_id}/routes/{route_id} | 
*SchedulesApi* | [**getAccountSchedule**](docs/SchedulesApi.md#getAccountSchedule) | **GET** /accounts/{account_id}/schedules/{schedule_id} | Show details of an individual schedule
*SchedulesApi* | [**listAccountSchedules**](docs/SchedulesApi.md#listAccountSchedules) | **GET** /accounts/{account_id}/schedules | Get a list of schedules for an account
*SmsApi* | [**createAccountSms**](docs/SmsApi.md#createAccountSms) | **POST** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
*SmsApi* | [**getAccountSms**](docs/SmsApi.md#getAccountSms) | **GET** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
*SmsApi* | [**listAccountSms**](docs/SmsApi.md#listAccountSms) | **GET** /accounts/{account_id}/sms | Get a list of SMS messages for an account
*SubaccountsApi* | [**createAccountSubaccount**](docs/SubaccountsApi.md#createAccountSubaccount) | **POST** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
*SubaccountsApi* | [**listAccountSubaccounts**](docs/SubaccountsApi.md#listAccountSubaccounts) | **GET** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client
*TrunksApi* | [**createAccountTrunk**](docs/TrunksApi.md#createAccountTrunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
*TrunksApi* | [**deleteAccountTrunk**](docs/TrunksApi.md#deleteAccountTrunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
*TrunksApi* | [**getAccountTrunk**](docs/TrunksApi.md#getAccountTrunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
*TrunksApi* | [**listAccountTrunks**](docs/TrunksApi.md#listAccountTrunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
*TrunksApi* | [**replaceAccountTrunk**](docs/TrunksApi.md#replaceAccountTrunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


## Documentation for Models

 - [AccountFull](docs/AccountFull.md)
 - [AccountSummary](docs/AccountSummary.md)
 - [Address](docs/Address.md)
 - [AddressListContacts](docs/AddressListContacts.md)
 - [Addresses](docs/Addresses.md)
 - [ApplicationFull](docs/ApplicationFull.md)
 - [ApplicationSummary](docs/ApplicationSummary.md)
 - [AvailableNumberFull](docs/AvailableNumberFull.md)
 - [AvailableNumberSummary](docs/AvailableNumberSummary.md)
 - [AvailableNumbersFull](docs/AvailableNumbersFull.md)
 - [AvailableNumbersSummary](docs/AvailableNumbersSummary.md)
 - [CallDetails](docs/CallDetails.md)
 - [CallLogFull](docs/CallLogFull.md)
 - [CallLogSummary](docs/CallLogSummary.md)
 - [CallNotifications](docs/CallNotifications.md)
 - [CallerIdFull](docs/CallerIdFull.md)
 - [CallerIdPhoneNumber](docs/CallerIdPhoneNumber.md)
 - [ContactAccount](docs/ContactAccount.md)
 - [ContactFull](docs/ContactFull.md)
 - [ContactSubaccount](docs/ContactSubaccount.md)
 - [ContactSummary](docs/ContactSummary.md)
 - [CreateContactParams](docs/CreateContactParams.md)
 - [CreateDeviceParams](docs/CreateDeviceParams.md)
 - [CreateExtensionParams](docs/CreateExtensionParams.md)
 - [CreateGroupParams](docs/CreateGroupParams.md)
 - [CreateMenuParams](docs/CreateMenuParams.md)
 - [CreatePhoneNumberParams](docs/CreatePhoneNumberParams.md)
 - [CreateQueueParams](docs/CreateQueueParams.md)
 - [CreateRouteParams](docs/CreateRouteParams.md)
 - [CreateSmsParams](docs/CreateSmsParams.md)
 - [CreateSubaccountParams](docs/CreateSubaccountParams.md)
 - [CreateTrunkParams](docs/CreateTrunkParams.md)
 - [DeleteContact](docs/DeleteContact.md)
 - [DeleteGroup](docs/DeleteGroup.md)
 - [DeleteMenu](docs/DeleteMenu.md)
 - [DeleteQueue](docs/DeleteQueue.md)
 - [DeleteRoute](docs/DeleteRoute.md)
 - [DeleteTrunk](docs/DeleteTrunk.md)
 - [DeviceFull](docs/DeviceFull.md)
 - [DeviceMembership](docs/DeviceMembership.md)
 - [DeviceSummary](docs/DeviceSummary.md)
 - [DevicesFull](docs/DevicesFull.md)
 - [DevicesSummary](docs/DevicesSummary.md)
 - [Email](docs/Email.md)
 - [Emails](docs/Emails.md)
 - [ExpressServiceCodeFull](docs/ExpressServiceCodeFull.md)
 - [ExpressServiceCodeSummary](docs/ExpressServiceCodeSummary.md)
 - [ExpressServiceCodesFull](docs/ExpressServiceCodesFull.md)
 - [ExpressServiceCodesSummary](docs/ExpressServiceCodesSummary.md)
 - [ExtensionFull](docs/ExtensionFull.md)
 - [ExtensionSummary](docs/ExtensionSummary.md)
 - [FilterCallLogs](docs/FilterCallLogs.md)
 - [FilterIdArray](docs/FilterIdArray.md)
 - [FilterIdDirectionFrom](docs/FilterIdDirectionFrom.md)
 - [FilterIdExtensionNameArray](docs/FilterIdExtensionNameArray.md)
 - [FilterIdGroupIdUpdatedAtArray](docs/FilterIdGroupIdUpdatedAtArray.md)
 - [FilterIdNameArray](docs/FilterIdNameArray.md)
 - [FilterIdNamePhoneNumberArray](docs/FilterIdNamePhoneNumberArray.md)
 - [FilterListAvailableNumbers](docs/FilterListAvailableNumbers.md)
 - [FilterListPhoneNumbersRegions](docs/FilterListPhoneNumbersRegions.md)
 - [FilterNameNumberArray](docs/FilterNameNumberArray.md)
 - [Greeting](docs/Greeting.md)
 - [GroupByListPhoneNumbersRegions](docs/GroupByListPhoneNumbersRegions.md)
 - [GroupFull](docs/GroupFull.md)
 - [GroupListContacts](docs/GroupListContacts.md)
 - [GroupSummary](docs/GroupSummary.md)
 - [GroupsFull](docs/GroupsFull.md)
 - [GroupsSummary](docs/GroupsSummary.md)
 - [HoldMusic](docs/HoldMusic.md)
 - [Line](docs/Line.md)
 - [ListAccountsFull](docs/ListAccountsFull.md)
 - [ListAccountsSummary](docs/ListAccountsSummary.md)
 - [ListApplicationsFull](docs/ListApplicationsFull.md)
 - [ListApplicationsSummary](docs/ListApplicationsSummary.md)
 - [ListAvailableNumbersFull](docs/ListAvailableNumbersFull.md)
 - [ListAvailableNumbersSummary](docs/ListAvailableNumbersSummary.md)
 - [ListCallLogsFull](docs/ListCallLogsFull.md)
 - [ListCallLogsSummary](docs/ListCallLogsSummary.md)
 - [ListCallerIdsFull](docs/ListCallerIdsFull.md)
 - [ListCallerIdsSummary](docs/ListCallerIdsSummary.md)
 - [ListContactsFull](docs/ListContactsFull.md)
 - [ListContactsSummary](docs/ListContactsSummary.md)
 - [ListDevicesFull](docs/ListDevicesFull.md)
 - [ListDevicesSummary](docs/ListDevicesSummary.md)
 - [ListExpressServiceCodesFull](docs/ListExpressServiceCodesFull.md)
 - [ListExpressServiceCodesSummary](docs/ListExpressServiceCodesSummary.md)
 - [ListExtensionsFull](docs/ListExtensionsFull.md)
 - [ListExtensionsSummary](docs/ListExtensionsSummary.md)
 - [ListGroupsFull](docs/ListGroupsFull.md)
 - [ListGroupsSummary](docs/ListGroupsSummary.md)
 - [ListMediaFull](docs/ListMediaFull.md)
 - [ListMediaSummary](docs/ListMediaSummary.md)
 - [ListMenusFull](docs/ListMenusFull.md)
 - [ListMenusSummary](docs/ListMenusSummary.md)
 - [ListPhoneNumbersFull](docs/ListPhoneNumbersFull.md)
 - [ListPhoneNumbersRegionsFull](docs/ListPhoneNumbersRegionsFull.md)
 - [ListPhoneNumbersRegionsSummary](docs/ListPhoneNumbersRegionsSummary.md)
 - [ListPhoneNumbersSummary](docs/ListPhoneNumbersSummary.md)
 - [ListQueuesFull](docs/ListQueuesFull.md)
 - [ListQueuesSummary](docs/ListQueuesSummary.md)
 - [ListRoutesFull](docs/ListRoutesFull.md)
 - [ListRoutesSummary](docs/ListRoutesSummary.md)
 - [ListSchedulesFull](docs/ListSchedulesFull.md)
 - [ListSchedulesSummary](docs/ListSchedulesSummary.md)
 - [ListSmsFull](docs/ListSmsFull.md)
 - [ListSmsSummary](docs/ListSmsSummary.md)
 - [ListTrunksFull](docs/ListTrunksFull.md)
 - [ListTrunksSummary](docs/ListTrunksSummary.md)
 - [MediaFull](docs/MediaFull.md)
 - [MediaSummary](docs/MediaSummary.md)
 - [Member](docs/Member.md)
 - [Members](docs/Members.md)
 - [MenuFull](docs/MenuFull.md)
 - [MenuSummary](docs/MenuSummary.md)
 - [MenusFull](docs/MenusFull.md)
 - [MenusSummary](docs/MenusSummary.md)
 - [Notification](docs/Notification.md)
 - [Option](docs/Option.md)
 - [OptionsListMenus](docs/OptionsListMenus.md)
 - [PhoneNumberContact](docs/PhoneNumberContact.md)
 - [PhoneNumberFull](docs/PhoneNumberFull.md)
 - [PhoneNumberSummary](docs/PhoneNumberSummary.md)
 - [PhoneNumbersFull](docs/PhoneNumbersFull.md)
 - [PhoneNumbersRegionFull](docs/PhoneNumbersRegionFull.md)
 - [PhoneNumbersRegionsFull](docs/PhoneNumbersRegionsFull.md)
 - [PhoneNumbersSummary](docs/PhoneNumbersSummary.md)
 - [QueueFull](docs/QueueFull.md)
 - [QueueSummary](docs/QueueSummary.md)
 - [Queues](docs/Queues.md)
 - [QueuesSummary](docs/QueuesSummary.md)
 - [Recipient](docs/Recipient.md)
 - [ReplaceExtensionParams](docs/ReplaceExtensionParams.md)
 - [ReplaceMenuParams](docs/ReplaceMenuParams.md)
 - [ReplacePhoneNumberParams](docs/ReplacePhoneNumberParams.md)
 - [RouteFull](docs/RouteFull.md)
 - [RouteSummary](docs/RouteSummary.md)
 - [RoutesFull](docs/RoutesFull.md)
 - [RoutesSummary](docs/RoutesSummary.md)
 - [RuleSet](docs/RuleSet.md)
 - [RuleSetAction](docs/RuleSetAction.md)
 - [RuleSetFilter](docs/RuleSetFilter.md)
 - [RuleSetForwardItem](docs/RuleSetForwardItem.md)
 - [ScheduleFull](docs/ScheduleFull.md)
 - [ScheduleSummary](docs/ScheduleSummary.md)
 - [SipAuthentication](docs/SipAuthentication.md)
 - [SmsForwarding](docs/SmsForwarding.md)
 - [SmsFull](docs/SmsFull.md)
 - [SmsSummary](docs/SmsSummary.md)
 - [SortCallLogs](docs/SortCallLogs.md)
 - [SortId](docs/SortId.md)
 - [SortIdCreatedAt](docs/SortIdCreatedAt.md)
 - [SortIdExtensionName](docs/SortIdExtensionName.md)
 - [SortIdName](docs/SortIdName.md)
 - [SortIdNamePhoneNumber](docs/SortIdNamePhoneNumber.md)
 - [SortIdUpdatedAt](docs/SortIdUpdatedAt.md)
 - [SortListAvailableNumbers](docs/SortListAvailableNumbers.md)
 - [SortListPhoneNumbersRegions](docs/SortListPhoneNumbersRegions.md)
 - [SortNameNumber](docs/SortNameNumber.md)
 - [TrunkFull](docs/TrunkFull.md)
 - [TrunkSummary](docs/TrunkSummary.md)
 - [Voicemail](docs/Voicemail.md)


## Documentation for Authorization

Authentication schemes defined for the API:
### apiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header


## Recommendation

It's recommended to create an instance of `ApiClient` per thread in a multithreaded environment to avoid any potential issue.

## Author

apisupport@phone.com

