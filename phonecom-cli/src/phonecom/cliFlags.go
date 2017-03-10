package main

import "github.com/urfave/cli"

var defaultId = ""
var defaultIdSecondary = ""
var defaultInput = ""
var defaultFrom = ""
var defaultTo = ""
var defaultText = ""
var defaultLimit = 25
var defaultOffset = 0

const defaultCommand = "https://api.phone.com/v4/ping"
const defaultAccountId = 1315091

const commandLong = "command"
const idLong = "id"
const idSecondaryLong = "id-secondary"
const limitLong = "limit"
const offsetLong = "offset"
const accountIdLong = "account"
const dryrunLong = "dryrun"
const inputLong = "input"
const verboseLong = "verbose-mode"
const contactLong = "contact"
const billingContactLong = "billing-contact"
const fromLong = "from"
const toLong = "to"
const textLong = "text"
const trunkNameLong = "name"
const trunkUriLong = "uri"
const maxConcurrentCallsLong = "max-concurrent-calls"
const maxMinutesPerMonthLong = "max-minutes-per-month"
const filtersTypeLong = "filtersType"
const filtersValueLong = "filtersValue"
const sortTypeLong = "sortType"
const sortValueLong = "sortValue"
const sampleInLong = "samplein"
const sampleOutLong = "sampleout"
const fullListLong = "fullList"
const inputFormatLong = "inputFormat"
const outputFormatLong = "outputFormat"
const apiKeyLong = "api-key"
const apiKeyPrefixLong = "api-key-prefix"

func getCliFlags() []cli.Flag {

	return []cli.Flag{

		cli.StringFlag{
			Name:  commandLong + ", c",
			Value: defaultCommand,
			Usage: "Phone.com API command that you want to execute. List of API endpoints: " + getAllCommands(),
		},
		cli.StringFlag{
			Name:  idLong,
			Value: defaultId,
			Usage: "ID of entity you want to operate",
		},
		cli.StringFlag{
			Name:  idSecondaryLong + ", is",
			Value: defaultIdSecondary,
			Usage: "Secondary ID of entity you want to operate",
		},
		cli.IntFlag{
			Name:  limitLong + ", l",
			Value: defaultLimit,
			Usage: "Upper limit of results you want to fetch",
		},
		cli.IntFlag{
			Name:  offsetLong + ", o",
			Value: defaultOffset,
			Usage: "Offset of results you want to fetch",
		},
		cli.IntFlag{
			Name:  accountIdLong + ", a",
			Value: defaultAccountId,
			Usage: "Phone.com API account to use in API calls",
		},
		cli.BoolFlag{
			Name:  dryrunLong + ", d",
			Usage: "Print the expected action without executing the API command",
		},
		cli.StringFlag{
			Name:  inputLong + ", i",
			Value: defaultInput,
			Usage: "Specify the path to the JSON file for making the API call",
		},
		cli.BoolFlag{
			Name:  verboseLong + ", vrm",
			Usage: "Activate verbose mode",
		},
		cli.StringFlag{
			Name:  contactLong,
			Usage: "Path to the local JSON descriptor to a contact object",
		},
		cli.StringFlag{
			Name:  billingContactLong,
			Usage: "Path to the local JSON descriptor to a billing contact object",
		},
		cli.StringFlag{
			Name:  fromLong,
			Value: defaultFrom,
			Usage: "The phone number of the SMS sender",
		},
		cli.StringFlag{
			Name:  toLong,
			Value: defaultTo,
			Usage: "The phone number of the SMS receiver",
		},
		cli.StringFlag{
			Name:  textLong,
			Value: defaultText,
			Usage: "SMS body",
		},
		cli.StringFlag{
			Name:  trunkNameLong,
			Usage: "Trunk name",
		},
		cli.StringFlag{
			Name:  trunkUriLong,
			Usage: "Trunk uri",
		},
		cli.IntFlag{
			Name:  maxConcurrentCallsLong,
			Value: 60,
			Usage: "Maximum concurrent calls for the trunk",
		},
		cli.IntFlag{
			Name:  maxMinutesPerMonthLong,
			Value: 800,
			Usage: "Maximum of minutes per month for the trunk",
		},
		cli.StringFlag{
			Name:  filtersTypeLong + ", ft",
			Usage: "Type of filter",
		},
		cli.StringFlag{
			Name:  filtersValueLong + ", fv",
			Usage: "The filter value",
		},
		cli.StringFlag{
			Name:  sortTypeLong + ", st",
			Usage: "Type of sort",
		},
		cli.StringFlag{
			Name:  sortValueLong + ", sv",
			Usage: "The sort value",
		},
		cli.StringFlag{
			Name:  sampleInLong + ", si",
			Usage: "Generate sample input json",
		},
		cli.StringFlag{
			Name:  sampleOutLong + ", so",
			Usage: "Generate sample output json",
		},
		cli.BoolFlag{
			Name:  fullListLong + ", fl",
			Usage: "Full list flag. If set, aditional information about the list is given",
		},
		cli.StringFlag{
			Name:  inputFormatLong + ", if",
			Usage: "Input format. Can be 'json' or 'xml'. Default is json",
		},
		cli.StringFlag{
			Name:  outputFormatLong + ", of",
			Usage: "Output format. Can be 'json' or 'csv'. Default is json",
		},
		cli.StringFlag{
			Name:  apiKeyLong + ", ak",
			Usage: "The API key for Phone.com",
		},
		cli.StringFlag{
			Name:  apiKeyPrefixLong + ", akp",
			Usage: "The API key prefix for Phone.com",
		},
	}
}

func getAllCommands() string {

	return "\n\n" +
		createCall + ", " +
		createDevice + ", " +
		createExtension + ", " +
		createMenu + ", " +
		createPhoneNumber + ", " +
		createQueue + ", " +
		createRoute + ", " +
		createSms + ", " +
		createSubaccount + ", " +
		createTrunk + ", " +
		createContact + ", " +
		createGroup + "\n\n" +

		deleteMenu + ", " +
		deleteQueue + ", " +
		deleteRoute + ", " +
		deleteTrunk + ", " +
		deleteContact + ", " +
		deleteGroup + "\n\n" +

		getAccount + ", " +
		getApplication + ", " +
		getCallLog + ", " +
		getDevice + ", " +
		getExpressServiceCode + ", " +
		getExtension + ", " +
		getRecording + ", " +
		getMenu + ", " +
		getPhoneNumber + ", " +
		getQueue + ", " +
		getRoute + ", " +
		getSchedule + ", " +
		getSms + ", " +
		getTrunk + ", " +
		getContact + ", " +
		getGroup + "\n\n" +

		listAccounts + ", " +
		listApplications + ", " +
		listCallLogs + ", " +
		listDevices + ", " +
		listExpressServiceCodes + ", " +
		listExtensions + ", " +
		listMedia + ", " +
		listMenus + ", " +
		listPhoneNumbers + ", " +
		listQueues + ", " +
		listRoutes + ", " +
		listSchedules + ", " +
		listSms + ", " +
		listSubaccounts + ", " +
		listTrunks + ", " +
		listCallerIds + ", " +
		listContacts + ", " +
		listGroups + ", " +
		listAvailablePhoneNumbers + ", " +
		listAvailablePhoneNumberRegions + "\n\n" +

		replaceDevice + ", " +
		replaceExtension + ", " +
		replaceMenu + ", " +
		replacePhoneNumber + ", " +
		replaceQueue + ", " +
		replaceRoute + ", " +
		replaceTrunk + ", " +
		replaceContact + ", " +
		replaceGroup + "\n\n"
}
