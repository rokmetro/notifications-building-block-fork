/*
 *   Copyright (c) 2020 Board of Trustees of the University of Illinois.
 *   All rights reserved.

 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at

 *   http://www.apache.org/licenses/LICENSE-2.0

 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package main

import (
	"log"
	"notifications/core"
	"notifications/driven/firebase"
	storage "notifications/driven/storage"
	driver "notifications/driver/web"
	"os"
)

var (
	// Version : version of this executable
	Version string
	// Build : build date of this executable
	Build string
)

func main() {
	if len(Version) == 0 {
		Version = "dev"
	}

	port := getEnvKey("PORT", true)

	// mongoDB adapter
	mongoDBAuth := getEnvKey("MONGO_AUTH", true)
	mongoDBName := getEnvKey("MONGO_DATABASE", true)
	mongoTimeout := getEnvKey("MONGO_TIMEOUT", false)
	storageAdapter := storage.NewStorageAdapter(mongoDBAuth, mongoDBName, mongoTimeout)
	err := storageAdapter.Start()
	if err != nil {
		log.Fatal("Cannot start the mongoDB adapter - " + err.Error())
	}

	// firebase credentials
	firebaseProjectID := getEnvKey("FIREBASE_PROJECT_ID", true)
	firebaseAuth := getEnvKey("FIREBASE_AUTH", true)
	firebaseAdapter := firebase.NewFirebaseAdapter(firebaseAuth, firebaseProjectID)
	err = firebaseAdapter.Start()
	if err != nil {
		log.Fatal("Cannot start the Firebase adapter - " + err.Error())
	}

	// application
	application := core.NewApplication(Version, Build, storageAdapter, firebaseAdapter)
	application.Start()

	// web adapter
	host := getEnvKey("HOST", true)
	internalAPIKey := getEnvKey("INTERNAL_API_KEY", true)
	coreAuthPrivateKey := getEnvKey("ROKWIRE_CORE_AUTH_PRIVATE_KEY", true)

	webAdapter := driver.NewWebAdapter(host, port, application, internalAPIKey, coreAuthPrivateKey)

	webAdapter.Start()
}

func getEnvKey(key string, required bool) string {
	//get from the environment
	value, exist := os.LookupEnv(key)
	if !exist {
		if required {
			log.Fatal("No provided environment variable for " + key)
		} else {
			log.Printf("No provided environment variable for " + key)
		}
	}
	printEnvVar(key, value)
	return value
}

func printEnvVar(name string, value string) {
	if Version == "dev" {
		log.Printf("%s=%s", name, value)
	}
}
