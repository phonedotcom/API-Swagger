package main

import (
	"fmt"
	"os"

	"errors"
	"github.com/igorsimevski/phonecom-goclient"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	app.Flags = getCliFlags()
	var cliConfig CliConfig

	app.Action = func(c *cli.Context) error {
		var err error
		err, _ = execute(c, cliConfig)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println()
		}

		return err
	}

	app.Run(os.Args)
}

func execute(
	context *cli.Context,
	cliConfig CliConfig) (error, map[string]interface{}) {

	param, err := createCliParams(context)

	if err != nil {
		return err, nil
	}

	showDryRunVerbose(param)
	if param.dryRun {
		return nil, nil
	}

	cliConfig, err = cliConfig.createOrReadCliConfig(param)

	if err != nil {
		return err, nil
	}

	if param.verbose {
		fmt.Printf("CLI configuration: %+v\n\n", cliConfig)
	}

	var sampleJsonCreator SampleJsonCreator
	sampleJsonCreator.createSampleInOutIfNeeded(param)

	if param.samplein != "" || param.sampleout != "" {
		return nil, nil
	}

	swaggerConfig := cliConfig.createSwaggerConfig(cliConfig)
	apiResolver := ApiResolver{swaggerConfig, param.command}
	api := apiResolver.resolve()

	param.accountId = cliConfig.AccountId

	responseHandler := ResponseHandler{param}

	if api == nil {
		if param.command == defaultCommand {
			defaultApi := swagger.DefaultApi{Configuration: swaggerConfig}
			return responseHandler.handle(defaultApi.Ping())
		} else {
			return errors.New(fmt.Sprintf(invalidCommand, param.command, getAllCommands())), nil
		}
	}

	return invokeCommand(responseHandler, param, api)
}

