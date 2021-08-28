//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListAllTemplates(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/templates/gitlab_ci_ymls", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
			   "content":"5-Minute-Production-App",
			   "name":"5-Minute-Production-App"
			},
			{
			   "content":"Android",
			   "name":"Android"
			},
			{
			   "content":"Android-Fastlane",
			   "name":"Android-Fastlane"
			},
			{
			   "content":"Auto-DevOps",
			   "name":"Auto-DevOps"
			}
		 ]`)
	})

	templates, _, err := client.CIYMLTemplate.ListAllTemplates(&ListCIYMLTemplatesOptions{})
	if err != nil {
		t.Errorf("CIYMLTemplates.ListAllTemplates returned error: %v", err)
	}

	want := []*CIYMLTemplate{
		{Name: "5-Minute-Production-App", Content: "5-Minute-Production-App"},
		{Name: "Android", Content: "Android"},
		{Name: "Android-Fastlane", Content: "Android-Fastlane"},
		{Name: "Auto-DevOps", Content: "Auto-DevOps"},
	}
	if !reflect.DeepEqual(want, templates) {
		t.Errorf("CIYMLTemplates.ListAllTemplates returned %+v, want %+v", templates, want)
	}
}

func TestGetTemplate(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/templates/gitlab_ci_ymls/Ruby", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{
			"name": "Ruby",
			"content": "# This file is a template, and might need editing before it works on your project.\n# Official language image. Look for the different tagged releases at:\n# https://hub.docker.com/r/library/ruby/tags/\nimage: \"ruby:2.5\"\n\n# Pick zero or more services to be used on all builds.\n# Only needed when using a docker container to run your tests in.\n# Check out: http://docs.gitlab.com/ee/ci/services/index.html\n  - mysql:latest\n  - redis:latest\n  - postgres:latest\n\nvariables:\n  POSTGRES_DB: database_name\n\n# Cache gems in between builds\ncache:\n  paths:\n    - vendor/ruby\n\n# This is a basic example for a gem or script which doesn't use\n# services such as redis or postgres\nbefore_script:\n  - ruby -v  # Print out ruby version for debugging\n  # Uncomment next line if your rails app needs a JS runtime:\n  # - apt-get update -q && apt-get install nodejs -yqq\n  - bundle install -j $(nproc) --path vendor  # Install dependencies into ./vendor/ruby\n\n# Optional - Delete if not using \nrubocop:\n  script:\n    - rubocop\n\nrspec:\n  script:\n    - rspec spec\n\nrails:\n  variables:\n    DATABASE_URL: \"postgresql://postgres:postgres@postgres:5432/$POSTGRES_DB\"\n  script:\n    - rails db:migrate\n    - rails db:seed\n    - rails test\n\n# This deploy job uses a simple deploy flow to Heroku, other providers, e.g. AWS Elastic Beanstalk\n# are supported too: https://github.com/travis-ci/dpl\ndeploy:\n  type: deploy\n  environment: production\n  script:\n    - gem install dpl\n    - dpl --provider=heroku --app=$HEROKU_APP_NAME --api-key=$HEROKU_PRODUCTION_KEY\n"
		  }`)
	})

	template, _, err := client.CIYMLTemplate.GetTemplate("Ruby")
	if err != nil {
		t.Errorf("CIYMLTemplates.GetTemplate returned error: %v", err)
	}

	want := &CIYMLTemplate{
		Name:    "Ruby",
		Content: "# This file is a template, and might need editing before it works on your project.\n# Official language image. Look for the different tagged releases at:\n# https://hub.docker.com/r/library/ruby/tags/\nimage: \"ruby:2.5\"\n\n# Pick zero or more services to be used on all builds.\n# Only needed when using a docker container to run your tests in.\n# Check out: http://docs.gitlab.com/ee/ci/services/index.html\n  - mysql:latest\n  - redis:latest\n  - postgres:latest\n\nvariables:\n  POSTGRES_DB: database_name\n\n# Cache gems in between builds\ncache:\n  paths:\n    - vendor/ruby\n\n# This is a basic example for a gem or script which doesn't use\n# services such as redis or postgres\nbefore_script:\n  - ruby -v  # Print out ruby version for debugging\n  # Uncomment next line if your rails app needs a JS runtime:\n  # - apt-get update -q && apt-get install nodejs -yqq\n  - bundle install -j $(nproc) --path vendor  # Install dependencies into ./vendor/ruby\n\n# Optional - Delete if not using \nrubocop:\n  script:\n    - rubocop\n\nrspec:\n  script:\n    - rspec spec\n\nrails:\n  variables:\n    DATABASE_URL: \"postgresql://postgres:postgres@postgres:5432/$POSTGRES_DB\"\n  script:\n    - rails db:migrate\n    - rails db:seed\n    - rails test\n\n# This deploy job uses a simple deploy flow to Heroku, other providers, e.g. AWS Elastic Beanstalk\n# are supported too: https://github.com/travis-ci/dpl\ndeploy:\n  type: deploy\n  environment: production\n  script:\n    - gem install dpl\n    - dpl --provider=heroku --app=$HEROKU_APP_NAME --api-key=$HEROKU_PRODUCTION_KEY\n",
	}
	if !reflect.DeepEqual(want, template) {
		t.Errorf("CIYMLTemplates.GetTemplate returned %+v, want %+v", template, want)
	}
}
