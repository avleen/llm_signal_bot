package main

func setupTestEnv() {
	test_config := map[string]string{
		"IMAGEDIR":          "test_images",
		"STATEDB":           "test_messages.db",
		"PHONE":             "+123456789",
		"URL":               "ws://localhost:8080",
		"REST_URL":          "http://localhost:8080",
		"MAX_AGE":           "168",
		"SUMMARY_PROVIDER":  "google",
		"LOCATION":          "us-central1",
		"PROJECT_ID":        "tmp-k8s-tutorial",
		"GOOGLE_TEXT_MODEL": "gemini-1.5-flash-001",
	}

	for envName, envValue := range test_config {
		config[envName] = envValue
	}
}