func invokeCommand(rh ResponseHandler, param CliParams, api interface{}) (error, map[string]interface{}) {

	var command = param.command
	var accountId = param.accountId
	var id = param.id
	var filtersId = param.filtersId
	var filterParams = param.filterParams
	var sortParams = param.sortParams
	var limit = param.limit
	var offset = param.offset
	var fields = param.fields
	var input = param.input
	var idSecondary = param.idSecondary
	var idString = param.idString

	switch api := api.(type) {

	case swagger.MediaApi:

		if param.otherParams.recordingId > 0 {
			id = param.otherParams.recordingId
		}

		switch command {

		case listMedia:

			return rh.handle(api.ListAccountMedia(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getRecording:

			return rh.handle(api.GetAccountMedia(accountId, id))
		}

	case swagger.MenusApi:

		if param.otherParams.menuId > 0 {
			id = param.otherParams.menuId
		}

		switch command {

		case listMenus:

			return rh.handle(api.ListAccountMenus(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getMenu:

			return rh.handle(api.GetAccountMenu(accountId, id))

		case createMenu:

			params := createMenuParams(input)
			return rh.handle(api.CreateAccountMenu(accountId, params))

		case replaceMenu:

			params := replaceMenuParams(input)
			return rh.handle(api.ReplaceAccountMenu(accountId, id, params))

		case deleteMenu:

			return rh.handle(api.DeleteAccountMenu(accountId, id))
		}

	case swagger.QueuesApi:

		if param.otherParams.queueId > 0 {
			id = param.otherParams.queueId
		}

		switch command {

		case listQueues:

			return rh.handle(api.ListAccountQueues(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getQueue:

			return rh.handle(api.GetAccountQueue(accountId, id))

		case createQueue:

			params := createQueueParams(input)
			return rh.handle(api.CreateAccountQueue(accountId, params))

		case replaceQueue:

			params := createQueueParams(input)
			return rh.handle(api.ReplaceAccountQueue(accountId, id, params))

		case deleteQueue:

			return rh.handle(api.DeleteAccountQueue(accountId, id))
		}

	case swagger.RoutesApi:

		if param.otherParams.routeId > 0 {
			id = param.otherParams.routeId
		}

		switch command {

		case listRoutes:

			return rh.handle(api.ListAccountRoutes(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getRoute:

			return rh.handle(api.GetAccountRoute(accountId, id))

		case createRoute:

			params := createRouteParams(input)
			return rh.handle(api.CreateRoute(accountId, params))

		case replaceRoute:

			params := createRouteParams(input)
			return rh.handle(api.ReplaceAccountRoute(accountId, id, params))

		case deleteRoute:

			return rh.handle(api.DeleteAccountRoute(accountId, id))
		}

	case swagger.SchedulesApi:

		if param.otherParams.scheduleId > 0 {
			id = param.otherParams.scheduleId
		}

		switch command {

		case listSchedules:

			return rh.handle(api.ListAccountSchedules(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getSchedule:

			return rh.handle(api.GetAccountSchedule(accountId, id))
		}

	case swagger.SmsApi:

		if param.otherParams.smsId != "" {
			idString = param.otherParams.smsId
		}

		switch command {

		case listSms:

			return rh.handle(api.ListAccountSms(accountId, filtersId, filterParams.filtersDirection, filterParams.filtersFrom, sortParams.sortId, sortParams.sortCreatedAt, limit, offset, fields))

		case getSms:

			return rh.handle(api.GetAccountSms(accountId, idString))

		case createSms:

			if param.otherParams.extensionId > 0 {
				id = param.otherParams.extensionId
			}

			params := createSmsParams(param.from, param.to, param.text, id)
			return rh.handle(api.CreateAccountSms(accountId, params))
		}

	case swagger.AvailablenumbersApi:

		switch command {

		case listAvailablePhoneNumbers:

			return rh.handle(api.ListAvailablePhoneNumbers(filterParams.filtersPhoneNumber, filterParams.filtersCountryCode, filterParams.filtersNpa, filterParams.filtersNxx, filterParams.filtersXxxx, filterParams.filtersCity, filterParams.filtersProvince, filterParams.filtersCountry, filterParams.filtersPrice, filterParams.filtersCategory, sortParams.sortInternal, sortParams.sortPrice, sortParams.sortPhoneNumber, limit, offset, fields))
		}

	case swagger.SubaccountsApi:

		switch command {

		case listSubaccounts:

			return rh.handle(api.ListAccountSubaccounts(accountId, filtersId, sortParams.sortId, limit, offset, fields))

		case createSubaccount:

			params := createSubaccountParams(input, param.contact, param.billingContact)
			return rh.handle(api.CreateAccountSubaccount(accountId, params))
		}

	case swagger.AccountsApi:

		switch command {

		case listAccounts:

			return rh.handle(api.ListAccounts(filtersId, sortParams.sortId, limit, offset, fields))

		case getAccount:

			return rh.handle(api.GetAccount(id))
		}

	case swagger.NumberregionsApi:

		switch command {

		case listAvailablePhoneNumberRegions:

			return rh.handle(api.ListAvailablePhoneNumberRegions(filterParams.filtersCountryCode, filterParams.filtersNpa, filterParams.filtersNxx, filterParams.filtersIsTollFree, filterParams.filtersCity, filterParams.filtersProvincePostalCode, filterParams.filtersCountryPostalCode, sortParams.sortCountryCode, sortParams.sortNpa, sortParams.sortNxx, sortParams.sortIsTollFree, sortParams.sortCity, sortParams.sortProvincePostalCode, sortParams.sortCountryPostalCode, limit, offset, fields, param.otherParams.groupBy))
		}

	case swagger.ApplicationsApi:

		switch command {

		case listApplications:

			return rh.handle(api.ListAccountApplications(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getApplication:

			return rh.handle(api.GetAccountApplication(accountId, id))
		}

	case swagger.CalllogsApi:

		switch command {

		case listCallLogs:

			return rh.handle(api.ListAccountCallLogs(accountId, filtersId, filterParams.filtersStartTime, filterParams.filtersCreatedAt, filterParams.filtersDirection, filterParams.filtersCalledNumber, filterParams.filtersType, filterParams.filtersExtension, sortParams.sortId, sortParams.sortStartTime, sortParams.sortCreatedAt, limit, offset, fields))
		}

	case swagger.CallsApi:

		switch command {

		case createCall:

			params := createCallParams(input)
			return rh.handle(api.CreateAccountCalls(accountId, params))

		}

	case swagger.DevicesApi:

		if param.otherParams.deviceId > 0 {
			id = param.otherParams.deviceId
		}

		switch command {

		case listDevices:

			return rh.handle(api.ListAccountDevices(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getDevice:

			return rh.handle(api.GetAccountDevice(accountId, id))

		case createDevice:

			params := createDeviceParams(input)
			return rh.handle(api.CreateAccountDevice(accountId, params))

		case replaceDevice:

			params := createDeviceParams(input)
			return rh.handle(api.ReplaceAccountDevice(accountId, id, params))
		}

	case swagger.ExpressservicecodesApi:

		if param.otherParams.codeId > 0 {
			id = param.otherParams.codeId
		}

		switch command {

		case listExpressServiceCodes:

			return rh.handle(api.ListAccountExpressSrvCodes(accountId, filtersId))

		case getExpressServiceCode:

			return rh.handle(api.GetAccountExpressSrvCode(accountId, id))
		}

	case swagger.ExtensionsApi:

		if param.otherParams.extensionId > 0 {
			id = param.otherParams.extensionId
		}

		switch command {

		case listExtensions:

			return rh.handle(api.ListAccountExtensions(accountId, filtersId, filterParams.filtersExtension, filterParams.filtersName, sortParams.sortId, sortParams.sortExtension, sortParams.sortName, limit, offset, fields))

		case getExtension:

			return rh.handle(api.GetAccountExtension(accountId, id))

		case createExtension:

			params := createExtensionsParams(input)
			return rh.handle(api.CreateAccountExtension(accountId, params))

		case replaceExtension:

			params := replaceExtensionParams(input)
			return rh.handle(api.ReplaceAccountExtension(accountId, id, params))

		}

	case swagger.CalleridsApi:

		switch command {

		case listCallerIds:
			return rh.handle(api.GetCallerIds(accountId, id, filterParams.filtersNumber, filterParams.filtersName, sortParams.sortNumber, sortParams.sortName, limit, offset, fields))
		}

	case swagger.ContactsApi:

		if param.otherParams.extensionId > 0 {
			id = param.otherParams.extensionId
		}

		if param.otherParams.contactId > 0 {
			idSecondary = param.otherParams.contactId
		}

		switch command {

		case listContacts:

			return rh.handle(api.ListAccountExtensionContacts(accountId, id, filtersId, filterParams.filtersGroupId, filterParams.filtersUpdatedAt, sortParams.sortId, sortParams.sortUpdatedAt, limit, offset, fields))

		case getContact:

			return rh.handle(api.GetAccountExtensionContact(accountId, id, idSecondary))

		case createContact:

			params := createContactParams(input)
			return rh.handle(api.CreateAccountExtensionContact(accountId, id, params))

		case replaceContact:

			params := createContactParams(input)
			return rh.handle(api.ReplaceAccountExtensionContact(accountId, id, idSecondary, params))

		case deleteContact:

			return rh.handle(api.DeleteAccountExtensionContact(accountId, id, idSecondary))
		}

	case swagger.GroupsApi:

		if param.otherParams.extensionId > 0 {
			id = param.otherParams.extensionId
		}

		if param.otherParams.groupId > 0 {
			idSecondary = param.otherParams.groupId
		}

		switch command {

		case listGroups:

			return rh.handle(api.ListAccountExtensionContactGroups(accountId, id, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getGroup:

			return rh.handle(api.GetAccountExtensionContactGroup(accountId, id, idSecondary))

		case createGroup:

			params := createGroupParams(input)
			return rh.handle(api.CreateAccountExtensionContactGroup(accountId, id, params))

		case replaceGroup:

			params := createGroupParams(input)
			return rh.handle(api.ReplaceAccountExtensionContactGroup(accountId, id, idSecondary, params))
		case deleteGroup:

			return rh.handle(api.DeleteAccountExtensionContactGroup(accountId, id, idSecondary))
		}

	case swagger.PhonenumbersApi:

		if param.otherParams.numberId > 0 {
			id = param.otherParams.numberId
		}

		switch command {

		case listPhoneNumbers:

			return rh.handle(api.ListAccountPhoneNumbers(accountId, filtersId, filterParams.filtersName, filterParams.filtersPhoneNumber, sortParams.sortId, sortParams.sortName, sortParams.sortPhoneNumber, limit, offset, fields))

		case getPhoneNumber:

			return rh.handle(api.GetAccountPhoneNumber(accountId, id))

		case createPhoneNumber:

			params := createPhoneNumberParams(input)
			return rh.handle(api.CreateAccountPhoneNumber(accountId, params))

		case replacePhoneNumber:

			params := replacePhoneNumberParams(input)
			return rh.handle(api.ReplaceAccountPhoneNumber(accountId, id, params))
		}

	case swagger.TrunksApi:

		if param.otherParams.trunkId > 0 {
			id = param.otherParams.trunkId
		}

		switch command {

		case listTrunks:

			return rh.handle(api.ListAccountTrunks(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

		case getTrunk:

			return rh.handle(api.GetAccountTrunk(accountId, id))

		case createTrunk:
			params := createTrunkParams(param.trunkName, param.trunkUri, param.trunkConcurrentCalls, param.trunkMaxMinutes)
			return rh.handle(api.CreateAccountTrunk(accountId, params))

		case replaceTrunk:

			params := createTrunkParams(param.trunkName, param.trunkUri, param.trunkConcurrentCalls, param.trunkMaxMinutes)
			return rh.handle(api.ReplaceAccountTrunk(accountId, id, params))

		case deleteTrunk:

			return rh.handle(api.DeleteAccountTrunk(accountId, id))
		}

	default:
		return nil, nil
	}

	return nil, nil
}
