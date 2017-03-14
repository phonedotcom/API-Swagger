package main

import "github.com/waiyuen/Phone.com-API-SDK-go"

type ApiResolver struct {
	config  *swagger.Configuration
	command string
}

func (r *ApiResolver) resolve() interface{} {

	var api interface{}

	switch r.command {

	case listMedia, getMedia, createMedia:

		api = swagger.MediaApi{Configuration: r.config}

	case listMenus, getMenu, createMenu, replaceMenu, deleteMenu:

		api = swagger.MenusApi{Configuration: r.config}

	case listQueues, getQueue, createQueue, replaceQueue, deleteQueue:

		api = swagger.QueuesApi{Configuration: r.config}

	case listRoutes, getRoute, createRoute, replaceRoute, deleteRoute:

		api = swagger.RoutesApi{Configuration: r.config}

	case listSchedules, getSchedule:

		api = swagger.SchedulesApi{Configuration: r.config}

	case listSms, getSms, createSms:

		api = swagger.SmsApi{Configuration: r.config}

	case listAvailablePhoneNumbers:

		api = swagger.AvailablenumbersApi{Configuration: r.config}

	case listAvailablePhoneNumberRegions:

		api = swagger.NumberregionsApi{Configuration: r.config}

	case listSubaccounts, createSubaccount:

		api = swagger.SubaccountsApi{Configuration: r.config}

	case listAccounts, getAccount:

		api = swagger.AccountsApi{Configuration: r.config}

	case listApplications, getApplication:

		api = swagger.ApplicationsApi{Configuration: r.config}

	case listCallLogs, getCallLog:

		api = swagger.CalllogsApi{Configuration: r.config}

	case listDevices, getDevice, createDevice, replaceDevice:

		api = swagger.DevicesApi{Configuration: r.config}

	case listExpressServiceCodes, getExpressServiceCode:

		api = swagger.ExpressservicecodesApi{Configuration: r.config}

	case listExtensions, getExtension, createExtension, replaceExtension:

		api = swagger.ExtensionsApi{Configuration: r.config}

	case listCallerIds:

		api = swagger.CalleridsApi{Configuration: r.config}

	case listContacts, getContact, createContact, replaceContact, deleteContact:

		api = swagger.ContactsApi{Configuration: r.config}

	case listGroups, getGroup, createGroup, replaceGroup, deleteGroup:

		api = swagger.GroupsApi{Configuration: r.config}

	case listPhoneNumbers, getPhoneNumber, createPhoneNumber, replacePhoneNumber:

		api = swagger.PhonenumbersApi{Configuration: r.config}

	case listTrunks, getTrunk, createTrunk, deleteTrunk, replaceTrunk:

		api = swagger.TrunksApi{Configuration: r.config}

	case createCall:

		api = swagger.CallsApi{Configuration: r.config}

	default:
		return nil
	}

	return api
}
