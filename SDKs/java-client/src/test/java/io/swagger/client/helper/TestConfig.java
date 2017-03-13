package io.swagger.client.helper;

import java.math.BigInteger;
import java.security.SecureRandom;
import java.util.ArrayList;
import java.util.List;
import java.util.Random;

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
        apiKey.setApiKey("Place your API key here");
	}

	public static List<String> createDefaultFilter() {

		List<String> filter = new ArrayList<>();
		filter.add(LESS_THAN_FILTER);
		filter.add(GREATER_THAN_FILTER);
		
		return filter;
	}
	
	public static String nextRandom() {
		SecureRandom random = new SecureRandom();
		return new BigInteger(80, random).toString(32);
	}
	
	public static int randomNumber(int min, int max) {
		Random rand = new Random();
		int randomNum = rand.nextInt((max - min) + 1) + min;
		return randomNum;
	}
}
