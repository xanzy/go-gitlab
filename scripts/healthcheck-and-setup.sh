#!/usr/bin/env sh

# This script is intended to be used as a Docker HEALTHCHECK for the GitLab container.
# It prepares GitLab prior to running acceptance tests.
#
# This is a known workaround for docker-compose lacking lifecycle hooks.
# See: https://github.com/docker/compose/issues/1809#issuecomment-657815188

set -e

# Check for a successful HTTP status code from GitLab.
curl --silent --show-error --fail --output /dev/null 127.0.0.1:80

# Because this script runs on a regular health check interval,
# this file functions as a marker that tells us if initialization already finished.
done=/var/gitlab-acctest-initialized

test -f $done || {
  echo 'Initializing GitLab for acceptance tests'

  echo 'Creating access token'
  (
    printf 'go_gitlab_token = PersonalAccessToken.create('
    printf 'user_id: 1, '
    printf 'scopes: [:api, :read_user], '
    printf 'name: :go-gitlab);'
    printf "go_gitlab_token.set_token('ACCTEST1234567890123');"
    printf 'go_gitlab_token.save!;'
  ) | gitlab-rails console

  echo 'Creating an instance level template group with a simple template based on rails'
  (
    printf 'group_template = Group.new('
    printf 'name: :go-gitlab, '
    printf 'path: :go-gitlab);'
    printf 'group_template.save!;'
    printf 'application_settings = ApplicationSetting.find_by "";'
    printf 'application_settings.custom_project_templates_group_id = group_template.id;'
    printf 'application_settings.save!;'
    printf 'attrs = {'
    printf 'name: :myrails, '
    printf 'path: :myrails, '
    printf 'namespace_id: group_template.id, '
    printf 'template_name: :rails, '
    printf 'id: 999};'
    printf 'project = ::Projects::CreateService.new(User.find_by_username("root"), attrs).execute;'
    printf 'project.saved?;'
  ) | gitlab-rails console

  touch $done
}

echo 'GitLab is ready for acceptance tests'
