package io.swagger.client.helper;

import java.util.ArrayList;
import java.util.List;

import io.swagger.client.ApiClient;
import io.swagger.client.Configuration;
import io.swagger.client.auth.ApiKeyAuth;

public final class TestConfig {

	private TestConfig() { }
	
	public static String LESS_THAN_FILTER = "lt:100000";
	public static String GREATER_THAN_FILTER = "gt:1";

	public static void setAuthorization() {

		ApiClient defaultClient = Configuration.getDefaultApiClient();
        
        // Configure API key authorization: apiKey
        ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
        apiKey.setApiKeyPrefix("Bearer");
        apiKey.setApiKey("Place your access token here");
	}

	public static List<String> createDefaultFilter() {

		List<String> filter = new ArrayList<>();
		filter.add(LESS_THAN_FILTER);
		filter.add(GREATER_THAN_FILTER);
		
		return filter;
	}
}
